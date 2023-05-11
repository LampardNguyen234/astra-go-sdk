package cosmos

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/cosmos/msg_params"
	"math/big"
	"testing"
	"time"
)

func TestCosmosClient_TxDelegate(t *testing.T) {
	p := msg_params.TxDelegateParams{
		TxParams:   *defaultTxParams,
		ValAddress: valAddr,
		Amount:     new(big.Int).SetUint64(testAmt),
	}

	resp, err := c.TxDelegate(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(*resp)
}

func TestCosmosClient_AutoCompound(t *testing.T) {
	p := msg_params.TxWithdrawRewardParams{
		TxParams:   *defaultTxParams,
		ValAddress: valAddr,
	}

	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		resp, err := c.TxWithdrawReward(p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("txHash: %v\n", resp.TxHash)
	}
}
