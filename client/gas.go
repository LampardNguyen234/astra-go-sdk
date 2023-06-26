package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkTx "github.com/cosmos/cosmos-sdk/types/tx"
)

const (
	maxGasForEstimation = uint64(1000000)
)

// EstimateGas simulates the execution of a transaction and returns the
// simulation response obtained by the query and the adjusted gas amount.
func (c *CosmosClient) EstimateGas(tx Tx, msgs ...sdk.Msg) (uint64, error) {
	if tx.params.GasLimit != 0 {
		tx.txf = tx.txf.WithGas(tx.params.GasLimit)
	} else {
		tx.txf = tx.txf.WithGas(maxGasForEstimation)
	}

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
