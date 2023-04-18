package account

import (
	"encoding/hex"
	_ "github.com/LampardNguyen234/astra-go-sdk/common"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	ethermintHd "github.com/evmos/ethermint/crypto/hd"
)

const (
	privateKeySize = 32
)

// KeyInfo holds necessary information of an account on the Astra blockchain.
type KeyInfo struct {
	// PrivateKey is the secret key of the account.
	PrivateKey string `json:"PrivateKey"`

	// CosmosAddress is the cosmos version of the address.
	CosmosAddress string `json:"CosmosAddress"`

	// EthAddress is the Eth version of the address.
	EthAddress string `json:"EthAddress"`
}

// NewKeyInfoFromPrivateKey recovers the KeyInfo from the given private key.
func NewKeyInfoFromPrivateKey(privateKeyStr string) (*KeyInfo, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyStr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidPrivateKey, "not a hex string")
	}
	if len(privateKeyBytes) != privateKeySize {
		return nil, errors.Wrapf(ErrInvalidPrivateKey, "expected key size %v, got %v", privateKeySize, len(privateKeyBytes))
	}

	privateKey := ethermintHd.EthSecp256k1.Generate()(privateKeyBytes)

	publicKey := privateKey.PubKey()
	cosmosAddr := types.AccAddress(publicKey.Address())

	return &KeyInfo{
		PrivateKey:    privateKeyStr,
		CosmosAddress: cosmosAddr.String(),
		EthAddress:    MustCosmosAddressToHexAddress(cosmosAddr.String()),
	}, nil
}
