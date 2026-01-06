package validators

import (
	"errors"
	"math/big"
	"strings"

	"kokka.com/kokka/internal/applications/dtos"
)

type IVNDXValidator interface {
	ValidateMintVNDXRequest(req *dtos.MintVNDXRequest) error
	ValidateBurnVNDXRequest(req *dtos.BurnVNDXRequest) error
	ValidateTransferVNDXRequest(req *dtos.TransferVNDXRequest) error
	ValidateGetVNDXBalanceRequest(req *dtos.GetVNDXBalanceRequest) error
}

type vndxValidator struct{}

func NewVNDXValidator() *vndxValidator {
	return &vndxValidator{}
}

// ValidateMintVNDXRequest validates a mint VNDX request
func (v *vndxValidator) ValidateMintVNDXRequest(req *dtos.MintVNDXRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
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

// ValidateBurnVNDXRequest validates a burn VNDX request
func (v *vndxValidator) ValidateBurnVNDXRequest(req *dtos.BurnVNDXRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.Amount == "" {
		return errors.New("amount is required")
	}

	if !isValidAmount(req.Amount) {
		return errors.New("invalid amount format")
	}

	return nil
}

// ValidateTransferVNDXRequest validates a transfer VNDX request
func (v *vndxValidator) ValidateTransferVNDXRequest(req *dtos.TransferVNDXRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
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

// ValidateGetVNDXBalanceRequest validates a get VNDX balance request
func (v *vndxValidator) ValidateGetVNDXBalanceRequest(req *dtos.GetVNDXBalanceRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
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
