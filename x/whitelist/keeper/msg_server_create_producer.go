package keeper

import (
	"context"

	"crow/x/whitelist/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateProducer(goCtx context.Context, msg *types.MsgCreateProducer) (*types.MsgCreateProducerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateProducerResponse{}, nil
}
