package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func TestCosmosClient_AllDelegationRewards(t *testing.T) {
	resp, err := c.AllDelegationRewards(addr)
	if err != nil {
		panic(err)
	}

	for val, amt := range resp {
		fmt.Printf("%v: %v\n", val, amt.String())
	}
}

func TestCosmosClient_DelegationRewards(t *testing.T) {
	resp, err := c.DelegationRewards(addr, valAddr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.String())
}

func TestCosmosClient_ValidatorOutstandingRewards(t *testing.T) {
	resp, err := c.ValidatorOutstandingRewards(valAddr)
	if err != nil {
		panic(err)
	}
	respDec := common.ParseAmountToDec(sdk.NewCoin(common.BaseDenom, resp))
	tmp := respDec.MulInt64(25).QuoInt64(950)

	fmt.Println(tmp.String(), respDec.String())
}
