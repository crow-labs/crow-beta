package keeper

import (
	"crow/x/whitelist/types"
)

var _ types.QueryServer = Keeper{}
