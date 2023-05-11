package cosmos

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/client/cosmos/msg_params"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCosmosClient_TxGrantAuthorization(t *testing.T) {
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
