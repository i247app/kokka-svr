package validators

import (
	"errors"

	"kokka.com/kokka/internal/applications/dtos"
)

type ISwapValidator interface {
	ValidateSwapTokenRequest(req *dtos.SwapTokenRequest) error
	ValidateGetSwapQuoteRequest(req *dtos.GetSwapQuoteRequest) error
	ValidateGetSwapInfoRequest(req *dtos.GetSwapInfoRequest) error
}

type swapValidator struct{}

func NewSwapValidator() *swapValidator {
	return &swapValidator{}
}

// ValidateSwapTokenRequest validates a swap token request
func (v *swapValidator) ValidateSwapTokenRequest(req *dtos.SwapTokenRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.AmountIn == "" {
		return errors.New("amount_in is required")
	}

	if !isValidAmount(req.AmountIn) {
		return errors.New("invalid amount_in format")
	}

	if req.Direction == "" {
		return errors.New("direction is required")
	}

	if req.Direction != "AtoB" && req.Direction != "BtoA" {
		return errors.New("direction must be either 'AtoB' or 'BtoA'")
	}

	if req.EncryptedPrivateKey == "" {
		return errors.New("encrypted_private_key is required")
	}

	return nil
}

// ValidateGetSwapQuoteRequest validates a get swap quote request
func (v *swapValidator) ValidateGetSwapQuoteRequest(req *dtos.GetSwapQuoteRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.AmountIn == "" {
		return errors.New("amount_in is required")
	}

	if !isValidAmount(req.AmountIn) {
		return errors.New("invalid amount_in format")
	}

	if req.Direction == "" {
		return errors.New("direction is required")
	}

	if req.Direction != "AtoB" && req.Direction != "BtoA" {
		return errors.New("direction must be either 'AtoB' or 'BtoA'")
	}

	return nil
}

// ValidateGetSwapInfoRequest validates a get swap info request
func (v *swapValidator) ValidateGetSwapInfoRequest(req *dtos.GetSwapInfoRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	return nil
}
