package dtos

// MintTokenRequest represents a request to mint tokens
type MintTokenRequest struct {
	ContractAddress     string `json:"contract_address"`
	To                  string `json:"to"`
	Amount              string `json:"amount"`
	EncryptedPrivateKey string `json:"encrypted_private_key"`
}

// MintTokenResponse represents the response from minting tokens
type MintTokenResponse struct {
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	To              string `json:"to"`
	Amount          string `json:"amount"`
	NewBalance      string `json:"new_balance,omitempty"`
}

// BurnTokenRequest represents a request to burn tokens
type BurnTokenRequest struct {
	ContractAddress     string `json:"contract_address"`
	Amount              string `json:"amount"`
	EncryptedPrivateKey string `json:"encrypted_private_key"`
}

// BurnTokenResponse represents the response from burning tokens
type BurnTokenResponse struct {
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	Amount          string `json:"amount"`
	NewBalance      string `json:"new_balance,omitempty"`
}

// TransferTokenRequest represents a request to transfer tokens
type TransferTokenRequest struct {
	ContractAddress     string `json:"contract_address"`
	To                  string `json:"to"`
	Amount              string `json:"amount"`
	EncryptedPrivateKey string `json:"encrypted_private_key"`
}

// TransferTokenResponse represents the response from transferring tokens
type TransferTokenResponse struct {
	TxHash          string `json:"tx_hash"`
	ContractAddress string `json:"contract_address"`
	From            string `json:"from"`
	To              string `json:"to"`
	Amount          string `json:"amount"`
}

// GetTokenBalanceRequest represents a request to get token balance
type GetTokenBalanceRequest struct {
	ContractAddress string `json:"contract_address"`
	Address         string `json:"address"`
}

// GetTokenBalanceResponse represents the response with token balance
type GetTokenBalanceResponse struct {
	ContractAddress string `json:"contract_address"`
	Address         string `json:"address"`
	Balance         string `json:"balance"`
}
