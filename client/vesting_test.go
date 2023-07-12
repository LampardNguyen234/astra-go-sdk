package client

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
	"testing"
)

func TestCosmosClient_GetAvailableVestingBalance(t *testing.T) {
	resp, err := c.GetAvailableVestingBalance(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.String())
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
