package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// TransactionSigner handles transaction signing operations
type TransactionSigner struct {
	privateKey *ecdsa.PrivateKey
	client     *Client
}

// NewTransactionSigner creates a new transaction signer
func NewTransactionSigner(privateKeyHex string, client *Client) (*TransactionSigner, error) {
	// Remove 0x prefix if present
	if len(privateKeyHex) > 2 && privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	return &TransactionSigner{
		privateKey: privateKey,
		client:     client,
	}, nil
}

// GetAddress returns the Ethereum address associated with the private key
func (s *TransactionSigner) GetAddress() string {
	publicKey := s.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return ""
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

// SignAndSendTransaction signs a transaction and sends it to the blockchain
func (s *TransactionSigner) SignAndSendTransaction(ctx context.Context, req *SignTransactionRequest) (string, error) {
	// Get chain ID
	chainIDHex, err := s.client.GetChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %w", err)
	}
	chainID := new(big.Int)
	chainID.SetString(chainIDHex[2:], 16) // Remove 0x and parse as hex

	// Parse addresses and values
	toAddress := common.HexToAddress(req.To)

	// Parse value (amount to send in wei)
	value := new(big.Int)
	if req.Value != "" {
		// Remove 0x prefix if present
		valueStr := req.Value
		if len(valueStr) > 2 && valueStr[:2] == "0x" {
			valueStr = valueStr[2:]
		}
		value.SetString(valueStr, 16)
	}

	// Parse nonce
	var nonce uint64
	if req.Nonce != "" {
		nonceStr := req.Nonce
		if len(nonceStr) > 2 && nonceStr[:2] == "0x" {
			nonceStr = nonceStr[2:]
		}
		nonceBig := new(big.Int)
		nonceBig.SetString(nonceStr, 16)
		nonce = nonceBig.Uint64()
	} else {
		// Get nonce from blockchain
		nonceHex, err := s.client.GetTransactionCount(ctx, s.GetAddress(), "pending")
		if err != nil {
			return "", fmt.Errorf("failed to get nonce: %w", err)
		}
		nonceStr := nonceHex
		if len(nonceStr) > 2 && nonceStr[:2] == "0x" {
			nonceStr = nonceStr[2:]
		}
		nonceBig := new(big.Int)
		nonceBig.SetString(nonceStr, 16)
		nonce = nonceBig.Uint64()
	}

	// Parse gas limit
	var gasLimit uint64 = 21000 // Default gas limit for simple transfers
	if req.GasLimit != "" {
		gasLimitStr := req.GasLimit
		if len(gasLimitStr) > 2 && gasLimitStr[:2] == "0x" {
			gasLimitStr = gasLimitStr[2:]
		}
		gasLimitBig := new(big.Int)
		gasLimitBig.SetString(gasLimitStr, 16)
		gasLimit = gasLimitBig.Uint64()
	} else if req.Data != "" {
		// If data is present, estimate gas
		estimatedGasHex, err := s.client.EstimateGas(ctx, s.GetAddress(), req.To, req.Data)
		if err != nil {
			return "", fmt.Errorf("failed to estimate gas: %w", err)
		}
		gasLimitStr := estimatedGasHex
		if len(gasLimitStr) > 2 && gasLimitStr[:2] == "0x" {
			gasLimitStr = gasLimitStr[2:]
		}
		gasLimitBig := new(big.Int)
		gasLimitBig.SetString(gasLimitStr, 16)
		gasLimit = gasLimitBig.Uint64()
	}

	// Parse gas price
	var gasPrice *big.Int
	if req.GasPrice != "" {
		gasPrice = new(big.Int)
		gasPriceStr := req.GasPrice
		if len(gasPriceStr) > 2 && gasPriceStr[:2] == "0x" {
			gasPriceStr = gasPriceStr[2:]
		}
		gasPrice.SetString(gasPriceStr, 16)
	} else {
		// Get gas price from blockchain
		gasPriceHex, err := s.client.GetGasPrice(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to get gas price: %w", err)
		}
		gasPrice = new(big.Int)
		gasPriceStr := gasPriceHex
		if len(gasPriceStr) > 2 && gasPriceStr[:2] == "0x" {
			gasPriceStr = gasPriceStr[2:]
		}
		gasPrice.SetString(gasPriceStr, 16)
	}

	// Parse data
	var data []byte
	if req.Data != "" {
		data = common.FromHex(req.Data)
	}

	// Create transaction
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), s.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Encode signed transaction to raw hex
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("failed to encode transaction: %w", err)
	}
	rawTxHex := "0x" + common.Bytes2Hex(rawTxBytes)

	// Send transaction
	txHash, err := s.client.SendRawTransaction(ctx, rawTxHex)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	return txHash, nil
}

// SignTransactionRequest represents the parameters needed to sign a transaction
type SignTransactionRequest struct {
	To       string `json:"to"`        // Recipient address
	Value    string `json:"value"`     // Amount in wei (hex string)
	Data     string `json:"data"`      // Optional: contract data (hex string)
	GasLimit string `json:"gas_limit"` // Optional: gas limit (hex string)
	GasPrice string `json:"gas_price"` // Optional: gas price (hex string)
	Nonce    string `json:"nonce"`     // Optional: transaction nonce (hex string)
}
