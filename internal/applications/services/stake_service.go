package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

// GetTotalStaked retrieves the total staked amount for a contract and token
func (s *StakeService) GetTotalStaked(ctx context.Context, req *dtos.GetTotalStakedRequest) (*dtos.GetTotalStakedResponse, error) {

	// Validate request
	if err := s.validator.ValidateGetTotalStakedRequest(req); err != nil {
		return nil, err
	}

	totalStaked, err := s.readOnlyStakeClient.TotalStaked(ctx, req.ContractAddress, req.TokenAddress)
	if err != nil {
		return nil, err
	}

	return &dtos.GetTotalStakedResponse{
		ContractAddress: req.ContractAddress,
		TotalStaked:     totalStaked,
	}, nil
}

// GetApyRates retrieves the APY rates for a contract and token
func (s *StakeService) GetApyRates(ctx context.Context, req *dtos.GetApyRatesRequest) (*dtos.GetApyRatesResponse, error) {

	// Validate request
	if err := s.validator.ValidateGetApyRatesRequest(req); err != nil {
		return nil, err
	}

	apyRates, err := s.readOnlyStakeClient.ApyRates(ctx, req.ContractAddress, req.TokenAddress)
	if err != nil {
		return nil, err
	}

	// Convert APY string to float64
	apyBigInt := new(big.Int)
	apyBigInt.SetString(apyRates, 10)
	apyFloat := new(big.Float).SetInt(apyBigInt)
	apyFloat64, _ := apyFloat.Float64()

	return &dtos.GetApyRatesResponse{
		ContractAddress: req.ContractAddress,
		ApyRates:        apyFloat64,
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

	// Create clients
	tokenClient, err := blockchain.NewTokenClient(s.client, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client: %w", err)
	}

	stakeClient, err := blockchain.NewStakeClient(s.client, signer)
	if err != nil {
		return nil, err
	}

	// Get addresses for gas estimation
	ownerAddr := signer.GetAddressAsCommon()
	spenderAddr := common.HexToAddress(req.ContractAddress)

	// Step 1: Estimate approve gas (doesn't depend on anything)
	approveGasLimit, err := tokenClient.EstimateApproveGas(ctx, req.TokenAddress, req.ContractAddress, amount, nonceHex, ownerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate approve gas: %w", err)
	}

	// Step 2: Estimate stake gas with state override (tries multiple storage slots)
	stakeGasLimit, err := stakeClient.EstimateStakeGasWithAllowanceOverride(ctx, req.ContractAddress, req.TokenAddress, amount, nextNonceHex, ownerAddr, spenderAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate stake gas: %w", err)
	}

	// Step 3: Send approve transaction (nonce N) with estimated gas
	approveTxHash, err := tokenClient.ApproveWithGasLimit(ctx, req.TokenAddress, req.ContractAddress, amount, nonceHex, approveGasLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to send approve transaction: %w", err)
	}

	// Step 4: Send stake transaction (nonce N+1) immediately with estimated gas
	stakeTxHash, err := stakeClient.StakeWithGasLimit(ctx, req.ContractAddress, req.TokenAddress, amount, nextNonceHex, stakeGasLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to send stake transaction (approve tx: %s): %w", approveTxHash, err)
	}

	// Wait for stake transaction receipt (approve will confirm first due to nonce order)
	err = s.client.WaitForTransaction(ctx, stakeTxHash)
	if err != nil {
		return nil, fmt.Errorf("stake transaction not confirmed (approve tx: %s, stake tx: %s): %w", approveTxHash, stakeTxHash, err)
	}

	return &dtos.StakeTokenResponse{
		TxHash:          stakeTxHash,
		ApproveTxHash:   approveTxHash,
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

	txHash, err := stakeClient.Withdraw(ctx, req.ContractAddress, req.TokenAddress, amount)
	if err != nil {
		return nil, err
	}

	return &dtos.WithdrawTokenResponse{
		Amount:          req.Amount,
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

	txHash, err := stakeClient.ClaimRewards(ctx, req.ContractAddress, req.TokenAddress)
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
