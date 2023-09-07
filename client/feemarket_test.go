package client

import (
	"fmt"
	"testing"
)

func TestCosmosClient_FeemarketParams(t *testing.T) {
	resp, err := c.FeemarketParams()
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
