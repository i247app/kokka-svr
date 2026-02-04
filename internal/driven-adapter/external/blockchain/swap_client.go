package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain/gen/swap"
)

// SwapClient handles interactions with swap contracts
type SwapClient struct {
	client *Client
	signer *TransactionSigner
	abi    abi.ABI
}

// NewSwapClient creates a new swap client
// signer can be nil for read-only operations (e.g., GetQuote, GetSwapInfo)
func NewSwapClient(client *Client, signer *TransactionSigner) (*SwapClient, error) {
	if client == nil {
		return nil, fmt.Errorf("blockchain client is required")
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(swap.SwapMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse swap contract ABI: %w", err)
	}

	return &SwapClient{
		client: client,
		signer: signer,
		abi:    parsedABI,
	}, nil
}

// SwapAforB executes a swap from token A to token B
func (s *SwapClient) SwapAforB(ctx context.Context, contractAddress string, amountIn *big.Int, nonce string) (string, error) {
	if s.signer == nil {
		return "", fmt.Errorf("signer is required for swap operations")
	}

	// Encode the swapAforB function call
	data, err := s.abi.Pack("swapAforB", amountIn)
	if err != nil {
		return "", fmt.Errorf("failed to encode swapAforB call: %w", err)
	}

	// Prepare transaction request
	txReq := &SignTransactionRequest{
		To:    contractAddress,
		Data:  hexutil.Encode(data),
		Nonce: nonce,
	}

	// Sign and send the transaction
	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send swapAforB transaction: %w", err)
	}

	return txHash, nil
}

// SwapBforA executes a swap from token B to token A
func (s *SwapClient) SwapBforA(ctx context.Context, contractAddress string, amountIn *big.Int, nonce string) (string, error) {
	if s.signer == nil {
		return "", fmt.Errorf("signer is required for swap operations")
	}

	// Encode the swapBforA function call
	data, err := s.abi.Pack("swapBforA", amountIn)
	if err != nil {
		return "", fmt.Errorf("failed to encode swapBforA call: %w", err)
	}

	// Prepare transaction request
	txReq := &SignTransactionRequest{
		To:    contractAddress,
		Data:  hexutil.Encode(data),
		Nonce: nonce,
	}

	// Sign and send the transaction
	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send swapBforA transaction: %w", err)
	}

	return txHash, nil
}

// GetAmountOutAforB returns the expected output amount for swapping A to B
func (s *SwapClient) GetAmountOutAforB(ctx context.Context, contractAddress string, amountIn *big.Int) (*big.Int, error) {
	// Encode the getAmountOutAforB function call
	data, err := s.abi.Pack("getAmountOutAforB", amountIn)
	if err != nil {
		return nil, fmt.Errorf("failed to encode getAmountOutAforB call: %w", err)
	}

	// Call the contract (read-only)
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return nil, fmt.Errorf("failed to call getAmountOutAforB: %w", err)
	}

	// Decode the result
	var amountOut *big.Int
	err = s.abi.UnpackIntoInterface(&amountOut, "getAmountOutAforB", common.FromHex(result))
	if err != nil {
		return nil, fmt.Errorf("failed to decode getAmountOutAforB result: %w", err)
	}

	return amountOut, nil
}

// GetAmountOutBforA returns the expected output amount for swapping B to A
func (s *SwapClient) GetAmountOutBforA(ctx context.Context, contractAddress string, amountIn *big.Int) (*big.Int, error) {
	// Encode the getAmountOutBforA function call
	data, err := s.abi.Pack("getAmountOutBforA", amountIn)
	if err != nil {
		return nil, fmt.Errorf("failed to encode getAmountOutBforA call: %w", err)
	}

	// Call the contract (read-only)
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return nil, fmt.Errorf("failed to call getAmountOutBforA: %w", err)
	}

	// Decode the result
	var amountOut *big.Int
	err = s.abi.UnpackIntoInterface(&amountOut, "getAmountOutBforA", common.FromHex(result))
	if err != nil {
		return nil, fmt.Errorf("failed to decode getAmountOutBforA result: %w", err)
	}

	return amountOut, nil
}

