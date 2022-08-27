package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/onchainengineer/tenacity/testutil/keeper"
	"github.com/onchainengineer/tenacity/x/tenacity/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.tenacityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
