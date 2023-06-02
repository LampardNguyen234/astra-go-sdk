package client

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	vestingTypes "github.com/evmos/evmos/v6/x/vesting/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/pkg/errors"
)

type VestingClient struct {
	vestingTypes.QueryClient
}

// NewVestingClient creates a new client for interacting with the `vesting` module.
func NewVestingClient(conn grpc.ClientConn) *VestingClient {
	return &VestingClient{
		QueryClient: vestingTypes.NewQueryClient(conn),
	}
}

// GetVestingBalance returns the detail balance of a vesting account.
func (c *CosmosClient) GetVestingBalance(strAddr string) (*vestingTypes.QueryBalancesResponse, error) {
	addr, err := account.ParseCosmosAddress(strAddr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}

	resp, err := c.vesting.Balances(c.ctx, &vestingTypes.QueryBalancesRequest{Address: addr.String()})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
