package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/gogo/protobuf/grpc"
	"github.com/pkg/errors"
)

type AuthzClient struct {
	authz.QueryClient
}

// NewAuthzClient creates a new client for interacting with the `authz` module.
func NewAuthzClient(conn grpc.ClientConn) *AuthzClient {
	return &AuthzClient{
		QueryClient: authz.NewQueryClient(conn),
	}
}

// Grants returns a list of `Authorization`, granted to the grantee by the granter.
// If msgTypeURL is provided, it will only return the grants matching with the msgTypeURL.
func (c *CosmosClient) Grants(granterStr, granteeStr, msgTypeURL string) ([]*Grant, error) {
	granter, err := account.ParseCosmosAddress(granterStr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, fmt.Sprintf("%v: %v", granterStr, err.Error()))
	}
	grantee, err := account.ParseCosmosAddress(granteeStr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, fmt.Sprintf("%v: %v", granteeStr, err.Error()))
	}

	resp, err := c.authz.Grants(c.ctx, &authz.QueryGrantsRequest{
		Granter:    granter.String(),
		Grantee:    grantee.String(),
		MsgTypeUrl: msgTypeURL,
	})
	if err != nil {
		return nil, err
	}

	return newGrants(c.Codec, resp.Grants...)
}

// GranterGrants returns a list of `Authorization` granted by the granter.
func (c *CosmosClient) GranterGrants(granterStr string) ([]*GrantAuthorization, error) {
	granter, err := account.ParseCosmosAddress(granterStr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, fmt.Sprintf("%v: %v", granterStr, err.Error()))
	}

	resp, err := c.authz.GranterGrants(c.ctx, &authz.QueryGranterGrantsRequest{
		Granter: granter.String(),
	})
	if err != nil {
		return nil, err
	}

	return newGrantAuthorizations(c.Codec, resp.Grants...)
}

// GranteeGrants returns a list of `Authorization` granted to the grantee.
func (c *CosmosClient) GranteeGrants(granteeStr string) ([]*GrantAuthorization, error) {
	grantee, err := account.ParseCosmosAddress(granteeStr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, fmt.Sprintf("%v: %v", granteeStr, err.Error()))
	}

	resp, err := c.authz.GranteeGrants(c.ctx, &authz.QueryGranteeGrantsRequest{
		Grantee: grantee.String(),
	})
	if err != nil {
		return nil, err
	}

	return newGrantAuthorizations(c.Codec, resp.Grants...)
}
