package blockchain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"kokka.com/kokka/internal/driven-adapter/external/provider/sign"
	"kokka.com/kokka/internal/shared/http_client"
)

// Client is a JSON-RPC client for blockchain interactions
type Client struct {
	httpClient *http_client.Client
	config     *Config
	requestID  int64
}

// NewClient creates a new blockchain JSON-RPC client
func NewClient(config *Config) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	// Create HTTP client with blockchain-specific configuration
	httpClient := http_client.NewClient(
		http_client.WithBaseURL(config.BaseURL),
		http_client.WithTimeout(config.Timeout),
		http_client.WithRetry(config.MaxRetries, config.RetryDelay, 500, 502, 503, 504),
		http_client.WithHeader("Content-Type", "application/json"),
	)

	return &Client{
		httpClient: httpClient,
		config:     config,
		requestID:  0,
	}
}

// Call executes a JSON-RPC method call
func (c *Client) Call(ctx context.Context, method string, params interface{}) (*JSONRPCResponse, error) {
	// Build JSON-RPC request
	request := JSONRPCRequest{
		ID:      atomic.AddInt64(&c.requestID, 1),
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	// Execute HTTP POST request
	resp, err := c.httpClient.Post(ctx, "", request)
	if err != nil {
		return nil, fmt.Errorf("failed to execute JSON-RPC request: %w", err)
	}

	// Check HTTP status
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("JSON-RPC request failed with status %d: %s", resp.StatusCode, resp.String())
	}

	// Parse JSON-RPC response
	var jsonRPCResp JSONRPCResponse
	if err := resp.JSON(&jsonRPCResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON-RPC response: %w", err)
	}

	// Check for JSON-RPC errors
	if jsonRPCResp.IsError() {
		return &jsonRPCResp, fmt.Errorf("JSON-RPC error %d: %s", jsonRPCResp.Error.Code, jsonRPCResp.Error.Message)
	}

	return &jsonRPCResp, nil
}

// GetBlockNumber returns the current block number
func (c *Client) GetBlockNumber(ctx context.Context) (string, error) {
	resp, err := c.Call(ctx, "eth_blockNumber", []interface{}{})
	if err != nil {
		return "", fmt.Errorf("failed to get block number: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse block number: %w", err)
	}

	return result, nil
}

// GetBalance returns the balance of an address at a given block
func (c *Client) GetBalance(ctx context.Context, address string, block string) (string, error) {
	params := []interface{}{address, block}
	resp, err := c.Call(ctx, "eth_getBalance", params)
	if err != nil {
		return "", fmt.Errorf("failed to get balance: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse balance: %w", err)
	}

	return result, nil
}

// GetTransactionByHash returns transaction details by hash
func (c *Client) GetTransactionByHash(ctx context.Context, txHash string) (*JSONRPCResponse, error) {
	params := []interface{}{txHash}
	resp, err := c.Call(ctx, "eth_getTransactionByHash", params)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	return resp, nil
}

// GetBlockByNumber returns block details by number
func (c *Client) GetBlockByNumber(ctx context.Context, blockNumber string, fullTx bool) (*JSONRPCResponse, error) {
	params := []interface{}{blockNumber, fullTx}
	resp, err := c.Call(ctx, "eth_getBlockByNumber", params)
	if err != nil {
		return nil, fmt.Errorf("failed to get block: %w", err)
	}

	return resp, nil
}

// CallContract calls a contract method (read-only)
func (c *Client) CallContract(ctx context.Context, to string, value string, block string) (string, error) {
	callObject := map[string]interface{}{
		"to":    to,
		"value": value,
	}
	params := []interface{}{callObject, block}

	resp, err := c.Call(ctx, "eth_call", params)
	if err != nil {
		return "", fmt.Errorf("failed to call contract: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse contract call result: %w", err)
	}

	return result, nil
}

