package msg_params

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/pkg/errors"
	"math/big"
)

type TxCreateValidatorParams struct {
	TxParams
	stakingTypes.Description

	// ValAddress is the address of the validator to whom the DelAddress delegates tokens.
	ValAddress string

	// SelfDelegation is the amount (measured in aastra) for self delegation. (default: 0).
	SelfDelegationAmt *big.Int

	// MinSelfDelegationAmt is the minimum amount (measured in aastra) for self delegation. (default: 0).
	MinSelfDelegationAmt *big.Int

	// CommissionRate is the commission rate, minimum of 5% (default: 0.05).
	CommissionRate float64

	// MaxCommissionRate is the maximum commission rate which validator can ever charge (default: 0.2).
	MaxCommissionRate float64

	// MaxChangeCommissionRate is the maximum daily increase of the validator commission (default: 0.01).
	MaxChangeCommissionRate float64
}

func (p TxCreateValidatorParams) IsValid() (bool, error) {
	if _, err := p.TxParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}

	if p.ValAddress != "" {
		if _, err := account.ParseCosmosValidatorAddress(p.ValAddress); err != nil {
			return false, errors.Wrapf(err, "invalid ValAddress")
		}
	}

	if p.CommissionRate < 0 && p.CommissionRate > 1 {
		return false, fmt.Errorf("expect 0 <= NewCommissionRate <= 1, got %v", p.MaxCommissionRate)
	}
	if p.MaxCommissionRate < 0 && p.MaxCommissionRate > 1 {
		return false, fmt.Errorf("expect 0 <= MaxCommissionRate <= 1, got %v", p.MaxCommissionRate)
	}
	if p.MaxChangeCommissionRate < 0 && p.MaxChangeCommissionRate > 1 {
		return false, fmt.Errorf("expect 0 <= MaxChangeCommissionRate <= 1, got %v", p.MaxChangeCommissionRate)
	}

	return true, nil
}

func (p TxCreateValidatorParams) ValidatorAddress() sdk.ValAddress {
	if p.ValAddress != "" {
		return account.MustParseCosmosValidatorAddress(p.ValAddress)
	}

	return p.MustGetPrivateKey().ValAddress()
}

func (p TxCreateValidatorParams) PubKey() cryptoTypes.PubKey {
	return p.MustGetPrivateKey().ConsensusKey().CosmosPubKey()
}

func (p TxCreateValidatorParams) SelfDelegation() sdk.Coin {
	amt := p.SelfDelegationAmt
	if amt == nil {
		amt = big.NewInt(0)
	}
	return sdk.NewCoin(common.BaseDenom, sdk.NewIntFromBigInt(amt))
}

func (p TxCreateValidatorParams) Desc() stakingTypes.Description {
	return p.Description
}

func (p TxCreateValidatorParams) CommissionRates() stakingTypes.CommissionRates {
	rate := sdk.MustNewDecFromStr(fmt.Sprintf("%.4f", p.CommissionRate))
	if rate.IsZero() {
		rate = sdk.NewDecWithPrec(5, 2)
	}
	maxRate := sdk.MustNewDecFromStr(fmt.Sprintf("%.4f", p.MaxCommissionRate))
	if maxRate.IsZero() {
		maxRate = sdk.NewDecWithPrec(20, 2)
	}
	maxRateChange := sdk.MustNewDecFromStr(fmt.Sprintf("%.4f", p.MaxChangeCommissionRate))
	if maxRateChange.IsZero() {
		maxRateChange = sdk.NewDecWithPrec(1, 2)
	}

	return stakingTypes.CommissionRates{
		Rate:          rate,
		MaxRate:       maxRate,
		MaxChangeRate: maxRateChange,
	}
}

func (p TxCreateValidatorParams) MinSelfDelegation() sdk.Int {
	if p.MinSelfDelegationAmt == nil {
		return sdk.ZeroInt()
	}
	return sdk.NewIntFromBigInt(p.MinSelfDelegationAmt)
}

func (p TxCreateValidatorParams) IsGrantExec() bool {
	return p.ValAddress != "" && p.MustGetPrivateKey().ValAddress().String() != p.ValAddress
}
