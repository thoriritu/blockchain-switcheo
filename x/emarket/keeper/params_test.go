package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "emarket/testutil/keeper"
	"emarket/x/emarket/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.EmarketKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
