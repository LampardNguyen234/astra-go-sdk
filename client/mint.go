package client

import (
	"context"
	"fmt"
	mintTypes "github.com/AstraProtocol/astra/v2/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/grpc"
	"time"
)

type MintClient struct {
	mintTypes.QueryClient
}

// NewMintClient creates a new client for interacting with the `mint` module.
func NewMintClient(conn grpc.ClientConn) *MintClient {
	return &MintClient{
		QueryClient: mintTypes.NewQueryClient(conn),
	}
}

type ProvisionInfo struct {
	Params               *mintTypes.Params
	Height               int64
	AnnualProvisions     sdk.Dec
	Inflation            sdk.Dec
	TotalMintedProvision sdk.Int
	BlockProvision       sdk.Int
	Supply               sdk.Int
	BondedRatio          sdk.Dec
	FoundationBalance    sdk.Int
}

func (c *CosmosClient) MintInfo() (*ProvisionInfo, error) {
	type msg struct {
		k int
		v interface{}
	}
	ch := make(chan msg)
	const (
		errKey = iota
		paramsKey
		heightKey
		annualProvisionKey
		inflationKey
		tmpKey
		blockProvisionKey
		supKey
		bondedRatioKey
		fBalKey
	)
	go func() {
		ret, err := c.MintParams()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{paramsKey, ret}
		}
	}()
	go func() {
		ret, err := c.LatestBlockHeight()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{heightKey, ret.Int64()}
		}
	}()
	go func() {
		ret, err := c.AnnualProvisions()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{annualProvisionKey, ret}
		}
	}()
	go func() {
		ret, err := c.Inflation()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{inflationKey, ret}
		}
	}()
	go func() {
		ret, err := c.TotalMintedProvision()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{tmpKey, ret}
		}
	}()
	go func() {
		ret, err := c.BlockProvision()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{blockProvisionKey, ret}
		}
	}()
	go func() {
		ret, err := c.CirculatingSupply()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{supKey, ret}
		}
	}()
	go func() {
		ret, err := c.GetBondedRatio()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{bondedRatioKey, ret}
		}
	}()
	go func() {
		ret, err := c.GetFoundationBalance()
		if err != nil {
			ch <- msg{errKey, err}
		} else {
			ch <- msg{fBalKey, ret.Total}
		}
	}()

	ctx, cancel := context.WithTimeout(c.ctx, 2*time.Second)
	defer cancel()

	ret := new(ProvisionInfo)
	count := 0
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timed-out")
		case tmp := <-ch:
			switch tmp.k {
			case errKey:
				return nil, tmp.v.(error)
			case heightKey:
				ret.Height = tmp.v.(int64)
				count++
			case paramsKey:
				ret.Params = tmp.v.(*mintTypes.Params)
				count++
			case annualProvisionKey:
				ret.AnnualProvisions = tmp.v.(sdk.Dec)
				count++
			case inflationKey:
				ret.Inflation = tmp.v.(sdk.Dec)
				count++
			case tmpKey:
				ret.TotalMintedProvision = tmp.v.(sdk.Int)
				count++
			case blockProvisionKey:
				ret.BlockProvision = tmp.v.(sdk.Int)
				count++
			case supKey:
				ret.Supply = tmp.v.(sdk.Int)
				count++
			case bondedRatioKey:
				ret.BondedRatio = tmp.v.(sdk.Dec)
				count++
			case fBalKey:
				ret.FoundationBalance = tmp.v.(sdk.Int)
				count++
			}
		default:
			if count == 9 {
				return ret, nil
			}
		}
	}

}

// MintParams returns the parameters for the mint module.
func (c *CosmosClient) MintParams() (*mintTypes.Params, error) {
	resp, err := c.mint.Params(c.ctx, &mintTypes.QueryParamsRequest{})
	if err != nil {
		return nil, err
	}

	return &(resp.Params), nil
}

// AnnualProvisions returns the current minting annual provisions.
func (c *CosmosClient) AnnualProvisions() (sdk.Dec, error) {
	resp, err := c.mint.AnnualProvisions(c.ctx, &mintTypes.QueryAnnualProvisionsRequest{})
	if err != nil {
		return sdk.ZeroDec(), err
	}

	return resp.AnnualProvisions, nil
}

// Inflation returns the current inflation.
func (c *CosmosClient) Inflation() (sdk.Dec, error) {
	resp, err := c.mint.Inflation(c.ctx, &mintTypes.QueryInflationRequest{})
	if err != nil {
		return sdk.ZeroDec(), err
	}

	return resp.Inflation, nil
}

// TotalMintedProvision returns the total amount of tokens minted via the `mint` module.
func (c *CosmosClient) TotalMintedProvision() (sdk.Int, error) {
	resp, err := c.mint.TotalMintedProvision(c.ctx, &mintTypes.QueryTotalMintedProvisionRequest{})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return resp.TotalMintedProvision.Amount.TruncateInt(), nil
}

// BlockProvision returns the current block reward amount.
func (c *CosmosClient) BlockProvision() (sdk.Int, error) {
	resp, err := c.mint.BlockProvision(c.ctx, &mintTypes.QueryBlockProvisionRequest{})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return resp.Provision.Amount, nil
}

// CirculatingSupply returns the current circulating supply.
func (c *CosmosClient) CirculatingSupply() (sdk.Int, error) {
	resp, err := c.mint.CirculatingSupply(c.ctx, &mintTypes.QueryCirculatingSupplyRequest{})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	return resp.CirculatingSupply.Amount.TruncateInt(), nil
}

// GetBondedRatio returns the current staking ratio.
func (c *CosmosClient) GetBondedRatio() (sdk.Dec, error) {
	resp, err := c.mint.GetBondedRatio(c.ctx, &mintTypes.QueryBondedRatioRequest{})
	if err != nil {
		return sdk.ZeroDec(), err
	}

	return resp.BondedRatio, nil
}

// GetFoundationBalance returns the balance of the foundation address.
func (c *CosmosClient) GetFoundationBalance() (*AccountBalance, error) {
	p, err := c.MintParams()
	if err != nil {
		return nil, fmt.Errorf("failed to get mint params: %v", err)
	}

	return c.Balance(p.FoundationAddress)
}
