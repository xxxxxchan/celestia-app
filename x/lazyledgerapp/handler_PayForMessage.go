package lazyledgerapp

import (
    "fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/types"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/keeper"
)

func handleMsgCreatePayForMessage(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePayForMessage) (*sdk.Result, error) {
	k.CreatePayForMessage(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdatePayForMessage(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdatePayForMessage) (*sdk.Result, error) {
	var PayForMessage = types.PayForMessage{
		Creator: msg.Creator,
		Id:      msg.Id,
    	TxType: msg.TxType,
    	TxFee: msg.TxFee,
    	Nonce: msg.Nonce,
    	MsgNamespaceID: msg.MsgNamespaceID,
    	MsgSize: msg.MsgSize,
    	MsgShareCommitment: msg.MsgShareCommitment,
	}

    // Checks that the element exists
    if !k.HasPayForMessage(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
    }

    // Checks if the the msg sender is the same as the current owner
    if msg.Creator != k.GetPayForMessageOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }          

	k.SetPayForMessage(ctx, PayForMessage)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeletePayForMessage(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeletePayForMessage) (*sdk.Result, error) {
    if !k.HasPayForMessage(ctx, msg.Id) {                                                                                                                                                                    
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
    }
    if msg.Creator != k.GetPayForMessageOwner(ctx, msg.Id) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    } 

	k.DeletePayForMessage(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
