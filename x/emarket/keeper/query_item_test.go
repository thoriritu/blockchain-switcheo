package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "emarket/testutil/keeper"
	"emarket/testutil/nullify"
	"emarket/x/emarket/types"
)

func TestItemQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EmarketKeeper(t)
	msgs := createNItem(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetItemRequest
		response *types.QueryGetItemResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetItemRequest{Id: msgs[0].Id},
			response: &types.QueryGetItemResponse{Item: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetItemRequest{Id: msgs[1].Id},
			response: &types.QueryGetItemResponse{Item: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetItemRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Item(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestItemQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EmarketKeeper(t)
	msgs := createNItem(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllItemRequest {
		return &types.QueryAllItemRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ItemAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Item), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Item),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ItemAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Item), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Item),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ItemAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Item),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ItemAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
