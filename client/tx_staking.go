package client

import (
	"fmt"
	msg_params2 "github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// TxDelegate creates a delegation transaction.
func (c *CosmosClient) TxDelegate(p msg_params2.TxDelegateParams) (*sdk.TxResponse, error) {
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

// TxGrantDelegate creates a staking delegation grant.
func (c *CosmosClient) TxGrantDelegate(p msg_params2.TxStakingGrantParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	auth, err := stakingTypes.NewStakeAuthorization(
		p.Allowed(),
		p.Denied(),
		stakingTypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE,
		p.Amount,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build authorization: %v", err)
	}

	return c.txGrantAuthorization(p.TxGrantParams, auth)
}
