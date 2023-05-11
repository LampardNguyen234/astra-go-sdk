package msg_params

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type TxStakingGrantParams struct {
	TxGrantParams

	// AllowedList is a list of validator addresses that are allowed to interact with.
	AllowedList []string

	// DeniedList is a list of validator addresses that are NOT allowed to interact with.
	DeniedList []string

	// Amount is the amount
	Amount *sdk.Coin
}

func (p TxStakingGrantParams) IsValid() (bool, error) {
	if _, err := p.TxGrantParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}
	for _, addr := range p.AllowedList {
		if _, err := account.ParseCosmosValidatorAddress(addr); err != nil {
			return false, errors.Wrapf(err, addr)
		}
	}
	for _, addr := range p.DeniedList {
		if _, err := account.ParseCosmosValidatorAddress(addr); err != nil {
			return false, errors.Wrapf(err, addr)
		}
	}
	return true, nil
}

func (p TxStakingGrantParams) Allowed() []sdk.ValAddress {
	ret := make([]sdk.ValAddress, 0)
	for _, addr := range p.AllowedList {
		ret = append(ret, account.MustParseCosmosValidatorAddress(addr))
	}

	return ret
}

func (p TxStakingGrantParams) Denied() []sdk.ValAddress {
	ret := make([]sdk.ValAddress, 0)
	for _, addr := range p.DeniedList {
		ret = append(ret, account.MustParseCosmosValidatorAddress(addr))
	}

	return ret
}
