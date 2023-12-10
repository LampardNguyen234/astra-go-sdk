package wallet

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/rand"
	"testing"
)

func TestNewXMRMnemonicFromEntropy(t *testing.T) {
	entropy := rand.Bytes(32)
	mnemonic, err := NewXMRMnemonicFromEntropy(entropy[:])
	if err != nil {
		t.Error(err)
	}
	fmt.Println(mnemonic)
}
