package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ignite-tutorial-nel/testutil/keeper"
	"ignite-tutorial-nel/x/ignitetutorialnel/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IgnitetutorialnelKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
