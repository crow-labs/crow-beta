package keeper_test

import (
	"context"
	"testing"

	keepertest "crow/testutil/keeper"
	"crow/x/escrow/keeper"
	"crow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EscrowKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
