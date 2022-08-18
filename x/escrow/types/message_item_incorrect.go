package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgItemIncorrect = "item_incorrect"

var _ sdk.Msg = &MsgItemIncorrect{}

func NewMsgItemIncorrect(userAddress string, escrowId string, description string) *MsgItemIncorrect {
	return &MsgItemIncorrect{
		UserAddress: userAddress,
		EscrowId:    escrowId,
		Description: description,
	}
}

func (msg *MsgItemIncorrect) Route() string {
	return RouterKey
}

func (msg *MsgItemIncorrect) Type() string {
	return TypeMsgItemIncorrect
}

func (msg *MsgItemIncorrect) GetSigners() []sdk.AccAddress {
	userAddress, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{userAddress}
}

func (msg *MsgItemIncorrect) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgItemIncorrect) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid userAddress address (%s)", err)
	}
	return nil
}
