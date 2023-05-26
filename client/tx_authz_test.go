package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	distrType "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/rand"
	"math/big"
	"testing"
	"time"
)

var (
	granterAddr string
	granteeAddr string
)

func randomAuth() authz.Authorization {
	i := rand.Int() % 6
	switch i {
	case 0:
		return authz.NewGenericAuthorization(sdk.MsgTypeURL(&types.MsgMultiSend{}))
	case 1:
		return authz.NewGenericAuthorization(sdk.MsgTypeURL(&distrType.MsgWithdrawDelegatorReward{}))
	case 2:
		return &stakingTypes.StakeAuthorization{AuthorizationType: stakingTypes.AuthorizationType_AUTHORIZATION_TYPE_DELEGATE}
	case 3:
		return &stakingTypes.StakeAuthorization{AuthorizationType: stakingTypes.AuthorizationType_AUTHORIZATION_TYPE_UNDELEGATE}
	case 4:
		return &stakingTypes.StakeAuthorization{AuthorizationType: stakingTypes.AuthorizationType_AUTHORIZATION_TYPE_REDELEGATE}
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

func TestCosmosClient_TxGrantMsgExec(t *testing.T) {
	expired := time.Now().Add(10 * time.Minute)
	grantee2, err := account.NewPrivateKeyFromString("88278f72156a52013a03cf97a6ecd857a7294b4a1ca99dbdf8ecf64177b0e287")
	if err != nil {
		panic(err)
	}

	amt := new(big.Int).SetUint64(testAmt)

	// granter grant MsgSend to grantee
	p := msg_params.TxGrantParams{
		TxParams:    *defaultTxParams,
		Grantee:     toAddr,
		ExpiredTime: expired,
	}
	resp, err := c.TxGrantSend(p, amt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("txHash: %v\n", resp.TxHash)
	time.Sleep(20 * time.Second)

	// grantee grants MsgExec to grantee2
	p2 := msg_params.TxGrantParams{
		TxParams: msg_params.TxParams{
			PrivateKey: granteePrivateKey,
		},
		Grantee:     grantee2.AccAddress().String(),
		ExpiredTime: expired,
	}
	resp, err = c.txGrantAuthorization(p2, authz.NewGenericAuthorization(sdk.MsgTypeURL(&authz.MsgExec{})))
	if err != nil {
		panic(err)
	}
	fmt.Printf("txHash: %v\n", resp.TxHash)
	time.Sleep(20 * time.Second)

	granterBalance, err := c.Balance(granterAddr)
	if err != nil {
		panic(err)
	}
	granteeBalance, err := c.Balance(granteeAddr)
	if err != nil {
		panic(err)
	}
	grantee2Balance, err := c.Balance(grantee2.AccAddress().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("BalanceBefore: %v, %v, %v\n",
		granterBalance.Total.String(), granteeBalance.Total.String(), grantee2Balance.Total.String(),
	)

	// grantee2 performs MsgSend(granter, grantee, amt/2)
	msg := types.NewMsgSend(
		account.MustParseCosmosAddress(granterAddr),
		account.MustParseCosmosAddress(granteeAddr),
		sdk.NewCoins(sdk.NewCoin(common.BaseDenom,
			sdk.NewIntFromBigInt(new(big.Int).Div(amt, new(big.Int).SetUint64(2))))),
	)
	tmpMsg := authz.NewMsgExec(
		account.MustParseCosmosAddress(granteeAddr), []sdk.Msg{msg})
	tmpMsg = authz.NewMsgExec(grantee2.AccAddress(), []sdk.Msg{&tmpMsg})

	txParams := msg_params.TxParams{
		PrivateKey: grantee2.String(),
	}
	resp, err = c.TxGrantExec(txParams, &tmpMsg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("txHash: %v\n", resp.TxHash)
	time.Sleep(20 * time.Second)

	granterBalance, err = c.Balance(granterAddr)
	if err != nil {
		panic(err)
	}
	granteeBalance, err = c.Balance(granteeAddr)
	if err != nil {
		panic(err)
	}
	grantee2Balance, err = c.Balance(grantee2.AccAddress().String())
	if err != nil {
		panic(err)
	}
	fmt.Printf("BalanceAfter: %v, %v, %v\n",
		granterBalance.Total.String(), granteeBalance.Total.String(), grantee2Balance.Total.String(),
	)
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
			fmt.Printf("resp: %v\n", resp.TxHash)
		}

		time.Sleep(10 * time.Second)
	}
}

func TestCosmosClientWithdrawAndStake(t *testing.T) {
	p := msg_params.TxParams{
		PrivateKey: granteePrivateKey,
	}

	from := account.MustParseCosmosAddress(addr)
	val := account.MustParseCosmosValidatorAddress("astravaloper1u7gf4z49v53yrxy6ggrzhxfqj46c3ap4tzku46")

	resp, err := c.TxGrantExec(
		p,
		distrType.NewMsgWithdrawDelegatorReward(
			from,
			val,
		),
		stakingTypes.NewMsgDelegate(
			from,
			val,
			sdk.NewCoin(common.BaseDenom, sdk.NewInt(1000000000000)),
		),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("txHash: %v\n", resp.TxHash)
}
