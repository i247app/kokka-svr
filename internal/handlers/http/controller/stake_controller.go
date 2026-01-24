package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kokka.com/kokka/internal/applications/dtos"
	diSvc "kokka.com/kokka/internal/core/di/services"
	"kokka.com/kokka/internal/shared/constant/status"
	"kokka.com/kokka/internal/shared/response"
)

type StakeController struct {
	stakeService diSvc.IStakeService
}

func NewStakeController(stakeService diSvc.IStakeService) *StakeController {
	return &StakeController{
		stakeService: stakeService,
	}
}

// HandleGetUserStake handles POST /stake/user-stake
func (c *StakeController) HandleGetUserStake(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.stakeService == nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("stake service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.GetUserStakeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.stakeService.GetUserStake(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.SUCCESS)
}

// HandleGetPendingRewards handles POST /stake/pending-rewards
func (c *StakeController) HandleGetPendingRewards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.stakeService == nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("stake service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.GetPendingRewardsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.stakeService.GetPendingRewards(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.SUCCESS)
}

// HandleStakeToken handles POST /stake/stake-token
func (c *StakeController) HandleStakeToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.stakeService == nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("stake service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.StakeTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.stakeService.StakeToken(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.SUCCESS)
}

// HandleWithdrawToken handles POST /stake/withdraw-token
func (c *StakeController) HandleWithdrawToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.stakeService == nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("stake service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.WithdrawTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.stakeService.WithdrawToken(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.SUCCESS)
}

// HandleClaimRewards handles POST /stake/claim-rewards
func (c *StakeController) HandleClaimRewards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.stakeService == nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("stake service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.ClaimRewardsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.stakeService.ClaimRewards(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.SUCCESS)
}

// HandleClaimRewards handles POST /stake/claim-rewards
func (c *StakeController) HandleClaimAllRewards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if c.stakeService == nil {
		response.WriteJson(w, r.Context(), nil, fmt.Errorf("stake service is not configured"), status.INTERNAL)
		return
	}

	var req dtos.ClaimAllRewardsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJson(w, ctx, nil, fmt.Errorf("invalid parameters"), status.FAIL)
		return
	}

	result, err := c.stakeService.ClaimAllRewards(ctx, &req)
	if err != nil {
		response.WriteJson(w, ctx, nil, err, status.INTERNAL)
		return
	}

	response.WriteJson(w, ctx, result, nil, status.SUCCESS)
}
