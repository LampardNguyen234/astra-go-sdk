package account

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
)

func TestNewPrivateKeyFromString(t *testing.T) {
	testCases := []struct {
		name          string
		privateKeyStr string
		err           error
	}{
		{
			"empty private key",
			"",
			ErrInvalidPrivateKey,
		},
		{
			"not a hex-string private key",
			"this is a non-hex string",
			ErrInvalidPrivateKey,
		},
		{
			"private key with less than 32 bytes",
			hex.EncodeToString(rand.Bytes(31)),
			ErrInvalidPrivateKey,
		},
		{
			"private key with more than 32 bytes",
			hex.EncodeToString(rand.Bytes(33)),
			ErrInvalidPrivateKey,
		},
		{
			"valid private key - 00",
			"c32b5551dd1acd26a17b53c1a67840961e6948f7ffd92fe6c38a5e5fd64a3893",
			nil,
		},
		{
			"valid private key - 01",
			"e6da0804ca8f25f4e0269109b054f3afefaf2585b54321a8890e3f744854f6c5",
			nil,
		},
	}

	for _, tc := range testCases {
		_, err := NewKeyInfoFromPrivateKey(tc.privateKeyStr)
		if err != nil {
			assert.ErrorContains(t, err, tc.err.Error(), "tc: %v", tc.name)
		} else {
			assert.NoError(t, err, "tc: %v", tc.name)
		}

	}
}
