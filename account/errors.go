package account

import "fmt"

var (
	ErrInvalidPrivateKey = fmt.Errorf("invalid private key")

	ErrInvalidCosmosAddress = fmt.Errorf("invalid cosmos address")
)
