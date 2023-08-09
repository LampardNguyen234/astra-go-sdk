package client

import (
	"fmt"
	"testing"
)

func TestCosmosClient_Proposals(t *testing.T) {
	resp, err := c.Proposals()
	if err != nil {
		panic(err)
	}

	for i, p := range resp {
		fmt.Println(i, p.FinalTallyResult)
	}
}
