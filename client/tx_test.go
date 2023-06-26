package client

import (
	"fmt"
	"testing"
)

func TestCosmosClient_TxByHash(t *testing.T) {
	resp, err := c.TxByHash("")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
