package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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

func TestCosmosClient_TxStakingGrant(t *testing.T) {
	p := msg_params.TxStakingGrantParams{
		TxGrantParams: msg_params.TxGrantParams{
			TxParams:    *defaultTxParams,
			Grantee:     toAddr,
			ExpiredTime: time.Now().Add(60 * time.Minute),
		},
	}

	resp, err := c.TxStakingGrant(p, stakingTypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE)
	if err != nil {
		panic(err)
	}
	fmt.Printf("txHash: %v\n", resp.TxHash)
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
