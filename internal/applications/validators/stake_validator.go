package validators

import (
	"errors"

	"kokka.com/kokka/internal/applications/dtos"
)

type IStakeValidator interface {
	ValidateStakeTokenRequest(req *dtos.StakeTokenRequest) error
	ValidateGetUserStakeRequest(req *dtos.GetUserStakeRequest) error
	ValidateGetPendingRewardsRequest(req *dtos.GetPendingRewardsRequest) error
	ValidateGetTotalStakedRequest(req *dtos.GetTotalStakedRequest) error
	ValidateGetApyRatesRequest(req *dtos.GetApyRatesRequest) error
	ValidateWithdrawTokenRequest(req *dtos.WithdrawTokenRequest) error
	ValidateClaimRewardsRequest(req *dtos.ClaimRewardsRequest) error
	ValidateClaimAllRewardsRequest(req *dtos.ClaimAllRewardsRequest) error
}

type stakeValidator struct{}

func NewStakeValidator() *stakeValidator {
	return &stakeValidator{}
}

func (v *stakeValidator) ValidateStakeTokenRequest(req *dtos.StakeTokenRequest) error {
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

	if req.EncryptedPrivateKey == "" {
		return errors.New("encrypted_private_key is required")
	}
	return nil
}

func (v *stakeValidator) ValidateGetUserStakeRequest(req *dtos.GetUserStakeRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.UserAddress == "" {
		return errors.New("user_address is required")
	}

	if !isValidEthereumAddress(req.UserAddress) {
		return errors.New("invalid user_address format")
	}
	return nil
}

func (v *stakeValidator) ValidateGetPendingRewardsRequest(req *dtos.GetPendingRewardsRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.UserAddress == "" {
		return errors.New("user_address is required")
	}

	if !isValidEthereumAddress(req.UserAddress) {
		return errors.New("invalid user_address format")
	}

	return nil
}

func (v *stakeValidator) ValidateGetTotalStakedRequest(req *dtos.GetTotalStakedRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if req.TokenAddress == "" {
		return errors.New("token_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if !isValidEthereumAddress(req.TokenAddress) {
		return errors.New("invalid token_address format")
	}

	return nil
}

func (v *stakeValidator) ValidateGetApyRatesRequest(req *dtos.GetApyRatesRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if req.TokenAddress == "" {
		return errors.New("token_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if !isValidEthereumAddress(req.TokenAddress) {
		return errors.New("invalid token_address format")
	}

	return nil
}

func (v *stakeValidator) ValidateWithdrawTokenRequest(req *dtos.WithdrawTokenRequest) error {
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

	if req.EncryptedPrivateKey == "" {
		return errors.New("encrypted_private_key is required")
	}

	return nil
}

func (v *stakeValidator) ValidateClaimRewardsRequest(req *dtos.ClaimRewardsRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.EncryptedPrivateKey == "" {
		return errors.New("encrypted_private_key is required")
	}

	return nil
}

func (v *stakeValidator) ValidateClaimAllRewardsRequest(req *dtos.ClaimAllRewardsRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}

	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	if req.EncryptedPrivateKey == "" {
		return errors.New("encrypted_private_key is required")
	}

	return nil
}
