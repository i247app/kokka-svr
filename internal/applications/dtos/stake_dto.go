package dtos

// GetUserStakeRequest represents a request to get a user's stake amount
type GetUserStakeRequest struct {
	ContractAddress string `json:"contract_address"`
	UserAddress     string `json:"user_address"`
}

// GetUserStakeResponse represents the response with a user's stake amount
type GetUserStakeResponse struct {
	ContractAddress string `json:"contract_address"`
	StakeAmount     string `json:"stake_amount"`
}

// GetPendingRewardsRequest represents a request to get a user's pending rewards
type GetPendingRewardsRequest struct {
	ContractAddress string `json:"contract_address"`
	UserAddress     string `json:"user_address"`
}

// GetPendingRewardsResponse represents the response with a user's pending rewards
type GetPendingRewardsResponse struct {
	ContractAddress string `json:"contract_address"`
	RewardsAmount   string `json:"rewards_amount"`
}

// GetTotalStakedRequest represents a request to get the total staked amount
type GetTotalStakedRequest struct {
	ContractAddress string `json:"contract_address"`
	TokenAddress    string `json:"token_address"`
}

// GetTotalStakedResponse represents the response with the total staked amount
type GetTotalStakedResponse struct {
	ContractAddress string `json:"contract_address"`
	TotalStaked     string `json:"total_staked"`
}

// GetApyRatesRequest represents a request to get APY rates
type GetApyRatesRequest struct {
	ContractAddress string `json:"contract_address"`
	TokenAddress    string `json:"token_address"`
}

// GetApyRatesResponse represents the response with APY rates
type GetApyRatesResponse struct {
	ContractAddress string  `json:"contract_address"`
	ApyRates        float64 `json:"apy_rates"`
}

// StakeTokenRequest represents a request to stake tokens
type StakeTokenRequest struct {
	ContractAddress     string `json:"contract_address"`
	TokenAddress        string `json:"token_address"`
	Amount              string `json:"amount"`
	EncryptedPrivateKey string `json:"encrypted_private_key"`
}

// StakeTokenResponse represents the response from staking tokens
type StakeTokenResponse struct {
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
}

// WithdrawTokenRequest represents a request to withdraw staked tokens
type WithdrawTokenRequest struct {
	ContractAddress     string `json:"contract_address"`
	TokenAddress        string `json:"token_address"`
	Amount              string `json:"amount"`
	EncryptedPrivateKey string `json:"encrypted_private_key"`
}

// WithdrawTokenResponse represents the response from withdrawing staked tokens
type WithdrawTokenResponse struct {
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Amount          string `json:"amount"`
}

// ClaimRewardsRequest represents a request to claim staking rewards
type ClaimRewardsRequest struct {
	ContractAddress     string `json:"contract_address"`
	TokenAddress        string `json:"token_address"`
	EncryptedPrivateKey string `json:"encrypted_private_key"`
}

// ClaimRewardsResponse represents the response from claiming staking rewards
type ClaimRewardsResponse struct {
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
}

// ClaimAllRewardsRequest represents a request to claim all staking rewards
type ClaimAllRewardsRequest struct {
	ContractAddress     string `json:"contract_address"`
	EncryptedPrivateKey string `json:"encrypted_private_key"`
}

// ClaimAllRewardsResponse represents the response from claiming all staking rewards
type ClaimAllRewardsResponse struct {
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
}
