package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgItemReceived = "item_received"

var _ sdk.Msg = &MsgItemReceived{}

func NewMsgItemReceived(userAddress string, escrowId uint64) *MsgItemReceived {
	return &MsgItemReceived{
		UserAddress: userAddress,
		EscrowId:    escrowId,
	}
}

func (msg *MsgItemReceived) Route() string {
	return RouterKey
}

func (msg *MsgItemReceived) Type() string {
	return TypeMsgItemReceived
}

func (msg *MsgItemReceived) GetSigners() []sdk.AccAddress {
	userAddress, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{userAddress}
}

func (msg *MsgItemReceived) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgItemReceived) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid userAddress address (%s)", err)
	}
	return nil
}
