package simulation

import (
	"math/rand"

	"crow/x/whitelist/keeper"
	"crow/x/whitelist/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateProducer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateProducer{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateProducer simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateProducer simulation not implemented"), nil, nil
	}
}
