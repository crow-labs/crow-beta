package keeper

import (
	"context"

	"crow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ItemShipped(goCtx context.Context, msg *types.MsgItemShipped) (*types.MsgItemShippedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgItemShippedResponse{}, nil
}
