package services

import (
	"context"
	"fmt"
	"math/big"

	"kokka.com/kokka/internal/applications/dtos"
	"kokka.com/kokka/internal/applications/validators"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain"
	"kokka.com/kokka/internal/shared/utils"
)

type StakeService struct {
	validator           validators.IStakeValidator
	client              *blockchain.Client
	decryptionKey       string
	readOnlyStakeClient *blockchain.StakeClient
}

// NewStakeService creates a new stake service
func NewStakeService(
	validator validators.IStakeValidator,
	client *blockchain.Client,
	decryptionKey string,
) (*StakeService, error) {
	// Create read-only stake client for stake queries (no signer needed)
	readOnlyClient, err := blockchain.NewStakeClient(client, nil)
	if err != nil {
		return nil, err
	}

	return &StakeService{
		validator:           validator,
		client:              client,
		decryptionKey:       decryptionKey,
		readOnlyStakeClient: readOnlyClient,
	}, nil
}

// GetUserStake retrieves the stake amount for a user
func (s *StakeService) GetUserStake(ctx context.Context, req *dtos.GetUserStakeRequest) (*dtos.GetUserStakeResponse, error) {

	// Validate request
	if err := s.validator.ValidateGetUserStakeRequest(req); err != nil {
		return nil, err
	}

	stakeAmount, err := s.readOnlyStakeClient.GetUserStake(ctx, req.ContractAddress, req.UserAddress)
	if err != nil {
		return nil, err
	}

	return &dtos.GetUserStakeResponse{
		ContractAddress: req.ContractAddress,
		StakeAmount:     stakeAmount,
	}, nil
}

// GetPendingRewards retrieves the pending rewards for a user
func (s *StakeService) GetPendingRewards(ctx context.Context, req *dtos.GetPendingRewardsRequest) (*dtos.GetPendingRewardsResponse, error) {

	// Validate request
	if err := s.validator.ValidateGetPendingRewardsRequest(req); err != nil {
		return nil, err
	}

	rewardsAmount, err := s.readOnlyStakeClient.PendingRewards(ctx, req.ContractAddress, req.UserAddress)
	if err != nil {
		return nil, err
	}

	return &dtos.GetPendingRewardsResponse{
		ContractAddress: req.ContractAddress,
		RewardsAmount:   rewardsAmount,
	}, nil
}

// StakeToken stakes tokens for the user
func (s *StakeService) StakeToken(ctx context.Context, req *dtos.StakeTokenRequest) (*dtos.StakeTokenResponse, error) {
	// Validate request
	if err := s.validator.ValidateStakeTokenRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amount, err := parseAmount(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}

	// Decrypt private key
	privateKey, err := utils.DecryptCrypto(req.EncryptedPrivateKey, s.decryptionKey)
	if err != nil {
		return nil, err
	}

	// Create transaction signer
	signer, err := blockchain.NewTransactionSigner(privateKey, s.client)
	if err != nil {
		return nil, err
	}

	// Get the current nonce for manual nonce management
	address := signer.GetAddress()
	nonceHex, err := s.client.GetTransactionCount(ctx, address, "pending")
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Calculate nonce N+1 for second transaction
	nonceStr := nonceHex
	if len(nonceStr) > 2 && nonceStr[:2] == "0x" {
		nonceStr = nonceStr[2:]
	}
	nonceBig := new(big.Int)
	nonceBig.SetString(nonceStr, 16)
	nextNonce := new(big.Int).Add(nonceBig, big.NewInt(1))
	nextNonceHex := "0x" + nextNonce.Text(16)

	// Approve the staking contract to spend tokens (uses nonce N)
	tokenClient, err := blockchain.NewTokenClient(s.client, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client: %w", err)
	}

	approveTxHash, err := tokenClient.Approve(ctx, req.TokenAddress, req.ContractAddress, amount, nonceHex)
	if err != nil {
		return nil, fmt.Errorf("failed to approve tokens: %w", err)
	}

	// Wait for approve transaction to be mined before proceeding
	err = s.client.WaitForTransaction(ctx, approveTxHash)
	if err != nil {
		return nil, fmt.Errorf("approve transaction not confirmed (tx: %s): %w", approveTxHash, err)
	}

	// Execute the stake transaction (uses nonce N+1)
	stakeClient, err := blockchain.NewStakeClient(s.client, signer)
	if err != nil {
		return nil, err
	}

	txHash, err := stakeClient.Stake(ctx, req.ContractAddress, req.TokenAddress, amount, nextNonceHex)
	if err != nil {
		return nil, fmt.Errorf("failed to stake tokens (approve tx: %s): %w", approveTxHash, err)
	}

	return &dtos.StakeTokenResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
	}, nil
}

