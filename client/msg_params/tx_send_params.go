package msg_params

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"math/big"
)

type TxSendRequestParams struct {
	TxParams

	// FromAddr is the address of the sender (either in hex format or cosmos format).
	// If empty, it will be retrieved from the private key.
	//
	// If FromAddr != Operator => GrantExecMsg.
	FromAddr string

	// ToAddr is the address of the recipient (either in hex format or cosmos format).
	ToAddr string

	// Amount is the amount of aASA to send.
	Amount *big.Int
}

func (p TxSendRequestParams) IsValid() (bool, error) {
	if _, err := p.TxParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}
	if p.Amount == nil {
		return false, fmt.Errorf("empty Amount")
	}
	if p.FromAddr != "" {
		if _, err := account.ParseCosmosAddress(p.FromAddr); err != nil {
			return false, errors.Wrapf(err, "invalid FromAddr")
		}
	}
	if _, err := account.ParseCosmosAddress(p.ToAddr); err != nil {
		return false, errors.Wrapf(err, "invalid ToAddr")
	}

	return true, nil
}

func (p TxSendRequestParams) From() sdk.AccAddress {
	if p.FromAddr != "" {
		return account.MustParseCosmosAddress(p.FromAddr)
	}

	return p.Operator()
}

func (p TxSendRequestParams) To() sdk.AccAddress {
	return account.MustParseCosmosAddress(p.ToAddr)
}

func (p TxSendRequestParams) SendAmount() sdk.Coins {
	return sdk.NewCoins(sdk.NewCoin(common.BaseDenom, sdk.NewIntFromBigInt(p.Amount)))
}
