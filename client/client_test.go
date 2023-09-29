package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"os"
)

var (
	c                 *CosmosClient
	privateKey        = "58814f8e23b9ed9c0079c33610e4db38c56f38f5915ed29bb9346d94f413acb4"
	granteePrivateKey = "7da1d31c1a25cb038d50b5a38c899cc95c4ca6f49455035a171afa8fd9a25793"
	addr              string
	toAddr            string
	valAddr           string

	proposer = "f7eb361f93d0238b7d6c14bb6cc736e7d3ae91f11c269d6cb7031a2122df0b35"
	voters   = []string{
		"A541136491FA9CA0318ADAE1D0438CE4C19AD545695822794656E39FA7CA5D5E",
		"7B51BC562200978679650C3B538F53BACF8B004C4B974D455390AA8208593DAC",
	}

	testAmt         = uint64(1000000000000000000)
	defaultTxParams *msg_params.TxParams
)

func init() {
	cfg := DefaultMainnetConfig()
	var err error
	c, err = NewCosmosClient(cfg)
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

	granterAddr = addr
	granteeAddr = toAddr

	allValidators, err := c.AllValidators(stakingTypes.Bonded)
	if err != nil {
		panic(err)
	}
	if len(allValidators) != 0 {
		valAddr = allValidators[0].OperatorAddress
	}

	fmt.Println(addr, toAddr, valAddr)

	common.Init()
}
