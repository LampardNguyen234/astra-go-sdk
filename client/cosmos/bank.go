package cosmos

import (
	"context"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	_ "github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/grpc"
)

type BankClient struct {
	bankTypes.QueryClient
}

// NewBankClient creates a new client for interacting with the `bank` module.
func NewBankClient(conn grpc.ClientConn) *BankClient {
	return &BankClient{
		QueryClient: bankTypes.NewQueryClient(conn),
	}
}

type AccountBalance struct {
	Total    sdk.Int `json:"Total"`
	Locked   sdk.Int `json:"Locked"`
	Unlocked sdk.Int `json:"Unlocked"`
}

// Balance retrieves the balances of an account.
func (c *CosmosClient) Balance(addr string) (*AccountBalance, error) {
	total, err := c.BankClient.QueryClient.Balance(c.ctx, &bankTypes.QueryBalanceRequest{
		Address: addr,
		Denom:   common.BaseDenom,
	})
	if err != nil {
		return nil, err
	}

	unlocked, err := c.BankClient.QueryClient.SpendableBalances(context.Background(), &bankTypes.QuerySpendableBalancesRequest{
		Address:    addr,
		Pagination: nil,
	})

	tmpTotal := total.GetBalance().Amount
	tmpUnlocked := unlocked.GetBalances().AmountOf(common.BaseDenom)

	return &AccountBalance{
		Total:    tmpTotal,
		Locked:   tmpTotal.Sub(tmpUnlocked),
		Unlocked: tmpUnlocked,
	}, nil
}
