package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
)

// TxDelegate creates a delegation transaction.
func (c *CosmosClient) TxDelegate(p msg_params.TxDelegateParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	msg := stakingTypes.NewMsgDelegate(p.DelegatorAddress(), p.ValidatorAddress(), p.DelegateAmount())
	if p.IsGrantExec() {
		return c.TxGrantExec(p.TxParams, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msg)
}

func (c *CosmosClient) TxCreateValidator(p msg_params.TxCreateValidatorParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	msg, err := stakingTypes.NewMsgCreateValidator(
		p.ValidatorAddress(),
		p.PubKey(),
		p.SelfDelegation(),
		p.Desc(),
		p.CommissionRates(),
		p.MinSelfDelegation(),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create msg")
	}
	if p.IsGrantExec() {
		return c.TxGrantExec(p.TxParams, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msg)
}

func (c *CosmosClient) TxEditValidator(p msg_params.TxEditValidatorParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	msg := stakingTypes.NewMsgEditValidator(
		p.ValidatorAddress(),
		p.Desc(),
		p.NewRate(),
		p.MinSelfDelegation(),
	)
	if p.IsGrantExec() {
		return c.TxGrantExec(p.TxParams, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msg)
}

// TxStakingGrant creates a staking grant.
func (c *CosmosClient) TxStakingGrant(p msg_params.TxStakingGrantParams,
	authzType stakingTypes.AuthorizationType,
) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	if len(p.AllowedList) == 0 && len(p.DeniedList) == 0 {
		validators, err := c.AllValidators(stakingTypes.Bonded)
		if err != nil {
			return nil, err
		}
		for _, val := range validators {
			p.AllowedList = append(p.AllowedList, val.OperatorAddress)
		}
	}

	auth, err := stakingTypes.NewStakeAuthorization(
		p.Allowed(),
		p.Denied(),
		authzType,
		p.Amount,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build authorization: %v", err)
	}

	return c.TxGrantAuthorization(p.TxGrantParams, auth)
}
