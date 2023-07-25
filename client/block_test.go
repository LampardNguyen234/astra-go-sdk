package client

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/json"
	"testing"
)

func TestCosmosClient_GetBlockByHeight(t *testing.T) {
	blk, err := c.GetBlockByHeight(4578909)
	if err != nil {
		panic(err)
	}

	jsb, _ := json.MarshalIndent(blk.Block.LastCommit, "", "\t")
	fmt.Println(string(jsb))
}
