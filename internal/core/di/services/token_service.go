package di

import (
	"context"

	"kokka.com/kokka/internal/applications/dtos"
)

type ITokenService interface {
	Mint(ctx context.Context, req *dtos.MintTokenRequest) (*dtos.MintTokenResponse, error)
	Burn(ctx context.Context, req *dtos.BurnTokenRequest) (*dtos.BurnTokenResponse, error)
	Transfer(ctx context.Context, req *dtos.TransferTokenRequest) (*dtos.TransferTokenResponse, error)
	GetBalance(ctx context.Context, req *dtos.GetTokenBalanceRequest) (*dtos.GetTokenBalanceResponse, error)
}
