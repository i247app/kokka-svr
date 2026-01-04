package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"kokka.com/kokka/internal/applications/dtos"
	"kokka.com/kokka/internal/applications/validators"
	diSvc "kokka.com/kokka/internal/core/di/services"
	"kokka.com/kokka/internal/shared/constant/status"
	"kokka.com/kokka/internal/shared/utils/response"
)

// BlockchainController handles blockchain-related HTTP requests
type BlockchainController struct {
	blockchainService diSvc.IBlockChainService
}

// NewBlockchainController creates a new blockchain controller
func NewBlockchainController(blockchainService diSvc.IBlockChainService) *BlockchainController {
	return &BlockchainController{
		blockchainService: blockchainService,
	}
}

// GetBlockNumber handles GET /api/blockchain/block-number
func (c *BlockchainController) GetBlockNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := c.blockchainService.GetBlockNumber(ctx)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GetBalance handles POST /api/blockchain/balance
func (c *BlockchainController) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.GetBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	// Validate request
	if err := validators.ValidateGetBalanceRequest(&req); err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}

	// Call service
	result, err := c.blockchainService.GetBalance(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GetBlock handles POST /api/blockchain/block
func (c *BlockchainController) GetBlock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.GetBlockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	// Validate request
	if err := validators.ValidateGetBlockRequest(&req); err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}

	// Call service
	result, err := c.blockchainService.GetBlock(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GetTransaction handles POST /api/blockchain/transaction
func (c *BlockchainController) GetTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.GetTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	// Validate request
	if err := validators.ValidateGetTransactionRequest(&req); err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}

	// Call service
	result, err := c.blockchainService.GetTransaction(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// CallContract handles POST /api/blockchain/call
func (c *BlockchainController) CallContract(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.CallContractRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	// Validate request
	if err := validators.ValidateCallContractRequest(&req); err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}

	// Call service
	result, err := c.blockchainService.CallContract(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// EstimateGas handles POST /api/blockchain/estimate-gas
func (c *BlockchainController) EstimateGas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.EstimateGasRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	// Validate request
	if err := validators.ValidateEstimateGasRequest(&req); err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}

	// Call service
	result, err := c.blockchainService.EstimateGas(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// SendRawTransaction handles POST /api/blockchain/send-transaction
func (c *BlockchainController) SendRawTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.SendRawTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	// Validate request
	if err := validators.ValidateSendRawTransactionRequest(&req); err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}

	// Call service
	result, err := c.blockchainService.SendRawTransaction(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GetGasPrice handles GET /api/blockchain/gas-price
func (c *BlockchainController) GetGasPrice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := c.blockchainService.GetGasPrice(ctx)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GetChainID handles GET /api/blockchain/chain-id
func (c *BlockchainController) GetChainID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := c.blockchainService.GetChainID(ctx)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GenericRPCCall handles POST /api/blockchain/rpc
// This allows calling any JSON-RPC method directly
func (c *BlockchainController) GenericRPCCall(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}
	defer r.Body.Close()

	var req dtos.GenericRPCRequest
	if err := json.Unmarshal(body, &req); err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}

	// Validate request
	if err := validators.ValidateGenericRPCRequest(&req); err != nil {
		response.WriteJson(w, ctx, nil, err, status.FAIL)
		return
	}

	// Call service
	result, err := c.blockchainService.GenericRPCCall(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}