// GetReserves returns the reserves of both tokens in the swap contract
func (s *SwapClient) GetReserves(ctx context.Context, contractAddress string) (reserveA *big.Int, reserveB *big.Int, err error) {
	// Encode the getReserves function call
	data, err := s.abi.Pack("getReserves")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode getReserves call: %w", err)
	}

	// Call the contract (read-only)
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to call getReserves: %w", err)
	}

	// Decode the result - getReserves returns (uint256 reserveA, uint256 reserveB)
	var reserves struct {
		ReserveA *big.Int
		ReserveB *big.Int
	}
	err = s.abi.UnpackIntoInterface(&reserves, "getReserves", common.FromHex(result))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode getReserves result: %w", err)
	}

	return reserves.ReserveA, reserves.ReserveB, nil
}

// GetExchangeRate returns the current exchange rate
func (s *SwapClient) GetExchangeRate(ctx context.Context, contractAddress string) (*big.Int, error) {
	// Encode the exchangeRate function call
	data, err := s.abi.Pack("exchangeRate")
	if err != nil {
		return nil, fmt.Errorf("failed to encode exchangeRate call: %w", err)
	}

	// Call the contract (read-only)
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return nil, fmt.Errorf("failed to call exchangeRate: %w", err)
	}

	// Decode the result
	var exchangeRate *big.Int
	err = s.abi.UnpackIntoInterface(&exchangeRate, "exchangeRate", common.FromHex(result))
	if err != nil {
		return nil, fmt.Errorf("failed to decode exchangeRate result: %w", err)
	}

	return exchangeRate, nil
}

// GetTokenA returns the address of token A
func (s *SwapClient) GetTokenA(ctx context.Context, contractAddress string) (string, error) {
	// Encode the tokenA function call
	data, err := s.abi.Pack("tokenA")
	if err != nil {
		return "", fmt.Errorf("failed to encode tokenA call: %w", err)
	}

	// Call the contract (read-only)
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return "", fmt.Errorf("failed to call tokenA: %w", err)
	}

	// Decode the result
	var tokenAddress common.Address
	err = s.abi.UnpackIntoInterface(&tokenAddress, "tokenA", common.FromHex(result))
	if err != nil {
		return "", fmt.Errorf("failed to decode tokenA result: %w", err)
	}

	return tokenAddress.Hex(), nil
}

// GetTokenB returns the address of token B
func (s *SwapClient) GetTokenB(ctx context.Context, contractAddress string) (string, error) {
	// Encode the tokenB function call
	data, err := s.abi.Pack("tokenB")
	if err != nil {
		return "", fmt.Errorf("failed to encode tokenB call: %w", err)
	}

	// Call the contract (read-only)
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return "", fmt.Errorf("failed to call tokenB: %w", err)
	}

	// Decode the result
	var tokenAddress common.Address
	err = s.abi.UnpackIntoInterface(&tokenAddress, "tokenB", common.FromHex(result))
	if err != nil {
		return "", fmt.Errorf("failed to decode tokenB result: %w", err)
	}

	return tokenAddress.Hex(), nil
}

// EstimateSwapAforBGas estimates the gas required for a swapAforB transaction
func (s *SwapClient) EstimateSwapAforBGas(ctx context.Context, contractAddress string, amountIn *big.Int, nonce string, owner common.Address) (string, error) {
	// Encode the swapAforB function call
	data, err := s.abi.Pack("swapAforB", amountIn)
	if err != nil {
		return "", fmt.Errorf("failed to encode swapAforB call: %w", err)
	}

	// Estimate gas
	gasLimit, err := s.client.EstimateGas(
		ctx,
		owner.Hex(),
		contractAddress,
		"",
		hexutil.Encode(data),
		nonce,
	)
	if err != nil {
		return "", fmt.Errorf("failed to estimate swapAforB gas: %w", err)
	}

	return gasLimit, nil
}

func (s *SwapClient) EstimateSwapBforAGas(ctx context.Context, contractAddress string, amountIn *big.Int, nonce string, owner common.Address) (string, error) {
	// Encode the swapBforA function call
	data, err := s.abi.Pack("swapBforA", amountIn)
	if err != nil {
		return "", fmt.Errorf("failed to encode swapBforA call: %w", err)
	}

	// Estimate gas
	gasLimit, err := s.client.EstimateGas(
		ctx,
		owner.Hex(),
		contractAddress,
		"",
		hexutil.Encode(data),
		nonce,
	)
	if err != nil {
		return "", fmt.Errorf("failed to estimate swapBforA gas: %w", err)
	}

	return gasLimit, nil
}

