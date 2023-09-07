package account

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
)

func TestNewKeyInfoFromPrivateKey(t *testing.T) {
	emptyAccount := KeyInfo{}

	testCases := []struct {
		name           string
		privateKey     string
		expectedErr    error
		expectedResult KeyInfo
	}{
		{
			"empty private key",
			"",
			ErrInvalidPrivateKey,
			emptyAccount,
		},
		{
			"not a hex-string private key",
			"this is a non-hex string",
			ErrInvalidPrivateKey,
			emptyAccount,
		},
		{
			"private key with less than 32 bytes",
			hex.EncodeToString(rand.Bytes(31)),
			ErrInvalidPrivateKey,
			emptyAccount,
		},
		{
			"private key with more than 32 bytes",
			hex.EncodeToString(rand.Bytes(33)),
			ErrInvalidPrivateKey,
			emptyAccount,
		},
		{
			"valid private key - 00",
			"c32b5551dd1acd26a17b53c1a67840961e6948f7ffd92fe6c38a5e5fd64a3893",
			nil,
			KeyInfo{
				PrivateKey:    "c32b5551dd1acd26a17b53c1a67840961e6948f7ffd92fe6c38a5e5fd64a3893",
				CosmosAddress: "astra15c2ma63s5lr0kdpardkvmuns30meh4g3my85ll",
				EthAddress:    "0xa615BeeA30A7c6fb343D1B6CcDf2708BF79Bd511",
				ValAddress:    "astravaloper15c2ma63s5lr0kdpardkvmuns30meh4g37ax9y3",
			},
		},
		{
			"valid private key - 01",
			"e6da0804ca8f25f4e0269109b054f3afefaf2585b54321a8890e3f744854f6c5",
			nil,
			KeyInfo{
				PrivateKey:    "e6da0804ca8f25f4e0269109b054f3afefaf2585b54321a8890e3f744854f6c5",
				CosmosAddress: "astra1d400r9vtegrel9gtag9ekuwzxyrcmawpsuc45a",
				EthAddress:    "0x6D5ef1958Bca079f950BEa0B9b71c231078df5c1",
				ValAddress:    "astravaloper1d400r9vtegrel9gtag9ekuwzxyrcmawp49ey0n",
			},
		},
		{
			"valid private key - 02",
			"d0c712d42c0a2cf24d97fb4bfec99b498c0e50bced770f7b567a6301ff7fa65b",
			nil,
			KeyInfo{
				PrivateKey:    "d0c712d42c0a2cf24d97fb4bfec99b498c0e50bced770f7b567a6301ff7fa65b",
				CosmosAddress: "astra1y7zcfvnvf6g9e93ayj9jhvfy58f676ag20zwsd",
				EthAddress:    "0x278584b26C4E905c963D248b2BB124A1d3aF6bA8",
				ValAddress:    "astravaloper1y7zcfvnvf6g9e93ayj9jhvfy58f676ag0krltr",
			},
		},
	}

	for _, tc := range testCases {
		acc, err := NewKeyInfoFromPrivateKey(tc.privateKey)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error(), "tc: %v", tc.name)
		} else {
			assert.Equal(t, tc.expectedResult, *acc, "tc: %v", tc.name)
		}
	}
}
