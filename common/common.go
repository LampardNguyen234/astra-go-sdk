package common

import "math/big"

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

	amt := new(big.Float).SetInt(rawValue)
	amt = amt.Quo(amt, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), AsaDecimalsBigInt, nil)))
	amtFloat, _ := amt.Float64()

	return amtFloat
}

// Float64ToBigInt returns the raw big int value including the decimals.
func Float64ToBigInt(amt float64) *big.Int {
	if amt == 0 {
		return new(big.Int).SetInt64(0)
	}

	tmp := new(big.Int).Exp(AsaDecimalsBigInt, big.NewInt(10), nil)
	tmp, _ = new(big.Float).Mul(big.NewFloat(amt), new(big.Float).SetInt(tmp)).Int(nil)
	return tmp
}
