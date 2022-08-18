package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgItemDamaged = "item_damaged"

var _ sdk.Msg = &MsgItemDamaged{}

func NewMsgItemDamaged(userAddress string, escrowId uint64, description string) *MsgItemDamaged {
	return &MsgItemDamaged{
		UserAddress: userAddress,
		EscrowId:    escrowId,
		Description: description,
	}
}

func (msg *MsgItemDamaged) Route() string {
	return RouterKey
}

func (msg *MsgItemDamaged) Type() string {
	return TypeMsgItemDamaged
}

func (msg *MsgItemDamaged) GetSigners() []sdk.AccAddress {
	userAddress, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{userAddress}
}

func (msg *MsgItemDamaged) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgItemDamaged) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid userAddress address (%s)", err)
	}
	return nil
}
