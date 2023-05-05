package account

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethermintHd "github.com/evmos/ethermint/crypto/hd"
	"github.com/pkg/errors"
)

type PrivateKey struct {
	types.PrivKey
}

func NewPrivateKeyFromString(privateKeyStr string) (*PrivateKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateKeyStr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidPrivateKey, "not a hex string")
	}
	if len(privateKeyBytes) != privateKeySize {
		return nil, errors.Wrapf(ErrInvalidPrivateKey, "expected key size %v, got %v", privateKeySize, len(privateKeyBytes))
	}

	privateKey := ethermintHd.EthSecp256k1.Generate()(privateKeyBytes)

	return &PrivateKey{privateKey}, nil
}

func (k *PrivateKey) AccAddress() sdk.AccAddress {
	return sdk.AccAddress(k.PubKey().Address())
}

func (k *PrivateKey) PubKey() types.PubKey {
	return k.PrivKey.PubKey()
}

func (k *PrivateKey) HexAddress() string {
	return common.BytesToAddress(k.AccAddress().Bytes()).String()
}
