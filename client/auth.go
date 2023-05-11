package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/cosmos/cosmos-sdk/types/query"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gogo/protobuf/grpc"
)

type AuthClient struct {
	authTypes.QueryClient
}

// NewAuthClient creates a new client for interacting with the `auth` module.
func NewAuthClient(conn grpc.ClientConn) *AuthClient {
	return &AuthClient{
		QueryClient: authTypes.NewQueryClient(conn),
	}
}

// AccountInfo returns a authTypes.AccountI for the given address.
func (c *CosmosClient) AccountInfo(addr string) (AccountInfoI, error) {
	resp, err := c.AuthClient.QueryClient.Account(c.ctx, &authTypes.QueryAccountRequest{Address: addr})
	if err != nil {
		return nil, err
	}

	if resp.Account == nil {
		return nil, fmt.Errorf("no account found")
	}

	var ret authTypes.AccountI
	err = c.Codec.UnpackAny(resp.Account, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// AccountExists checks if an account exists given its address.
func (c *CosmosClient) AccountExists(addr string) error {
	accAddr, err := account.NewCosmosAddressFromStr(addr)
	if err != nil {
		return err
	}

	err = c.BaseClient.AccountRetriever.EnsureExists(c.BaseClient.Context, accAddr)
	if err != nil {
		return err
	}

	return nil
}

// TotalAccounts returns the number of accounts on the blockchain.
func (c *CosmosClient) TotalAccounts() (uint64, error) {
	resp, err := c.AuthClient.Accounts(c.ctx, &authTypes.QueryAccountsRequest{
		Pagination: &query.PageRequest{
			Limit:      1,
			CountTotal: true,
		},
	})
	if err != nil {
		return 0, err
	}

	return resp.Pagination.Total, nil
}
