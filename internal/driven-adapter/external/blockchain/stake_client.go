package blockchain

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain/gen/stake"
)

type StakeClient struct {
	client *Client
	signer *TransactionSigner
	abi    abi.ABI
}

// NewStakeClient creates a new stake client
// signer can be nil for read-only operations
func NewStakeClient(client *Client, signer *TransactionSigner) (*StakeClient, error) {
	if client == nil {
		return nil, fmt.Errorf("bockchain client is required")
	}

	parsedABI, err := abi.JSON(strings.NewReader(stake.StakeMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse stake contract ABI: %w", err)
	}

	return &StakeClient{
		client: client,
		signer: signer,
		abi:    parsedABI,
	}, nil

}

// GetUserStake retrieves the stake amount for a user
// Signer is not required for this read-only operation
// Returns stake amount as string
func (s *StakeClient) GetUserStake(ctx context.Context, contractAddress string, userAddress string) (string, error) {
	// Encode the getUserStake function call
	data, err := s.abi.Pack("getUserStake", userAddress)
	if err != nil {
		return "", fmt.Errorf("failed to encode getUserStake call: %w", err)
	}

	// Execute the call
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return "", fmt.Errorf("failed to call getUserStake: %w", err)
	}

	// Parse the result
	var stakeAmount string
	err = s.abi.UnpackIntoInterface(&stakeAmount, "getUserStake", common.FromHex(result))
	if err != nil {
		return "", fmt.Errorf("failed to unpack getUserStake result: %w", err)
	}
	return stakeAmount, nil
}

// PendingRewards retrieves the pending rewards for a user
// Signer is not required for this read-only operation
// Returns rewards amount as string
func (s *StakeClient) PendingRewards(ctx context.Context, contractAddress string, userAddress string) (string, error) {
	// Encode the pendingRewards function call
	data, err := s.abi.Pack("pendingRewards", userAddress)
	if err != nil {
		return "", fmt.Errorf("failed to encode pendingRewards call: %w", err)
	}

	// Execute the call
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return "", fmt.Errorf("failed to call pendingRewards: %w", err)
	}

	// Parse the result
	var rewardsAmount string
	err = s.abi.UnpackIntoInterface(&rewardsAmount, "pendingRewards", common.FromHex(result))
	if err != nil {
		return "", fmt.Errorf("failed to unpack pendingRewards result: %w", err)
	}
	return rewardsAmount, nil
}

func (s *StakeClient) TotalStaked(ctx context.Context, contractAddress string, tokenAddress string) (string, error) {
	// Encode the totalStaked function call
	data, err := s.abi.Pack("totalStaked", common.HexToAddress(tokenAddress))
	if err != nil {
		return "", fmt.Errorf("failed to encode totalStaked call: %w", err)
	}

	// Execute the call
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return "", fmt.Errorf("failed to call totalStaked: %w", err)
	}

	// Parse the result
	var totalStaked *big.Int

	var totalStakedStr string

	err = s.abi.UnpackIntoInterface(&totalStaked, "totalStaked", common.FromHex(result))
	if err != nil {
		return "", fmt.Errorf("failed to unpack totalStaked result: %w", err)
	}
	if totalStaked != nil {
		totalStakedStr = totalStaked.String()
	} else {
		return "", fmt.Errorf("totalStaked result is nil")
	}

	return totalStakedStr, nil
}

func (s *StakeClient) ApyRates(ctx context.Context, contractAddress string, tokenAddress string) (string, error) {
	// Encode the apyRates function call
	data, err := s.abi.Pack("apyRates", common.HexToAddress(tokenAddress))
	if err != nil {
		return "", fmt.Errorf("failed to encode apyRates call: %v", err)
	}

	// Execute the call
	result, err := s.client.CallContract(ctx, contractAddress, hexutil.Encode(data), "latest")
	if err != nil {
		return "", fmt.Errorf("failed to call apyRates: %v", err)
	}

	// Parse the result
	var apyRates *big.Int

	var apyRatesStr string

	err = s.abi.UnpackIntoInterface(&apyRates, "apyRates", common.FromHex(result))
	if err != nil {
		return "", fmt.Errorf("failed to unpack apyRates result: %v", err)
	}

	if apyRates != nil {
		apyRatesStr = apyRates.String()
	} else {
		return "", fmt.Errorf("apyRates result is nil")
	}

	return apyRatesStr, nil
}

// Stake stakes tokens with a specified gas limit
func (s *StakeClient) Stake(ctx context.Context, contractAddress, tokenAddress string, amount *big.Int, nonce, gasLimit string) (string, error) {
	if s.signer == nil {
		return "", fmt.Errorf("signer is required for stake operations")
	}

	data, err := s.abi.Pack("stake", common.HexToAddress(tokenAddress), amount)
	if err != nil {
		return "", fmt.Errorf("failed to encode stake call: %w", err)
	}

	txReq := &SignTransactionRequest{
		To:       contractAddress,
		Data:     hexutil.Encode(data),
		Nonce:    nonce,
		GasLimit: gasLimit,
	}

	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send stake transaction: %w", err)
	}

	return txHash, nil
}

