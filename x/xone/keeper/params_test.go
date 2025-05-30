package keeper_test

import (
	"testing"

	testkeeper "github.com/hello-xone/xone-chain/testutil/keeper"
	"github.com/hello-xone/xone-chain/x/xone/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.XoneKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
