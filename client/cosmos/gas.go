package cosmos

import (
	clientTx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkTx "github.com/cosmos/cosmos-sdk/types/tx"
)

const (
	maxGasForEstimation = uint64(10000000)
)

// EstimateGas simulates the execution of a transaction and returns the
// simulation response obtained by the query and the adjusted gas amount.
func (c *CosmosClient) EstimateGas(tx Tx, msgs ...sdk.Msg) (uint64, error) {
	tx.txf = tx.txf.WithGas(maxGasForEstimation)
	txBuilder, err := tx.Build(msgs...)
	if err != nil {
		return 0, err
	}
	txBytes, err := c.EncodeTx(txBuilder)
	if err != nil {
		return 0, err
	}

	txSvcClient := sdkTx.NewServiceClient(c.BaseClient)
	simRes, err := txSvcClient.Simulate(c.ctx, &sdkTx.SimulateRequest{
		TxBytes: txBytes,
	})
	if err != nil {
		return 0, err
	}

	estimatedGas := uint64(tx.txf.GasAdjustment() * float64(simRes.GasInfo.GasUsed))
	tx.txf = tx.txf.WithGas(estimatedGas)
	return estimatedGas, nil
}

// BuildSimTx creates an unsigned tx with an empty single signature and returns
// the encoded transaction or an error if the unsigned transaction cannot be
// built.
func BuildSimTx(txf clientTx.Factory, msgs ...sdk.Msg) ([]byte, error) {
	return txf.BuildSimTx(msgs...)
}