// Withdraw withdraws a specified amount
// Returns transaction hash
func (s *StakeClient) Withdraw(ctx context.Context, contractAddress string, tokenAddress string, amount *big.Int) (string, error) {
	if s.signer == nil {
		return "", fmt.Errorf("signer is required for withdraw operations")
	}

	data, err := s.abi.Pack("withdraw", common.HexToAddress(tokenAddress), amount)
	if err != nil {
		return "", fmt.Errorf("failed to encode withdraw call: %w", err)
	}

	txReq := &SignTransactionRequest{
		To:   contractAddress,
		Data: hexutil.Encode(data),
	}

	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send withdraw transaction: %w", err)
	}

	return txHash, nil
}

// ClaimRewards claims pending rewards
// Returns transaction hash
func (s *StakeClient) ClaimRewards(ctx context.Context, contractAddress string, tokenAddress string) (string, error) {
	if s.signer == nil {
		return "", fmt.Errorf("signer is required for claim rewards operations")
	}

	data, err := s.abi.Pack("claimRewards", common.HexToAddress(tokenAddress))
	if err != nil {
		return "", fmt.Errorf("failed to encode claimRewards call: %w", err)
	}

	txReq := &SignTransactionRequest{
		To:   contractAddress,
		Data: hexutil.Encode(data),
	}

	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send claimRewards transaction: %w", err)
	}

	return txHash, nil
}

// ClaimAllRewards claims all pending rewards
func (s *StakeClient) ClaimAllRewards(ctx context.Context, contractAddress string) (string, error) {
	if s.signer == nil {
		return "", fmt.Errorf("signer is required for claim all rewards operations")
	}

	data, err := s.abi.Pack("claimAllRewards")
	if err != nil {
		return "", fmt.Errorf("failed to encode claimAllRewards call: %w", err)
	}

	txReq := &SignTransactionRequest{
		To:   contractAddress,
		Data: hexutil.Encode(data),
	}

	txHash, err := s.signer.SignAndSendTransaction(ctx, txReq)
	if err != nil {
		return "", fmt.Errorf("failed to send claimAllRewards transaction: %w", err)
	}

	return txHash, nil
}

// EstimateStakeGas estimates gas for stake with pending block
func (s *StakeClient) EstimateStakeGas(ctx context.Context, contractAddress string, tokenAddress string, amount *big.Int, nonce string, owner common.Address) (string, error) {
	// Encode the stake function call
	data, err := s.abi.Pack("stake", common.HexToAddress(tokenAddress), amount)
	if err != nil {
		return "", fmt.Errorf("failed to encode stake call: %w", err)
	}

	// Estimate gas using pending block (will include pending approve transaction)
	gasLimit, err := s.client.EstimateGas(
		ctx,
		owner.Hex(),
		contractAddress,
		"",
		hexutil.Encode(data),
		nonce,
	)
	if err != nil {
		return "", fmt.Errorf("failed to estimate stake gas: %w", err)
	}

	return gasLimit, nil
}

// FindAllowanceStorageSlot tries to find the correct storage slot for ERC20 allowances
// by testing common slot positions (0-10)
func (s *StakeClient) FindAllowanceStorageSlot(ctx context.Context, tokenAddress string, owner common.Address, spender common.Address) (uint64, error) {
	// Try common slots (typically 1-4 for most ERC20 contracts)
	// Standard ERC20: _balances at slot 0, _allowances at slot 1 or 2
	commonSlots := []uint64{1, 2, 0, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, slot := range commonSlots {
		calculatedSlot := s.client.CalculateAllowanceSlot(ctx, owner, spender, slot)

		// Check if this slot has any value
		value, err := s.client.GetStorageAt(ctx, tokenAddress, calculatedSlot.Hex(), "latest")
		if err == nil && value != "" && value != "0x" && value != "0x0" && value != "0x0000000000000000000000000000000000000000000000000000000000000000" {
			// Found a slot with value - this might be the allowance slot
			return slot, nil
		}
	}

	// If not found, return default slot 2
	return 2, nil
}

// EstimateStakeGasWithAllowanceOverride estimates gas for stake with simulated allowance
// Simulates that the approve transaction is already confirmed by overriding storage
func (s *StakeClient) EstimateStakeGasWithAllowanceOverride(ctx context.Context, contractAddress string, tokenAddress string, amount *big.Int, nonce string, owner common.Address, spender common.Address) (string, error) {
	// Encode the stake function call
	data, err := s.abi.Pack("stake", common.HexToAddress(tokenAddress), amount)
	if err != nil {
		return "", fmt.Errorf("failed to encode stake call: %w", err)
	}

	// Try multiple storage slots to find the correct one
	slotCandidates := []uint64{2, 1, 3, 0, 4, 5}

	var lastErr error
	for _, slot := range slotCandidates {
		// Calculate storage slot for allowance mapping
		allowanceSlot := s.client.CalculateAllowanceSlot(ctx, owner, spender, slot)

		// Convert amount to hex (32 bytes padded)
		amountHex := common.BytesToHash(common.LeftPadBytes(amount.Bytes(), 32)).Hex()

		// Build state override to simulate allowance
		stateOverride := map[string]interface{}{
			tokenAddress: map[string]interface{}{
				"stateDiff": map[string]string{
					allowanceSlot.Hex(): amountHex,
				},
			},
		}

		// Try to estimate gas with this slot
		gasLimit, err := s.client.EstimateGasWithStateOverride(
			ctx,
			owner.Hex(),
			contractAddress,
			"",
			hexutil.Encode(data),
			nonce,
			stateOverride,
		)
		if err == nil {
			return gasLimit, nil
		}

		lastErr = err

	}

	// All slots failed
	return "", fmt.Errorf("failed to estimate gas with all allowance overrides: %w", lastErr)
}
