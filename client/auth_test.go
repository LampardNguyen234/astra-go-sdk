package client

import (
	"fmt"
	distrType "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCosmosClient_AccountInfo(t *testing.T) {
	resp, err := c.AccountInfo(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.GetAccountNumber(), resp.GetSequence())
}

func TestCosmosClient_CountAccounts(t *testing.T) {
	resp, err := c.TotalAccounts()
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func TestCosmosClient_GetModuleAccount(t *testing.T) {
	resp, err := c.GetModuleAccount(distrType.ModuleName)
	assert.NoError(t, err)

	fmt.Println(resp.GetAddress())
}
