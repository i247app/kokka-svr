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
// signer can be nil for read-only operations (e.g., BalanceOf)
func NewTokenClient(client *Client, signer *TransactionSigner) (*TokenClient, error) {
	if client == nil {
		return nil, fmt.Errorf("blockchain client is required")
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

type AddressInfoResponse struct {
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Decimals     uint8  `json:"decimals"`
	TotalSupply  string `json:"total_supply"`
	OwnerAddress string `json:"owner_address"`
	OwnerBalance string `json:"owner_balance"`
}

// AddressInfo retrieves basic information about the token contract at the given address
func (v *TokenClient) AddressInfo(ctx context.Context, address string) (*AddressInfoResponse, error) {

	// ERC20 ABI
	const erc20ABI = `[
  		{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"type":"function"},
  		{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"type":"function"},
  		{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"type":"function"},
  		{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"type":"function"},
  		{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"type":"function"},
  		{"constant":true,"inputs":[],"name":"owner","outputs":[{"name":"","type":"address"}],"type":"function"}
	]`

	contractAddress := common.HexToAddress(address)

	parsedABI, _ := abi.JSON(strings.NewReader(erc20ABI))

	call := func(method string, args ...interface{}) string {
		data, _ := parsedABI.Pack(method, args...)
		res, _ := v.client.CallContract(ctx, contractAddress.Hex(), hexutil.Encode(data), "latest")
		return res
	}

	// Decode the results
	var name string
	err := parsedABI.UnpackIntoInterface(&name, "name", common.FromHex(call("name")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode name: %w", err)
	}

	var symbol string
	err = parsedABI.UnpackIntoInterface(&symbol, "symbol", common.FromHex(call("symbol")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode symbol: %w", err)
	}

	var decimals uint8
	err = parsedABI.UnpackIntoInterface(&decimals, "decimals", common.FromHex(call("decimals")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode decimals: %w", err)
	}

	var totalSupply *big.Int
	err = parsedABI.UnpackIntoInterface(&totalSupply, "totalSupply", common.FromHex(call("totalSupply")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode totalSupply: %w", err)
	}

	var owner common.Address
	err = parsedABI.UnpackIntoInterface(&owner, "owner", common.FromHex(call("owner")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode owner: %w", err)
	}

	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", common.FromHex(call("balanceOf", owner)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode balanceOf: %w", err)
	}

	return &AddressInfoResponse{
		Name:         name,
		Symbol:       symbol,
		Decimals:     decimals,
		TotalSupply:  totalSupply.String(),
		OwnerAddress: owner.Hex(),
		OwnerBalance: balance.String(),
	}, nil
}
