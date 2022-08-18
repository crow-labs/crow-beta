package keeper

import (
	"context"

	"crow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CancelEscrow(goCtx context.Context, msg *types.MsgCancelEscrow) (*types.MsgCancelEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCancelEscrowResponse{}, nil
}
