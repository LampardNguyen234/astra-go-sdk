package common

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethermintTypes "github.com/evmos/ethermint/types"
)

func init() {
	initConfig()
	return
}

func Init() {
	initConfig()
}

// initConfig should only be called once.
func initConfig() {
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetPurpose(44)
	sdkConfig.SetCoinType(ethermintTypes.Bip44CoinType)

	bech32PrefixAccAddr := fmt.Sprintf("%v", ChainPrefix)
	bech32PrefixAccPub := fmt.Sprintf("%vpub", ChainPrefix)
	bech32PrefixValAddr := fmt.Sprintf("%vvaloper", ChainPrefix)
	bech32PrefixValPub := fmt.Sprintf("%vvaloperpub", ChainPrefix)
	bech32PrefixConsAddr := fmt.Sprintf("%vvalcons", ChainPrefix)
	bech32PrefixConsPub := fmt.Sprintf("%vvalconspub", ChainPrefix)

	sdkConfig.SetBech32PrefixForAccount(bech32PrefixAccAddr, bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(bech32PrefixValAddr, bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(bech32PrefixConsAddr, bech32PrefixConsPub)
}
