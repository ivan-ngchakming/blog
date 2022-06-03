package keeper

import (
	"github.com/ivan-ngchakming/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
