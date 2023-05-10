package cosmos

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	"os"
)

var (
	c          *CosmosClient
	privateKey = ""
	addr       string

	testAmt = uint64(1000000000000000000)
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

	common.Init()
}
