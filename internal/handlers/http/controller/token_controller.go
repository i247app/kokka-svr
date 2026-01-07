package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kokka.com/kokka/internal/applications/dtos"
	diSvc "kokka.com/kokka/internal/core/di/services"
	"kokka.com/kokka/internal/shared/constant/status"
	"kokka.com/kokka/internal/shared/utils/response"
)

type TokenController struct {
	tokenService diSvc.ITokenService
}

func NewTokenController(tokenService diSvc.ITokenService) *TokenController {
	return &TokenController{
		tokenService: tokenService,
	}
}

// HandleMintToken handles POST /token/mint
func (c *TokenController) HandleMintToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.tokenService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("token service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.MintTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.tokenService.Mint(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleBurnToken handles POST /token/burn
func (c *TokenController) HandleBurnToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.tokenService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("token service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.BurnTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.tokenService.Burn(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleGetTokenBalance handles POST /token/balance
func (c *TokenController) HandleGetTokenBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.tokenService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("token service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.GetTokenBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.tokenService.GetBalance(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleGetMintRequest retrieves mint requests (not implemented yet)
func (c *TokenController) HandleGetMintRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response.WriteJson(w, ctx, nil, fmt.Errorf("not implemented"), status.FAIL)
}

// HandleGetBurnRequest retrieves burn requests (not implemented yet)
func (c *TokenController) HandleGetBurnRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response.WriteJson(w, ctx, nil, fmt.Errorf("not implemented"), status.FAIL)
}

// HandleTransferToken handles POST /token/transfer
func (c *TokenController) HandleTransferToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.tokenService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("token service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.TransferTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.tokenService.Transfer(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleGetTokenTransactionHistory retrieves the transaction history for tokens (not implemented yet)
func (c *TokenController) HandleGetTokenTransactionHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response.WriteJson(w, ctx, nil, fmt.Errorf("not implemented"), status.FAIL)
}

// HandleGetAddressInfo handles POST /token/address-info
func (c *TokenController) HandleGetAddressInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if c.tokenService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("token service is not configured"), status.INTERNAL)
		return
	}
	var req dtos.GetAddressInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.tokenService.GetAddressInfo(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}
