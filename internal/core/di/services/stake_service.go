package di

import (
	"context"

	"kokka.com/kokka/internal/applications/dtos"
)

type IStakeService interface {
	GetUserStake(ctx context.Context, req *dtos.GetUserStakeRequest) (*dtos.GetUserStakeResponse, error)
	GetPendingRewards(ctx context.Context, req *dtos.GetPendingRewardsRequest) (*dtos.GetPendingRewardsResponse, error)
	StakeToken(ctx context.Context, req *dtos.StakeTokenRequest) (*dtos.StakeTokenResponse, error)
	WithdrawToken(ctx context.Context, req *dtos.WithdrawTokenRequest) (*dtos.WithdrawTokenResponse, error)
	ClaimRewards(ctx context.Context, req *dtos.ClaimRewardsRequest) (*dtos.ClaimRewardsResponse, error)
	ClaimAllRewards(ctx context.Context, req *dtos.ClaimAllRewardsRequest) (*dtos.ClaimAllRewardsResponse, error)
}
