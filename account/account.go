package account

import (
	_ "github.com/LampardNguyen234/astra-go-sdk/common"
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
	privateKey, err := NewPrivateKeyFromString(privateKeyStr)
	if err != nil {
		return nil, err
	}

	return &KeyInfo{
		PrivateKey:    privateKeyStr,
		CosmosAddress: privateKey.AccAddress().String(),
		EthAddress:    privateKey.HexAddress().String(),
	}, nil
}
