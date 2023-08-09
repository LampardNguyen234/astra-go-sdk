package client

import (
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (c *CosmosClient) TxSubmitProposal(p msg_params.TxParams, content govTypes.Content, initialDeposit ...sdk.Coin) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}
	msg, err := govTypes.NewMsgSubmitProposal(content, initialDeposit, p.MustGetPrivateKey().AccAddress())
	if err != nil {
		return nil, err
	}

	return c.BuildAndSendTx(p, msg)
}

func (c *CosmosClient) TxVoteProposal(p msg_params.TxParams, id uint64, option govTypes.VoteOption) (*sdk.TxResponse, error) {
	if _, err := p.IsValid(); err != nil {
		return nil, err
	}
	return c.BuildAndSendTx(p, govTypes.NewMsgVote(p.MustGetPrivateKey().AccAddress(), id, option))
}
