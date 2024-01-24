package action

import (
	"github.com/hnlxhzw/go-behave/core"
)

//Work tick 时候执行执行的func
func Work(f func()) core.Node {
	base := core.NewLeaf("work")
	return &work{Leaf: base, cbFunc: f}
}

// work ...
type work struct {
	*core.Leaf
	cbFunc func()
}

// Enter ...
func (a *work) Enter(ctx *core.Context) {}

// Tick ...
func (a *work) Tick(ctx *core.Context) core.Status {
	if a.cbFunc != nil {
		a.cbFunc()
	}
	return core.StatusSuccess
}

// Leave ...
func (a *work) Leave(ctx *core.Context) {}
