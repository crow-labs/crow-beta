package keeper_test

import (
	"testing"

	testkeeper "crow/testutil/keeper"
	"crow/x/whitelist/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WhitelistKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
