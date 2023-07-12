package msg_params

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type TxWithdrawCommissionParams struct {
	TxParams

	// ValAddress is the address of the validator to whom the DelAddress delegates tokens.
	ValAddress string
}

func (p TxWithdrawCommissionParams) IsValid() (bool, error) {
	if _, err := p.TxParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}

	if p.ValAddress != "" {
		if _, err := account.ParseCosmosValidatorAddress(p.ValAddress); err != nil {
			return false, errors.Wrapf(err, "invalid ValAddress")
		}
	}

	return true, nil
}

func (p TxWithdrawCommissionParams) ValidatorAddress() sdk.ValAddress {
	if p.ValAddress != "" {
		return account.MustParseCosmosValidatorAddress(p.ValAddress)
	}
	return p.TxParams.MustGetPrivateKey().ValAddress()
}

func (p TxWithdrawCommissionParams) IsGrantExec() bool {
	return p.ValAddress != "" && p.TxParams.MustGetPrivateKey().ValAddress().String() != p.ValAddress
}
