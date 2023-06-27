package client

import (
	"context"
	"fmt"
	"github.com/AstraProtocol/astra/v2/app"
	"github.com/LampardNguyen234/astra-go-sdk/common"
	_ "github.com/LampardNguyen234/astra-go-sdk/common"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/evmos/ethermint/encoding"
	"github.com/tendermint/tendermint/rpc/client/http"
	"google.golang.org/grpc"
)

type BaseClient struct {
	client.Context
	ctx  context.Context
	grpc *grpc.ClientConn
}

type CosmosClient struct {
	*BaseClient
	bank    *BankClient
	auth    *AuthClient
	authz   *AuthzClient
	distr   *DistrClient
	staking *StakingClient
	vesting *VestingClient
	mint    *MintClient
}

// NewCosmosClient creates a new cosmos client.
func NewCosmosClient(cfg CosmosClientConfig) (*CosmosClient, error) {
	if _, err := cfg.IsValid(); err != nil {
		return nil, err
	}

	encCfg := encoding.MakeConfig(app.ModuleBasics)
	rpcHttp, err := http.New(fmt.Sprintf("%v:%v", cfg.Endpoint, cfg.TendermintPort), "/websocket")
	if err != nil {
		return nil, err
	}
	clientCtx := client.Context{}.WithClient(rpcHttp).
		WithCodec(encCfg.Marshaler).
		WithInterfaceRegistry(encCfg.InterfaceRegistry).
		WithTxConfig(encCfg.TxConfig).
		WithLegacyAmino(encCfg.Amino).
		WithChainID(cfg.ChainID).
		WithAccountRetriever(authTypes.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastSync)

	baseClient := &BaseClient{
		Context: clientCtx,
		ctx:     context.Background(),
	}

	common.Init()

	return &CosmosClient{
		BaseClient: baseClient,
		auth:       NewAuthClient(clientCtx),
		authz:      NewAuthzClient(clientCtx),
		bank:       NewBankClient(clientCtx),
		distr:      NewDistrClient(clientCtx),
		staking:    NewStakingClient(clientCtx),
		vesting:    NewVestingClient(clientCtx),
		mint:       NewMintClient(clientCtx),
	}, nil
}
