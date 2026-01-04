package services

import (
	"kokka.com/kokka/internal/app/resources"
	"kokka.com/kokka/internal/applications/services"
	diSvc "kokka.com/kokka/internal/core/di/services"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain"
)

type ServiceContainer struct {
	BlockchainService diSvc.IBlockChainService
}

func SetupServiceContainer(res *resources.AppResource) (*ServiceContainer, error) {
	// Initialize blockchain client
	blockchainConfig := blockchain.DefaultConfig()
	if res.Env.BlockchainConfig != nil && res.Env.BlockchainConfig.RPCURL != "" {
		blockchainConfig = blockchainConfig.WithBaseURL(res.Env.BlockchainConfig.RPCURL)
	}
	blockchainClient := blockchain.NewClient(blockchainConfig)

	// Initialize blockchain service
	blockchainService := services.NewBlockchainService(blockchainClient)

	// Initialize blockchain controller

	return &ServiceContainer{
		BlockchainService: blockchainService,
	}, nil
}
