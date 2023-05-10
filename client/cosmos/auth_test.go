package cosmos

import (
	"fmt"
	"testing"
)

func TestCosmosClient_AccountInfo(t *testing.T) {
	addr := "astra1vcf8dwxgxtdqd3cfm0ptpsrrutcayhhex84e5k"

	resp, err := c.AccountInfo(addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.GetAccountNumber(), resp.GetSequence())
}
