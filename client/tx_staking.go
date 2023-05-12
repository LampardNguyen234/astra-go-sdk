package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// TxDelegate creates a delegation transaction.
func (c *CosmosClient) TxDelegate(p msg_params.TxDelegateParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	delegator := p.DelegatorAddress()
	msg := stakingTypes.NewMsgDelegate(p.DelegatorAddress(), p.ValidatorAddress(), p.DelegateAmount())
	if delegator.String() != p.Operator().String() { // grant execution
		return c.txGrantExec(p.TxParams, msg)
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

	auth, err := stakingTypes.NewStakeAuthorization(
		p.Allowed(),
		p.Denied(),
		authzType,
		p.Amount,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build authorization: %v", err)
	}

	return c.txGrantAuthorization(p.TxGrantParams, auth)
}