// WithdrawToken withdraws staked tokens for the user
func (s *StakeService) WithdrawToken(ctx context.Context, req *dtos.WithdrawTokenRequest) (*dtos.WithdrawTokenResponse, error) {
	// Validate request
	if err := s.validator.ValidateWithdrawTokenRequest(req); err != nil {
		return nil, err
	}

	// Parse amount
	amount, err := parseAmount(req.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to parse amount: %w", err)
	}

	// Decrypt private key
	privateKey, err := utils.DecryptCrypto(req.EncryptedPrivateKey, s.decryptionKey)
	if err != nil {
		return nil, err
	}

	// Create transaction signer
	signer, err := blockchain.NewTransactionSigner(privateKey, s.client)
	if err != nil {
		return nil, err
	}

	// Create stake client
	stakeClient, err := blockchain.NewStakeClient(s.client, signer)
	if err != nil {
		return nil, err
	}

	txHash, err := stakeClient.Withdraw(ctx, req.ContractAddress, amount)
	if err != nil {
		return nil, err
	}

	return &dtos.WithdrawTokenResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
	}, nil
}

// ClaimRewards claims pending rewards for the user
func (s *StakeService) ClaimRewards(ctx context.Context, req *dtos.ClaimRewardsRequest) (*dtos.ClaimRewardsResponse, error) {
	// Validate request
	if err := s.validator.ValidateClaimRewardsRequest(req); err != nil {
		return nil, err
	}

	// Decrypt private key
	privateKey, err := utils.DecryptCrypto(req.EncryptedPrivateKey, s.decryptionKey)
	if err != nil {
		return nil, err
	}

	// Create transaction signer
	signer, err := blockchain.NewTransactionSigner(privateKey, s.client)
	if err != nil {
		return nil, err
	}

	// Create stake client
	stakeClient, err := blockchain.NewStakeClient(s.client, signer)
	if err != nil {
		return nil, err
	}

	txHash, err := stakeClient.ClaimRewards(ctx, req.ContractAddress)
	if err != nil {
		return nil, err
	}

	return &dtos.ClaimRewardsResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
	}, nil
}

// ClaimAllRewards claims all pending rewards for the user
func (s *StakeService) ClaimAllRewards(ctx context.Context, req *dtos.ClaimAllRewardsRequest) (*dtos.ClaimAllRewardsResponse, error) {
	// Validate request
	if err := s.validator.ValidateClaimAllRewardsRequest(req); err != nil {
		return nil, err
	}

	// Decrypt private key
	privateKey, err := utils.DecryptCrypto(req.EncryptedPrivateKey, s.decryptionKey)
	if err != nil {
		return nil, err
	}

	// Create transaction signer
	signer, err := blockchain.NewTransactionSigner(privateKey, s.client)
	if err != nil {
		return nil, err
	}

	// Create stake client
	stakeClient, err := blockchain.NewStakeClient(s.client, signer)
	if err != nil {
		return nil, err
	}

	txHash, err := stakeClient.ClaimAllRewards(ctx, req.ContractAddress)
	if err != nil {
		return nil, err
	}

	return &dtos.ClaimAllRewardsResponse{
		TxHash:          txHash,
		ContractAddress: req.ContractAddress,
	}, nil
}
