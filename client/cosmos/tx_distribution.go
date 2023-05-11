package cosmos

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/cosmos/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrType "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// TxWithdrawReward creates a staking reward withdrawal transaction.
func (c *CosmosClient) TxWithdrawReward(p msg_params.TxWithdrawRewardParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	delegator := p.DelegatorAddress()
	from := p.Operator()

	var msg sdk.Msg
	msg = distrType.NewMsgWithdrawDelegatorReward(
		delegator, p.ValidatorAddress(),
	)
	if delegator.String() != from.String() { // grant execution
		return c.txGrantExec(p.TxParams, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msg)
}

// TxGrantWithdrawReward creates a staking reward withdrawal grant.
func (c *CosmosClient) TxGrantWithdrawReward(p msg_params.TxGrantParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	auth, err := msg_params.NewAuthorization(sdk.MsgTypeURL(&distrType.MsgWithdrawDelegatorReward{}), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build authorization: %v", err)
	}

	return c.txGrantAuthorization(p, auth)
}
