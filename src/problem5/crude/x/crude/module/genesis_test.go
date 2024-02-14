package crude_test

import (
	"testing"

	keepertest "crude/testutil/keeper"
	"crude/testutil/nullify"
	crude "crude/x/crude/module"
	"crude/x/crude/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TransactionList: []types.Transaction{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TransactionCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrudeKeeper(t)
	crude.InitGenesis(ctx, k, genesisState)
	got := crude.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TransactionList, got.TransactionList)
	require.Equal(t, genesisState.TransactionCount, got.TransactionCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
