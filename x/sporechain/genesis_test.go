package sporechain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "sporechain/testutil/keeper"
	"sporechain/testutil/nullify"
	"sporechain/x/sporechain"
	"sporechain/x/sporechain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SporechainKeeper(t)
	sporechain.InitGenesis(ctx, *k, genesisState)
	got := sporechain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
