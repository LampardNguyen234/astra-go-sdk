package account

import (
	_ "github.com/LampardNguyen234/astra-go-sdk/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v6/types"
	"strings"
)

// NewCosmosAddressFromStr creates a new sdk.AccAddress from the given bech32-encoded string.
func NewCosmosAddressFromStr(strAddr string) (sdk.AccAddress, error) {
	return types.GetEvmosAddressFromBech32(strAddr)
}

// MustNewCosmosAddressFromStr is the same as NewCosmosAddressFromStr except that it will panic upon encountering errors.
func MustNewCosmosAddressFromStr(strAddr string) sdk.AccAddress {
	acc, err := NewCosmosAddressFromStr(strAddr)
	if err != nil {
		panic(err)
	}

	return acc
}

// CosmosAddressToHexAddress converts an address in the cosmos version to the hex version.
func CosmosAddressToHexAddress(cosmosAddr string) (string, error) {
	acc, err := NewCosmosAddressFromStr(cosmosAddr)
	if err != nil {
		return "", err
	}

	return common.BytesToAddress(acc.Bytes()).String(), nil
}

// MustCosmosAddressToHexAddress is the same as CosmosAddressToHexAddress except that it will panic upon encountering errors.
func MustCosmosAddressToHexAddress(cosmosAddr string) string {
	addr, err := CosmosAddressToHexAddress(cosmosAddr)
	if err != nil {
		panic(err)
	}

	return addr
}

// HexToCosmosAddress creates a new Cosmos version address from the given hex address.
func HexToCosmosAddress(hexAddr string) (string, error) {
	hexAddr = strings.Replace(hexAddr, "0x", "", -1)
	hexAddr = strings.Replace(hexAddr, "0X", "", -1)
	acc, err := sdk.AccAddressFromHex(hexAddr)
	if err != nil {
		return "", err
	}

	return acc.String(), nil
}

// MustHexToCosmosAddress is the same as HexToCosmosAddress except that it will panic upon encountering errors.
func MustHexToCosmosAddress(hexAddr string) string {
	cosmosAddr, err := HexToCosmosAddress(hexAddr)
	if err != nil {
		panic(err)
	}

	return cosmosAddr
}

// ParseCosmosAddress converts the given address (either in hex format or bech32 format) to a valid sdk.AccAddress.
func ParseCosmosAddress(strAddr string) (sdk.AccAddress, error) {
	cosmosAddr, err := HexToCosmosAddress(strAddr)
	if err == nil {
		return NewCosmosAddressFromStr(cosmosAddr)
	}

	return NewCosmosAddressFromStr(strAddr)
}

// MustParseCosmosAddress converts the given address (either in hex format or bech32 format) to a valid sdk.AccAddress.
// It will panic if any error occurs.
func MustParseCosmosAddress(strAddr string) sdk.AccAddress {
	ret, err := ParseCosmosAddress(strAddr)
	if err != nil {
		panic(err)
	}

	return ret
}

// ParseCosmosValidatorAddress converts the given address to a valid sdk.ValAddress.
func ParseCosmosValidatorAddress(strAddr string) (sdk.ValAddress, error) {
	addr, err := sdk.ValAddressFromHex(strAddr)
	if err == nil {
		return addr, nil
	}

	return sdk.ValAddressFromBech32(strAddr)
}

// MustParseCosmosValidatorAddress converts the given address to a valid sdk.ValAddress.
// It will panic if any error occurs.
func MustParseCosmosValidatorAddress(strAddr string) sdk.ValAddress {
	ret, err := ParseCosmosValidatorAddress(strAddr)
	if err != nil {
		panic(err)
	}

	return ret
}
