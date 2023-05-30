package blackdist

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elysium-station/black/x/blackdist/keeper"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.MintPeriodInflation(ctx)
	if err != nil {
		panic(err)
	}
}
