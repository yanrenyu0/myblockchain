package myblockchain_test

import (
	"testing"

	keepertest "myblockchain/testutil/keeper"
	"myblockchain/testutil/nullify"
	myblockchain "myblockchain/x/myblockchain/module"
	"myblockchain/x/myblockchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MyblockchainKeeper(t)
	myblockchain.InitGenesis(ctx, k, genesisState)
	got := myblockchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
