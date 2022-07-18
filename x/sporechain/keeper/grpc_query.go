package keeper

import (
	"sporechain/x/sporechain/types"
)

var _ types.QueryServer = Keeper{}
