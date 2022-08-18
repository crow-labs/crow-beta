package keeper

import (
	"crow/x/escrow/types"
)

var _ types.QueryServer = Keeper{}
