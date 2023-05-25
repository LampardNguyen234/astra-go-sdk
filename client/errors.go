package client

import "fmt"

var (
	ErrAccountNotExisted = fmt.Errorf("account not existed")

	ErrBuildTx     = fmt.Errorf("failed to build transaction")
	ErrSignTx      = fmt.Errorf("failed to sign transaction")
	ErrBroadcastTx = fmt.Errorf("failed to broadcast transaction")

	ErrInsufficientBalance = fmt.Errorf("insufficient balance")

	ErrInvalidAccAddress = fmt.Errorf("invalid AccAddress")
	ErrInvalidValAddress = fmt.Errorf("invalid ValAddress")
)
