package msg_params

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type TxWithdrawRewardParams struct {
	TxParams

	// DelAddress is the address of the delegator.
	//
	// If empty, it will be retrieved from the private key.
	//
	// If DelAddress != Operator => GrantExecMsg.
	DelAddress string

	// ValAddress is the address of the validator to whom the DelAddress delegates tokens.
	ValAddress string
}

func (p TxWithdrawRewardParams) IsValid() (bool, error) {
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

	return true, nil
}

func (p TxWithdrawRewardParams) DelegatorAddress() sdk.AccAddress {
	if p.DelAddress == "" {
		return p.TxParams.Operator()
	}

	return account.MustParseCosmosAddress(p.DelAddress)
}

func (p TxWithdrawRewardParams) ValidatorAddress() sdk.ValAddress {
	return account.MustParseCosmosValidatorAddress(p.ValAddress)
}
