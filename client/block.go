package client

import (
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"math/big"
)

// LatestBlockHeight returns the latest block height from the current chain.
func (c *CosmosClient) LatestBlockHeight() (*big.Int, error) {
	ret, err := c.Client.Block(c.ctx, nil)
	if err != nil {
		return nil, err
	}

	return new(big.Int).SetUint64(uint64(ret.Block.Height)), nil
}

func (c *CosmosClient) GetBlockByHeight(height int64) (*coretypes.ResultBlock, error) {
	return c.Client.Block(c.ctx, &height)
}
