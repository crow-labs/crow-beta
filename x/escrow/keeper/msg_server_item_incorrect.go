package keeper

import (
	"context"

	"crow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ItemIncorrect(goCtx context.Context, msg *types.MsgItemIncorrect) (*types.MsgItemIncorrectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgItemIncorrectResponse{}, nil
}
