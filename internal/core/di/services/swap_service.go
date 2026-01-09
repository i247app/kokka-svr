package di

import (
	"context"

	"kokka.com/kokka/internal/applications/dtos"
)

type ISwapService interface {
	Swap(ctx context.Context, req *dtos.SwapTokenRequest) (*dtos.SwapTokenResponse, error)
	GetQuote(ctx context.Context, req *dtos.GetSwapQuoteRequest) (*dtos.GetSwapQuoteResponse, error)
	GetSwapInfo(ctx context.Context, req *dtos.GetSwapInfoRequest) (*dtos.GetSwapInfoResponse, error)
}
