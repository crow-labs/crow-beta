package types

import (
	"testing"

	"crow/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgItemShipped_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgItemShipped
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgItemShipped{
				ProducerAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgItemShipped{
				ProducerAddress: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
