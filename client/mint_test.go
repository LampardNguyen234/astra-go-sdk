package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/json"
	"testing"
)

func TestCosmosClient_MintParams(t *testing.T) {
	p, err := c.MintParams()
	assert.NoError(t, err)

	fmt.Println(p)
}

func TestCosmosClient_Inflation(t *testing.T) {
	ret, err := c.Inflation()
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestCosmosClient_GetFoundationBalance(t *testing.T) {
	b, err := c.GetFoundationBalance()
	assert.NoError(t, err)

	jsb, _ := json.Marshal(b)
	fmt.Println(string(jsb))
}

func TestCosmosClient_MintInfo(t *testing.T) {
	resp, err := c.MintInfo()
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(jsb))
}
