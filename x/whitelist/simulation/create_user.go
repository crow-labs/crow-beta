package simulation

import (
	"math/rand"

	"crow/x/whitelist/keeper"
	"crow/x/whitelist/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateUser(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateUser{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateUser simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateUser simulation not implemented"), nil, nil
	}
}
