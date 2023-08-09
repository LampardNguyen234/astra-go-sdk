package client

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
	"testing"
)

func TestCosmosClient_SigningInfos(t *testing.T) {
	infos, err := c.SigningInfos()
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(infos, "", "\t")
	fmt.Println(string(jsb))
}
