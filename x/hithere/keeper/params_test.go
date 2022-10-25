package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "hithere/testutil/keeper"
	"hithere/x/hithere/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HithereKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
