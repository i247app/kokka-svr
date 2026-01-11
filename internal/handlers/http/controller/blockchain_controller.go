package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"kokka.com/kokka/internal/applications/dtos"
	diSvc "kokka.com/kokka/internal/core/di/services"
	"kokka.com/kokka/internal/shared/constant/status"
	"kokka.com/kokka/internal/shared/response"
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

// GetBlockNumber handles GET /blockchain/block-number
func (c *BlockchainController) GetBlockNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := c.blockchainService.GetBlockNumber(ctx)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GetGasPrice handles GET /blockchain/gas-price
func (c *BlockchainController) GetGasPrice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := c.blockchainService.GetGasPrice(ctx)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GetChainID handles GET /blockchain/chain-id
func (c *BlockchainController) GetChainID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := c.blockchainService.GetChainID(ctx)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GetBalance handles POST /blockchain/balance
func (c *BlockchainController) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.GetBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
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

// GetBlock handles POST /blockchain/block
func (c *BlockchainController) GetBlock(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.GetBlockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
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

// GetTransaction handles POST /blockchain/transaction
func (c *BlockchainController) GetTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.GetTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
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

// CallContract handles POST /blockchain/call
func (c *BlockchainController) CallContract(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.CallContractRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
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

// EstimateGas handles POST /blockchain/estimate-gas
func (c *BlockchainController) EstimateGas(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.EstimateGasRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
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

// SendRawTransaction handles POST /blockchain/send-transaction
func (c *BlockchainController) SendRawTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.SendRawTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
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

// SignAndSendTransaction handles POST /blockchain/sign-and-send
func (c *BlockchainController) SignAndSendTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dtos.SignAndSendTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	// Call service
	result, err := c.blockchainService.SignAndSendTransaction(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// GenericRPCCall handles POST /blockchain/rpc
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

	// Call service
	result, err := c.blockchainService.GenericRPCCall(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}
