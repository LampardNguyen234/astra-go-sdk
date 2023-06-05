package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/pkg/errors"
	"math/big"
)

// TxSend performs a transfer transaction.
// If sender is not the owner, it performs a TxGrantExec transaction.
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
		return c.TxGrantExec(p.TxParams, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msg)
}

func (c *CosmosClient) TxGrantSend(p msg_params.TxGrantParams, maxSpent *big.Int) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	if maxSpent == nil {
		return nil, fmt.Errorf("empty Amount")
	}

	auth := bankTypes.NewSendAuthorization(sdk.NewCoins(sdk.NewCoin(common.BaseDenom, sdk.NewIntFromBigInt(maxSpent))))

	return c.TxGrantAuthorization(p, auth)
}
