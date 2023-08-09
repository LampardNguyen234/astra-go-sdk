package client

import (
	"fmt"
	"math/big"
	"testing"
)

func TestCosmosClient_TxByHash(t *testing.T) {
	resp, err := c.TxByHash("1B22B59E64B7F4BC36EA1D4291189828E3922A07264537E27E4C587A71B39C76")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

func TestCosmosClient_BlockTxsByHeight(t *testing.T) {
	resp, err := c.BlockTxsByHeight(c.ctx, big.NewInt(627))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
