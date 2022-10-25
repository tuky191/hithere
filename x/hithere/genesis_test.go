package hithere_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hithere/testutil/keeper"
	"hithere/testutil/nullify"
	"hithere/x/hithere"
	"hithere/x/hithere/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HithereKeeper(t)
	hithere.InitGenesis(ctx, *k, genesisState)
	got := hithere.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
