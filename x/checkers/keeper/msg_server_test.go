package keeper_test

import (
	"context"
	"testing"

	"github.com/carmonasl/checkers/x/checkers/keeper"
	"github.com/carmonasl/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := setupKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
