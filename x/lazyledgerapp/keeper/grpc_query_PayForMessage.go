package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PayForMessageAll(c context.Context, req *types.QueryAllPayForMessageRequest) (*types.QueryAllPayForMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var PayForMessages []*types.PayForMessage
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	PayForMessageStore := prefix.NewStore(store, types.KeyPrefix(types.PayForMessageKey))

	pageRes, err := query.Paginate(PayForMessageStore, req.Pagination, func(key []byte, value []byte) error {
		var PayForMessage types.PayForMessage
		if err := k.cdc.UnmarshalBinaryBare(value, &PayForMessage); err != nil {
			return err
		}

		PayForMessages = append(PayForMessages, &PayForMessage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPayForMessageResponse{PayForMessage: PayForMessages, Pagination: pageRes}, nil
}

func (k Keeper) PayForMessage(c context.Context, req *types.QueryGetPayForMessageRequest) (*types.QueryGetPayForMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var PayForMessage types.PayForMessage
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PayForMessageKey + req.Id)), &PayForMessage)

	return &types.QueryGetPayForMessageResponse{PayForMessage: &PayForMessage}, nil
}
