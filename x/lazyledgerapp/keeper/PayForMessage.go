package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetPayForMessageCount get the total number of PayForMessage
func (k Keeper) GetPayForMessageCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageCountKey))
	byteKey := types.KeyPrefix(types.PayForMessageCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetPayForMessageCount set the total number of PayForMessage
func (k Keeper) SetPayForMessageCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageCountKey))
	byteKey := types.KeyPrefix(types.PayForMessageCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreatePayForMessage creates a PayForMessage with a new id and update the count
func (k Keeper) CreatePayForMessage(ctx sdk.Context, msg types.MsgCreatePayForMessage) {
	// Create the PayForMessage
    count := k.GetPayForMessageCount(ctx)
    var PayForMessage = types.PayForMessage{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        TxType: msg.TxType,
        TxFee: msg.TxFee,
        Nonce: msg.Nonce,
        MsgNamespaceID: msg.MsgNamespaceID,
        MsgSize: msg.MsgSize,
        MsgShareCommitment: msg.MsgShareCommitment,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageKey))
    key := types.KeyPrefix(types.PayForMessageKey + PayForMessage.Id)
    value := k.cdc.MustMarshalBinaryBare(&PayForMessage)
    store.Set(key, value)

    // Update PayForMessage count
    k.SetPayForMessageCount(ctx, count+1)
}

// SetPayForMessage set a specific PayForMessage in the store
func (k Keeper) SetPayForMessage(ctx sdk.Context, PayForMessage types.PayForMessage) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageKey))
	b := k.cdc.MustMarshalBinaryBare(&PayForMessage)
	store.Set(types.KeyPrefix(types.PayForMessageKey + PayForMessage.Id), b)
}

// GetPayForMessage returns a PayForMessage from its id
func (k Keeper) GetPayForMessage(ctx sdk.Context, key string) types.PayForMessage {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageKey))
	var PayForMessage types.PayForMessage
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PayForMessageKey + key)), &PayForMessage)
	return PayForMessage
}

// HasPayForMessage checks if the PayForMessage exists
func (k Keeper) HasPayForMessage(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageKey))
	return store.Has(types.KeyPrefix(types.PayForMessageKey + id))
}

// GetPayForMessageOwner returns the creator of the PayForMessage
func (k Keeper) GetPayForMessageOwner(ctx sdk.Context, key string) string {
    return k.GetPayForMessage(ctx, key).Creator
}

// DeletePayForMessage deletes a PayForMessage
func (k Keeper) DeletePayForMessage(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageKey))
	store.Delete(types.KeyPrefix(types.PayForMessageKey + key))
}

// GetAllPayForMessage returns all PayForMessage
func (k Keeper) GetAllPayForMessage(ctx sdk.Context) (msgs []types.PayForMessage) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PayForMessageKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PayForMessageKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.PayForMessage
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
