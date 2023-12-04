package action

import (
	"github.com/hnlxhzw/go-behave/core"
)

// Fail returns a new fail node, which always fails in one tick.
func Fail() core.Node {
	base := core.NewLeaf("Fail")
	return &fail{Leaf: base}
}

// fail ...
type fail struct {
	*core.Leaf
}

// Enter ...
func (a *fail) Enter(ctx *core.Context) {}

// Tick ...
func (a *fail) Tick(ctx *core.Context) core.Status {
	return core.StatusFailure
}

// Leave ...
func (a *fail) Leave(ctx *core.Context) {}
