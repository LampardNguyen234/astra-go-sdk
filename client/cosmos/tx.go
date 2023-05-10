package cosmos

import (
	"github.com/LampardNguyen234/astra-go-sdk/client/cosmos/msg_params"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type Tx struct {
	txf         tx.Factory
	txConfig    client.TxConfig
	chainID     string
	params      msg_params.TxParams
	senderCheck TxSenderAccountI
}

// NewTx creates a new Tx from the given parameters.
func (c *CosmosClient) NewTx(txParams msg_params.TxParams) *Tx {
	gasPrice := txParams.GasPrice
	if gasPrice == "" {
		gasPrice = msg_params.DefaultTxParams().GasPrice
	}
	gasLimit := txParams.GasLimit
	if gasLimit == 0 {
		gasLimit = msg_params.DefaultTxParams().GasLimit
	}
	gasAdjustment := txParams.GasAdjustment
	if gasAdjustment == 0 {
		gasAdjustment = msg_params.DefaultTxParams().GasAdjustment
	}

	txf := tx.Factory{}.
		WithChainID(c.BaseClient.ChainID).
		WithTxConfig(c.BaseClient.TxConfig).
		WithGasAdjustment(gasAdjustment).
		WithGasPrices(gasPrice).
		WithGas(gasLimit).
		WithSignMode(c.BaseClient.TxConfig.SignModeHandler().DefaultMode())

	return &Tx{
		txf:         txf,
		params:      txParams,
		txConfig:    c.BaseClient.TxConfig,
		chainID:     c.BaseClient.ChainID,
		senderCheck: c,
	}
}

// BuildAndSendTx creates and sends a transaction with given parameters.
func (c *CosmosClient) BuildAndSendTx(txParams msg_params.TxParams, msgs ...types.Msg) (*types.TxResponse, error) {
	if _, err := txParams.IsValid(); err != nil {
		return nil, errors.Wrapf(err, "invalid txParams")
	}
	tmpTx := c.NewTx(txParams)
	if tmpTx.txf.Gas() == 0 {
		estimatedGas, err := c.EstimateGas(*tmpTx, msgs...)
		if err != nil {
			return nil, err
		}

		tmpTx.txf = tmpTx.txf.WithGas(estimatedGas)
	}

	txBuilder, err := tmpTx.Build(msgs...)
	if err != nil {
		return nil, errors.Wrapf(ErrBuildTx, err.Error())
	}

	return c.buildAndBroadcastTx(txBuilder)
}

// EncodeTx generates the raw transaction given a client.TxBuilder.
func (c *CosmosClient) EncodeTx(txBuilder client.TxBuilder) ([]byte, error) {
	rawTx, err := c.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, errors.Wrapf(ErrBuildTx, err.Error())
	}

	return rawTx, nil
}

// buildAndBroadcastTx performs EncodeTx and then BroadcastTx.
func (c *CosmosClient) buildAndBroadcastTx(txBuilder client.TxBuilder) (*types.TxResponse, error) {
	rawTx, err := c.EncodeTx(txBuilder)
	if err != nil {
		return nil, err
	}

	resp, err := c.BroadcastTx(rawTx)
	if resp.Code != 0 {
		err = errors.Wrapf(ErrBroadcastTx, resp.RawLog)
	}

	return resp, err
}
