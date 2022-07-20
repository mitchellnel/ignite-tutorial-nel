package ignitetutorialnel_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ignite-tutorial-nel/testutil/keeper"
	"ignite-tutorial-nel/testutil/nullify"
	"ignite-tutorial-nel/x/ignitetutorialnel"
	"ignite-tutorial-nel/x/ignitetutorialnel/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitetutorialnelKeeper(t)
	ignitetutorialnel.InitGenesis(ctx, *k, genesisState)
	got := ignitetutorialnel.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
