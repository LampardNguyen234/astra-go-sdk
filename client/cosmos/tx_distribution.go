package cosmos

import (
	"github.com/LampardNguyen234/astra-go-sdk/client/cosmos/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrType "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// TxWithdrawReward creates a staking reward withdrawal transaction.
func (c *CosmosClient) TxWithdrawReward(p *msg_params.TxWithdrawRewardParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	msg := distrType.NewMsgWithdrawDelegatorReward(
		p.DelegatorAddress(), p.ValidatorAddress(),
	)

	return c.BuildAndSendTx(p.TxParams, msg)
}

// TxDelegate creates a delegation transaction.
func (c *CosmosClient) TxDelegate(p *msg_params.TxDelegateParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	msg := stakingTypes.NewMsgDelegate(p.DelegatorAddress(), p.ValidatorAddress(), p.DelegateAmount())

	return c.BuildAndSendTx(p.TxParams, msg)
}
