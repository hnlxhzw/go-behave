package action

import (
	"github.com/hnlxhzw/go-behave/core"
)

// Succeed returns a new succeed node, which always succeeds in one tick.
func Succeed1() core.Node {
	base := core.NewLeaf("Succeed")
	return &succeed1{Leaf: base}
}

// succeed ...
type succeed1 struct {
	*core.Leaf
}

// Enter ...
func (a *succeed1) Enter(ctx *core.Context) {}

// Tick ...
func (a *succeed1) Tick(ctx *core.Context) core.Status {
	return core.StatusSuccess
}

// Leave ...
func (a *succeed1) Leave(ctx *core.Context) {}
