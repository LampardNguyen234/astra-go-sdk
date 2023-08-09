package client

import (
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/gogo/protobuf/grpc"
)

type GovClient struct {
	govTypes.QueryClient
}

// NewGovClient creates a new client for interacting with the `gov` module.
func NewGovClient(conn grpc.ClientConn) *GovClient {
	return &GovClient{
		QueryClient: govTypes.NewQueryClient(conn),
	}
}

func (c *CosmosClient) Proposal(id uint64) (*govTypes.Proposal, error) {
	resp, err := c.gov.Proposal(c.ctx, &govTypes.QueryProposalRequest{ProposalId: id})
	if err != nil {
		return nil, err
	}

	return &(resp.Proposal), nil
}

func (c *CosmosClient) Proposals() ([]govTypes.Proposal, error) {
	resp, err := c.gov.Proposals(c.ctx, &govTypes.QueryProposalsRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Proposals, nil
}

func (c *CosmosClient) Vote(id uint64, voter string) (*govTypes.Vote, error) {
	resp, err := c.gov.Vote(c.ctx, &govTypes.QueryVoteRequest{ProposalId: id, Voter: voter})
	if err != nil {
		return nil, err
	}

	return &(resp.Vote), nil
}

func (c *CosmosClient) Votes(id uint64) ([]govTypes.Vote, error) {
	resp, err := c.gov.Votes(c.ctx, &govTypes.QueryVotesRequest{
		ProposalId: id,
	})
	if err != nil {
		return nil, err
	}

	return resp.Votes, nil
}
