package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain/vndx"
)

// TokenClient handles interactions with ERC20 token contracts
type TokenClient struct {
	client *Client
	signer *TransactionSigner
	abi    abi.ABI
}

// NewTokenClient creates a new token client
func NewTokenClient(client *Client, signer *TransactionSigner) (*TokenClient, error) {
	if client == nil {
		return nil, fmt.Errorf("blockchain client is required")
	}

	if signer == nil {
		return nil, fmt.Errorf("transaction signer is required")
	}

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(vndx.VNDXMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	return &TokenClient{
		client: client,
		signer: signer,
		abi:    parsedABI,
	}, nil
}

// Mint mints new tokens to a specified address
func (v *TokenClient) Mint(ctx context.Context, contractAddress string, to string, amount *big.Int) (string, error) {
	// Encode the mint function call
	data, err := v.abi.Pack("mint", common.HexToAddress(to), amount)
	if err != nil {
		return "", fmt.Errorf("failed to encode mint call: %w", err)
	}

	// Prepare transaction request
	txReq := &SignTransactionRequest{
		To:   contractAddress,
		Data: hexutil.Encode(data),
	}

	// Sign and send the transaction
	txHash, err := v.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send mint transaction: %w", err)
	}

	return txHash, nil
}

// Burn burns tokens from the caller's account
func (v *TokenClient) Burn(ctx context.Context, contractAddress string, amount *big.Int) (string, error) {
	// Encode the burn function call
	data, err := v.abi.Pack("burn", amount)
	if err != nil {
		return "", fmt.Errorf("failed to encode burn call: %w", err)
	}

	// Prepare transaction request
	txReq := &SignTransactionRequest{
		To:   contractAddress,
		Data: hexutil.Encode(data),
	}

	// Sign and send the transaction
	txHash, err := v.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send burn transaction: %w", err)
	}

	return txHash, nil
}

// Transfer transfers tokens to a specified address
func (v *TokenClient) Transfer(ctx context.Context, contractAddress string, to string, amount *big.Int) (string, error) {
	// Encode the transfer function call
	data, err := v.abi.Pack("transfer", common.HexToAddress(to), amount)
	if err != nil {
		return "", fmt.Errorf("failed to encode transfer call: %w", err)
	}

	// Prepare transaction request
	txReq := &SignTransactionRequest{
		To:   contractAddress,
		Data: hexutil.Encode(data),
	}

	// Sign and send the transaction
	txHash, err := v.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send transfer transaction: %w", err)
	}

	return txHash, nil
}

// BalanceOf returns the token balance of an address
func (v *TokenClient) BalanceOf(ctx context.Context, contractAddress string, address string) (*big.Int, error) {
	// Encode the balanceOf function call
	data, err := v.abi.Pack("balanceOf", common.HexToAddress(address))
	if err != nil {
		return nil, fmt.Errorf("failed to encode balanceOf call: %w", err)
	}

	// Call the contract (read-only)
	result, err := v.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return nil, fmt.Errorf("failed to call balanceOf: %w", err)
	}

	// Decode the result
	var balance *big.Int
	err = v.abi.UnpackIntoInterface(&balance, "balanceOf", common.FromHex(result))
	if err != nil {
		return nil, fmt.Errorf("failed to decode balanceOf result: %w", err)
	}

	return balance, nil
}
