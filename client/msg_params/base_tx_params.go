package msg_params

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type TxParams struct {
	PrivateKey    string
	GasLimit      uint64
	GasAdjustment float64
	GasPrice      string
	Memo          string
}

// NewTxParams creates a new TxParams from the given parameters.
func NewTxParams(privateKeyStr string, gasLimit uint64, gasPrice string, gasAdjustment float64) (*TxParams, error) {
	p := &TxParams{
		PrivateKey:    privateKeyStr,
		GasLimit:      gasLimit,
		GasAdjustment: gasAdjustment,
		GasPrice:      gasPrice,
	}
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	return p, nil
}

func DefaultTxParams() TxParams {
	return TxParams{
		GasLimit:      0,
		GasAdjustment: 1.01,
		GasPrice:      "100000000000aastra",
	}
}

func (p TxParams) MustGetPrivateKey() account.PrivateKey {
	k, err := account.NewPrivateKeyFromString(p.PrivateKey)
	if err != nil {
		panic(err)
	}

	return *k
}

func (p TxParams) IsValid() (bool, error) {
	_, err := account.NewPrivateKeyFromString(p.PrivateKey)
	if err != nil {
		return false, errors.Wrapf(err, "invalid privateKey")
	}

	return true, nil
}

func (p TxParams) Operator() sdk.AccAddress {
	return p.MustGetPrivateKey().AccAddress()
}