func (s *SwapClient) EstimateSwapAforBGasWithAllowanceOverride(ctx context.Context, contractAddress string, tokenAAddress string, amountIn *big.Int, nonce string, owner common.Address, spender common.Address) (string, error) {
	// Encode the swapAforB function call
	data, err := s.abi.Pack("swapAforB", amountIn)
	if err != nil {
		return "", fmt.Errorf("failed to encode swapAforB call: %w", err)
	}

	slotCandidates := []uint64{2, 1, 3, 0, 4, 5}

	var lastErr error
	for _, slot := range slotCandidates {
		// Calculate allowance storage slot for allowance mapping
		allowanceSlot := s.client.CalculateAllowanceSlot(ctx, owner, spender, slot)

		// Convert amount to hex
		amountHex := common.BytesToHash(common.LeftPadBytes(amountIn.Bytes(), 32)).Hex()

		stateOverride := map[string]interface{}{
			tokenAAddress: map[string]interface{}{
				"stateDiff": map[string]interface{}{
					allowanceSlot.Hex(): amountHex,
				},
			},
		}

		// Estimate gas with state override
		gasLimit, err := s.client.EstimateGasWithStateOverride(ctx, owner.Hex(), contractAddress, "", hexutil.Encode(data), nonce, stateOverride)
		if err == nil {
			return gasLimit, nil
		}

		lastErr = err
	}

	return "", fmt.Errorf("failed to estimate swapAforB gas with allowance override: %w", lastErr)
}

func (s *SwapClient) EstimateSwapBforAGasWithAllowanceOverride(ctx context.Context, contractAddress string, tokenBAddress string, amountIn *big.Int, nonce string, owner common.Address, spender common.Address) (string, error) {
	// Encode the swapBforA function call
	data, err := s.abi.Pack("swapBforA", amountIn)
	if err != nil {
		return "", fmt.Errorf("failed to encode swapBforA call: %w", err)
	}

	slotCandidates := []uint64{2, 1, 3, 0, 4, 5}

	var lastErr error
	for _, slot := range slotCandidates {
		// Calculate allowance storage slot for allowance mapping
		allowanceSlot := s.client.CalculateAllowanceSlot(ctx, owner, spender, slot)

		// Convert amount to hex
		amountHex := common.BytesToHash(common.LeftPadBytes(amountIn.Bytes(), 32)).Hex()

		stateOverride := map[string]interface{}{
			tokenBAddress: map[string]interface{}{
				"stateDiff": map[string]interface{}{
					allowanceSlot.Hex(): amountHex,
				},
			},
		}

		// Estimate gas with state override
		gasLimit, err := s.client.EstimateGasWithStateOverride(ctx, owner.Hex(), contractAddress, "", hexutil.Encode(data), nonce, stateOverride)
		if err == nil {
			return gasLimit, nil
		}

		lastErr = err
	}

	return "", fmt.Errorf("failed to estimate swapBforA gas with allowance override: %w", lastErr)
}

func (s *SwapClient) SwapAforBWithGasLimit(ctx context.Context, contractAddress string, tokenAAddress string, amountIn *big.Int, nonce string, gasLimit string) (string, error) {
	if s.signer == nil {
		return "", fmt.Errorf("signer is required for swap operations")
	}

	data, err := s.abi.Pack("swapAforB", amountIn)
	if err != nil {
		return "", fmt.Errorf("failed to encode swapAforB call: %w", err)
	}

	txReq := &SignTransactionRequest{
		To:       contractAddress,
		Data:     hexutil.Encode(data),
		Nonce:    nonce,
		GasLimit: gasLimit,
	}

	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send swapAforB transaction: %w", err)
	}

	return txHash, nil
}

func (s *SwapClient) SwapBforAWithGasLimit(ctx context.Context, contractAddress string, tokenBAddress string, amountIn *big.Int, nonce string, gasLimit string) (string, error) {
	if s.signer == nil {
		return "", fmt.Errorf("signer is required for swap operations")
	}

	data, err := s.abi.Pack("swapBforA", amountIn)
	if err != nil {
		return "", fmt.Errorf("failed to encode swapBforA call: %w", err)
	}

	txReq := &SignTransactionRequest{
		To:       contractAddress,
		Data:     hexutil.Encode(data),
		Nonce:    nonce,
		GasLimit: gasLimit,
	}

	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send swapBforA transaction: %w", err)
	}

	return txHash, nil
}
