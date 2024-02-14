package keeper_test

import (
	"context"
	"testing"

	keepertest "crude/testutil/keeper"
	"crude/testutil/nullify"
	"crude/x/crude/keeper"
	"crude/x/crude/types"

	"github.com/stretchr/testify/require"
)

func createNTransaction(keeper keeper.Keeper, ctx context.Context, n int) []types.Transaction {
	items := make([]types.Transaction, n)
	for i := range items {
		items[i].Id = keeper.AppendTransaction(ctx, items[i])
	}
	return items
}

func TestTransactionGet(t *testing.T) {
	keeper, ctx := keepertest.CrudeKeeper(t)
	items := createNTransaction(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTransaction(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTransactionRemove(t *testing.T) {
	keeper, ctx := keepertest.CrudeKeeper(t)
	items := createNTransaction(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTransaction(ctx, item.Id)
		_, found := keeper.GetTransaction(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTransactionGetAll(t *testing.T) {
	keeper, ctx := keepertest.CrudeKeeper(t)
	items := createNTransaction(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTransaction(ctx)),
	)
}

func TestTransactionCount(t *testing.T) {
	keeper, ctx := keepertest.CrudeKeeper(t)
	items := createNTransaction(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTransactionCount(ctx))
}
