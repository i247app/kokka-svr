package di

import (
	"context"

	"kokka.com/kokka/internal/applications/dtos"
)

type IBlockChainService interface {
	GetBlockNumber(ctx context.Context) (*dtos.GetBlockNumberResponse, error)
	GetBalance(ctx context.Context, req *dtos.GetBalanceRequest) (*dtos.GetBalanceResponse, error)
	GetBlock(ctx context.Context, req *dtos.GetBlockRequest) (*dtos.GetBlockResponse, error)
	GetTransaction(ctx context.Context, req *dtos.GetTransactionRequest) (*dtos.GetTransactionResponse, error)
	CallContract(ctx context.Context, req *dtos.CallContractRequest) (*dtos.CallContractResponse, error)
	EstimateGas(ctx context.Context, req *dtos.EstimateGasRequest) (*dtos.EstimateGasResponse, error)
	SendRawTransaction(ctx context.Context, req *dtos.SendRawTransactionRequest) (*dtos.SendRawTransactionResponse, error)
	GetGasPrice(ctx context.Context) (*dtos.GetGasPriceResponse, error)
	GetChainID(ctx context.Context) (*dtos.GetChainIDResponse, error)
	GenericRPCCall(ctx context.Context, req *dtos.GenericRPCRequest) (*dtos.GenericRPCResponse, error)
}
