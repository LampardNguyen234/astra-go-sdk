package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
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

	return common.ParseAmount(vestingBalance.Vested).Sub(
		common.ParseAmount(totalVesting).Sub(balance.Total),
	), nil
}

// VestingBalances reports the balances of a vesting account.
type VestingBalances struct {
	// Total is the original vesting amount.
	Total sdk.Coins

	// Locked is the is amount of coins which are locked.
	Locked sdk.Coins

	// Unlocked is the is amount of coins which are unlocked.
	Unlocked sdk.Coins

	// Vested is the is amount of coins which are vested.
	Vested sdk.Coins

	// Unvested is the is amount of coins which are in the vesting period.
	Unvested sdk.Coins
}

// GetVestingBalance returns the detail balance of a vesting account.
func (c *CosmosClient) GetVestingBalance(strAddr string) (*VestingBalances, error) {
	va, err := c.GetVestingAccount(strAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get vesting account %v: %v", strAddr, err)
	}

	target := time.Now().Unix()
	latestBlk, err := c.LatestBlockHeight()
	if err == nil {
		blk, err := c.GetBlockByHeight(latestBlk.Int64())
		if err == nil {
			target = blk.Block.Time.Unix()
		}
	}

	resp := new(VestingBalances)
	resp.Unlocked = vestingTypes.ReadSchedule(
		va.GetStartTime(), va.EndTime, va.LockupPeriods, va.OriginalVesting, target,
	)
	resp.Locked = va.OriginalVesting.Sub(resp.Unlocked)

	resp.Vested = vestingTypes.ReadSchedule(
		va.GetStartTime(), va.EndTime, va.VestingPeriods, va.OriginalVesting, target,
	)
	resp.Unvested = va.OriginalVesting.Sub(resp.Vested)
	resp.Total = va.OriginalVesting

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

	var ret *vestingTypes.ClawbackVestingAccount
	err = c.Codec.UnpackAny(resp.Account, &ret)
	if err != nil {
		return nil, fmt.Errorf("not a vesting account: %v", err)
	}

	return ret, nil
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

	return next, common.ParseAmount(acc.VestingPeriods[count+1].Amount), nil
}
