package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "sporechain/testutil/keeper"
	"sporechain/x/sporechain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SporechainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
