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
	VNDXService       diSvc.IVNDXService
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

	// Initialize VNDX service
	var vndxService diSvc.IVNDXService
	if res.Env.BlockchainConfig.VNDXIssuerAddress != "" && txSigner != nil {
		vndxClient, err := blockchain.NewVNDXClient(
			res.Env.BlockchainConfig.VNDXIssuerAddress,
			blockchainClient,
			txSigner,
		)
		if err != nil {
			println("Warning: Failed to initialize VNDX client:", err.Error())
		} else {
			vndxValidator := validators.NewVNDXValidator()
			vndxService = services.NewVNDXService(vndxValidator, vndxClient)
		}
	}

	return &ServiceContainer{
		BlockchainService: blockchainService,
		VNDXService:       vndxService,
	}, nil
}
