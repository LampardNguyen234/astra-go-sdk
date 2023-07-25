package msg_params

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/pkg/errors"
	"math/big"
	"time"
)

type TxCreateVestingParams struct {
	TxParams `json:"TxParams"`

	// FunderAddr is the address of the funder.
	//
	// If empty, it will be retrieved from the private key.
	//
	// If FunderAddr != Operator => GrantExecMsg.
	FunderAddr string `json:"FunderAddr,omitempty"`

	// ToAddr is the vesting recipient.
	ToAddr string `json:"ToAddr"`

	// Start is the starting timestamp of the vesting.
	// If not given, startTime will be set to time.Now().
	Start time.Time `json:"Start,omitempty"`

	// VestingDuration is the vesting duration (in terms of seconds).
	VestingDuration int64 `json:"VestingDuration,omitempty"`

	// Amount is the amount of aASA to delegate.
	Amount *big.Int `json:"Amount"`

	// Vesting is the periods to vest.
	Vesting types.Periods `json:"Vesting,omitempty"`

	// VestingLength is the number of periods to vest.
	// If used, each period will have the same amount to vest.
	// This is unused if Vesting is provided.
	VestingLength uint `json:"VestingLength,omitempty"`

	// Lockup is the periods to lock.
	Lockup types.Periods `json:"Lockup,omitempty"`
}

func (p *TxCreateVestingParams) IsValid() (bool, error) {
	if _, err := p.TxParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}
	if p.Amount == nil {
		return false, fmt.Errorf("empty Amount")
	}
	if p.FunderAddr != "" {
		if _, err := account.ParseCosmosAddress(p.FunderAddr); err != nil {
			return false, errors.Wrapf(err, "invalid FunderAddr")
		}
	}
	if _, err := account.ParseCosmosAddress(p.ToAddr); err != nil {
		return false, errors.Wrapf(err, "invalid ToAddr")
	}
	if len(p.Vesting) == 0 && p.VestingLength == 0 {
		return false, fmt.Errorf("require either `Vesting` or `VestingLength`")
	}
	if len(p.Vesting) == 0 && p.VestingDuration == 0 {
		return false, fmt.Errorf("require either `Vesting` or `VestingDuration`")
	}
	if p.VestingDuration != 0 && p.VestingDuration < int64(p.VestingLength) {
		return false, fmt.Errorf("require `VestingLength <= VestingDuration`")
	}

	return true, nil
}

func (p *TxCreateVestingParams) StartTime() time.Time {
	if p.Start.IsZero() {
		p.Start = time.Now()
	}

	return p.Start
}

func (p *TxCreateVestingParams) VestingPeriods() types.Periods {
	if len(p.Vesting) > 0 {
		return p.Vesting
	}
	ret := make([]types.Period, 0)
	remaining := sdk.NewCoins(sdk.NewCoin(common.BaseDenom, sdk.NewIntFromBigInt(p.Amount)))
	each := sdk.NewCoins(sdk.NewCoin(common.BaseDenom, sdk.NewDecFromBigInt(p.Amount).QuoInt64(int64(p.VestingLength)).TruncateInt()))
	duration := p.VestingDuration / int64(p.VestingLength)
	if duration == 0 {
		duration = 1
	}
	for i := uint(0); i < p.VestingLength; i++ {
		tmp := types.Period{}
		if i == p.VestingLength-1 {
			tmp.Amount = remaining
			tmp.Length = p.VestingDuration - int64(p.VestingLength-1)*duration
			if tmp.Length == 0 {
				tmp.Length = 1
			}
		} else {
			tmp.Length = duration
			tmp.Amount = each
		}
		ret = append(ret, tmp)
		remaining = remaining.Sub(each)
	}

	return ret
}

func (p *TxCreateVestingParams) LockupPeriods() types.Periods {
	if len(p.Lockup) > 0 {
		return p.Lockup
	}

	return nil
}

func (p *TxCreateVestingParams) Funder() sdk.AccAddress {
	if p.FunderAddr != "" {
		return account.MustParseCosmosAddress(p.FunderAddr)
	}

	return p.MustGetPrivateKey().AccAddress()
}

func (p *TxCreateVestingParams) VestingAddress() sdk.AccAddress {
	return account.MustParseCosmosAddress(p.ToAddr)
}
