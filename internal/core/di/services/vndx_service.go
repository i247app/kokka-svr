package di

import (
	"context"

	"kokka.com/kokka/internal/applications/dtos"
)

type IVNDXService interface {
	Mint(ctx context.Context, req *dtos.MintVNDXRequest) (*dtos.MintVNDXResponse, error)
	Burn(ctx context.Context, req *dtos.BurnVNDXRequest) (*dtos.BurnVNDXResponse, error)
	Transfer(ctx context.Context, req *dtos.TransferVNDXRequest) (*dtos.TransferVNDXResponse, error)
	GetBalance(ctx context.Context, req *dtos.GetVNDXBalanceRequest) (*dtos.GetVNDXBalanceResponse, error)
}
