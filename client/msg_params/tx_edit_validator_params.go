package msg_params

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
	"math/big"
)

type TxEditValidatorParams struct {
	TxParams
	stakingTypes.Description

	// ValAddress is the address of the validator to whom the DelAddress delegates tokens.
	ValAddress string

	// MinSelfDelegationAmt is the minimum amount (measured in aastra) for self delegation. (default: 0).
	MinSelfDelegationAmt *big.Int

	// NewCommissionRate is the commission rate.
	NewCommissionRate float64
}

func (p TxEditValidatorParams) IsValid() (bool, error) {
	if _, err := p.TxParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}

	if p.ValAddress != "" {
		if _, err := account.ParseCosmosValidatorAddress(p.ValAddress); err != nil {
			return false, errors.Wrapf(err, "invalid ValAddress")
		}
	}

	if p.NewCommissionRate < 0 && p.NewCommissionRate > 1 {
		return false, fmt.Errorf("expect 0 <= NewCommissionRate <= 1, got %v", p.NewCommissionRate)
	}

	return true, nil
}

func (p TxEditValidatorParams) ValidatorAddress() sdk.ValAddress {
	if p.ValAddress != "" {
		return account.MustParseCosmosValidatorAddress(p.ValAddress)
	}

	return p.MustGetPrivateKey().ValAddress()
}

func (p TxEditValidatorParams) Desc() stakingTypes.Description {
	return p.Description
}

func (p TxEditValidatorParams) NewRate() *sdk.Dec {
	if p.NewCommissionRate == 0 {
		return nil
	}

	rate := sdk.MustNewDecFromStr(fmt.Sprintf("%.04f", p.NewCommissionRate))
	return &rate
}

func (p TxEditValidatorParams) MinSelfDelegation() *sdk.Int {
	if p.MinSelfDelegationAmt == nil {
		return nil
	}
	ret := sdk.NewIntFromBigInt(p.MinSelfDelegationAmt)
	return &ret
}

func (p TxEditValidatorParams) IsGrantExec() bool {
	return p.ValAddress != "" && p.MustGetPrivateKey().ValAddress().String() != p.ValAddress
}
