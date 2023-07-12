package client

import (
	"fmt"
	"testing"
)

func TestCosmosClient_TxByHash(t *testing.T) {
	resp, err := c.TxByHash("0D150E8CA571CF32B82583490067C5E9858072F84FFDA2AC74B244556730FCCB")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
