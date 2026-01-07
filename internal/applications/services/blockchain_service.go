package services

import (
	"context"
	"encoding/json"
	"fmt"

	"kokka.com/kokka/internal/applications/dtos"
	"kokka.com/kokka/internal/applications/validators"
	"kokka.com/kokka/internal/driven-adapter/external/blockchain"
)

// BlockchainService handles blockchain-related business logic
type BlockchainService struct {
	validator validators.IBlockchainValidator
	client    *blockchain.Client
	signer    *blockchain.TransactionSigner
}

// NewBlockchainService creates a new blockchain service
func NewBlockchainService(
	validator validators.IBlockchainValidator,
	client *blockchain.Client,
	signer *blockchain.TransactionSigner,
) *BlockchainService {
	return &BlockchainService{
		validator: validator,
		client:    client,
		signer:    signer,
	}
}

// GetBlockNumber returns the current block number
func (s *BlockchainService) GetBlockNumber(ctx context.Context) (*dtos.GetBlockNumberResponse, error) {
	blockNumber, err := s.client.GetBlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get block number: %w", err)
	}

	return &dtos.GetBlockNumberResponse{
		BlockNumber: blockNumber,
	}, nil
}

// GetGasPrice returns the current gas price
func (s *BlockchainService) GetGasPrice(ctx context.Context) (*dtos.GetGasPriceResponse, error) {
	gasPrice, err := s.client.GetGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	return &dtos.GetGasPriceResponse{
		GasPrice: gasPrice,
	}, nil
}

// GetChainID returns the chain ID
func (s *BlockchainService) GetChainID(ctx context.Context) (*dtos.GetChainIDResponse, error) {
	chainID, err := s.client.GetChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	return &dtos.GetChainIDResponse{
		ChainID: chainID,
	}, nil
}

// GetBalance returns the balance of an address
func (s *BlockchainService) GetBalance(ctx context.Context, req *dtos.GetBalanceRequest) (*dtos.GetBalanceResponse, error) {
	// Validate request
	if err := s.validator.ValidateGetBalanceRequest(req); err != nil {
		return nil, err
	}

	// Default to "latest" if block is not specified
	block := req.Block
	if block == "" {
		block = "latest"
	}

	balance, err := s.client.GetBalance(ctx, req.Address, block)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return &dtos.GetBalanceResponse{
		Address: req.Address,
		Balance: balance,
		Block:   block,
	}, nil
}

// GetBlock returns block details by number
func (s *BlockchainService) GetBlock(ctx context.Context, req *dtos.GetBlockRequest) (*dtos.GetBlockResponse, error) {
	// Validate request
	if err := s.validator.ValidateGetBlockRequest(req); err != nil {
		return nil, err
	}

	resp, err := s.client.GetBlockByNumber(ctx, req.BlockNumber, req.FullTx)
	if err != nil {
		return nil, fmt.Errorf("failed to get block: %w", err)
	}

	// Unmarshal the result into a generic interface
	var blockData interface{}
	if err := json.Unmarshal(resp.Result, &blockData); err != nil {
		return nil, fmt.Errorf("failed to parse block data: %w", err)
	}

	return &dtos.GetBlockResponse{
		Block: blockData,
	}, nil
}

// GetTransaction returns transaction details by hash
func (s *BlockchainService) GetTransaction(ctx context.Context, req *dtos.GetTransactionRequest) (*dtos.GetTransactionResponse, error) {
	// Validate request
	if err := s.validator.ValidateGetTransactionRequest(req); err != nil {
		return nil, err
	}

	resp, err := s.client.GetTransactionByHash(ctx, req.TxHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	// Unmarshal the result into a generic interface
	var txData interface{}
	if err := json.Unmarshal(resp.Result, &txData); err != nil {
		return nil, fmt.Errorf("failed to parse transaction data: %w", err)
	}

	return &dtos.GetTransactionResponse{
		Transaction: txData,
	}, nil
}

// CallContract calls a smart contract method (read-only)
func (s *BlockchainService) CallContract(ctx context.Context, req *dtos.CallContractRequest) (*dtos.CallContractResponse, error) {
	// Validate request
	if err := s.validator.ValidateCallContractRequest(req); err != nil {
		return nil, err
	}

	// Default to "latest" if block is not specified
	block := req.Block
	if block == "" {
		block = "latest"
	}

	result, err := s.client.CallContract(ctx, req.To, req.Data, block)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %w", err)
	}

	return &dtos.CallContractResponse{
		Result: result,
	}, nil
}

// EstimateGas estimates the gas required for a transaction
func (s *BlockchainService) EstimateGas(ctx context.Context, req *dtos.EstimateGasRequest) (*dtos.EstimateGasResponse, error) {
	// Validate request
	if err := s.validator.ValidateEstimateGasRequest(req); err != nil {
		return nil, err
	}

	gasEstimate, err := s.client.EstimateGas(ctx, req.From, req.To, req.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}

	return &dtos.EstimateGasResponse{
		GasEstimate: gasEstimate,
	}, nil
}

// SendRawTransaction broadcasts a signed transaction to the network
func (s *BlockchainService) SendRawTransaction(ctx context.Context, req *dtos.SendRawTransactionRequest) (*dtos.SendRawTransactionResponse, error) {
	// Validate request
	if err := s.validator.ValidateSendRawTransactionRequest(req); err != nil {
		return nil, err
	}

	txHash, err := s.client.SendRawTransaction(ctx, req.SignedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return &dtos.SendRawTransactionResponse{
		TxHash: txHash,
	}, nil
}

// SignAndSendTransaction signs a transaction on the server and sends it to the blockchain
func (s *BlockchainService) SignAndSendTransaction(ctx context.Context, req *dtos.SignAndSendTransactionRequest) (*dtos.SignAndSendTransactionResponse, error) {
	// Validate request
	if err := s.validator.ValidateSignAndSendTransactionRequest(req); err != nil {
		return nil, err
	}

	if s.signer == nil {
		return nil, fmt.Errorf("transaction signer is not configured (missing private key)")
	}

	// Convert DTO to signer request
	signerReq := &blockchain.SignTransactionRequest{
		To:       req.To,
		Value:    req.Value,
		Data:     req.Data,
		GasLimit: req.GasLimit,
		GasPrice: req.GasPrice,
		Nonce:    req.Nonce,
	}

	// Sign and send transaction
	txHash, err := s.signer.SignAndSendTransaction(ctx, signerReq)
	if err != nil {
		return nil, fmt.Errorf("failed to sign and send transaction: %w", err)
	}

	return &dtos.SignAndSendTransactionResponse{
		TxHash:      txHash,
		FromAddress: s.signer.GetAddress(),
	}, nil
}

// GenericRPCCall allows calling any JSON-RPC method directly
func (s *BlockchainService) GenericRPCCall(ctx context.Context, req *dtos.GenericRPCRequest) (*dtos.GenericRPCResponse, error) {
	// Validate request
	if err := s.validator.ValidateGenericRPCRequest(req); err != nil {
		return nil, err
	}

	resp, err := s.client.Call(ctx, req.Method, req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to execute RPC call: %w", err)
	}

	// Unmarshal the result into a generic interface
	var result interface{}
	if err := json.Unmarshal(resp.Result, &result); err != nil {
		return nil, fmt.Errorf("failed to parse RPC result: %w", err)
	}

	return &dtos.GenericRPCResponse{
		Result: result,
	}, nil
}
