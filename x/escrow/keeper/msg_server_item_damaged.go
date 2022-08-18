package keeper

import (
	"context"

	"crow/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ItemDamaged(goCtx context.Context, msg *types.MsgItemDamaged) (*types.MsgItemDamagedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgItemDamagedResponse{}, nil
}
