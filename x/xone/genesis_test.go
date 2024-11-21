package xone_test

import (
	"testing"

	keepertest "github.com/huione-labs/xone/testutil/keeper"
	"github.com/huione-labs/xone/testutil/nullify"
	"github.com/huione-labs/xone/x/xone"
	"github.com/huione-labs/xone/x/xone/types"
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
