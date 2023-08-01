package client

import (
	"context"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"math/big"
	"time"
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

func (c *CosmosClient) StreamBlocks(ctx context.Context, blkChan chan interface{}, startBlk *big.Int) {
	var currentBlk *big.Int
	if startBlk != nil {
		currentBlk = new(big.Int).SetUint64(startBlk.Uint64())
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			head, err := c.LatestBlockHeight()
			if err != nil {
				time.Sleep(3 * time.Second)
				continue
			}
			if currentBlk == nil || currentBlk.Cmp(new(big.Int).SetUint64(0)) <= 0 {
				currentBlk = big.NewInt(head.Int64())
			}
			if head.Cmp(currentBlk) < 0 {
				time.Sleep(3 * time.Second)
				continue
			}

			block, err := c.GetBlockByHeight(currentBlk.Int64())
			if err != nil {
				time.Sleep(3 * time.Second)
				continue
			}
			blkChan <- block

			currentBlk = currentBlk.Add(currentBlk, big.NewInt(1))
		}
	}
}
