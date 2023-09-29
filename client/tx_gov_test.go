package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"testing"
	"time"
)

func proposeAndVote(t *testing.T, content govTypes.Content) {
	resp, err := c.TxSubmitProposal(msg_params.TxParams{PrivateKey: proposer},
		content,
		sdk.NewCoin(common.BaseDenom, sdk.NewIntWithDecimal(1, 18)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("proposalTx: %v\n", resp.TxHash)
	time.Sleep(20 * time.Second)

	allProposals, err := c.Proposals()
	if err != nil {
		panic(err)
	}
	toBeVoted := uint64(1)
	for _, p := range allProposals {
		if p.ProposalId > toBeVoted {
			toBeVoted = p.ProposalId
		}
	}

	// Vote the proposal
	for _, voter := range voters {
		resp, err = c.TxVoteProposal(msg_params.TxParams{PrivateKey: voter},
			toBeVoted, govTypes.OptionYes)
		if err != nil {
			panic(err)
		}
		fmt.Printf("vote of %v: %v\n", voter, resp.TxHash)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(60 * time.Second)
	p, err := c.Proposal(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}

func TestCosmosClient_TxSoftwareUpdate(t *testing.T) {
	latestBlk, err := c.LatestBlockHeight()
	if err != nil {
		panic(err)
	}
	fmt.Printf("latestBlk: %v\n", latestBlk.Uint64())

	//https://github.com/AstraProtocol/astra/releases/download/v0.0.1/astra_3.0.0_Darwin_arm64.zip
	content := &types.SoftwareUpgradeProposal{
		Title:       "Astra v3.0.0-rc",
		Description: "Upgrade to cosmos-sdk v0.46",
		Plan: types.Plan{
			Name:   "v3.0.0",
			Height: latestBlk.Int64() + 70,
			Info:   `{"binaries":{"darwin/arm64":"https://github.com/AstraProtocol/astra/releases/download/v3.0.0-rc/astra_3.0.0-rc_Darwin_arm64.tar.gz", "darwin/amd64":"https://github.com/AstraProtocol/astra/releases/download/v3.0.0-rc/astra_3.0.0-rc_Darwin_amd64.tar.gz", "linux/amd64":"https://github.com/AstraProtocol/astra/releases/download/v3.0.0-rc/astra_3.0.0-rc_Linux_amd64.tar.gz", "linux/arm64":"https://github.com/AstraProtocol/astra/releases/download/v3.0.0-rc/astra_3.0.0-rc_Linux_arm64.tar.gz"}}`,
		},
	}

	proposeAndVote(t, content)
}

func TestCosmosClient_TxVoteProposal(t *testing.T) {
	resp, err := c.TxVoteProposal(msg_params.TxParams{
		PrivateKey: proposer,
	},
		2,
		govTypes.OptionYes,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.TxHash)
}

func TestCosmosClient_TxDepositProposal(t *testing.T) {
	resp, err := c.TxDepositProposal(msg_params.TxParams{PrivateKey: proposer},
		53,
		sdk.NewCoin(common.BaseDenom, sdk.NewIntWithDecimal(19999, 18)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("depositTx: %v\n", resp.TxHash)
}

func TestCosmosClient_FeemarketParamsUpdate(t *testing.T) {
	content := &proposal.ParameterChangeProposal{
		Title:       "Increase Min Feemarket GasPrice",
		Description: "If successful, this parameter-change governance proposal will increase feemarket MinGasPrice from `0` to `100000000000` (100 billion)",
		Changes: []proposal.ParamChange{
			{
				Subspace: "feemarket",
				Key:      "MinGasPrice",
				Value:    "100000000000.00",
			},
		},
	}

	proposeAndVote(t, content)
}

func TestCosmosClient_Proposal(t *testing.T) {
	p, err := c.Proposal(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
