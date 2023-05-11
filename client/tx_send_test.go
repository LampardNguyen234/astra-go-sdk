package client

import (
	"fmt"
	msg_params2 "github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"math/big"
	"testing"
)

func TestCosmosClient_TxSend(t *testing.T) {
	txParams := &msg_params2.TxParams{
		PrivateKey: privateKey,
	}

	p := msg_params2.TxSendRequestParams{
		TxParams: *txParams,
		ToAddr:   toAddr,
		Amount:   new(big.Int).SetUint64(testAmt),
	}

	resp, err := c.TxSend(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(*resp)
}
