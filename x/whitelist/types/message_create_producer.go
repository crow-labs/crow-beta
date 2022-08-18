package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateProducer = "create_producer"

var _ sdk.Msg = &MsgCreateProducer{}

func NewMsgCreateProducer(creator string, name string) *MsgCreateProducer {
	return &MsgCreateProducer{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgCreateProducer) Route() string {
	return RouterKey
}

func (msg *MsgCreateProducer) Type() string {
	return TypeMsgCreateProducer
}

func (msg *MsgCreateProducer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProducer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProducer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
