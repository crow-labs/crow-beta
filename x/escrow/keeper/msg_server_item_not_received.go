package keeper

import (
	"context"

	"crow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ItemNotReceived(goCtx context.Context, msg *types.MsgItemNotReceived) (*types.MsgItemNotReceivedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgItemNotReceivedResponse{}, nil
}
