package keeper

import (
	"context"

    "sporechain/x/sporechain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) CreateSporechainEntry(goCtx context.Context,  msg *types.MsgCreateSporechainEntry) (*types.MsgCreateSporechainEntryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgCreateSporechainEntryResponse{}, nil
}
