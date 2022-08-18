package cli

import (
	"strconv"

	"crow/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdItemNotReceived() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "item-not-received [escrow-id] [description]",
		Short: "Broadcast message itemNotReceived",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argEscrowId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argDescription := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgItemNotReceived(
				clientCtx.GetFromAddress().String(),
				argEscrowId,
				argDescription,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
