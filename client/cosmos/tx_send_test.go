package cosmos

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/cosmos/msg_params"
	"math/big"
	"testing"
)

func TestCosmosClient_TxSend(t *testing.T) {
	txParams := &msg_params.TxParams{
		PrivateKey: privateKey,
	}

	p := &msg_params.TxSendRequestParams{
		TxParams: *txParams,
		ToAddr:   "astra1rlpfcs404zdrxhztwyh5jwuntelwypxk4w33jq",
		Amount:   new(big.Int).SetUint64(testAmt),
	}

	resp, err := c.TxSend(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(*resp)
}
