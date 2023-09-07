package client

import (
	"context"
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

func TestCosmosClient_StreamBlocks(t *testing.T) {
	ret := make(chan interface{})

	ctx := context.Background()
	go c.StreamBlocks(ctx, ret, nil)
	for {
		select {
		case <-ctx.Done():
			panic(fmt.Sprintf("ctx.Done() with err: %v", ctx.Err()))
		case tmp := <-ret:
			jsb, _ := json.Marshal(tmp)
			fmt.Println(string(jsb))
		}
	}
}
