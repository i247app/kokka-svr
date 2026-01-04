package blockchain

import "encoding/json"

// JSONRPCRequest represents a JSON-RPC 2.0 request
type JSONRPCRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int64       `json:"id"`
}

// JSONRPCResponse represents a JSON-RPC 2.0 response
type JSONRPCResponse struct {
	JsonRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *JSONRPCError   `json:"error,omitempty"`
	ID      int64           `json:"id"`
}

// JSONRPCError represents a JSON-RPC 2.0 error
type JSONRPCError struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

// IsError checks if the response contains an error
func (r *JSONRPCResponse) IsError() bool {
	return r.Error != nil
}

// GetResultAsString returns the result as a string
func (r *JSONRPCResponse) GetResultAsString() (string, error) {
	var result string
	if err := json.Unmarshal(r.Result, &result); err != nil {
		return "", err
	}
	return result, nil
}

// GetResultAsInt returns the result as an int
func (r *JSONRPCResponse) GetResultAsInt() (int64, error) {
	var result int64
	if err := json.Unmarshal(r.Result, &result); err != nil {
		return 0, err
	}
	return result, nil
}

// UnmarshalResult unmarshals the result into the provided interface
func (r *JSONRPCResponse) UnmarshalResult(v interface{}) error {
	return json.Unmarshal(r.Result, v)
}
