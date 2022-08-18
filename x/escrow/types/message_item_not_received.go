package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgItemNotReceived = "item_not_received"

var _ sdk.Msg = &MsgItemNotReceived{}

func NewMsgItemNotReceived(userAddress string, escrowId uint64, description string) *MsgItemNotReceived {
	return &MsgItemNotReceived{
		UserAddress: userAddress,
		EscrowId:    escrowId,
		Description: description,
	}
}

func (msg *MsgItemNotReceived) Route() string {
	return RouterKey
}

func (msg *MsgItemNotReceived) Type() string {
	return TypeMsgItemNotReceived
}

func (msg *MsgItemNotReceived) GetSigners() []sdk.AccAddress {
	userAddress, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{userAddress}
}

func (msg *MsgItemNotReceived) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgItemNotReceived) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid userAddress address (%s)", err)
	}
	return nil
}
