package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateListing = "create_listing"

var _ sdk.Msg = &MsgCreateListing{}

func NewMsgCreateListing(creator string, title string, description string, minPrice string, producerId uint64) *MsgCreateListing {
	return &MsgCreateListing{
		Creator:     creator,
		Title:       title,
		Description: description,
		MinPrice:    minPrice,
		ProducerId:  producerId,
	}
}

func (msg *MsgCreateListing) Route() string {
	return RouterKey
}

func (msg *MsgCreateListing) Type() string {
	return TypeMsgCreateListing
}

func (msg *MsgCreateListing) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateListing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateListing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
