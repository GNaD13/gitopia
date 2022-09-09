package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/gitopia/gitopia/testutil/keeper"
	"github.com/gitopia/gitopia/testutil/nullify"
	"github.com/gitopia/gitopia/x/gitopia/keeper"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/stretchr/testify/require"
)

func createNTag(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Tag {
	items := make([]types.Tag, n)
	for i := range items {
		items[i].RepositoryId = 0
		items[i].Name = fmt.Sprintf("tag-%d", i)
		items[i].Id = keeper.AppendTag(ctx, items[i])
	}
	return items
}

func TestTagGet(t *testing.T) {
	keeper, ctx := keepertest.GitopiaKeeper(t)
	items := createNTag(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRepositoryTag(ctx, item.RepositoryId, item.Name)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTagRemove(t *testing.T) {
	keeper, ctx := keepertest.GitopiaKeeper(t)
	items := createNTag(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRepositoryTag(ctx, item.RepositoryId, item.Name)
		_, found := keeper.GetRepositoryTag(ctx, item.RepositoryId, item.Name)
		require.False(t, found)
	}
}

func TestTagGetAll(t *testing.T) {
	keeper, ctx := keepertest.GitopiaKeeper(t)
	items := createNTag(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTag(ctx)),
	)
}

func TestRepositoryTagGetAll(t *testing.T) {
	keeper, ctx := keepertest.GitopiaKeeper(t)
	items := createNTag(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRepositoryTag(ctx, 0)),
	)
}

func TestTagCount(t *testing.T) {
	keeper, ctx := keepertest.GitopiaKeeper(t)
	items := createNTag(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTagCount(ctx))
}
