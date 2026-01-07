package services

import (
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

	// Initialize transaction signer (if private key is configured)
	var txSigner *blockchain.TransactionSigner
	if res.Env.BlockchainConfig != nil && res.Env.BlockchainConfig.PrivateKey != "" {
		var err error
		txSigner, err = blockchain.NewTransactionSigner(res.Env.BlockchainConfig.PrivateKey, blockchainClient)
		if err != nil {
			// Log warning but don't fail - signing features just won't be available
			println("Warning: Failed to initialize transaction signer:", err.Error())
		}
	}

	println("RPCURL", res.Env.BlockchainConfig.RPCURL)
	println("PrivateKey", res.Env.BlockchainConfig.PrivateKey)

	// Initialize blockchain service
	blockChainValidator := validators.NewBlockChainValidator()
	blockchainService := services.NewBlockchainService(blockChainValidator, blockchainClient, txSigner)

	// Initialize Token service
	var tokenService diSvc.ITokenService
	if txSigner != nil {
		tokenClient, err := blockchain.NewTokenClient(
			blockchainClient,
			txSigner,
		)
		if err != nil {
			println("Warning: Failed to initialize token client:", err.Error())
		} else {
			tokenValidator := validators.NewTokenValidator()
			tokenService = services.NewTokenService(tokenValidator, tokenClient)
		}
	}

	return &ServiceContainer{
		BlockchainService: blockchainService,
		TokenService:      tokenService,
	}, nil
}
