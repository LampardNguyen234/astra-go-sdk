package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"testing"
	"time"
)

func TestCosmosClient_TxSoftwareUpdate(t *testing.T) {
	proposer := "5DDDA62142D4616E9CA70EF24EF52B35CCD5F54B69042023F6B73B1EF1EFCADB"
	voters := []string{
		//"ACA0E44BBE27494E1D169B46B08CE92C29A7A2069CCABC0574E8899F5D877658",
		"B5BDCB4769D99DB1EB00CFD2117C543404B92E26B1C1FBD899DE41743D36614F",
		"24BCAA7F388EB1B8B3F6A46530CE0DCB2CCD4177F945320A0C0CE9DF544469A2",
		"5DDDA62142D4616E9CA70EF24EF52B35CCD5F54B69042023F6B73B1EF1EFCADB",
	}
	//https://github.com/LampardNguyen234/astra-release/releases/download/v0.0.1/astra_3.0.0_Darwin_arm64.zip
	content := &types.SoftwareUpgradeProposal{
		Title:       "Astra v3.0.0",
		Description: "Upgrade to cosmos-sdk v0.46",
		Plan: types.Plan{
			Name:   "v3.0.0",
			Height: 150,
			Info:   `{"binaries":{"darwin/arm64":"https://github.com/LampardNguyen234/astra-release/releases/download/v0.0.1/astra_3.0.0_Darwin_arm64.tar.gz"}}`,
		},
	}

	resp, err := c.TxSubmitProposal(msg_params.TxParams{PrivateKey: proposer},
		content,
		sdk.NewCoin(common.BaseDenom, sdk.NewIntWithDecimal(1, 18)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("proposalTx: %v\n", resp.TxHash)
	time.Sleep(20 * time.Second)

	// Vote the proposal
	for _, voter := range voters {
		resp, err = c.TxVoteProposal(msg_params.TxParams{PrivateKey: voter},
			1, govTypes.OptionYes)
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

func TestCosmosClient_Proposal(t *testing.T) {
	p, err := c.Proposal(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
