package whitelist_test

import (
	"testing"

	keepertest "crow/testutil/keeper"
	"crow/testutil/nullify"
	"crow/x/whitelist"
	"crow/x/whitelist/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WhitelistKeeper(t)
	whitelist.InitGenesis(ctx, *k, genesisState)
	got := whitelist.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
