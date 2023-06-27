package client

import (
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingTypes "github.com/evmos/evmos/v6/x/vesting/types"
	"github.com/pkg/errors"
)

// TxCreateVesting creates a new-vesting-account transaction.
// If sender is not the owner, it performs a TxGrantExec transaction.
func (c *CosmosClient) TxCreateVesting(p msg_params.TxCreateVestingParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	from := p.Funder()
	to := p.VestingAddress()

	// check sufficient balance
	balance, err := c.Balance(from.String())
	if err != nil {
		return nil, err
	}
	if balance.Unlocked.BigInt().Cmp(p.Amount) < 0 {
		return nil, errors.Wrapf(ErrInsufficientBalance, "expected at least %v, got %v", p.Amount.String(), balance.Unlocked.BigInt().String())
	}

	msg := vestingTypes.NewMsgCreateClawbackVestingAccount(
		from, to, p.StartTime(), p.LockupPeriods(), p.VestingPeriods(),
	)
	err = msg.ValidateBasic()
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidParameters, err.Error())
	}
	if from.String() != p.Operator().String() {
		return c.TxGrantExec(p.TxParams, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msg)
}

func (c *CosmosClient) TxClawBackVesting(p msg_params.TxClawBackVestingParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	msg := vestingTypes.NewMsgClawback(p.Funder(), p.VestingAddress(), p.To())

	err := msg.ValidateBasic()
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidParameters, err.Error())
	}
	if p.Funder().String() != p.Operator().String() {
		return c.TxGrantExec(p.TxParams, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msg)
}
