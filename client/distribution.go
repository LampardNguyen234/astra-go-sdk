package client

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrType "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/pkg/errors"
)

type DistrClient struct {
	distrType.QueryClient
}

// NewDistrClient creates a new client for interacting with the `distribution` module.
func NewDistrClient(conn grpc.ClientConn) *DistrClient {
	return &DistrClient{
		QueryClient: distrType.NewQueryClient(conn),
	}
}

// AllDelegationRewards returns the delegating reward of a delegator for all validators.
func (c *CosmosClient) AllDelegationRewards(delAddr string) (map[string]sdk.Int, error) {
	delegator, err := account.ParseCosmosAddress(delAddr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}

	ret := make(map[string]sdk.Int)
	delDetail, err := c.DelegationDetail(delegator.String())
	if err != nil {
		return nil, err
	}

	for val := range delDetail {
		tmpAmt, err := c.DelegationRewards(delegator.String(), val)
		if err != nil {
			return nil, err
		}

		ret[val] = tmpAmt
	}

	return ret, nil
}

// DelegationRewards returns the delegating reward of a delegator w.r.t to a validator address.
func (c *CosmosClient) DelegationRewards(delAddr, valAddr string) (sdk.Int, error) {
	delegator, err := account.ParseCosmosAddress(delAddr)
	if err != nil {
		return sdk.ZeroInt(), errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}

	_, err = account.ParseCosmosValidatorAddress(valAddr)
	if err != nil {
		return sdk.ZeroInt(), errors.Wrapf(ErrInvalidValAddress, err.Error())
	}

	resp, err := c.distr.DelegationRewards(c.ctx, &distrType.QueryDelegationRewardsRequest{
		DelegatorAddress: delegator.String(),
		ValidatorAddress: valAddr,
	})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return common.ParseAmount(resp.Rewards), nil
}

// GetCommunityPoolBalance returns the balance of the community pool
func (c *CosmosClient) GetCommunityPoolBalance() (sdk.Int, error) {
	resp, err := c.distr.CommunityPool(c.ctx, &distrType.QueryCommunityPoolRequest{})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return common.ParseAmount(resp.Pool), nil
}

func (c *CosmosClient) ValidatorOutstandingRewards(valAddr string) (sdk.Int, error) {
	if _, err := account.ParseCosmosValidatorAddress(valAddr); err != nil {
		return sdk.ZeroInt(), errors.Wrapf(ErrInvalidValAddress, err.Error())
	}

	resp, err := c.distr.ValidatorOutstandingRewards(c.ctx, &distrType.QueryValidatorOutstandingRewardsRequest{ValidatorAddress: valAddr})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return common.ParseAmount(resp.Rewards.Rewards), nil
}
