package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgItemShipped = "item_shipped"

var _ sdk.Msg = &MsgItemShipped{}

func NewMsgItemShipped(producerAddress string, escrowId uint64, description string) *MsgItemShipped {
	return &MsgItemShipped{
		ProducerAddress: producerAddress,
		EscrowId:        escrowId,
		Description:     description,
	}
}

func (msg *MsgItemShipped) Route() string {
	return RouterKey
}

func (msg *MsgItemShipped) Type() string {
	return TypeMsgItemShipped
}

func (msg *MsgItemShipped) GetSigners() []sdk.AccAddress {
	producerAddress, err := sdk.AccAddressFromBech32(msg.ProducerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{producerAddress}
}

func (msg *MsgItemShipped) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgItemShipped) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ProducerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid producerAddress address (%s)", err)
	}
	return nil
}
