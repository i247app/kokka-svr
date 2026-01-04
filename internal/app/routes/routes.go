package routes

import (
	"net/http"

	"github.com/i247app/gex"
	"kokka.com/kokka/internal/app/resources"
	"kokka.com/kokka/internal/app/services"
	"kokka.com/kokka/internal/handlers/http/controller"
	"kokka.com/kokka/internal/shared/constant/status"
	"kokka.com/kokka/internal/shared/utils/response"
)

func SetUpHttpRoutes(server *gex.Server, res *resources.AppResource, services *services.ServiceContainer) {
	// Test route
	server.AddRoute("POST /test", func(w http.ResponseWriter, r *http.Request) {
		response.WriteJson(w, r.Context(), "hehe cho test thoi", nil, status.OK)
	})

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
	server.AddRoute("POST /blockchain/rpc", bc.GenericRPCCall)
}
