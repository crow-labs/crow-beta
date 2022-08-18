package keeper

import (
	"context"

	"crow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ItemReceived(goCtx context.Context, msg *types.MsgItemReceived) (*types.MsgItemReceivedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgItemReceivedResponse{}, nil
}
