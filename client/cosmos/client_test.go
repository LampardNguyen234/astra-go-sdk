package cosmos

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/client/cosmos/msg_params"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	"os"
)

var (
	c                 *CosmosClient
	privateKey        = ""
	granteePrivateKey = ""
	addr              string
	toAddr            string

	testAmt         = uint64(1000000000000000000)
	defaultTxParams *msg_params.TxParams
)

func init() {
	var err error
	c, err = NewCosmosClient(DefaultTestnetConfig())
	if err != nil {
		panic(err)
	}

	if privateKey == "" {
		privateKey = os.Getenv("SDK_KEY")
	}

	prvKey, err := account.NewPrivateKeyFromString(privateKey)
	if err != nil {
		panic(err)
	}
	addr = prvKey.AccAddress().String()

	granteePrv, err := account.NewPrivateKeyFromString(granteePrivateKey)
	if err != nil {
		panic(err)
	}
	toAddr = granteePrv.AccAddress().String()

	defaultTxParams, err = msg_params.NewTxParams(
		privateKey,
		0, "", 0,
	)
	if err != nil {
		panic(err)
	}

	common.Init()
}
