package client

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestCosmosClient_DecodeTxBytes(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString("CosBCogBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEmgKLGFzdHJhMXNxYXBwN2h0cmFjZmFzMDJkNTM0Z2Z2YWV3a3Y5M3F4d2t2OXF6Eixhc3RyYTE3Zm41eDVuZWZ4Z25wZ3lmZDU5MDMzbDY0N3h0NDRjZWxxMHg5ORoKCgZhYXN0cmESABJ9ClcKTwooL2V0aGVybWludC5jcnlwdG8udjEuZXRoc2VjcDI1NmsxLlB1YktleRIjCiECfVWtUMuEci20+bWCqxPUmZnNy2vSiy3O4GBU9LrtpaoSAgoAGAISIgocCgZhYXN0cmESEjEwMDAwMDAwMDAwMDAwMDAwMBCgjQYaAA==")
	if err != nil {
		panic(err)
	}

	tx, err := c.DecodeTxBytes(data)
	if err != nil {
		panic(err)
	}
	jsonData, _ := c.TxConfig.TxJSONEncoder()(tx)
	fmt.Println(string(jsonData))
}
