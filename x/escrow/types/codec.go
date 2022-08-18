package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgItemReceived{}, "escrow/ItemReceived", nil)
	cdc.RegisterConcrete(&MsgItemDamaged{}, "escrow/ItemDamaged", nil)
	cdc.RegisterConcrete(&MsgItemIncorrect{}, "escrow/ItemIncorrect", nil)
	cdc.RegisterConcrete(&MsgItemShipped{}, "escrow/ItemShipped", nil)
	cdc.RegisterConcrete(&MsgCancelEscrow{}, "escrow/CancelEscrow", nil)
	cdc.RegisterConcrete(&MsgItemNotReceived{}, "escrow/ItemNotReceived", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgItemReceived{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgItemDamaged{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgItemIncorrect{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgItemShipped{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelEscrow{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgItemNotReceived{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
