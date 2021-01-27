package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(&MsgCreatePayForMessage{}, "lazyledgerapp/CreatePayForMessage", nil)
cdc.RegisterConcrete(&MsgUpdatePayForMessage{}, "lazyledgerapp/UpdatePayForMessage", nil)
cdc.RegisterConcrete(&MsgDeletePayForMessage{}, "lazyledgerapp/DeletePayForMessage", nil)

} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgCreatePayForMessage{},
	&MsgUpdatePayForMessage{},
	&MsgDeletePayForMessage{},
)
}

var (
	amino = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
