package cosmos

import (
	"fmt"
	"testing"
)

func TestCosmosClient_AccountInfo(t *testing.T) {
	resp, err := c.AccountInfo(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.GetAccountNumber(), resp.GetSequence())
}

func TestCosmosClient_CountAccounts(t *testing.T) {
	resp, err := c.TotalAccounts()
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
