package msg_params

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"math/big"
)

type TxDelegateParams struct {
	TxParams

	// DelAddress is the address of the delegator.
	// If empty, it will be retrieved from the private key.
	DelAddress string

	// ValAddress is the address of the validator to whom the DelAddress delegates tokens.
	ValAddress string

	// Amount is the amount of aASA to delegate.
	Amount *big.Int
}

func (p TxDelegateParams) IsValid() (bool, error) {
	if _, err := p.TxParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}
	if p.DelAddress != "" {
		if _, err := account.ParseCosmosAddress(p.DelAddress); err != nil {
			return false, errors.Wrapf(err, "invalid DelAddress")
		}
	}
	if _, err := account.ParseCosmosValidatorAddress(p.ValAddress); err != nil {
		return false, errors.Wrapf(err, "invalid ValAddress")
	}
	if p.Amount == nil {
		return false, fmt.Errorf("empty Amount")
	}

	return true, nil
}

func (p TxDelegateParams) DelegatorAddress() sdk.AccAddress {
	if p.DelAddress == "" {
		return p.TxParams.From()
	}

	return account.MustParseCosmosAddress(p.DelAddress)
}

func (p TxDelegateParams) ValidatorAddress() sdk.ValAddress {
	return account.MustParseCosmosValidatorAddress(p.ValAddress)
}

func (p TxDelegateParams) DelegateAmount() sdk.Coin {
	return sdk.NewCoin(common.BaseDenom, sdk.NewIntFromBigInt(p.Amount))
}
