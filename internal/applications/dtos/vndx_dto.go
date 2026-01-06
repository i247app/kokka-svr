package dtos

// MintVNDXRequest represents a request to mint VNDX tokens
type MintVNDXRequest struct {
	To     string `json:"to"`
	Amount string `json:"amount"`
}

// MintVNDXResponse represents the response from minting VNDX tokens
type MintVNDXResponse struct {
	TxHash     string `json:"tx_hash"`
	To         string `json:"to"`
	Amount     string `json:"amount"`
	NewBalance string `json:"new_balance,omitempty"`
}

// BurnVNDXRequest represents a request to burn VNDX tokens
type BurnVNDXRequest struct {
	Amount string `json:"amount"`
}

// BurnVNDXResponse represents the response from burning VNDX tokens
type BurnVNDXResponse struct {
	TxHash     string `json:"tx_hash"`
	Amount     string `json:"amount"`
	NewBalance string `json:"new_balance,omitempty"`
}

// TransferVNDXRequest represents a request to transfer VNDX tokens
type TransferVNDXRequest struct {
	To     string `json:"to"`
	Amount string `json:"amount"`
}

// TransferVNDXResponse represents the response from transferring VNDX tokens
type TransferVNDXResponse struct {
	TxHash string `json:"tx_hash"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

// GetVNDXBalanceRequest represents a request to get VNDX balance
type GetVNDXBalanceRequest struct {
	Address string `json:"address"`
}

// GetVNDXBalanceResponse represents the response with VNDX balance
type GetVNDXBalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}
