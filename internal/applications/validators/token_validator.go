package validators

import (
	"errors"
	"math/big"

	"kokka.com/kokka/internal/applications/dtos"
)

type ITokenValidator interface {
	ValidateMintTokenRequest(req *dtos.MintTokenRequest) error
	ValidateBurnTokenRequest(req *dtos.BurnTokenRequest) error
	ValidateTransferTokenRequest(req *dtos.TransferTokenRequest) error
	ValidateGetTokenBalanceRequest(req *dtos.GetTokenBalanceRequest) error
	ValidateGetAddressInfoRequest(req *dtos.GetAddressInfoRequest) error
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

	if req.EncryptedPrivateKey == "" {
		return errors.New("encrypted_private_key is required")
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

	if req.EncryptedPrivateKey == "" {
		return errors.New("encrypted_private_key is required")
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

	if req.EncryptedPrivateKey == "" {
		return errors.New("encrypted_private_key is required")
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

// ValidateGetAddressInfoRequest validates a get address info request
func (v *tokenValidator) ValidateGetAddressInfoRequest(req *dtos.GetAddressInfoRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}
	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	return nil
}

// isValidAmount checks if a string is a valid token amount (decimal number like "2" or "2.5")
func isValidAmount(amount string) bool {
	// Parse as floating point number
	amountFloat := new(big.Float)
	_, ok := amountFloat.SetString(amount)
	if !ok {
		return false
	}

	// Ensure amount is positive
	return amountFloat.Cmp(big.NewFloat(0)) > 0
}
