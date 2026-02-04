package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"kokka.com/kokka/internal/applications/dtos"
	"kokka.com/kokka/internal/applications/validators"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain"
	"kokka.com/kokka/internal/shared/utils"
)

// SwapService handles swap business logic
type SwapService struct {
	validator          validators.ISwapValidator
	client             *blockchain.Client
	decryptionKey      string
	readOnlySwapClient *blockchain.SwapClient
}

// NewSwapService creates a new swap service
func NewSwapService(
	validator validators.ISwapValidator,
	client *blockchain.Client,
	decryptionKey string,
) (*SwapService, error) {
	// Create read-only swap client for quote queries (no signer needed)
	readOnlyClient, err := blockchain.NewSwapClient(client, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create read-only swap client: %w", err)
	}

	return &SwapService{
		validator:          validator,
		client:             client,
		decryptionKey:      decryptionKey,
		readOnlySwapClient: readOnlyClient,
	}, nil
}

// Swap executes a token swap
func (s *SwapService) Swap(ctx context.Context, req *dtos.SwapTokenRequest) (*dtos.SwapTokenResponse, error) {
	// Validate request
	if err := s.validator.ValidateSwapTokenRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amountIn, err := parseAmount(req.AmountIn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount_in: %w", err)
	}

	// Decrypt private key
	privateKey, err := utils.DecryptCrypto(req.EncryptedPrivateKey, s.decryptionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt private key: %w", err)
	}

	// Create transaction signer
	signer, err := blockchain.NewTransactionSigner(privateKey, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction signer: %w", err)
	}

	// Get the current nonce for manual nonce management
	address := signer.GetAddress()
	nonceHex, err := s.client.GetTransactionCount(ctx, address, "pending")
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction count: %w", err)
	}

	// Calculate nonce N+1 for second transaction
	nonceStr := nonceHex
	if len(nonceStr) > 2 && nonceStr[:2] == "0x" {
		nonceStr = nonceStr[2:]
	}
	firstNonce := new(big.Int)
	firstNonce.SetString(nonceStr, 16)
	secondNonce := new(big.Int).Add(firstNonce, big.NewInt(1))
	secondNonceHex := "0x" + secondNonce.Text(16)

	// Create swap client (read-only first to get token addresses)
	swapClient, err := blockchain.NewSwapClient(s.client, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create swap client: %w", err)
	}

	// Get token addresses from swap contract
	tokenA, err := swapClient.GetTokenA(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokenA address: %w", err)
	}

	tokenB, err := swapClient.GetTokenB(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokenB address: %w", err)
	}

	// Determine which token to approve based on direction
	var tokenToApprove string
	if req.Direction == "AtoB" {
		tokenToApprove = tokenA // Swapping FROM tokenA
	} else {
		tokenToApprove = tokenB // Swapping FROM tokenB
	}

	// Create token client
	tokenClient, err := blockchain.NewTokenClient(s.client, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client: %w", err)
	}

	// Get addresses for gas estimation
	ownerAddr := signer.GetAddressAsCommon()
	spenderAddr := common.HexToAddress(req.ContractAddress)

	// Estimate approve gas
	approveGasLimit, err := tokenClient.EstimateApproveGas(ctx, tokenToApprove, req.ContractAddress, amountIn, nonceHex, ownerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate approve gas: %w", err)
	}

	// Estimate swap gas and get quote based on direction
	var swapGasLimit string
	var amountOut *big.Int

	if req.Direction == "AtoB" {
		// Get expected output amount
		amountOut, err = swapClient.GetAmountOutAforB(ctx, req.ContractAddress, amountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to get quote for AtoB swap: %w", err)
		}

		// Estimate swap gas with state override
		swapGasLimit, err = swapClient.EstimateSwapAforBGasWithAllowanceOverride(ctx, req.ContractAddress, tokenA, amountIn, secondNonceHex, ownerAddr, spenderAddr)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate swapAforB gas: %w", err)
		}
	} else { // BtoA
		// Get expected output amount
		amountOut, err = swapClient.GetAmountOutBforA(ctx, req.ContractAddress, amountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to get quote for BtoA swap: %w", err)
		}

		// Estimate swap gas with state override
		swapGasLimit, err = swapClient.EstimateSwapBforAGasWithAllowanceOverride(ctx, req.ContractAddress, tokenB, amountIn, secondNonceHex, ownerAddr, spenderAddr)
		if err != nil {
			return nil, fmt.Errorf("failed to estimate swapBforA gas: %w", err)
		}
	}

	// Send both transactions in parallel with pre-estimated gas
	type txResult struct {
		txHash string
		err    error
	}

	approveChan := make(chan txResult, 1)
	swapChan := make(chan txResult, 1)

	// Goroutine 1: Send approve (nonce N)
	go func() {
		txHash, err := tokenClient.Approve(ctx, tokenToApprove, req.ContractAddress, amountIn, nonceHex, approveGasLimit)
		approveChan <- txResult{txHash: txHash, err: err}
	}()

	// Goroutine 2: Send swap (nonce N+1)
	go func() {
		var txHash string
		var err error
		if req.Direction == "AtoB" {
			txHash, err = swapClient.SwapAforBWithGasLimit(ctx, req.ContractAddress, tokenA, amountIn, secondNonceHex, swapGasLimit)
		} else {
			txHash, err = swapClient.SwapBforAWithGasLimit(ctx, req.ContractAddress, tokenB, amountIn, secondNonceHex, swapGasLimit)
		}
		swapChan <- txResult{txHash: txHash, err: err}
	}()

	swapResult := <-swapChan
	if swapResult.err != nil {
		return nil, fmt.Errorf("failed to send swap transaction: %w", swapResult.err)
	}

	// Determine from/to tokens based on direction
	var fromToken, toToken string
	if req.Direction == "AtoB" {
		fromToken = tokenA
		toToken = tokenB
	} else {
		fromToken = tokenB
		toToken = tokenA
	}

	return &dtos.SwapTokenResponse{
		TxHash:          swapResult.txHash,
		ContractAddress: req.ContractAddress,
		AmountIn:        amountIn.String(),
		AmountOut:       amountOut.String(),
		FromToken:       fromToken,
		ToToken:         toToken,
		Direction:       req.Direction,
	}, nil
}

