package keeper

import (
	"context"

	"crow/x/marketplace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateOrderResponse{}, nil
}
