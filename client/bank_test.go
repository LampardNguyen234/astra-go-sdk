package client

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
	"testing"
)

func TestCosmosClient_Balance(t *testing.T) {
	ret, err := c.Balance(addr)
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(ret, "", "\t")
	fmt.Println(string(jsb))
}
