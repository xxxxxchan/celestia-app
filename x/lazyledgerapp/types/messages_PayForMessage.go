package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePayForMessage{}

func NewMsgCreatePayForMessage(creator string, TxType uint32, TxFee uint64, Nonce uint64, MsgNamespaceID []byte, MsgSize uint64, MsgShareCommitment []byte) *MsgCreatePayForMessage {
	return &MsgCreatePayForMessage{
		Creator:            creator,
		TxType:             TxType,
		TxFee:              TxFee,
		Nonce:              Nonce,
		MsgNamespaceID:     MsgNamespaceID,
		MsgSize:            MsgSize,
		MsgShareCommitment: MsgShareCommitment,
	}
}

func (msg *MsgCreatePayForMessage) Route() string {
	return RouterKey
}

func (msg *MsgCreatePayForMessage) Type() string {
	return "CreatePayForMessage"
}

func (msg *MsgCreatePayForMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePayForMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePayForMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePayForMessage{}

func NewMsgUpdatePayForMessage(creator string, id string, TxType uint32, TxFee uint64, Nonce uint64, MsgNamespaceID []byte, MsgSize uint64, MsgShareCommitment []byte) *MsgUpdatePayForMessage {
	return &MsgUpdatePayForMessage{
		Id:                 id,
		Creator:            creator,
		TxType:             TxType,
		TxFee:              TxFee,
		Nonce:              Nonce,
		MsgNamespaceID:     MsgNamespaceID,
		MsgSize:            MsgSize,
		MsgShareCommitment: MsgShareCommitment,
	}
}

func (msg *MsgUpdatePayForMessage) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePayForMessage) Type() string {
	return "UpdatePayForMessage"
}

func (msg *MsgUpdatePayForMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePayForMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePayForMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreatePayForMessage{}

func NewMsgDeletePayForMessage(creator string, id string) *MsgDeletePayForMessage {
	return &MsgDeletePayForMessage{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePayForMessage) Route() string {
	return RouterKey
}

func (msg *MsgDeletePayForMessage) Type() string {
	return "DeletePayForMessage"
}

func (msg *MsgDeletePayForMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePayForMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePayForMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
