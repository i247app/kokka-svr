package services

import (
	"fmt"

	"kokka.com/kokka/internal/app/resources"
	"kokka.com/kokka/internal/applications/services"
	"kokka.com/kokka/internal/applications/validators"
	diSvc "kokka.com/kokka/internal/core/di/services"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain"
)

type ServiceContainer struct {
	BlockchainService diSvc.IBlockChainService
	TokenService      diSvc.ITokenService
}

func SetupServiceContainer(res *resources.AppResource) (*ServiceContainer, error) {
	// Initialize blockchain client
	blockchainConfig := blockchain.DefaultConfig()
	if res.Env.BlockchainConfig != nil && res.Env.BlockchainConfig.RPCURL != "" {
		blockchainConfig = blockchainConfig.WithBaseURL(res.Env.BlockchainConfig.RPCURL)
	}
	blockchainClient := blockchain.NewClient(blockchainConfig)

	// Initialize blockchain service (no global signer - uses per-request signing)
	blockChainValidator := validators.NewBlockChainValidator()
	blockchainService := services.NewBlockchainService(blockChainValidator, blockchainClient, nil)

	// Initialize Token service (uses per-request signers, no global signer needed)
	tokenValidator := validators.NewTokenValidator()
	tokenService, err := services.NewTokenService(
		tokenValidator,
		blockchainClient,
		res.Env.BlockchainConfig.DecryptionKey,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize token service: %w", err)
	}

	return &ServiceContainer{
		BlockchainService: blockchainService,
		TokenService:      tokenService,
	}, nil
}
