package common

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"log"
	"testing"
)

func TestParseCoinsAmount(t *testing.T) {
	tcs := []struct {
		name        string
		coins       sdk.Coins
		denom       string
		expectedRet sdk.Dec
	}{
		{
			name: "all coins with same base denom - 00",
			coins: []sdk.Coin{
				{
					Denom:  BaseDenom,
					Amount: sdk.NewInt(1),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.NewInt(10),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.NewInt(100),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.NewInt(1000),
				},
			},
			denom:       BaseDenom,
			expectedRet: sdk.NewDec(1111),
		},
		{
			name: "all coins with same base denom - 01",
			coins: []sdk.Coin{
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("1000000000000000000").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("10000000000000000000").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("10000000000000000").TruncateInt(),
				},
			},
			denom:       Denom,
			expectedRet: sdk.MustNewDecFromStr("11.11"),
		},
		{
			name: "all coins with same denom - 00",
			coins: []sdk.Coin{
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("10").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("100").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1000").TruncateInt(),
				},
			},
			denom:       BaseDenom,
			expectedRet: sdk.MustNewDecFromStr("1111000000000000000000"),
		},
		{
			name: "all coins with same denom - 01",
			coins: []sdk.Coin{
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("10").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("100").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1000").TruncateInt(),
				},
			},
			denom:       Denom,
			expectedRet: sdk.MustNewDecFromStr("1111"),
		},
		{
			name: "all coins with mixed denoms - 00",
			coins: []sdk.Coin{
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("100").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000000").TruncateInt(),
				},
			},
			denom:       Denom,
			expectedRet: sdk.MustNewDecFromStr("201.1"),
		},
		{
			name: "all coins with mixed denoms - 01",
			coins: []sdk.Coin{
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("100").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000000").TruncateInt(),
				},
			},
			denom:       BaseDenom,
			expectedRet: sdk.MustNewDecFromStr("201100000000000000000"),
		},
		{
			name: "all coins with mixed denoms and unsupported denoms - 00",
			coins: []sdk.Coin{
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000").TruncateInt(),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("100").TruncateInt(),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000000").TruncateInt(),
				},
				{
					Denom:  "test-denom-1",
					Amount: sdk.MustNewDecFromStr("100000000000000000000").TruncateInt(),
				},
				{
					Denom:  "test-denom-2",
					Amount: sdk.MustNewDecFromStr("12348734133219043289").TruncateInt(),
				},
				{
					Denom:  "test-denom-3",
					Amount: sdk.MustNewDecFromStr("123248232314908328742340").TruncateInt(),
				},
			},
			denom:       BaseDenom,
			expectedRet: sdk.MustNewDecFromStr("201100000000000000000"),
		},
	}

	for _, tc := range tcs {
		msg := fmt.Sprintf("[tc: %v]", tc.name)

		actual := ParseCoinsAmount(tc.coins, tc.denom)
		if !actual.Equal(tc.expectedRet) {
			log.Panicf("%v expect %v, got %v", msg, tc.expectedRet, actual)
		}
	}
}

func TestParseDecCoinsAmount(t *testing.T) {
	tcs := []struct {
		name        string
		coins       sdk.DecCoins
		denom       string
		expectedRet sdk.Dec
	}{
		{
			name: "all coins with same base denom - 00",
			coins: []sdk.DecCoin{
				{
					Denom:  BaseDenom,
					Amount: sdk.NewDec(1),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.NewDec(10),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.NewDec(100),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.NewDec(1000),
				},
			},
			denom:       BaseDenom,
			expectedRet: sdk.NewDec(1111),
		},
		{
			name: "all coins with same base denom - 01",
			coins: []sdk.DecCoin{
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000.5"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("1000000000000000000"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("10000000000000000000"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("10000000000000000"),
				},
			},
			denom:       Denom,
			expectedRet: sdk.MustNewDecFromStr("11.11"),
		},
		{
			name: "all coins with same denom - 00",
			coins: []sdk.DecCoin{
				{
					Denom:  Denom,
					Amount: sdk.NewDecWithPrec(1, 1),
				},
				{
					Denom:  Denom,
					Amount: sdk.NewDecWithPrec(10, 0),
				},
				{
					Denom:  Denom,
					Amount: sdk.NewDecWithPrec(100, 0),
				},
				{
					Denom:  Denom,
					Amount: sdk.NewDecWithPrec(1, 2),
				},
			},
			denom:       BaseDenom,
			expectedRet: sdk.MustNewDecFromStr("110110000000000000000"),
		},
		{
			name: "all coins with same denom - 01",
			coins: []sdk.DecCoin{
				{
					Denom:  Denom,
					Amount: sdk.NewDecWithPrec(1, 1),
				},
				{
					Denom:  Denom,
					Amount: sdk.NewDecWithPrec(10, 0),
				},
				{
					Denom:  Denom,
					Amount: sdk.NewDecWithPrec(100, 0),
				},
				{
					Denom:  Denom,
					Amount: sdk.NewDecWithPrec(1, 2),
				},
			},
			denom:       Denom,
			expectedRet: sdk.NewDecWithPrec(11011, 2),
		},
		{
			name: "all coins with mixed denoms - 00",
			coins: []sdk.DecCoin{
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000"),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("100"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000000"),
				},
			},
			denom:       Denom,
			expectedRet: sdk.MustNewDecFromStr("201.1"),
		},
		{
			name: "all coins with mixed denoms - 01",
			coins: []sdk.DecCoin{
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000"),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("100"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000000"),
				},
			},
			denom:       BaseDenom,
			expectedRet: sdk.MustNewDecFromStr("201100000000000000000"),
		},
		{
			name: "all coins with mixed denoms and unsupported denoms - 00",
			coins: []sdk.DecCoin{
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("1"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000"),
				},
				{
					Denom:  Denom,
					Amount: sdk.MustNewDecFromStr("100"),
				},
				{
					Denom:  BaseDenom,
					Amount: sdk.MustNewDecFromStr("100000000000000000000"),
				},
				{
					Denom:  "test-denom-1",
					Amount: sdk.MustNewDecFromStr("100000000000000000000"),
				},
				{
					Denom:  "test-denom-2",
					Amount: sdk.MustNewDecFromStr("12348734133219043289"),
				},
				{
					Denom:  "test-denom-3",
					Amount: sdk.MustNewDecFromStr("123248232314908328742340"),
				},
			},
			denom:       BaseDenom,
			expectedRet: sdk.MustNewDecFromStr("201100000000000000000"),
		},
	}

	for _, tc := range tcs {
		msg := fmt.Sprintf("[tc: %v]", tc.name)

		actual := ParseDecCoinsAmount(tc.coins, tc.denom)
		if !actual.Equal(tc.expectedRet) {
			log.Panicf("%v expect %v, got %v", msg, tc.expectedRet, actual)
		}
	}
}
