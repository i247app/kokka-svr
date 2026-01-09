package dtos

// SwapTokenRequest represents a request to swap tokens
type SwapTokenRequest struct {
	ContractAddress     string `json:"contract_address"`     // Swap contract address
	AmountIn            string `json:"amount_in"`            // Amount of input token to swap
	Direction           string `json:"direction"`            // "AtoB" or "BtoA"
	EncryptedPrivateKey string `json:"encrypted_private_key"` // Encrypted private key for signing
}

// SwapTokenResponse represents the response from swapping tokens
type SwapTokenResponse struct {
	TxHash          string `json:"tx_hash"`           // Transaction hash
	ContractAddress string `json:"contract_address"`  // Swap contract address
	AmountIn        string `json:"amount_in"`         // Amount of input token swapped
	AmountOut       string `json:"amount_out"`        // Amount of output token received (estimated)
	FromToken       string `json:"from_token"`        // Address of token swapped from
	ToToken         string `json:"to_token"`          // Address of token swapped to
	Direction       string `json:"direction"`         // "AtoB" or "BtoA"
}

// GetSwapQuoteRequest represents a request to get a swap quote
type GetSwapQuoteRequest struct {
	ContractAddress string `json:"contract_address"` // Swap contract address
	AmountIn        string `json:"amount_in"`        // Amount of input token
	Direction       string `json:"direction"`        // "AtoB" or "BtoA"
}

// GetSwapQuoteResponse represents the response with swap quote
type GetSwapQuoteResponse struct {
	ContractAddress string `json:"contract_address"` // Swap contract address
	AmountIn        string `json:"amount_in"`        // Amount of input token
	AmountOut       string `json:"amount_out"`       // Expected amount of output token
	Direction       string `json:"direction"`        // "AtoB" or "BtoA"
	ExchangeRate    string `json:"exchange_rate"`    // Current exchange rate
}

// GetSwapInfoRequest represents a request to get swap contract info
type GetSwapInfoRequest struct {
	ContractAddress string `json:"contract_address"` // Swap contract address
}

// GetSwapInfoResponse represents the response with swap contract info
type GetSwapInfoResponse struct {
	ContractAddress string `json:"contract_address"` // Swap contract address
	TokenA          string `json:"token_a"`          // Address of token A
	TokenB          string `json:"token_b"`          // Address of token B
	ReserveA        string `json:"reserve_a"`        // Reserve of token A
	ReserveB        string `json:"reserve_b"`        // Reserve of token B
	ExchangeRate    string `json:"exchange_rate"`    // Current exchange rate
}
