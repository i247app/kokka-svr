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

type VndxController struct {
	vndxService diSvc.IVNDXService
}

func NewVndxController(vndxService diSvc.IVNDXService) *VndxController {
	return &VndxController{
		vndxService: vndxService,
	}
}

// HandleMintVNDX handles POST /vndx/mint
func (c *VndxController) HandleMintVNDX(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.vndxService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("VNDX service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.MintVNDXRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.vndxService.Mint(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleBurnVNDX handles POST /vndx/burn
func (c *VndxController) HandleBurnVNDX(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.vndxService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("VNDX service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.BurnVNDXRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.vndxService.Burn(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleGetVNDXBalance handles GET /vndx/balance
func (c *VndxController) HandleGetVNDXBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.vndxService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("VNDX service is not configured"), status.INTERNAL)
		return
	}

	// Get address from query parameter
	address := r.URL.Query().Get("address")
	if address == "" {
		response.WriteJson(w, ctx, nil, fmt.Errorf("address parameter is required"), status.FAIL)
		return
	}

	req := &dtos.GetVNDXBalanceRequest{
		Address: address,
	}

	result, err := c.vndxService.GetBalance(ctx, req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleGetMintRequest retrieves mint requests (not implemented yet)
func (c *VndxController) HandleGetMintRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response.WriteJson(w, ctx, nil, fmt.Errorf("not implemented"), status.FAIL)
}

// HandleGetBurnRequest retrieves burn requests (not implemented yet)
func (c *VndxController) HandleGetBurnRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response.WriteJson(w, ctx, nil, fmt.Errorf("not implemented"), status.FAIL)
}

// HandleTransferVNDX handles POST /vndx/transfer
func (c *VndxController) HandleTransferVNDX(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.vndxService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("VNDX service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.TransferVNDXRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.vndxService.Transfer(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleGetVNDXTransactionHistory retrieves the transaction history for VNDX tokens (not implemented yet)
func (c *VndxController) HandleGetVNDXTransactionHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response.WriteJson(w, ctx, nil, fmt.Errorf("not implemented"), status.FAIL)
}
