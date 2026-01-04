package dtos

// ========================================
// Request DTOs (from client to server)
// ========================================

// GetBalanceRequest represents a request to get an address balance
type GetBalanceRequest struct {
	Address string `json:"address"`
	Block   string `json:"block,omitempty"` // Optional, defaults to "latest"
}

// GetBlockRequest represents a request to get block details
type GetBlockRequest struct {
	BlockNumber string `json:"block_number"`
	FullTx      bool   `json:"full_tx,omitempty"` // If true, returns full transaction objects
}

// GetTransactionRequest represents a request to get transaction details
type GetTransactionRequest struct {
	TxHash string `json:"tx_hash"`
}

// CallContractRequest represents a request to call a contract method (read-only)
type CallContractRequest struct {
	To    string `json:"to"`
	Data  string `json:"data"`
	Block string `json:"block,omitempty"` // Optional, defaults to "latest"
}

// EstimateGasRequest represents a request to estimate gas for a transaction
type EstimateGasRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Data string `json:"data,omitempty"`
}

// SendRawTransactionRequest represents a request to broadcast a signed transaction
type SendRawTransactionRequest struct {
	SignedTx string `json:"signed_tx"`
}

// GenericRPCRequest represents a generic JSON-RPC request
// This allows clients to call any RPC method directly
type GenericRPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

// ========================================
// Response DTOs (from server to client)
// ========================================

// GetBalanceResponse represents the response for balance query
type GetBalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"` // Hex-encoded wei amount
	Block   string `json:"block"`
}

// GetBlockNumberResponse represents the response for current block number
type GetBlockNumberResponse struct {
	BlockNumber string `json:"block_number"` // Hex-encoded block number
}

// GetBlockResponse represents the response for block details
type GetBlockResponse struct {
	Block interface{} `json:"block"` // Full block object
}

// GetTransactionResponse represents the response for transaction details
type GetTransactionResponse struct {
	Transaction interface{} `json:"transaction"` // Full transaction object
}

// CallContractResponse represents the response for contract call
type CallContractResponse struct {
	Result string `json:"result"` // Hex-encoded return data
}

// EstimateGasResponse represents the response for gas estimation
type EstimateGasResponse struct {
	GasEstimate string `json:"gas_estimate"` // Hex-encoded gas amount
}

// SendRawTransactionResponse represents the response for transaction broadcast
type SendRawTransactionResponse struct {
	TxHash string `json:"tx_hash"` // Transaction hash
}

// GetGasPriceResponse represents the response for current gas price
type GetGasPriceResponse struct {
	GasPrice string `json:"gas_price"` // Hex-encoded gas price in wei
}

// GetChainIDResponse represents the response for chain ID
type GetChainIDResponse struct {
	ChainID string `json:"chain_id"` // Hex-encoded chain ID
}

// GenericRPCResponse represents a generic JSON-RPC response
type GenericRPCResponse struct {
	Result interface{} `json:"result"`
}
