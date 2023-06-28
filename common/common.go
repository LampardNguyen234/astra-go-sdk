package common

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

const (
	ChainPrefix = "astra"
	Denom       = "astra"
	BaseDenom   = "aastra"
	Decimals    = 18
)

var AsaDecimalsBigInt = big.NewInt(Decimals)

// BigIntToFloat64 returns the normalized amount without decimals.
func BigIntToFloat64(rawValue *big.Int) float64 {
	if rawValue == nil {
		return 0
	}

	tmp := sdk.NewDecFromBigInt(rawValue)
	tmp = tmp.QuoInt(sdk.NewIntWithDecimal(1, Decimals))

	return tmp.MustFloat64()
}

// Float64ToBigInt returns the raw big int value including the decimals.
func Float64ToBigInt(amt float64) *big.Int {
	if amt == 0 {
		return new(big.Int).SetInt64(0)
	}

	tmp := sdk.MustNewDecFromStr(fmt.Sprintf("%f", amt))
	tmp = tmp.MulInt(sdk.NewIntWithDecimal(1, Decimals))

	return tmp.TruncateInt().BigInt()
}

// ParseAmount gets the amounts of the given coins w.r.t to the BaseDenom.
func ParseAmount(coins interface{}) sdk.Int {
	return ParseAmountWithDenom(coins, BaseDenom).TruncateInt()
}

// ParseAmountWithDenom gets the amounts of the given coins w.r.t given denom.
func ParseAmountWithDenom(coins interface{}, denom string) sdk.Dec {
	if _, ok := coins.(sdk.Coins); ok {
		return ParseCoinsAmount(coins.(sdk.Coins), denom)
	}
	if _, ok := coins.(sdk.DecCoins); ok {
		return ParseDecCoinsAmount(coins.(sdk.DecCoins), denom)
	}

	return sdk.ZeroDec()
}

// ParseCoinsAmount returns the amount of coins w.r.t the given denom.
func ParseCoinsAmount(coins sdk.Coins, denom string) sdk.Dec {
	ret := sdk.ZeroDec()
	switch denom {
	case BaseDenom:
		for _, coin := range coins {
			switch coin.Denom {
			case BaseDenom:
				ret = ret.Add(sdk.NewDecFromInt(coin.Amount))
			case Denom:
				ret = ret.Add(sdk.NewDecFromBigInt(
					Float64ToBigInt(
						float64(coin.Amount.Int64()),
					),
				))
			default:
				continue
			}
		}
		return ret
	case Denom:
		for _, coin := range coins {
			switch coin.Denom {
			case Denom:
				ret = ret.Add(sdk.NewDecFromInt(coin.Amount))
			case BaseDenom:
				ret = ret.Add(sdk.NewDecFromBigIntWithPrec(coin.Amount.BigInt(), Decimals))
			default:
				continue
			}
		}
		return ret
	default:
		return sdk.ZeroDec()
	}
}

// ParseDecCoinsAmount returns the amount of dec-coins w.r.t the given denom.
func ParseDecCoinsAmount(coins sdk.DecCoins, denom string) sdk.Dec {
	ret := sdk.ZeroDec()
	switch denom {
	case BaseDenom:
		for _, coin := range coins {
			switch coin.Denom {
			case BaseDenom:
				ret = ret.Add(coin.Amount)
			case Denom:
				ret = ret.Add(sdk.NewDecFromBigInt(
					Float64ToBigInt(
						coin.Amount.MustFloat64(),
					),
				))
			default:
				continue
			}
		}
		return ret
	case Denom:
		for _, coin := range coins {
			switch coin.Denom {
			case Denom:
				ret = ret.Add(coin.Amount)
			case BaseDenom:
				ret = ret.Add(sdk.NewDecFromIntWithPrec(coin.Amount.TruncateInt(), Decimals))
			default:
				continue
			}
		}
		return ret
	default:
		return sdk.ZeroDec()
	}
}
