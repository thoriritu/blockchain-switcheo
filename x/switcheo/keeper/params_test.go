package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "switcheo/testutil/keeper"
	"switcheo/x/switcheo/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.switcheoKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
