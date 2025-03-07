package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dymensionxyz/dymension/v3/x/iro/types"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

type Keeper struct {
	rollapptypes.StubRollappCreatedHooks
	authority string // authority is the x/gov module account

	cdc      codec.BinaryCodec
	storeKey storetypes.Key

	AK types.AccountKeeper
	BK types.BankKeeper
	dk types.DenomMetadataKeeper
	rk types.RollappKeeper
	gk types.GammKeeper
	pm types.PoolManagerKeeper
	ik types.IncentivesKeeper
	tk types.TxFeesKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.Key,
	authority string,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	dk types.DenomMetadataKeeper,
	rk types.RollappKeeper,
	gk types.GammKeeper,
	ik types.IncentivesKeeper,
	pm types.PoolManagerKeeper,
	tk types.TxFeesKeeper,
) *Keeper {
	return &Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		authority: authority,
		AK:        ak,
		BK:        bk,
		dk:        dk,
		rk:        rk,
		gk:        gk,
		ik:        ik,
		pm:        pm,
		tk:        tk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetModuleAccountAddress returns the address of the module account
func (k Keeper) GetModuleAccountAddress() string {
	return k.AK.GetModuleAddress(types.ModuleName).String()
}
