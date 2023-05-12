package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	distrType "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
	"time"
)

var (
	granterAddr string
	granteeAddr string
)

func randomAuth() authz.Authorization {
	i := rand.Int() % 3
	switch i {
	case 0:
		return authz.NewGenericAuthorization(sdk.MsgTypeURL(&types.MsgMultiSend{}))
	case 1:
		return authz.NewGenericAuthorization(sdk.MsgTypeURL(&distrType.MsgWithdrawDelegatorReward{}))
	default:
		return authz.NewGenericAuthorization(types.SendAuthorization{}.MsgTypeURL())
	}
}

func TestCosmosClient_TxGrantRevokeAll(t *testing.T) {
	p := msg_params.TxGrantParams{
		TxParams:    *defaultTxParams,
		Grantee:     granteeAddr,
		ExpiredTime: time.Now().Add(5 * time.Minute),
	}

	numTests := 1 + rand.Int()%20

	auths := make([]authz.Authorization, 0)
	for i := 0; i < numTests; i++ {
		auths = append(auths, randomAuth())
	}
	resp, err := c.txGrantAuthorization(p, auths...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("txHash: %v\n", resp.TxHash)

	time.Sleep(20 * time.Second)

	grants, err := c.Grants(granterAddr, granteeAddr, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("grants:", len(grants))
	for _, grant := range grants {
		fmt.Printf("msg: %v\n", grant.Authorization.MsgTypeURL())
	}

	resp, err = c.TxGrantRevokeAll(*defaultTxParams, granteeAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("txHash:", resp.TxHash)

	time.Sleep(20 * time.Second)

	grants, err = c.Grants(granterAddr, granteeAddr, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("grants:", len(grants))
}

func TestCosmosClient_TxGranterGrantRevokeAll(t *testing.T) {
	p := msg_params.TxGrantParams{
		TxParams:    *defaultTxParams,
		Grantee:     granteeAddr,
		ExpiredTime: time.Now().Add(5 * time.Minute),
	}

	numTests := 1 + rand.Int()%20

	auths := make([]authz.Authorization, 0)
	for i := 0; i < numTests; i++ {
		auths = append(auths, randomAuth())
	}
	resp, err := c.txGrantAuthorization(p, auths...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("txHash: %v\n", resp.TxHash)

	time.Sleep(20 * time.Second)

	grants, err := c.Grants(granterAddr, granteeAddr, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("grants:", len(grants))
	for _, grant := range grants {
		fmt.Printf("msg: %v\n", grant.Authorization.MsgTypeURL())
	}

	resp, err = c.TxGrantRevokeAll(*defaultTxParams, granteeAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("txHash:", resp.TxHash)

	time.Sleep(20 * time.Second)

	grants, err = c.Grants(granterAddr, granteeAddr, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("grants:", len(grants))
}

func TestCosmosClient_txGrantAuthorization(t *testing.T) {
	defaultExpired := time.Now().Add(1 * time.Minute)

	auth := authz.NewGenericAuthorization(types.SendAuthorization{}.MsgTypeURL())

	testCases := []struct {
		name        string
		hook        func()
		params      msg_params.TxGrantParams
		auth        authz.Authorization
		expectedErr error
	}{
		{
			name: "valid grant",
			hook: nil,
			params: msg_params.TxGrantParams{
				TxParams:    *defaultTxParams,
				Grantee:     toAddr,
				ExpiredTime: defaultExpired,
			},
			auth:        auth,
			expectedErr: nil,
		},
		{
			name: "empty auth",
			hook: nil,
			params: msg_params.TxGrantParams{
				TxParams:    *defaultTxParams,
				Grantee:     toAddr,
				ExpiredTime: defaultExpired,
			},
			auth:        nil,
			expectedErr: fmt.Errorf("empty Auth"),
		},
	}

	for _, tc := range testCases {
		msgStr := fmt.Sprintf("tc: %v", tc.name)

		resp, err := c.txGrantAuthorization(tc.params, tc.auth)
		if tc.expectedErr != nil {
			assert.Error(t, err, msgStr)
			assert.ErrorContains(t, err, tc.expectedErr.Error(), msgStr)

		} else {
			fmt.Printf("resp: %v\n", resp.String())
		}
	}
}
