package services

import (
	"context"
	"fmt"
	"math/big"

	"kokka.com/kokka/internal/applications/dtos"
	"kokka.com/kokka/internal/applications/validators"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain"
	"kokka.com/kokka/internal/shared/utils/crypto"
)

// SwapService handles swap business logic
type SwapService struct {
	validator           validators.ISwapValidator
	client              *blockchain.Client
	decryptionKey       string
	readOnlySwapClient  *blockchain.SwapClient
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
	privateKey, err := crypto.DecryptCrypto(req.EncryptedPrivateKey, s.decryptionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt private key: %w", err)
	}

	// Create transaction signer
	signer, err := blockchain.NewTransactionSigner(privateKey, s.client)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction signer: %w", err)
	}

	// Create swap client
	swapClient, err := blockchain.NewSwapClient(s.client, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create swap client: %w", err)
	}

	// Execute swap transaction based on direction
	var txHash string
	var amountOut *big.Int

	if req.Direction == "AtoB" {
		// Get expected output amount before swapping
		amountOut, err = swapClient.GetAmountOutAforB(ctx, req.ContractAddress, amountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to get quote for AtoB swap: %w", err)
		}

		// Execute swap A for B
		txHash, err = swapClient.SwapAforB(ctx, req.ContractAddress, amountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to execute AtoB swap: %w", err)
		}
	} else { // BtoA
		// Get expected output amount before swapping
		amountOut, err = swapClient.GetAmountOutBforA(ctx, req.ContractAddress, amountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to get quote for BtoA swap: %w", err)
		}

		// Execute swap B for A
		txHash, err = swapClient.SwapBforA(ctx, req.ContractAddress, amountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to execute BtoA swap: %w", err)
		}
	}

	// Get token addresses
	tokenA, err := swapClient.GetTokenA(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokenA address: %w", err)
	}

	tokenB, err := swapClient.GetTokenB(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokenB address: %w", err)
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
		TxHash:          txHash,
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
