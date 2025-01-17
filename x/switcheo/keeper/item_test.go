package keeper_test

import (
	"context"
	"testing"

	keepertest "switcheo/testutil/keeper"
	"switcheo/testutil/nullify"
	"switcheo/x/switcheo/keeper"
	"switcheo/x/switcheo/types"

	"github.com/stretchr/testify/require"
)

func createNItem(keeper keeper.Keeper, ctx context.Context, n int) []types.Item {
	items := make([]types.Item, n)
	for i := range items {
		items[i].Id = keeper.AppendItem(ctx, items[i])
	}
	return items
}

func TestItemGet(t *testing.T) {
	keeper, ctx := keepertest.switcheoKeeper(t)
	items := createNItem(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetItem(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestItemRemove(t *testing.T) {
	keeper, ctx := keepertest.switcheoKeeper(t)
	items := createNItem(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveItem(ctx, item.Id)
		_, found := keeper.GetItem(ctx, item.Id)
		require.False(t, found)
	}
}

func TestItemGetAll(t *testing.T) {
	keeper, ctx := keepertest.switcheoKeeper(t)
	items := createNItem(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllItem(ctx)),
	)
}

func TestItemCount(t *testing.T) {
	keeper, ctx := keepertest.switcheoKeeper(t)
	items := createNItem(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetItemCount(ctx))
}
