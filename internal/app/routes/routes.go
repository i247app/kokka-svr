package routes

import (
	"github.com/i247app/gex"
	"kokka.com/kokka/internal/app/resources"
	"kokka.com/kokka/internal/app/services"
	"kokka.com/kokka/internal/handlers/http/controller"
)

func SetUpHttpRoutes(server *gex.Server, res *resources.AppResource, services *services.ServiceContainer) {

	// blockchain routes
	bc := controller.NewBlockchainController(services.BlockchainService)
	// GET endpoints
	server.AddRoute("GET /blockchain/block-number", bc.GetBlockNumber)
	server.AddRoute("GET /blockchain/gas-price", bc.GetGasPrice)
	server.AddRoute("GET /blockchain/chain-id", bc.GetChainID)

	// POST endpoints
	server.AddRoute("POST /blockchain/balance", bc.GetBalance)
	server.AddRoute("POST /blockchain/block", bc.GetBlock)
	server.AddRoute("POST /blockchain/transaction", bc.GetTransaction)
	server.AddRoute("POST /blockchain/call", bc.CallContract)
	server.AddRoute("POST /blockchain/estimate-gas", bc.EstimateGas)
	server.AddRoute("POST /blockchain/send-transaction", bc.SendRawTransaction)
	server.AddRoute("POST /blockchain/sign-and-send", bc.SignAndSendTransaction)
	server.AddRoute("POST /blockchain/rpc", bc.GenericRPCCall)

	// vndx token routes
	vndx := controller.NewVndxController(services.VNDXService)
	// POST endpoints
	server.AddRoute("POST /vndx/mint", vndx.HandleMintVNDX)
	server.AddRoute("POST /vndx/burn", vndx.HandleBurnVNDX)
	server.AddRoute("POST /vndx/transfer", vndx.HandleTransferVNDX)

	// GET endpoints
	server.AddRoute("GET /vndx/balance", vndx.HandleGetVNDXBalance)
	server.AddRoute("GET /vndx/mint-request", vndx.HandleGetMintRequest)
	server.AddRoute("GET /vndx/burn-request", vndx.HandleGetBurnRequest)
	server.AddRoute("GET /vndx/transaction-history", vndx.HandleGetVNDXTransactionHistory)
}
