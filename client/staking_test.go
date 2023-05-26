package client

import (
	"fmt"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"testing"
)

func TestStakingClient_DelegationDetail(t *testing.T) {
	ret, err := c.DelegationDetail("astra1re7nrzgtcnqc3j0duq3jrv94fwxfxfknyelfkg")
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}

func TestStakingClient_AllValidators(t *testing.T) {
	ret, err := c.AllValidators(stakingTypes.Bonded)
	if err != nil {
		panic(err)
	}

	fmt.Println("count:", len(ret))
	for _, val := range ret {
		fmt.Println(val)
	}
}
