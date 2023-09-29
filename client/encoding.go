package client

import sdk "github.com/cosmos/cosmos-sdk/types"

func (c *CosmosClient) DecodeTxBytes(data []byte) (sdk.Tx, error) {
	return c.TxConfig.TxDecoder()(data)
}
