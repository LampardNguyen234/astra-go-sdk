package cosmos

import (
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountInfoI interface {
	GetAddress() sdk.AccAddress

	GetPubKey() types.PubKey // can return nil.

	GetAccountNumber() uint64

	GetSequence() uint64

	String() string
}

// TxSenderAccountI specifies required methods for checking the sender account.
type TxSenderAccountI interface {
	AccountExists(address string) error
	AccountInfo(addr string) (AccountInfoI, error)
}
