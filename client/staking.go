package client

import (
	"context"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/pkg/errors"
	"strings"
)

type StakingClient struct {
	stakingTypes.QueryClient
	stakingTypes.MsgClient
}

// NewStakingClient creates a new client for interacting with the `staking` module.
func NewStakingClient(conn grpc.ClientConn) *StakingClient {
	return &StakingClient{
		QueryClient: stakingTypes.NewQueryClient(conn),
		MsgClient:   stakingTypes.NewMsgClient(conn),
	}
}

type DelegationDetail struct {
	Validator string  `json:"Validator"`
	Amount    sdk.Int `json:"Amount"`
}

// DelegationDetail returns the staking amounts of the given address.
func (c *CosmosClient) DelegationDetail(addr string) (map[string]DelegationDetail, error) {
	delegator, err := account.ParseCosmosAddress(addr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}

	ret := make(map[string]DelegationDetail)
	validators, err := c.AllValidators(stakingTypes.Bonded)
	if err != nil {
		return nil, err
	}

	for _, val := range validators {
		resp, err := c.staking.QueryClient.Delegation(c.ctx, &stakingTypes.QueryDelegationRequest{
			DelegatorAddr: delegator.String(),
			ValidatorAddr: val.OperatorAddress,
		})
		if err != nil {
			if strings.Contains(err.Error(), "not found for validator") {
				continue
			}
			return nil, err
		}

		ret[val.OperatorAddress] = DelegationDetail{
			Validator: val.OperatorAddress,
			Amount:    resp.DelegationResponse.Balance.Amount,
		}
	}

	return ret, nil
}

// AllValidators returns all validators matching with the given status.
func (c *CosmosClient) AllValidators(status stakingTypes.BondStatus) ([]stakingTypes.Validator, error) {
	ret, err := c.staking.QueryClient.Validators(context.Background(), &stakingTypes.QueryValidatorsRequest{
		Status: status.String(),
	})
	if err != nil {
		return nil, err
	}

	return ret.Validators, nil
}

// GetValidatorDetail returns the detail of a validator given its address.
func (c *CosmosClient) GetValidatorDetail(valAddress string) (*stakingTypes.Validator, error) {
	resp, err := c.staking.Validator(context.Background(), &stakingTypes.QueryValidatorRequest{ValidatorAddr: valAddress})
	if err != nil {
		return nil, err
	}

	return &resp.Validator, nil
}
