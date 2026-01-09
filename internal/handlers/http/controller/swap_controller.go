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

type SwapController struct {
	swapService diSvc.ISwapService
}

func NewSwapController(swapService diSvc.ISwapService) *SwapController {
	return &SwapController{
		swapService: swapService,
	}
}

// HandleSwap handles POST /swap
func (c *SwapController) HandleSwap(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.swapService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("swap service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.SwapTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.swapService.Swap(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleGetSwapQuote handles POST /swap/quote
func (c *SwapController) HandleGetSwapQuote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.swapService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("swap service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.GetSwapQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.swapService.GetQuote(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}

// HandleGetSwapInfo handles POST /swap/info
func (c *SwapController) HandleGetSwapInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.swapService == nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("swap service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.GetSwapInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.swapService.GetSwapInfo(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.OK)
}
