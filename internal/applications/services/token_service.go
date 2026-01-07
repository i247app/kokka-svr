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

// TokenService handles token business logic
type TokenService struct {
	validator           validators.ITokenValidator
	client              *blockchain.Client
	decryptionKey       string
	readOnlyTokenClient *blockchain.TokenClient
}

// NewTokenService creates a new token service
func NewTokenService(
	validator validators.ITokenValidator,
	client *blockchain.Client,
	decryptionKey string,
) (*TokenService, error) {
	// Create read-only token client for balance queries (no signer needed)
	readOnlyClient, err := blockchain.NewTokenClient(client, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create read-only token client: %w", err)
	}

	return &TokenService{
		validator:           validator,
		client:              client,
		decryptionKey:       decryptionKey,
		readOnlyTokenClient: readOnlyClient,
	}, nil
}

// Mint mints new tokens to a specified address
func (s *TokenService) Mint(ctx context.Context, req *dtos.MintTokenRequest) (*dtos.MintTokenResponse, error) {
	// Validate request
	if err := s.validator.ValidateMintTokenRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amount, err := parseAmount(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
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

	// Create token client
	tokenClient, err := blockchain.NewTokenClient(s.client, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client: %w", err)
	}

	// Execute mint transaction
	txHash, err := tokenClient.Mint(ctx, req.ContractAddress, req.To, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to mint tokens: %w", err)
	}

	// // Query new balance (best effort - don't fail if balance query fails)
	// newBalance, err := tokenClient.BalanceOf(ctx, req.ContractAddress, req.To)
	// var newBalanceStr string
	// if err == nil && newBalance != nil {
	// 	newBalanceStr = newBalance.String()
	// }

	return &dtos.MintTokenResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
		To:              req.To,
		Amount:          amount.String(),
		// NewBalance:      newBalanceStr,
	}, nil
}

// Burn burns tokens from the caller's account
func (s *TokenService) Burn(ctx context.Context, req *dtos.BurnTokenRequest) (*dtos.BurnTokenResponse, error) {
	// Validate request
	if err := s.validator.ValidateBurnTokenRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amount, err := parseAmount(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
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

	// Create token client
	tokenClient, err := blockchain.NewTokenClient(s.client, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client: %w", err)
	}

	// Execute burn transaction
	txHash, err := tokenClient.Burn(ctx, req.ContractAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to burn tokens: %w", err)
	}

	// // Query new balance (best effort - don't fail if balance query fails)
	// newBalance, err := tokenClient.BalanceOf(ctx, req.ContractAddress, signer.GetAddress())
	// var newBalanceStr string
	// if err == nil && newBalance != nil {
	// 	newBalanceStr = newBalance.String()
	// }

	return &dtos.BurnTokenResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
		Amount:          amount.String(),
		// NewBalance:      newBalanceStr,
	}, nil
}

// Transfer transfers tokens to a specified address
func (s *TokenService) Transfer(ctx context.Context, req *dtos.TransferTokenRequest) (*dtos.TransferTokenResponse, error) {
	// Validate request
	if err := s.validator.ValidateTransferTokenRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amount, err := parseAmount(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
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

	// Create token client
	tokenClient, err := blockchain.NewTokenClient(s.client, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client: %w", err)
	}

	// Execute transfer transaction
	txHash, err := tokenClient.Transfer(ctx, req.ContractAddress, req.To, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer tokens: %w", err)
	}

	return &dtos.TransferTokenResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
		From:            signer.GetAddress(),
		To:              req.To,
		Amount:          amount.String(),
	}, nil
}

// GetBalance returns the token balance of an address
func (s *TokenService) GetBalance(ctx context.Context, req *dtos.GetTokenBalanceRequest) (*dtos.GetTokenBalanceResponse, error) {
	// Validate request
	if err := s.validator.ValidateGetTokenBalanceRequest(req); err != nil {
		return nil, err
	}

	// Query balance using read-only client (no signing required)
	balance, err := s.readOnlyTokenClient.BalanceOf(ctx, req.ContractAddress, req.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return &dtos.GetTokenBalanceResponse{
		ContractAddress: req.ContractAddress,
		Address:         req.Address,
		Balance:         balance.String(),
	}, nil
}

// parseAmount parses an amount string and converts token units to wei (multiplies by 10^18)
// Input: "2" or "2.5" (token units)
// Output: big.Int representing wei (e.g., "2" -> 2000000000000000000)
func parseAmount(amount string) (*big.Int, error) {
	// Parse the amount as a floating point number
	amountFloat := new(big.Float)
	_, ok := amountFloat.SetString(amount)
	if !ok {
		return nil, fmt.Errorf("invalid amount format: %s", amount)
	}

	// Check if amount is positive
	if amountFloat.Cmp(big.NewFloat(0)) <= 0 {
		return nil, fmt.Errorf("amount must be positive: %s", amount)
	}

	// Multiply by 10^18 to convert to wei (18 decimals for ERC20 tokens)
	decimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	decimalFloat := new(big.Float).SetInt(decimals)
	weiFloat := new(big.Float).Mul(amountFloat, decimalFloat)

	// Convert to big.Int (truncate any remaining decimals)
	weiInt := new(big.Int)
	weiFloat.Int(weiInt)

	return weiInt, nil
}

// GetAddressInfo retrieves basic information about the token contract at the given address
func (s *TokenService) GetAddressInfo(ctx context.Context, req *dtos.GetAddressInfoRequest) (*dtos.GetAddressInfoResponse, error) {
	// Validate request
	if err := s.validator.ValidateGetAddressInfoRequest(req); err != nil {
		return nil, err
	}

	// Query address info using read-only client (no signing required)
	info, err := s.readOnlyTokenClient.AddressInfo(ctx, req.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get address info: %w", err)
	}

	return &dtos.GetAddressInfoResponse{
		Name:         info.Name,
		Symbol:       info.Symbol,
		Decimals:     info.Decimals,
		TotalSupply:  info.TotalSupply,
		OwnerAddress: info.OwnerAddress,
		OwnerBalance: info.OwnerBalance,
	}, nil
}
