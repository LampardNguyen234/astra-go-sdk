package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
	"testing"
)

func TestCosmosClient_TxSend(t *testing.T) {
	txParams := &msg_params.TxParams{
		PrivateKey: privateKey,
	}

	p := msg_params.TxSendRequestParams{
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

func TestCosmosClient_TxGrantSend(t *testing.T) {
	txParams := *defaultTxParams
	p := msg_params.TxGrantParams{
		TxParams: txParams,
		Grantee:  toAddr,
	}

	tx, err := c.TxGrantSend(p,
		common.ParseAmount(sdk.NewCoin(common.Denom, sdk.NewInt(1000))).BigInt(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(tx.TxHash)
}
