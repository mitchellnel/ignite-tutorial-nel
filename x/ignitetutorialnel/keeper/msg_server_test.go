package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "ignite-tutorial-nel/testutil/keeper"
	"ignite-tutorial-nel/x/ignitetutorialnel/keeper"
	"ignite-tutorial-nel/x/ignitetutorialnel/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IgnitetutorialnelKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
