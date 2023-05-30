package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elysium-station/black/app"
)

func init() {
	blackConfig := sdk.GetConfig()
	app.SetBech32AddressPrefixes(blackConfig)
	app.SetBip44CoinType(blackConfig)
	blackConfig.Seal()
}
