package keeper

import (
	"github.com/hello-xone/xone-chain/x/xone/types"
)

var _ types.QueryServer = Keeper{}
