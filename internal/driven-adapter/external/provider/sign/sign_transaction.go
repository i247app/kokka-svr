package sign

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type SignTxParams struct {
	ChainID    string
	PrivateKey string
	Nonce      uint64
	To         string
	Value      *big.Int
	Data       []byte
	GasLimit   uint64
	GasPrice   *big.Int
}

// SignTx signs a transaction with the given parameters.
func SignTx(params SignTxParams) (*types.Transaction, error) {

	pkHex := strings.TrimPrefix(params.PrivateKey, "0x")
	privateKey, err := crypto.HexToECDSA(pkHex)
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(params.To)

	txData := &types.LegacyTx{
		Nonce:    params.Nonce,
		Gas:      params.GasLimit,
		GasPrice: params.GasPrice,
		To:       &toAddress,
		Value:    params.Value,
		Data:     params.Data,
	}

	tx := types.NewTx(txData)

	chain_id_number, err := hexutil.DecodeBig(params.ChainID)
	if err != nil {
		return nil, fmt.Errorf("invalid chain ID: %s", params.ChainID)
	}

	signer := types.NewEIP155Signer(chain_id_number)

	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}
