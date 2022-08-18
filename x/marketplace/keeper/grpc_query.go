package keeper

import (
	"crow/x/marketplace/types"
)

var _ types.QueryServer = Keeper{}
