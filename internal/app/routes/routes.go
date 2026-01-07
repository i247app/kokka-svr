package routes

import (
	"github.com/i247app/gex"
	"kokka.com/kokka/internal/app/resources"
	"kokka.com/kokka/internal/app/services"
	"kokka.com/kokka/internal/handlers/http/controller"
)

func SetUpHttpRoutes(server *gex.Server, res *resources.AppResource, services *services.ServiceContainer) {
	// Static file routes for goboard UI
	staticCtrl := controller.NewStaticController("html/goboard")
	server.AddRoute("GET /goboard", staticCtrl.ServeFile("index.html"))
	server.AddRoute("GET /goboard/", staticCtrl.ServeFile("index.html"))
	server.AddRoute("GET /goboard/index.html", staticCtrl.ServeFile("index.html"))
	server.AddRoute("GET /goboard/index.css", staticCtrl.ServeFile("index.css"))
	server.AddRoute("GET /goboard/helper.js", staticCtrl.ServeFile("helper.js"))

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

	// token routes (supports VNDX, SGDX, YEXN, etc.)
	token := controller.NewTokenController(services.TokenService)
	// POST endpoints
	server.AddRoute("POST /token/mint", token.HandleMintToken)
	server.AddRoute("POST /token/burn", token.HandleBurnToken)
	server.AddRoute("POST /token/transfer", token.HandleTransferToken)
	server.AddRoute("POST /token/contract-address-info", token.HandleGetAddressInfo)

	// GET endpoints
	server.AddRoute("POST /token/balance", token.HandleGetTokenBalance)
	server.AddRoute("POST /token/mint-request", token.HandleGetMintRequest)
	server.AddRoute("POST /token/burn-request", token.HandleGetBurnRequest)
	server.AddRoute("POST /token/transaction-history", token.HandleGetTokenTransactionHistory)
}
