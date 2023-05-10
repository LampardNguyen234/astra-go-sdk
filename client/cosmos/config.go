package cosmos

import "fmt"

type CosmosClientConfig struct {
	Endpoint       string `json:"Endpoint"`
	TendermintPort string `json:"TendermintPort"`
	ChainID        string `json:"ChainID"`
}

// IsValid checks if the current CosmosClientConfig is valid.
func (cfg CosmosClientConfig) IsValid() (bool, error) {
	if cfg.Endpoint == "" {
		return false, fmt.Errorf("empty endpoint")
	}
	if cfg.ChainID == "" {
		return false, fmt.Errorf("empty chainID")
	}
	if cfg.TendermintPort == "" {
		return false, fmt.Errorf("empty TendermintPort")
	}

	return true, nil
}

func DefaultTestnetConfig() CosmosClientConfig {
	return CosmosClientConfig{
		TendermintPort: "26657",
		Endpoint:       "http://206.189.158.191",
		ChainID:        "astra_11115-1",
	}
}

func DefaultMainnetConfig() CosmosClientConfig {
	return CosmosClientConfig{
		TendermintPort: "26657",
		Endpoint:       "https://cosmos.astranaut.io",
		ChainID:        "astra_11110-1",
	}
}