// EstimateGas estimates gas for a transaction
func (c *Client) EstimateGas(ctx context.Context, from, to, value string) (string, error) {
	txObject := map[string]interface{}{
		"from":  from,
		"to":    to,
		"value": value,
	}
	params := []interface{}{txObject}

	resp, err := c.Call(ctx, "eth_estimateGas", params)
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse gas estimate: %w", err)
	}

	return result, nil
}

// SendRawTransaction broadcasts a signed transaction
func (c *Client) SendRawTransaction(ctx context.Context, signedTx string) (string, error) {
	params := []interface{}{signedTx}

	resp, err := c.Call(ctx, "eth_sendRawTransaction", params)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse transaction hash: %w", err)
	}

	return result, nil
}

// GetGasPrice returns the current gas price
func (c *Client) GetGasPrice(ctx context.Context) (string, error) {
	resp, err := c.Call(ctx, "eth_gasPrice", []interface{}{})
	if err != nil {
		return "", fmt.Errorf("failed to get gas price: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse gas price: %w", err)
	}

	return result, nil
}

// GetChainID returns the chain ID
func (c *Client) GetChainID(ctx context.Context) (string, error) {
	resp, err := c.Call(ctx, "eth_chainId", []interface{}{})
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse chain ID: %w", err)
	}

	return result, nil
}

// GetLatestNonce retrieves the latest nonce for an address
func (c *Client) GetLatestNonce(
	ctx context.Context,
	privateKeyHex string,
) (uint64, error) {

	pkHex := strings.TrimPrefix(privateKeyHex, "0x")

	privateKey, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	resp, err := c.Call(
		ctx,
		"eth_getTransactionCount",
		[]interface{}{fromAddress, "latest"},
	)
	if err != nil {
		return 0, err
	}

	var hexNonce string
	if err := json.Unmarshal(resp.Result, &hexNonce); err != nil {
		return 0, err
	}

	nonce, err := strconv.ParseUint(
		strings.TrimPrefix(hexNonce, "0x"),
		16,
		64,
	)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

// GetEstimatedGasPrice estimates gas for a transaction
func (c *Client) GetEstimatedGasPrice(ctx context.Context, from string, to string, nonce uint64, value *big.Int, data []byte) (*big.Int, error) {

	pkHex := strings.TrimPrefix(from, "0x")
	privKey, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot cast public key to ECDSA")
	}

	fromAddr := crypto.PubkeyToAddress(*pubKeyECDSA).Hex()

	hexNonce := hexutil.EncodeUint64(nonce)
	valueHex := hexutil.EncodeBig(value)
	txObject := map[string]interface{}{
		"from":  fromAddr,
		"to":    to,
		"nonce": hexNonce,
		"value": valueHex,
		"data":  data,
	}
	params := []interface{}{txObject}

	resp, err := c.Call(ctx, "eth_estimateGas", params)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return nil, fmt.Errorf("failed to parse gas estimate: %w", err)
	}

	resultAsInt, ok := new(big.Int).SetString(result[2:], 16)
	if !ok {
		log.Fatal("fail to convert gas estimate to big.Int")
	}

	return resultAsInt, nil
}

// SignAndSendTransaction signs and sends a transaction
func (c *Client) SignAndSendTransaction(ctx context.Context, privateKey string, contractAddress string, data []byte) (string, error) {

	// Get chain ID
	chain_id, err := c.GetChainID(ctx)
	if err != nil {
		log.Fatal("fail to get chain id")
	}

	// Get latest nonce
	nonce, err := c.GetLatestNonce(
		ctx,
		privateKey,
	)
	if err != nil {
		log.Fatal("fail to get latest nonce")
	}

	// Calculate gas price
	gasPriceEstimated, err := c.GetEstimatedGasPrice(ctx, privateKey, contractAddress, nonce, big.NewInt(0), data)
	if err != nil {
		log.Fatal("fail to estimate gas")
	}

	// Sign transaction
	sign_params := sign.SignTxParams{
		ChainID:    chain_id,
		PrivateKey: privateKey,
		Nonce:      nonce,
		To:         contractAddress,
		Value:      big.NewInt(0),
		Data:       data,
		GasLimit:   21000,
		GasPrice:   gasPriceEstimated,
	}

	signedTx, err := sign.SignTx(sign_params)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer

	if err := signedTx.EncodeRLP(&buff); err != nil {
		log.Fatal(err)
	}

	rawBytes := buff.Bytes()
	rawHex := hexutil.Encode(rawBytes)

	params := []interface{}{rawHex}

	resp, err := c.Call(ctx, "eth_sendRawTransaction", params)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse transaction hash: %w", err)
	}

	return result, nil
}

