package keeper

import (
	"context"
	"fmt"

	"crude/x/crude/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTransaction(goCtx context.Context, msg *types.MsgCreateTransaction) (*types.MsgCreateTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var transaction = types.Transaction{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Remarks: msg.Remarks,
	}

	id := k.AppendTransaction(
		ctx,
		transaction,
	)

	return &types.MsgCreateTransactionResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTransaction(goCtx context.Context, msg *types.MsgUpdateTransaction) (*types.MsgUpdateTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var transaction = types.Transaction{
		Creator: msg.Creator,
		Id:      msg.Id,
		Amount:  msg.Amount,
		Remarks: msg.Remarks,
	}

	// Checks that the element exists
	val, found := k.GetTransaction(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetTransaction(ctx, transaction)

	return &types.MsgUpdateTransactionResponse{}, nil
}

func (k msgServer) DeleteTransaction(goCtx context.Context, msg *types.MsgDeleteTransaction) (*types.MsgDeleteTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetTransaction(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTransaction(ctx, msg.Id)

	return &types.MsgDeleteTransactionResponse{}, nil
}
