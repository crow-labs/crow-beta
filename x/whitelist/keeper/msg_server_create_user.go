package keeper

import (
	"context"

	"crow/x/whitelist/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateUserResponse{}, nil
}
