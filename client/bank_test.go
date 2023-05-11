package client

import (
	"fmt"
	"testing"
)

func TestCosmosClient_Balance(t *testing.T) {
	ret, err := c.Balance(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}
