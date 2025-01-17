package keeper

import (
	"context"
	"fmt"

	"switcheo/x/switcheo/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateItem(goCtx context.Context, msg *types.MsgCreateItem) (*types.MsgCreateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var item = types.Item{
		Creator:     msg.Creator,
		Name:        msg.Name,
		ProductType: msg.ProductType,
		Amount:      msg.Amount,
		Price:       msg.Price,
		Discounted:  msg.Discounted,
	}

	id := k.AppendItem(
		ctx,
		item,
	)

	return &types.MsgCreateItemResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateItem(goCtx context.Context, msg *types.MsgUpdateItem) (*types.MsgUpdateItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var item = types.Item{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Name:        msg.Name,
		ProductType: msg.ProductType,
		Amount:      msg.Amount,
		Price:       msg.Price,
		Discounted:  msg.Discounted,
	}

	// Checks that the element exists
	val, found := k.GetItem(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetItem(ctx, item)

	return &types.MsgUpdateItemResponse{}, nil
}

func (k msgServer) DeleteItem(goCtx context.Context, msg *types.MsgDeleteItem) (*types.MsgDeleteItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetItem(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveItem(ctx, msg.Id)

	return &types.MsgDeleteItemResponse{}, nil
}
