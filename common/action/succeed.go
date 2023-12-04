package action

import (
	"github.com/hnlxhzw/go-behave/core"
)

// Succeed returns a new succeed node, which always succeeds in one tick.
func Succeed() core.Node {
	base := core.NewLeaf("Succeed")
	return &succeed{Leaf: base}
}

// succeed ...
type succeed struct {
	*core.Leaf
}

// Enter ...
func (a *succeed) Enter(ctx *core.Context) {}

// Tick ...
func (a *succeed) Tick(ctx *core.Context) core.Status {
	return core.StatusSuccess
}

// Leave ...
func (a *succeed) Leave(ctx *core.Context) {}
