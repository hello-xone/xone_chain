//go:build !rocksdb
// +build !rocksdb

package app

import (
	"errors"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// versionDB constant for 'versiondb'
const versionDB = "versiondb"

// setupVersionDB returns error on non-rocksdb build.
// To use versionDB, build with: go build -tags rocksdb
func (app *App) setupVersionDB(
	homePath string,
	keys map[string]*storetypes.KVStoreKey,
	tkeys map[string]*storetypes.TransientStoreKey,
	memKeys map[string]*storetypes.MemoryStoreKey,
) (sdk.MultiStore, error) {
	return nil, errors.New("versiondb is not supported in this build. Build with '-tags rocksdb' to enable versiondb support")
}
