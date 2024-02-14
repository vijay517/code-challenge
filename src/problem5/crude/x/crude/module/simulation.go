package crude

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"crude/testutil/sample"
	crudesimulation "crude/x/crude/simulation"
	"crude/x/crude/types"
)

// avoid unused import issue
var (
	_ = crudesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateTransaction = "op_weight_msg_transaction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTransaction int = 100

	opWeightMsgUpdateTransaction = "op_weight_msg_transaction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTransaction int = 100

	opWeightMsgDeleteTransaction = "op_weight_msg_transaction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTransaction int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	crudeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TransactionList: []types.Transaction{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TransactionCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&crudeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTransaction int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTransaction, &weightMsgCreateTransaction, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTransaction = defaultWeightMsgCreateTransaction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTransaction,
		crudesimulation.SimulateMsgCreateTransaction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTransaction int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTransaction, &weightMsgUpdateTransaction, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTransaction = defaultWeightMsgUpdateTransaction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTransaction,
		crudesimulation.SimulateMsgUpdateTransaction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTransaction int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteTransaction, &weightMsgDeleteTransaction, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTransaction = defaultWeightMsgDeleteTransaction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTransaction,
		crudesimulation.SimulateMsgDeleteTransaction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTransaction,
			defaultWeightMsgCreateTransaction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudesimulation.SimulateMsgCreateTransaction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTransaction,
			defaultWeightMsgUpdateTransaction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudesimulation.SimulateMsgUpdateTransaction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTransaction,
			defaultWeightMsgDeleteTransaction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				crudesimulation.SimulateMsgDeleteTransaction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
