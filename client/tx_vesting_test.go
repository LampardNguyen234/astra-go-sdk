package client

import (
	"encoding/hex"
	"fmt"
	"github.com/LampardNguyen234/astra-go-sdk/account"
	"github.com/LampardNguyen234/astra-go-sdk/client/msg_params"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/rand"
	"math/big"
	"testing"
	"time"
)

func TestCosmosClient_TxCreateVesting(t *testing.T) {
	txParams := &msg_params.TxParams{
		PrivateKey: privateKey,
	}

	recipient := account.MustNewPrivateKeyFromString(hex.EncodeToString(crypto.Sha256(rand.Bytes(32))))
	fmt.Printf("recipientKey: %v, %v\n", recipient.String(), recipient.AccAddress().String())

	p := msg_params.TxCreateVestingParams{
		TxParams:        *txParams,
		ToAddr:          recipient.AccAddress().String(),
		Amount:          new(big.Int).SetUint64(testAmt),
		VestingLength:   10,
		VestingDuration: 60000,
	}

	resp, err := c.TxCreateVesting(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(*resp)
}

func TestCosmosClient_TxClawBackVesting(t *testing.T) {
	txParams := &msg_params.TxParams{
		PrivateKey: privateKey,
	}

	recipient := account.MustNewPrivateKeyFromString(hex.EncodeToString(crypto.Sha256(rand.Bytes(32))))
	fmt.Printf("recipientKey: %v\n", recipient.String())

	p := msg_params.TxCreateVestingParams{
		TxParams:        *txParams,
		ToAddr:          recipient.AccAddress().String(),
		Amount:          new(big.Int).SetUint64(testAmt),
		VestingLength:   10,
		VestingDuration: 600,
	}

	resp, err := c.TxCreateVesting(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("txCreate: %v\n", resp.TxHash)

	time.Sleep(10 * time.Second)

	resp, err = c.TxClawBackVesting(msg_params.TxClawBackVestingParams{
		TxParams:    *txParams,
		FunderAddr:  txParams.MustGetPrivateKey().AccAddress().String(),
		AccountAddr: recipient.AccAddress().String(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("txClawBack: %v\n", resp.TxHash)
}
