package client

import (
	slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/gogo/protobuf/grpc"
)

type SlashingClient struct {
	slashingTypes.QueryClient
}

// NewSlashingClient creates a new client for interacting with the `slashing` module.
func NewSlashingClient(conn grpc.ClientConn) *SlashingClient {
	return &SlashingClient{
		QueryClient: slashingTypes.NewQueryClient(conn),
	}
}

func (c *CosmosClient) SigningInfos() ([]slashingTypes.ValidatorSigningInfo, error) {
	resp, err := c.slashing.SigningInfos(c.ctx, &slashingTypes.QuerySigningInfosRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Info, nil
}
