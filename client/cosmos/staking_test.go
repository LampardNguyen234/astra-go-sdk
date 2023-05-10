package cosmos

import (
	"fmt"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"testing"
)

func TestStakingClient_DelegationDetail(t *testing.T) {
	ret, err := c.DelegationDetail("astra1xdmyqh8td7663wu9sz3yult2f6dr3q028t8u3x")
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
