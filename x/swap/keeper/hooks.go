package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/elysium-station/black/x/swap/types"
)

// Implements SwapHooks interface
var _ types.SwapHooks = Keeper{}

// AfterPoolDepositCreated - call hook if registered
func (k Keeper) AfterPoolDepositCreated(ctx sdk.Context, poolID string, depositor sdk.AccAddress, sharesOwned sdkmath.Int) {
	if k.hooks != nil {
		k.hooks.AfterPoolDepositCreated(ctx, poolID, depositor, sharesOwned)
	}
}

// BeforePoolDepositModified - call hook if registered
func (k Keeper) BeforePoolDepositModified(ctx sdk.Context, poolID string, depositor sdk.AccAddress, sharesOwned sdkmath.Int) {
	if k.hooks != nil {
		k.hooks.BeforePoolDepositModified(ctx, poolID, depositor, sharesOwned)
	}
}
