package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"kokka.com/kokka/internal/applications/dtos"
	"kokka.com/kokka/internal/applications/validators"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain"
)

// TokenService handles token business logic
type TokenService struct {
	validator validators.ITokenValidator
	client    *blockchain.TokenClient
}

// NewTokenService creates a new token service
func NewTokenService(
	validator validators.ITokenValidator,
	client *blockchain.TokenClient,
) *TokenService {
	return &TokenService{
		validator: validator,
		client:    client,
	}
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

	// Execute mint transaction
	txHash, err := s.client.Mint(ctx, req.ContractAddress, req.To, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to mint tokens: %w", err)
	}

	// Query new balance (best effort - don't fail if balance query fails)
	newBalance, err := s.client.BalanceOf(ctx, req.ContractAddress, req.To)
	var newBalanceStr string
	if err == nil && newBalance != nil {
		newBalanceStr = newBalance.String()
	}

	return &dtos.MintTokenResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
		To:              req.To,
		Amount:          amount.String(),
		NewBalance:      newBalanceStr,
	}, nil
}

// Burn burns tokens from the server's account
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

	// Execute burn transaction
	txHash, err := s.client.Burn(ctx, req.ContractAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to burn tokens: %w", err)
	}

	// Query new balance (best effort - don't fail if balance query fails)
	// Note: We would need the server's address to query balance after burn
	// For now, we'll leave it empty
	var newBalanceStr string

	return &dtos.BurnTokenResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
		Amount:          amount.String(),
		NewBalance:      newBalanceStr,
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

	// Execute transfer transaction
	txHash, err := s.client.Transfer(ctx, req.ContractAddress, req.To, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer tokens: %w", err)
	}

	return &dtos.TransferTokenResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
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

	// Query balance
	balance, err := s.client.BalanceOf(ctx, req.ContractAddress, req.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return &dtos.GetTokenBalanceResponse{
		ContractAddress: req.ContractAddress,
		Address:         req.Address,
		Balance:         balance.String(),
	}, nil
}

// ========================================
// Helper functions
// ========================================

// parseAmount parses an amount string (decimal or hex) to *big.Int
func parseAmount(amount string) (*big.Int, error) {
	result := new(big.Int)

	// Handle hex format (0x prefix)
	if strings.HasPrefix(amount, "0x") {
		// Remove 0x prefix and parse as hex
		_, ok := result.SetString(amount[2:], 16)
		if !ok {
			return nil, fmt.Errorf("invalid hex amount: %s", amount)
		}
		return result, nil
	}

	// Handle decimal format
	_, ok := result.SetString(amount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid decimal amount: %s", amount)
	}

	return result, nil
}
