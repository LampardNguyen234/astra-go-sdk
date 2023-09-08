package account

import (
	_ "github.com/LampardNguyen234/astra-go-sdk/common"
	"github.com/LampardNguyen234/astra-go-sdk/wallet"
	"github.com/pkg/errors"
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

	// ValAddress is the validator version of the address.
	ValAddress string `json:"ValAddress"`
}

// NewKeyInfoFromMnemonic returns a KeyInfo from the given mnemonic and index.
func NewKeyInfoFromMnemonic(mnemonic string, idx int) (*KeyInfo, error) {
	w, err := wallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, errors.Wrap(err, "failed to recover wallet from mnemonic")
	}

	account, err := w.DeriveAtIndex(idx, false)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to derive account at index %v", idx)
	}
	privKey, err := w.PrivateKeyHex(account)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve private key")
	}

	return NewKeyInfoFromPrivateKey(privKey)
}

// RandKeyInfo generates a random KeyInfo.
func RandKeyInfo() (mnemonic string, key *KeyInfo, err error) {
	mnemonic, err = wallet.NewMnemonic(128)
	if err != nil {
		return
	}

	key, err = NewKeyInfoFromMnemonic(mnemonic, 0)
	if err != nil {
		return
	}

	return
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
		ValAddress:    privateKey.ValAddress().String(),
	}, nil
}
