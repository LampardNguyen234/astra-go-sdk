package client

import (
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/pkg/errors"
)

func (c *CosmosClient) TxSend(p msg_params.TxSendRequestParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	from := p.From()
	to := p.To()

	// check sufficient balance
	balance, err := c.Balance(from.String())
	if err != nil {
		return nil, err
	}
	if balance.Unlocked.BigInt().Cmp(p.Amount) < 0 {
		return nil, errors.Wrapf(ErrInsufficientBalance, "expected at least %v, got %v", p.Amount.String(), balance.Unlocked.BigInt().String())
	}

	msg := bankTypes.NewMsgSend(from, to, p.SendAmount())
	if from.String() != p.Operator().String() {
		return c.txGrantExec(p.TxParams, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msg)
}
