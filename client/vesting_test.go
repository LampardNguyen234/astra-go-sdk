package client

import (
	"fmt"
	"testing"
	"time"
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

	fmt.Println(resp.String())
	fmt.Println(resp.Locked.String())
}

func TestCosmosClient_GetVestingAccount(t *testing.T) {
	resp, err := c.GetVestingAccount(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.GetVestedOnly(time.Now()))
}

func TestCosmosClient_GetNextVestingPeriod(t *testing.T) {
	next, amt, err := c.GetNextVestingPeriod(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(next.String(), amt)
}
