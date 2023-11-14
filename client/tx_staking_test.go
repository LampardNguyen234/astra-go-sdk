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
	amt, _ := new(big.Int).SetString("30000000000000000000", 10)
	p := msg_params.TxDelegateParams{
		TxParams:   *defaultTxParams,
		ValAddress: valAddr,
		Amount:     amt,
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

func TestCosmosClient_DelegationDetail(t *testing.T) {
	resp, err := c.DelegationDetail(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func TestCosmosClient_TxCreateValidator(t *testing.T) {
	createParams := msg_params.TxCreateValidatorParams{
		TxParams: msg_params.TxParams{PrivateKey: ""},
		Description: stakingTypes.Description{
			Moniker: "SampleVal",
			Details: "Hello world",
		},
		SelfDelegationAmt:       big.NewInt(1000000000000000),
		MinSelfDelegationAmt:    big.NewInt(1),
		CommissionRate:          0.07,
		MaxCommissionRate:       0.5,
		MaxChangeCommissionRate: 0.01,
	}

	txRespond, err := c.TxCreateValidator(createParams)
	if err != nil {
		panic(err)
	}
	fmt.Println(txRespond.TxHash)
}

func TestCosmosClient_TxEditValidator(t *testing.T) {
	createParams := msg_params.TxEditValidatorParams{
		TxParams: msg_params.TxParams{PrivateKey: ""},
		Description: stakingTypes.Description{
			Moniker: "SampleVal",
			Details: "Hello world",
		},
		MinSelfDelegationAmt: big.NewInt(10000000000000000),
		NewCommissionRate:    0.07,
	}

	txRespond, err := c.TxEditValidator(createParams)
	if err != nil {
		panic(err)
	}
	fmt.Println(txRespond.TxHash)
}
