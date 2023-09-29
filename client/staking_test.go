package client

import (
	"fmt"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/json"
	"testing"
)

func TestStakingClient_DelegationDetail(t *testing.T) {
	ret, err := c.DelegationDetail(addr)
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(ret, "", "\t")
	fmt.Println(string(jsb))
}

func TestStakingClient_AllValidators(t *testing.T) {
	ret, err := c.AllValidators(stakingTypes.Bonded)
	if err != nil {
		panic(err)
	}

	fmt.Println("count:", len(ret))
	for _, val := range ret {
		fmt.Println(val.Description)
	}
}

func TestCosmosClient_ValidatorDelegations(t *testing.T) {
	resp, err := c.ValidatorDelegations("astravaloper1k4dtdgfqulhcq6t0kxvlhqlkl396xfcuwvp2xn")
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("total:", len(resp))
	fmt.Println(string(jsb))
}
