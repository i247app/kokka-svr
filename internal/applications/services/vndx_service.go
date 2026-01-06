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

// VNDXService handles VNDX token business logic
type VNDXService struct {
	validator validators.IVNDXValidator
	client    *blockchain.VNDXClient
}

// NewVNDXService creates a new VNDX service
func NewVNDXService(
	validator validators.IVNDXValidator,
	client *blockchain.VNDXClient,
) *VNDXService {
	return &VNDXService{
		validator: validator,
		client:    client,
	}
}

// Mint mints new VNDX tokens to a specified address
func (s *VNDXService) Mint(ctx context.Context, req *dtos.MintVNDXRequest) (*dtos.MintVNDXResponse, error) {
	// Validate request
	if err := s.validator.ValidateMintVNDXRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amount, err := parseAmount(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}

	// Execute mint transaction
	txHash, err := s.client.Mint(ctx, req.To, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to mint tokens: %w", err)
	}

	// Query new balance (best effort - don't fail if balance query fails)
	newBalance, err := s.client.BalanceOf(ctx, req.To)
	var newBalanceStr string
	if err == nil && newBalance != nil {
		newBalanceStr = newBalance.String()
	}

	return &dtos.MintVNDXResponse{
		TxHash:     txHash,
		To:         req.To,
		Amount:     amount.String(),
		NewBalance: newBalanceStr,
	}, nil
}

// Burn burns VNDX tokens from the server's account
func (s *VNDXService) Burn(ctx context.Context, req *dtos.BurnVNDXRequest) (*dtos.BurnVNDXResponse, error) {
	// Validate request
	if err := s.validator.ValidateBurnVNDXRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amount, err := parseAmount(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}

	// Execute burn transaction
	txHash, err := s.client.Burn(ctx, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to burn tokens: %w", err)
	}

	// Query new balance (best effort - don't fail if balance query fails)
	// Note: We would need the server's address to query balance after burn
	// For now, we'll leave it empty
	var newBalanceStr string

	return &dtos.BurnVNDXResponse{
		TxHash:     txHash,
		Amount:     amount.String(),
		NewBalance: newBalanceStr,
	}, nil
}

// Transfer transfers VNDX tokens to a specified address
func (s *VNDXService) Transfer(ctx context.Context, req *dtos.TransferVNDXRequest) (*dtos.TransferVNDXResponse, error) {
	// Validate request
	if err := s.validator.ValidateTransferVNDXRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amount, err := parseAmount(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}

	// Execute transfer transaction
	txHash, err := s.client.Transfer(ctx, req.To, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer tokens: %w", err)
	}

	return &dtos.TransferVNDXResponse{
		TxHash: txHash,
		To:     req.To,
		Amount: amount.String(),
	}, nil
}

// GetBalance returns the VNDX token balance of an address
func (s *VNDXService) GetBalance(ctx context.Context, req *dtos.GetVNDXBalanceRequest) (*dtos.GetVNDXBalanceResponse, error) {
	// Validate request
	if err := s.validator.ValidateGetVNDXBalanceRequest(req); err != nil {
		return nil, err
	}

	// Query balance
	balance, err := s.client.BalanceOf(ctx, req.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return &dtos.GetVNDXBalanceResponse{
		Address: req.Address,
		Balance: balance.String(),
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
