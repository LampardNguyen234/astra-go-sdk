package account

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCosmosAddressToHexAddress(t *testing.T) {
	testCases := []struct {
		name            string
		cosmosAddr      string
		expectedHexAddr string
		expectedErr     error
	}{
		{
			"empty cosmos addr",
			"",
			"",
			errors.ErrInvalidAddress,
		},
		{
			"invalid cosmos addr",
			"astra1d400r9vtegrel9gtag9ekuwzxyrcmawpsuc45b",
			"",
			errors.ErrInvalidAddress,
		},
		{
			"valid cosmos addr - 00",
			"astra1d400r9vtegrel9gtag9ekuwzxyrcmawpsuc45a",
			"0x6D5ef1958Bca079f950BEa0B9b71c231078df5c1",
			nil,
		},
		{
			"valid cosmos addr - 01",
			"astra1y7zcfvnvf6g9e93ayj9jhvfy58f676ag20zwsd",
			"0x278584b26C4E905c963D248b2BB124A1d3aF6bA8",
			nil,
		},
	}

	for _, tc := range testCases {
		hexAddr, err := CosmosAddressToHexAddress(tc.cosmosAddr)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error(), "tc: %v", tc.name)
		} else {
			assert.Equal(t, tc.expectedHexAddr, hexAddr, "tc: %v", tc.name)
		}
	}
}

func TestMustHexToCosmosAddress(t *testing.T) {
	testCases := []struct {
		name               string
		expectedCosmosAddr string
		hexaddr            string
		expectedErr        error
	}{
		{
			"empty hex addr",
			"",
			"",
			fmt.Errorf("must provide an address"),
		},
		{
			"invalid hex addr",
			"",
			"0x6D5ef1958Bca079f950BEa0B9b71c231078df5cg",
			fmt.Errorf("encoding/hex: invalid byte"),
		},
		{
			"valid hex addr - 00",
			"astra1d400r9vtegrel9gtag9ekuwzxyrcmawpsuc45a",
			"0x6D5ef1958Bca079f950BEa0B9b71c231078df5c1",
			nil,
		},
		{
			"valid hex addr - 01",
			"astra1y7zcfvnvf6g9e93ayj9jhvfy58f676ag20zwsd",
			"0x278584b26C4E905c963D248b2BB124A1d3aF6bA8",
			nil,
		},
	}

	for _, tc := range testCases {
		cosmosAddr, err := HexToCosmosAddress(tc.hexaddr)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error(), "tc: %v", tc.name)
		} else {
			assert.Equal(t, tc.expectedCosmosAddr, cosmosAddr, "tc: %v", tc.name)
		}
	}
}
