package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"testing"
	"time"
)

func TestCosmosClient_TxWithdrawReward(t *testing.T) {
	txParams, err := msg_params.NewTxParams(
		privateKey,
		0, "", 0,
	)
	if err != nil {
		panic(err)
	}

	p := msg_params.TxWithdrawRewardParams{
		TxParams:   *txParams,
		ValAddress: valAddr,
	}

	resp, err := c.TxWithdrawReward(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(*resp)
}

func TestCosmosClient_TxGrantWithdrawReward(t *testing.T) {
	p := msg_params.TxGrantParams{
		TxParams:    *defaultTxParams,
		Grantee:     toAddr,
		ExpiredTime: time.Now().Add(60 * time.Minute),
	}

	resp, err := c.TxGrantWithdrawReward(p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("txHash: %v\n", resp.TxHash)

	time.Sleep(20 * time.Second)

	newTxParams, err := msg_params.NewTxParams(
		granteePrivateKey,
		0, "", 0,
	)
	if err != nil {
		panic(err)
	}

	newParams := msg_params.TxWithdrawRewardParams{
		TxParams:   *newTxParams,
		DelAddress: addr,
		ValAddress: valAddr,
	}

	tmpResp, err := c.TxWithdrawReward(newParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("txHash: %v\n", tmpResp.TxHash)
}