// SignAndMint signs and sends a minting transaction
func (c *Client) SignAndMint(ctx context.Context, privateKey string, contractAddress string, data []byte) (string, error) {

	// Get chain ID
	chain_id, err := c.GetChainID(ctx)
	if err != nil {
		log.Fatal("fail to get chain id")
	}

	// Get latest nonce
	nonce, err := c.GetLatestNonce(
		ctx,
		privateKey,
	)
	if err != nil {
		log.Fatal("fail to get latest nonce")
	}

	// Derive sender address from private key for estimation
	pkHex := strings.TrimPrefix(privateKey, "0x")
	privKey, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot cast public key to ECDSA")
	}
	fromAddr := crypto.PubkeyToAddress(*pubKeyECDSA).Hex()

	// Calculate gas price
	gasPriceEstimated, err := c.GetEstimatedGasPrice(ctx, fromAddr, contractAddress, nonce, big.NewInt(2e18), data)
	if err != nil {
		log.Fatal("fail to estimate gas")
	}

	// Sign transaction
	sign_params := sign.SignTxParams{
		ChainID:    chain_id,
		PrivateKey: privateKey,
		Nonce:      nonce,
		To:         contractAddress,
		Value:      big.NewInt(0),
		Data:       data,
		GasLimit:   21000,
		GasPrice:   gasPriceEstimated,
	}

	signedTx, err := sign.SignTx(sign_params)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer

	if err := signedTx.EncodeRLP(&buff); err != nil {
		log.Fatal(err)
	}

	rawBytes := buff.Bytes()
	rawHex := hexutil.Encode(rawBytes)

	params := []interface{}{rawHex}

	resp, err := c.Call(ctx, "eth_sendRawTransaction", params)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse transaction hash: %w", err)
	}

	return result, nil
}

// SignAndBurn signs and sends a burning transaction
func (c *Client) SignAndBurn(ctx context.Context, privateKey string, contractAddress string, data []byte) (string, error) {

	// Get chain ID
	chain_id, err := c.GetChainID(ctx)
	if err != nil {
		log.Fatal("fail to get chain id")
	}

	// Get latest nonce
	nonce, err := c.GetLatestNonce(
		ctx,
		privateKey,
	)
	if err != nil {
		log.Fatal("fail to get latest nonce")
	}

	// Calculate gas price
	gasPriceEstimated, err := c.GetEstimatedGasPrice(ctx, privateKey, contractAddress, nonce, big.NewInt(2e18), data)
	if err != nil {
		log.Fatal("fail to estimate gas")
	}

	// Sign transaction
	sign_params := sign.SignTxParams{
		ChainID:    chain_id,
		PrivateKey: privateKey,
		Nonce:      nonce,
		To:         contractAddress,
		Value:      big.NewInt(0),
		Data:       data,
		GasLimit:   21000,
		GasPrice:   gasPriceEstimated,
	}

	signedTx, err := sign.SignTx(sign_params)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer

	if err := signedTx.EncodeRLP(&buff); err != nil {
		log.Fatal(err)
	}

	rawBytes := buff.Bytes()
	rawHex := hexutil.Encode(rawBytes)

	params := []interface{}{rawHex}

	resp, err := c.Call(ctx, "eth_sendRawTransaction", params)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	result, err := resp.GetResultAsString()
	if err != nil {
		return "", fmt.Errorf("failed to parse transaction hash: %w", err)
	}

	return result, nil
}
