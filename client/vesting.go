package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingTypes "github.com/evmos/evmos/v6/x/vesting/types"
	"github.com/gogo/protobuf/grpc"
	"github.com/pkg/errors"
	"time"
)

type VestingClient struct {
	vestingTypes.QueryClient
}

// NewVestingClient creates a new client for interacting with the `vesting` module.
func NewVestingClient(conn grpc.ClientConn) *VestingClient {
	return &VestingClient{
		QueryClient: vestingTypes.NewQueryClient(conn),
	}
}

// GetAvailableVestingBalance returns the spendable amount of a vesting address.
func (c *CosmosClient) GetAvailableVestingBalance(strAddr string) (sdk.Int, error) {
	vestingBalance, err := c.GetVestingBalance(strAddr)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	balance, err := c.Balance(strAddr)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	totalVesting := vestingBalance.Vested.Add(vestingBalance.Unvested...).Add(vestingBalance.Locked...)

	return vestingBalance.Vested.AmountOf(common.BaseDenom).Sub(
		totalVesting.AmountOf(common.BaseDenom).Sub(balance.Total),
	), nil
}

// GetVestingBalance returns the detail balance of a vesting account.
func (c *CosmosClient) GetVestingBalance(strAddr string) (*vestingTypes.QueryBalancesResponse, error) {
	addr, err := account.ParseCosmosAddress(strAddr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}

	resp, err := c.vesting.Balances(c.ctx, &vestingTypes.QueryBalancesRequest{Address: addr.String()})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetVestingAccount returns the vesting of the given address.
func (c *CosmosClient) GetVestingAccount(strAddr string) (*vestingTypes.ClawbackVestingAccount, error) {
	addr, err := account.ParseCosmosAddress(strAddr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAccAddress, err.Error())
	}

	resp, err := c.auth.Account(c.ctx, &authTypes.QueryAccountRequest{Address: addr.String()})
	if err != nil {
		return nil, err
	}
	if resp.Account == nil {
		return nil, fmt.Errorf("no account found")
	}

	var ret exported.VestingAccount
	err = c.Codec.UnpackAny(resp.Account, &ret)
	if err != nil {
		return nil, fmt.Errorf("not a vesting account: %v", err)
	}

	return ret.(*vestingTypes.ClawbackVestingAccount), nil
}

// GetNextVestingPeriod returns the next vesting information of the given address.
func (c *CosmosClient) GetNextVestingPeriod(strAddr string) (time.Time, sdk.Int, error) {
	acc, err := c.GetVestingAccount(strAddr)
	if err != nil {
		return time.Now(), sdk.ZeroInt(), err
	}

	next := acc.StartTime
	count := acc.GetPassedPeriodCount(time.Now())
	if count == len(acc.VestingPeriods) { // vesting DONE
		return time.Now(), sdk.ZeroInt(), nil
	}
	for i := 0; i < count+1; i++ {
		next = next.Add(acc.VestingPeriods[i].Duration())
	}

	return next, acc.VestingPeriods[count+1].Amount.AmountOf(common.BaseDenom), nil
}
