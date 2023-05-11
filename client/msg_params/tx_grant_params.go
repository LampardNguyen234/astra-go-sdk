package msg_params

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/pkg/errors"
	"time"
)

var (
	noExpiredTime = time.Date(9999, 12, 31, 23, 59, 59, 0, time.Now().Location())
)

// NewAuthorization creates a new authz.Authorization from the given data.
func NewAuthorization(msgTypeURL string, data interface{}) (authz.Authorization, error) {
	switch msgTypeURL {
	default:
		return authz.NewGenericAuthorization(msgTypeURL), nil
	}
}

type TxGrantParams struct {
	TxParams

	// Grantee is the address of the grantee to whom the Granter grants authorizations to.
	Grantee string

	// ExpiredTime is the time at which the grant is expired.
	ExpiredTime time.Time
}

func (p TxGrantParams) IsValid() (bool, error) {
	if _, err := p.TxParams.IsValid(); err != nil {
		return false, errors.Wrapf(err, "invalid TxParams")
	}
	if _, err := account.ParseCosmosAddress(p.Grantee); err != nil {
		return false, errors.Wrapf(err, "invalid ValAddress")
	}

	return true, nil
}

func (p TxGrantParams) GranterAddress() sdk.AccAddress {
	return p.Operator()
}

func (p TxGrantParams) GranteeAddress() sdk.AccAddress {
	return account.MustParseCosmosAddress(p.Grantee)
}

func (p TxGrantParams) Expiration() time.Time {
	if p.ExpiredTime.IsZero() {
		return noExpiredTime
	}

	return p.ExpiredTime
}
