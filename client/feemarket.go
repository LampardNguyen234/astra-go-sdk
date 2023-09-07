package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
	"github.com/gogo/protobuf/grpc"
)

type FeemarketClient struct {
	feemarkettypes.QueryClient
}

// NewFeemarketClient creates a new client for interacting with the `feemarket` module.
func NewFeemarketClient(conn grpc.ClientConn) *FeemarketClient {
	return &FeemarketClient{
		QueryClient: feemarkettypes.NewQueryClient(conn),
	}
}

func (c *CosmosClient) FeemarketParams() (*feemarkettypes.Params, error) {
	resp, err := c.feemarket.Params(c.ctx, &feemarkettypes.QueryParamsRequest{})
	if err != nil {
		return nil, err
	}

	return &(resp.Params), nil
}

func (c *CosmosClient) BlockGas() (int64, error) {
	resp, err := c.feemarket.BlockGas(c.ctx, &feemarkettypes.QueryBlockGasRequest{})
	if err != nil {
		return 0, err
	}

	return resp.Gas, nil
}

func (c *CosmosClient) BaseFee() (sdk.Int, error) {
	resp, err := c.feemarket.BaseFee(c.ctx, &feemarkettypes.QueryBaseFeeRequest{})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	if resp.BaseFee == nil {
		return sdk.ZeroInt(), nil
	}

	return *resp.BaseFee, nil
}
