package client

import (
	"context"
	mintTypes "github.com/AstraProtocol/astra/v2/x/mint/types"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/pkg/errors"
	"sort"
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
	ValidatorName string  `json:"ValidatorName,omitempty"`
	Validator     string  `json:"Validator"`
	Delegator     string  `json:"Delegator"`
	Amount        sdk.Int `json:"Amount"`
	AmountDec     sdk.Dec `json:"AmountDec"`
}

func (c *CosmosClient) Delegation(delAddr, valAddr string) (DelegationDetail, error) {
	ret := DelegationDetail{}
	delegator, err := account.ParseCosmosAddress(delAddr)
	if err != nil {
		return ret, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}
	validator, err := account.ParseCosmosValidatorAddress(valAddr)
	if err != nil {
		return ret, errors.Wrapf(ErrInvalidValAddress, err.Error())
	}

	resp, err := c.staking.QueryClient.Delegation(c.ctx, &stakingTypes.QueryDelegationRequest{
		DelegatorAddr: delegator.String(),
		ValidatorAddr: validator.String(),
	})
	if err != nil {
		return ret, err
	}

	ret = DelegationDetail{
		Validator: validator.String(),
		Delegator: delegator.String(),
		Amount:    resp.DelegationResponse.Balance.Amount,
		AmountDec: common.ParseAmountToDec(resp.DelegationResponse.Balance),
	}

	return ret, nil
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
		resp, err := c.Delegation(delegator.String(), val.OperatorAddress)
		if err != nil {
			if strings.Contains(err.Error(), "not found for validator") {
				continue
			}
			return nil, err
		}
		resp.ValidatorName = val.GetMoniker()

		ret[val.OperatorAddress] = resp
	}

	return ret, nil
}

// ValidatorDelegations returns the staking amounts of the given address.
func (c *CosmosClient) ValidatorDelegations(valAddr string) ([]DelegationDetail, error) {
	validator, err := account.ParseCosmosValidatorAddress(valAddr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}
	ret := make([]DelegationDetail, 0)
	page := &query.PageRequest{
		Key:        nil,
		Offset:     0,
		CountTotal: true,
	}
	count := 0
	for {
		delegations, err := c.staking.QueryClient.ValidatorDelegations(c.ctx, &stakingTypes.QueryValidatorDelegationsRequest{
			ValidatorAddr: validator.String(),
			Pagination:    page,
		})
		if err != nil {
			return nil, err
		}
		for _, delegation := range delegations.DelegationResponses {
			ret = append(ret, DelegationDetail{
				Validator: valAddr,
				Delegator: delegation.Delegation.DelegatorAddress,
				Amount:    common.ParseAmount(delegation.Balance),
				AmountDec: common.ParseAmountToDec(delegation.Balance),
			})
		}
		count += len(delegations.DelegationResponses)
		if uint64(count) >= delegations.Pagination.Total {
			break
		}
		page.Offset = uint64(count)
	}
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i].Amount.GTE(ret[j].Amount)
	})

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

// StakingSupply returns the total number of tokens which are bonded.
func (c *CosmosClient) StakingSupply() (sdk.Int, error) {
	resp, err := c.staking.Pool(c.ctx, &stakingTypes.QueryPoolRequest{})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return resp.Pool.BondedTokens, nil
}

// GetBondedRatio returns the current staking ratio.
func (c *CosmosClient) GetBondedRatio() (sdk.Dec, error) {
	resp, err := c.mint.GetBondedRatio(c.ctx, &mintTypes.QueryBondedRatioRequest{})
	if err != nil {
		return sdk.ZeroDec(), err
	}

	return resp.BondedRatio, nil
}
