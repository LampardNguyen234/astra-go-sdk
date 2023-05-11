package cosmos

import (
	"fmt"
	"testing"
)

func TestCosmosClient_Grants(t *testing.T) {
}

func TestCosmosClient_GranterGrants(t *testing.T) {
	resp, err := c.GranterGrants(addr)
	if err != nil {
		panic(err)
	}
	for _, grant := range resp {
		fmt.Println(grant.Granter, grant.Grantee, grant.Expiration.String(), grant.Authorization)
	}

}

func TestCosmosClient_GranteeGrants(t *testing.T) {
	resp, err := c.GranteeGrants(toAddr)
	if err != nil {
		panic(err)
	}
	for _, grant := range resp {
		fmt.Println(grant.Granter, grant.Grantee, grant.Expiration.String(), grant.Authorization)
	}
}
