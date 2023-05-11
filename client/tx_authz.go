package client

import (
	"fmt"
	msg_params2 "github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func (c *CosmosClient) txGrantAuthorization(p msg_params2.TxGrantParams, auth authz.Authorization) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}
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

	return c.BuildAndSendTx(p.TxParams, msg)
}

func (c *CosmosClient) txGrantExec(p msg_params2.TxParams, msgs ...sdk.Msg) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}
	if len(msgs) == 0 {
		return nil, fmt.Errorf("no message to execute")
	}

	msg := authz.NewMsgExec(p.Operator(), msgs)

	return c.BuildAndSendTx(p, &msg)
}
