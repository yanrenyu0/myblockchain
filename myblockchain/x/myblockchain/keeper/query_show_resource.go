package keeper

import (
	"context"

	"myblockchain/x/myblockchain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
	
)

func (k Keeper) ShowResource(goCtx context.Context, req *types.QueryShowResourceRequest) (*types.QueryShowResourceResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    ctx := sdk.UnwrapSDKContext(goCtx)
    resource, found := k.GetResource(ctx, req.Id)
    if !found {
        return nil, sdkerrors.ErrKeyNotFound
    }

    return &types.QueryShowResourceResponse{Resource: resource}, nil
}
