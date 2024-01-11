package action001

import (
	"github.com/hnlxhzw/go-behave/core"
)

// Succeed returns a new succeed node, which always succeeds in one tick.
func Succeed2() core.Node {
	base := core.NewLeaf("Succeed")
	return &succeed2{Leaf: base}
}

// succeed ...
type succeed2 struct {
	*core.Leaf
}

// Enter ...
func (a *succeed2) Enter(ctx *core.Context) {}

// Tick ...
func (a *succeed2) Tick(ctx *core.Context) core.Status {
	return core.StatusSuccess
}

// Leave ...
func (a *succeed2) Leave(ctx *core.Context) {}
