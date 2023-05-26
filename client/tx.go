package client

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"math/big"
	"time"
)

func (c *CosmosClient) TxByHash(hash string) (*sdk.TxResponse, error) {
	resp, err := tx.QueryTx(c.BaseClient.Context, hash)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// BlockTxsByHeight retrieves the receipts of all transaction in a block given its height.
func (c *CosmosClient) BlockTxsByHeight(ctx context.Context, blk *big.Int) ([]*sdk.TxResponse, error) {
	height := blk.Int64()
	block, err := c.Client.Block(ctx, &height)
	if err != nil {
		return nil, err
	}

	res := make([]*sdk.TxResponse, 0)
	for _, tx := range block.Block.Txs {
		receipt, err := c.TxByHash(fmt.Sprintf("%X", tx.Hash()))
		if err != nil {
			return nil, err
		}

		res = append(res, receipt)
	}

	return res, nil
}

func (c *CosmosClient) ListenToTxs(ctx context.Context, txResult chan interface{}, startBlk *big.Int) {
	var currentBlk *big.Int
	if startBlk != nil {
		currentBlk = new(big.Int).SetUint64(startBlk.Uint64())
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			head, err := c.LatestBlockHeight(ctx)
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

			txs, err := c.BlockTxsByHeight(ctx, currentBlk)
			if err != nil {
				continue
			}
			for _, tmpTx := range txs {
				txResult <- tmpTx
			}

			currentBlk = currentBlk.Add(currentBlk, big.NewInt(1))
		}
	}
}

