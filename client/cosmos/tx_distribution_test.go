package cosmos

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/cosmos/msg_params"
	"math/big"
	"testing"
)

func TestCosmosClient_TxWithdrawReward(t *testing.T) {
	txParams, err := msg_params.NewTxParams(
		privateKey,
		0, "", 0,
	)
	if err != nil {
		panic(err)
	}

	p := &msg_params.TxWithdrawRewardParams{
		TxParams:   *txParams,
		ValAddress: valAddr,
	}

	resp, err := c.TxWithdrawReward(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(*resp)
}

func TestCosmosClient_TxDelegate(t *testing.T) {
	txParams, err := msg_params.NewTxParams(
		privateKey,
		0, "", 0,
	)
	if err != nil {
		panic(err)
	}

	p := &msg_params.TxDelegateParams{
		TxParams:   *txParams,
		ValAddress: valAddr,
		Amount:     new(big.Int).SetUint64(testAmt),
	}

	resp, err := c.TxDelegate(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(*resp)
}
