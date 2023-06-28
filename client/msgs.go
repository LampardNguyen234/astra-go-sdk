package client

import (
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	vestingTypes "github.com/evmos/evmos/v6/x/vesting/types"
)

// ParseCosmosMsgValue returns the value of the given Cosmos message.
func ParseCosmosMsgValue(msg sdk.Msg) (sdk.Int, error) {
	v := sdk.ZeroInt()
	switch msg.(type) {
	case *bankTypes.MsgSend:
		v = common.ParseCoinsAmount(msg.(*bankTypes.MsgSend).Amount, common.BaseDenom).TruncateInt()
	case *bankTypes.MsgMultiSend:
		tmpMsg := msg.(*bankTypes.MsgMultiSend)
		for _, out := range tmpMsg.Outputs {
			v = v.Add(
				common.ParseCoinsAmount(out.Coins, common.BaseDenom).TruncateInt(),
			)
		}

	case *stakingTypes.MsgDelegate:
		v = msg.(*stakingTypes.MsgDelegate).Amount.Amount
	case *stakingTypes.MsgBeginRedelegate:
		v = msg.(*stakingTypes.MsgBeginRedelegate).Amount.Amount
	case *stakingTypes.MsgUndelegate:
		v = msg.(*stakingTypes.MsgUndelegate).Amount.Amount
	case *stakingTypes.MsgCreateValidator:
		v = msg.(*stakingTypes.MsgCreateValidator).Value.Amount

	case *vestingTypes.MsgCreateClawbackVestingAccount:
		tmpMsg := msg.(*vestingTypes.MsgCreateClawbackVestingAccount)
		v = common.ParseCoinsAmount(tmpMsg.VestingPeriods.TotalAmount(), common.BaseDenom).TruncateInt()
	default:
		return sdk.ZeroInt(), fmt.Errorf("msg `%v` not supported", sdk.MsgTypeURL(msg))
	}

	return v, nil
}

// ParseCosmosMsgSender returns the sender of the given Cosmos message.
func ParseCosmosMsgSender(msg sdk.Msg) (string, error) {
	ret := ""
	switch msg.(type) {
	case *bankTypes.MsgSend,
		*stakingTypes.MsgCreateValidator,
		*stakingTypes.MsgDelegate,
		*stakingTypes.MsgBeginRedelegate,
		*stakingTypes.MsgUndelegate,
		*vestingTypes.MsgCreateClawbackVestingAccount,
		*vestingTypes.MsgClawback:
		ret = msg.GetSigners()[0].String()
	default:
		return "", fmt.Errorf("msg `%v` not supported", sdk.MsgTypeURL(msg))
	}

	return ret, nil
}

// ParseCosmosMsgReceiver returns the recipients of the given Cosmos message.
func ParseCosmosMsgReceiver(msg sdk.Msg) (string, error) {
	ret := ""
	switch msg.(type) {
	case *bankTypes.MsgSend:
		ret = msg.(*bankTypes.MsgSend).ToAddress
	case *vestingTypes.MsgCreateClawbackVestingAccount:
		ret = msg.(*vestingTypes.MsgCreateClawbackVestingAccount).ToAddress
	case *stakingTypes.MsgCreateValidator,
		*stakingTypes.MsgDelegate,
		*stakingTypes.MsgBeginRedelegate,
		*stakingTypes.MsgUndelegate,
		*vestingTypes.MsgClawback:
	default:
		return "", fmt.Errorf("msg `%v` not supported", sdk.MsgTypeURL(msg))
	}

	return ret, nil
}
