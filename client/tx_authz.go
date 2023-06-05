package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/pkg/errors"
)

// TxGranterRevokeAll revokes all authz permissions granted by the sender.
func (c *CosmosClient) TxGranterRevokeAll(p msg_params.TxParams) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	grants, err := c.GranterGrants(p.Operator().String())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get granterGrants")
	}

	msgs := make([]sdk.Msg, 0)
	for _, grant := range grants {
		tmpMsg := authz.NewMsgRevoke(p.Operator(), account.MustParseCosmosAddress(grant.Grantee), grant.Authorization.MsgTypeURL())
		msgs = append(msgs, &tmpMsg)
	}

	if len(msgs) == 0 {
		return nil, fmt.Errorf("no grant to revoke")
	}

	return c.BuildAndSendTx(p, msgs...)
}

// TxGrantRevokeAll revokes all authz permissions granted by the sender to the grantee.
func (c *CosmosClient) TxGrantRevokeAll(p msg_params.TxParams, granteeStr string) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}
	grantee, err := account.ParseCosmosAddress(granteeStr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}

	grants, err := c.Grants(p.Operator().String(), granteeStr, "")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get granterGrants")
	}

	msgs := make([]sdk.Msg, 0)
	for _, grant := range grants {
		tmpMsg := authz.NewMsgRevoke(p.Operator(), grantee, grant.Authorization.MsgTypeURL())
		msgs = append(msgs, &tmpMsg)
	}

	if len(msgs) == 0 {
		return nil, fmt.Errorf("no grant to revoke")
	}

	return c.BuildAndSendTx(p, msgs...)
}

// TxGrantRevoke revoke a grant permission granted by to sender to the grantee.
func (c *CosmosClient) TxGrantRevoke(p msg_params.TxParams, granteeStr string, msgTypeURL string) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}

	grantee, err := account.ParseCosmosAddress(granteeStr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, "%v: %v", granteeStr, err.Error())
	}

	msg := authz.NewMsgRevoke(p.Operator(), grantee, msgTypeURL)
	return c.BuildAndSendTx(p, &msg)
}

func (c *CosmosClient) TxGrantAuthorization(p msg_params.TxGrantParams, auths ...authz.Authorization) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}
	if len(auths) == 0 {
		return nil, fmt.Errorf("no auth provided")
	}
	msgs := make([]sdk.Msg, 0)
	for _, auth := range auths {
		if auth == nil {
			return nil, fmt.Errorf("empty Auth")
		}
		if err := auth.ValidateBasic(); err != nil {
			return nil, err
		}

		msg, err := authz.NewMsgGrant(p.GranterAddress(), p.GranteeAddress(), auth, p.Expiration())
		if err != nil {
			return nil, err
		}

		msgs = append(msgs, msg)
	}

	return c.BuildAndSendTx(p.TxParams, msgs...)
}

func (c *CosmosClient) TxGrantExec(p msg_params.TxParams, msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}
	if len(msgs) == 0 {
		return nil, fmt.Errorf("no message to execute")
	}

	msg := authz.NewMsgExec(p.Operator(), msgs)

	return c.BuildAndSendTx(p, &msg)
}
