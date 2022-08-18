package simulation

import (
	"math/rand"

	"crow/x/escrow/keeper"
	"crow/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgItemShipped(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgItemShipped{
			ProducerAddress: simAccount.Address.String(),
		}

		// TODO: Handling the ItemShipped simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ItemShipped simulation not implemented"), nil, nil
	}
}
