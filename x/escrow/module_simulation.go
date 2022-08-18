package escrow

import (
	"math/rand"

	"crow/testutil/sample"
	escrowsimulation "crow/x/escrow/simulation"
	"crow/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = escrowsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgItemReceived = "op_weight_msg_item_received"
	// TODO: Determine the simulation weight value
	defaultWeightMsgItemReceived int = 100

	opWeightMsgItemDamaged = "op_weight_msg_item_damaged"
	// TODO: Determine the simulation weight value
	defaultWeightMsgItemDamaged int = 100

	opWeightMsgItemIncorrect = "op_weight_msg_item_incorrect"
	// TODO: Determine the simulation weight value
	defaultWeightMsgItemIncorrect int = 100

	opWeightMsgItemShipped = "op_weight_msg_item_shipped"
	// TODO: Determine the simulation weight value
	defaultWeightMsgItemShipped int = 100

	opWeightMsgCancelEscrow = "op_weight_msg_cancel_escrow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelEscrow int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	escrowGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&escrowGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgItemReceived int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgItemReceived, &weightMsgItemReceived, nil,
		func(_ *rand.Rand) {
			weightMsgItemReceived = defaultWeightMsgItemReceived
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgItemReceived,
		escrowsimulation.SimulateMsgItemReceived(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgItemDamaged int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgItemDamaged, &weightMsgItemDamaged, nil,
		func(_ *rand.Rand) {
			weightMsgItemDamaged = defaultWeightMsgItemDamaged
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgItemDamaged,
		escrowsimulation.SimulateMsgItemDamaged(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgItemIncorrect int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgItemIncorrect, &weightMsgItemIncorrect, nil,
		func(_ *rand.Rand) {
			weightMsgItemIncorrect = defaultWeightMsgItemIncorrect
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgItemIncorrect,
		escrowsimulation.SimulateMsgItemIncorrect(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgItemShipped int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgItemShipped, &weightMsgItemShipped, nil,
		func(_ *rand.Rand) {
			weightMsgItemShipped = defaultWeightMsgItemShipped
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgItemShipped,
		escrowsimulation.SimulateMsgItemShipped(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelEscrow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelEscrow, &weightMsgCancelEscrow, nil,
		func(_ *rand.Rand) {
			weightMsgCancelEscrow = defaultWeightMsgCancelEscrow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelEscrow,
		escrowsimulation.SimulateMsgCancelEscrow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
