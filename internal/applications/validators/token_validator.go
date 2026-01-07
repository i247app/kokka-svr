package validators

import (
	"errors"
	"math/big"
	"strings"

	"kokka.com/kokka/internal/applications/dtos"
)

type ITokenValidator interface {
	ValidateMintTokenRequest(req *dtos.MintTokenRequest) error
	ValidateBurnTokenRequest(req *dtos.BurnTokenRequest) error
	ValidateTransferTokenRequest(req *dtos.TransferTokenRequest) error
	ValidateGetTokenBalanceRequest(req *dtos.GetTokenBalanceRequest) error
}

type tokenValidator struct{}

func NewTokenValidator() *tokenValidator {
	return &tokenValidator{}
}

// ValidateMintTokenRequest validates a mint token request
func (v *tokenValidator) ValidateMintTokenRequest(req *dtos.MintTokenRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.To == "" {
		return errors.New("to address is required")
	}

	if !isValidEthereumAddress(req.To) {
		return errors.New("invalid to address format")
	}

	if req.Amount == "" {
		return errors.New("amount is required")
	}

	if !isValidAmount(req.Amount) {
		return errors.New("invalid amount format")
	}

	return nil
}

// ValidateBurnTokenRequest validates a burn token request
func (v *tokenValidator) ValidateBurnTokenRequest(req *dtos.BurnTokenRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.Amount == "" {
		return errors.New("amount is required")
	}

	if !isValidAmount(req.Amount) {
		return errors.New("invalid amount format")
	}

	return nil
}

// ValidateTransferTokenRequest validates a transfer token request
func (v *tokenValidator) ValidateTransferTokenRequest(req *dtos.TransferTokenRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.To == "" {
		return errors.New("to address is required")
	}

	if !isValidEthereumAddress(req.To) {
		return errors.New("invalid to address format")
	}

	if req.Amount == "" {
		return errors.New("amount is required")
	}

	if !isValidAmount(req.Amount) {
		return errors.New("invalid amount format")
	}

	return nil
}

// ValidateGetTokenBalanceRequest validates a get token balance request
func (v *tokenValidator) ValidateGetTokenBalanceRequest(req *dtos.GetTokenBalanceRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.Address == "" {
		return errors.New("address is required")
	}

	if !isValidEthereumAddress(req.Address) {
		return errors.New("invalid address format")
	}

	return nil
}

// ========================================
// Helper validation functions
// ========================================

// isValidAmount checks if a string is a valid token amount (decimal string or hex)
func isValidAmount(amount string) bool {
	// Allow hex format (0x prefix)
	if strings.HasPrefix(amount, "0x") {
		if len(amount) == 2 {
			return false // Just "0x" is not valid
		}
		return isValidHex(amount[2:])
	}

	// Allow decimal format (positive integer)
	_, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return false
	}

	// Ensure amount is positive
	amountBig := new(big.Int)
	amountBig.SetString(amount, 10)
	return amountBig.Cmp(big.NewInt(0)) > 0
}
