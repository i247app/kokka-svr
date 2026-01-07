package validators

import (
	"errors"
	"strings"

	"kokka.com/kokka/internal/applications/dtos"
)

// ValidateGetBalanceRequest validates a get balance request
func ValidateGetBalanceRequest(req *dtos.GetBalanceRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.Address == "" {
		return errors.New("address is required")
	}

	if !isValidEthereumAddress(req.Address) {
		return errors.New("invalid ethereum address format")
	}

	// Validate block parameter if provided
	if req.Block != "" && !isValidBlockParameter(req.Block) {
		return errors.New("invalid block parameter")
	}

	return nil
}

// ValidateGetBlockRequest validates a get block request
func ValidateGetBlockRequest(req *dtos.GetBlockRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.BlockNumber == "" {
		return errors.New("block_number is required")
	}

	if !isValidBlockParameter(req.BlockNumber) {
		return errors.New("invalid block_number format")
	}

	return nil
}

// ValidateGetTransactionRequest validates a get transaction request
func ValidateGetTransactionRequest(req *dtos.GetTransactionRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.TxHash == "" {
		return errors.New("tx_hash is required")
	}

	if !isValidHash(req.TxHash) {
		return errors.New("invalid transaction hash format")
	}

	return nil
}

// ValidateCallContractRequest validates a call contract request
func ValidateCallContractRequest(req *dtos.CallContractRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.To == "" {
		return errors.New("to address is required")
	}

	if !isValidEthereumAddress(req.To) {
		return errors.New("invalid to address format")
	}

	if req.Value == "" {
		return errors.New("data is required")
	}

	if !isValidHexData(req.Value) {
		return errors.New("invalid data format (must be hex string with 0x prefix)")
	}

	// Validate block parameter if provided
	if req.Block != "" && !isValidBlockParameter(req.Block) {
		return errors.New("invalid block parameter")
	}

	return nil
}

// ValidateEstimateGasRequest validates an estimate gas request
func ValidateEstimateGasRequest(req *dtos.EstimateGasRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.From != "" && !isValidEthereumAddress(req.From) {
		return errors.New("invalid from address format")
	}

	if req.To == "" {
		return errors.New("to address is required")
	}

	if !isValidEthereumAddress(req.To) {
		return errors.New("invalid to address format")
	}

	// Data is optional, but if provided, should be valid hex
	if req.Value != "" && !isValidHexData(req.Value) {
		return errors.New("invalid data format (must be hex string with 0x prefix)")
	}

	return nil
}

// ValidateSendRawTransactionRequest validates a send raw transaction request
func ValidateSendRawTransactionRequest(req *dtos.SendRawTransactionRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.SignedTx == "" {
		return errors.New("signed_tx is required")
	}

	if !isValidHexData(req.SignedTx) {
		return errors.New("invalid signed transaction format (must be hex string with 0x prefix)")
	}

	return nil
}

// ValidateGenericRPCRequest validates a generic RPC request
func ValidateGenericRPCRequest(req *dtos.GenericRPCRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	if req.Method == "" {
		return errors.New("method is required")
	}

	return nil
}

// ========================================
// Helper validation functions
// ========================================

// isValidEthereumAddress checks if a string is a valid Ethereum address
func isValidEthereumAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if !strings.HasPrefix(address, "0x") {
		return false
	}
	// Check if remaining characters are valid hex
	return isValidHex(address[2:])
}

// isValidHash checks if a string is a valid 32-byte hash
func isValidHash(hash string) bool {
	if len(hash) != 66 {
		return false
	}
	if !strings.HasPrefix(hash, "0x") {
		return false
	}
	return isValidHex(hash[2:])
}

// isValidHexData checks if a string is valid hex data (with 0x prefix)
func isValidHexData(data string) bool {
	if !strings.HasPrefix(data, "0x") {
		return false
	}
	// Data must have even length (excluding 0x prefix)
	if len(data[2:])%2 != 0 {
		return false
	}
	return isValidHex(data[2:])
}

// isValidHex checks if a string contains only valid hex characters
func isValidHex(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// isValidBlockParameter checks if a block parameter is valid
// Valid values: "latest", "earliest", "pending", or hex-encoded block number
func isValidBlockParameter(block string) bool {
	// Check for special values
	if block == "latest" || block == "earliest" || block == "pending" {
		return true
	}

	// Check for hex-encoded number
	if strings.HasPrefix(block, "0x") {
		return isValidHex(block[2:])
	}

	return false
}

// ValidateSignAndMintRequest validates a sign and mint request
func ValidateSignAndMintRequest(req *dtos.SignAndMintRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	return nil
}

// ValidateSignAndBurnRequest validates a sign and burn request
func ValidateSignAndBurnRequest(req *dtos.SignAndBurnRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}

	return nil
}

// ValidateSignAndSendTransactionRequest validates a sign and send transaction request
func ValidateSignAndSendTransactionRequest(req *dtos.SignAndSendTransactionRequest) error {
	if req == nil {
		return errors.New("request cannot be nil")
	}
	if req.PrivateKey == "" {
		return errors.New("private_key is required")
	}
	if req.ContractAddress == "" {
		return errors.New("contract_address is required")
	}
	if !isValidEthereumAddress(req.ContractAddress) {
		return errors.New("invalid contract_address format")
	}

	return nil
}
