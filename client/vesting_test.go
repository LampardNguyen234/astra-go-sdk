package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	"github.com/tendermint/tendermint/libs/json"
	"testing"
)

func TestCosmosClient_GetAvailableVestingBalance(t *testing.T) {
	resp, err := c.GetAvailableVestingBalance("0x1ADcaa8b10C781653c87C9e3b6962d14c0A8Adbe")
	if err != nil {
		panic(err)
	}

	fmt.Println(common.ParseAmountToDec(resp))
}

func TestCosmosClient_GetVestingBalance(t *testing.T) {
	resp, err := c.GetVestingBalance(addr)
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(jsb))
}

func TestCosmosClient_GetVestingAccount(t *testing.T) {
	resp, err := c.GetVestingAccount(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.LockupPeriods)
}

func TestCosmosClient_GetNextVestingPeriod(t *testing.T) {
	next, amt, err := c.GetNextVestingPeriod(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(next.String(), amt)
}
