package client

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/pkg/errors"
	"time"
)

type Grant struct {
	Authorization authz.Authorization
	Expiration    time.Time
}

func newGrants(codec codec.Codec, grants ...*authz.Grant) ([]*Grant, error) {
	ret := make([]*Grant, 0)
	for _, grant := range grants {
		tmpRet := &Grant{
			Expiration: grant.Expiration,
		}

		var tmpAuth authz.Authorization
		err := codec.UnpackAny(grant.Authorization, &tmpAuth)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unpack auth")
		}
		tmpRet.Authorization = tmpAuth
		ret = append(ret, tmpRet)
	}

	return ret, nil
}

type GrantAuthorization struct {
	Granter       string
	Grantee       string
	Authorization authz.Authorization
	Expiration    time.Time
}

func newGrantAuthorizations(codec codec.Codec, grants ...*authz.GrantAuthorization) ([]*GrantAuthorization, error) {
	ret := make([]*GrantAuthorization, 0)
	for _, grant := range grants {
		tmpRet := &GrantAuthorization{
			Granter:       grant.Granter,
			Grantee:       grant.Grantee,
			Authorization: nil,
			Expiration:    grant.Expiration,
		}

		var tmpAuth authz.Authorization
		err := codec.UnpackAny(grant.Authorization, &tmpAuth)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unpack auth")
		}
		tmpRet.Authorization = tmpAuth
		ret = append(ret, tmpRet)
	}

	return ret, nil
}
