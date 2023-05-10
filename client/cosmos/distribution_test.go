package cosmos

import (
	"fmt"
	"testing"
)

var (
	valAddr = "astravaloper1u7gf4z49v53yrxy6ggrzhxfqj46c3ap4tzku46"
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
