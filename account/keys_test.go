package account

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	cmtjson "github.com/tendermint/tendermint/libs/json"
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

func TestPrivateKey_ConsensusPubKey(t *testing.T) {
	k := MustNewPrivateKeyFromString("")

	tmp := k.ConsensusKey()
	jsonBytes, err := cmtjson.MarshalIndent(tmp, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes))

	nodeKey := k.NodeKey()
	jsonBytes, err = cmtjson.MarshalIndent(nodeKey, "", "  ")
	fmt.Println(string(jsonBytes))

	jsonBytes, err = cmtjson.MarshalIndent(k.ConsensusKey().CosmosPubKey(), "", "  ")
	fmt.Println(string(jsonBytes))
	fmt.Println(k.ConsensusKey().CosmosPubKey().Type(), fmt.Sprintf("%x", k.ConsensusKey().CosmosPubKey().Bytes()))

	//tmpPrvKey, err := NewConsensusKeyFromPrivateConsensusKey("hKwznM43l8BDkN9Qm9DAj/irzOITjt4u5RFSe7Wif2VjiA3lyLO/AMgy+hqi1OdIzcqhLo+XVDDbptO3iEOTDg==")
	//if err != nil {
	//	panic(err)
	//}
	//jsonBytes, err = cmtjson.MarshalIndent(tmpPrvKey, "", "  ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(jsonBytes))
}
