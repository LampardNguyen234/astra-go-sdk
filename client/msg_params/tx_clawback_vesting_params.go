package msg_params

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type TxClawBackVestingParams struct {
	TxParams `json:"TxParams"`

	// FunderAddr is the address of the funder.
	//
	// If not given, it will be retrieved from the private key.
	//
	// If FunderAddr != Operator => GrantExecMsg.
	FunderAddr string `json:"FunderAddr,omitempty"`

	// AccountAddr is the address of the vesting account to clawback.
	AccountAddr string `json:"AccountAddr"`

	// ToAddr is the address to which the remaining tokens will be transferred to.
	// If not given, ToAddr = FunderAddr.
	ToAddr string `json:"ToAddr,omitempty"`
}

func (p *TxClawBackVestingParams) IsValid() (bool, error) {
	if _, err := p.TxParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}

	if p.FunderAddr != "" {
		if _, err := account.ParseCosmosAddress(p.FunderAddr); err != nil {
			return false, errors.Wrapf(err, "invalid FunderAddr")
		}
	}

	if _, err := account.ParseCosmosAddress(p.AccountAddr); err != nil {
		return false, errors.Wrapf(err, "invalid AccountAddr")
	}

	if p.ToAddr != "" {
		if _, err := account.ParseCosmosAddress(p.ToAddr); err != nil {
			return false, errors.Wrapf(err, "invalid ToAddr")
		}
	}

	return true, nil
}

func (p *TxClawBackVestingParams) Funder() sdk.AccAddress {
	if p.FunderAddr != "" {
		return account.MustParseCosmosAddress(p.FunderAddr)
	}

	return p.MustGetPrivateKey().AccAddress()
}

func (p *TxClawBackVestingParams) VestingAddress() sdk.AccAddress {
	return account.MustParseCosmosAddress(p.AccountAddr)
}

func (p *TxClawBackVestingParams) To() sdk.AccAddress {
	if p.ToAddr != "" {
		return account.MustParseCosmosAddress(p.ToAddr)
	}

	return p.Funder()
}
