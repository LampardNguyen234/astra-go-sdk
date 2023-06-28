package client

import (
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	_ "github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/pkg/errors"
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
func (c *CosmosClient) Balance(strAddr string) (*AccountBalance, error) {
	addr, err := account.ParseCosmosAddress(strAddr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}

	total, err := c.bank.QueryClient.Balance(c.ctx, &bankTypes.QueryBalanceRequest{
		Address: addr.String(),
		Denom:   common.BaseDenom,
	})
	if err != nil {
		return nil, err
	}

	unlocked, err := c.bank.QueryClient.SpendableBalances(c.ctx, &bankTypes.QuerySpendableBalancesRequest{
		Address:    addr.String(),
		Pagination: nil,
	})

	tmpTotal := total.GetBalance().Amount
	tmpUnlocked := common.ParseAmount(unlocked.GetBalances())

	return &AccountBalance{
		Total:    tmpTotal,
		Locked:   tmpTotal.Sub(tmpUnlocked),
		Unlocked: tmpUnlocked,
	}, nil
}
