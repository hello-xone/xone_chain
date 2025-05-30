package xone_test

import (
	"testing"

	keepertest "github.com/hello-xone/xone-chain/testutil/keeper"
	"github.com/hello-xone/xone-chain/testutil/nullify"
	"github.com/hello-xone/xone-chain/x/xone"
	"github.com/hello-xone/xone-chain/x/xone/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.XoneKeeper(t)
	xone.InitGenesis(ctx, *k, genesisState)
	got := xone.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