// GetQuote returns a quote for a swap without executing it
func (s *SwapService) GetQuote(ctx context.Context, req *dtos.GetSwapQuoteRequest) (*dtos.GetSwapQuoteResponse, error) {
	// Validate request
	if err := s.validator.ValidateGetSwapQuoteRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amountIn, err := parseAmount(req.AmountIn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount_in: %w", err)
	}

	// Get expected output amount based on direction
	var amountOut *big.Int
	if req.Direction == "AtoB" {
		amountOut, err = s.readOnlySwapClient.GetAmountOutAforB(ctx, req.ContractAddress, amountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to get quote for AtoB: %w", err)
		}
	} else { // BtoA
		amountOut, err = s.readOnlySwapClient.GetAmountOutBforA(ctx, req.ContractAddress, amountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to get quote for BtoA: %w", err)
		}
	}

	// Get exchange rate
	exchangeRate, err := s.readOnlySwapClient.GetExchangeRate(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	return &dtos.GetSwapQuoteResponse{
		ContractAddress: req.ContractAddress,
		AmountIn:        amountIn.String(),
		AmountOut:       amountOut.String(),
		Direction:       req.Direction,
		ExchangeRate:    exchangeRate.String(),
	}, nil
}

// GetSwapInfo returns information about a swap contract
func (s *SwapService) GetSwapInfo(ctx context.Context, req *dtos.GetSwapInfoRequest) (*dtos.GetSwapInfoResponse, error) {
	// Validate request
	if err := s.validator.ValidateGetSwapInfoRequest(req); err != nil {
		return nil, err
	}

	// Get token addresses
	tokenA, err := s.readOnlySwapClient.GetTokenA(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokenA address: %w", err)
	}

	tokenB, err := s.readOnlySwapClient.GetTokenB(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokenB address: %w", err)
	}

	// Get reserves
	reserveA, reserveB, err := s.readOnlySwapClient.GetReserves(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get reserves: %w", err)
	}

	// Get exchange rate
	exchangeRate, err := s.readOnlySwapClient.GetExchangeRate(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	return &dtos.GetSwapInfoResponse{
		ContractAddress: req.ContractAddress,
		TokenA:          tokenA,
		TokenB:          tokenB,
		ReserveA:        reserveA.String(),
		ReserveB:        reserveB.String(),
		ExchangeRate:    exchangeRate.String(),
	}, nil
}
