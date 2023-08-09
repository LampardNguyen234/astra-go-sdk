package client

import (
	feeburnTypes "github.com/AstraProtocol/astra/v2/x/feeburn/types"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/grpc"
)

type FeeburnClient struct {
	feeburnTypes.QueryClient
}

// NewFeeburnClient creates a new client for interacting with the `feeburn` module.
func NewFeeburnClient(conn grpc.ClientConn) *FeeburnClient {
	return &FeeburnClient{
		QueryClient: feeburnTypes.NewQueryClient(conn),
	}
}

func (c *CosmosClient) TotalFeeBurn() (sdk.Int, error) {
	resp, err := c.feeburn.TotalFeeBurn(c.ctx, &feeburnTypes.QueryTotalFeeBurnRequest{})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return common.ParseAmount(resp.TotalFeeBurn), nil
}
