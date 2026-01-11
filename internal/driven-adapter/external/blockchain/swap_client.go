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
func (s *SwapClient) SwapAforB(ctx context.Context, contractAddress string, amountIn *big.Int) (string, error) {
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
		To:   contractAddress,
		Data: hexutil.Encode(data),
	}

	// Sign and send the transaction
	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send swapAforB transaction: %w", err)
	}

	return txHash, nil
}

// SwapBforA executes a swap from token B to token A
func (s *SwapClient) SwapBforA(ctx context.Context, contractAddress string, amountIn *big.Int) (string, error) {
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
		To:   contractAddress,
		Data: hexutil.Encode(data),
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
