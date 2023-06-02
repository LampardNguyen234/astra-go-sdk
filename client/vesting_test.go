package client

import (
	"fmt"
	"testing"
)

func TestCosmosClient_GetVestingBalance(t *testing.T) {
	resp, err := c.GetVestingBalance(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.String())
	fmt.Println(resp.Locked.String())
}
