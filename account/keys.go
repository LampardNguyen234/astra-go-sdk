package account

import (
	"encoding/base64"
	"encoding/hex"
	_ "github.com/LampardNguyen234/astra-go-sdk/common"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethermintHd "github.com/evmos/ethermint/crypto/hd"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

type PrivateKey struct {
	types.PrivKey
	privKeyString string
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

	return &PrivateKey{PrivKey: privateKey, privKeyString: privateKeyStr}, nil
}

func MustNewPrivateKeyFromString(privateKeyStr string) *PrivateKey {
	ret, err := NewPrivateKeyFromString(privateKeyStr)
	if err != nil {
		panic(err)
	}

	return ret
}

func (k PrivateKey) AccAddress() sdk.AccAddress {
	return sdk.AccAddress(k.PubKey().Address())
}

func (k PrivateKey) PubKey() types.PubKey {
	return k.PrivKey.PubKey()
}

func (k PrivateKey) HexAddress() common.Address {
	return common.BytesToAddress(k.AccAddress().Bytes())
}

func (k PrivateKey) ValAddress() sdk.ValAddress {
	return k.AccAddress().Bytes()
}

func (k PrivateKey) String() string {
	return k.privKeyString
}

type ConsensusKey struct {
	PrivateKey crypto.PrivKey `json:"priv_key"`
	PublicKey  crypto.PubKey  `json:"pub_key"`
	Address    crypto.Address `json:"address"`
}

func (k PrivateKey) ConsensusKey() ConsensusKey {
	seed := crypto.Sha256(k.Bytes())
	prvKey := ed25519.GenPrivKeyFromSecret(seed[:])

	return ConsensusKey{
		PrivateKey: prvKey,
		PublicKey:  prvKey.PubKey(),
		Address:    prvKey.PubKey().Address(),
	}
}

func (k ConsensusKey) CosmosPubKey() types.PubKey {
	sdkPK, _ := cryptocodec.FromTmPubKeyInterface(k.PublicKey)
	return sdkPK
}

func NewConsensusKeyFromPrivateConsensusKey(consPrvKeyStr string) (*ConsensusKey, error) {
	prvKeyBytes, err := base64.StdEncoding.DecodeString(consPrvKeyStr)
	if err != nil {
		return nil, err
	}
	prvKey := ed25519.PrivKey(prvKeyBytes)

	return &ConsensusKey{
		PrivateKey: prvKey,
		PublicKey:  prvKey.PubKey(),
		Address:    prvKey.PubKey().Address(),
	}, nil
}

type NodeKey struct {
	PrivateKey crypto.PrivKey `json:"priv_key"`
}

func (k PrivateKey) NodeKey() NodeKey {
	seed := crypto.Sha256(append(k.Bytes(), []byte("NODE_KEY")...))
	prvKey := ed25519.GenPrivKeyFromSecret(seed[:])

	return NodeKey{PrivateKey: prvKey}
}
