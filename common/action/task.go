package action

import (
	"github.com/hnlxhzw/go-behave/core"
)

//Task、tick 时候执行执行的func
func Task(f func() bool) core.Node {
	base := core.NewLeaf("task")
	return &task{Leaf: base, cbFunc: f}
}

// task ...
type task struct {
	*core.Leaf
	cbFunc func() bool
}

// Enter ...
func (a *task) Enter(ctx *core.Context) {}

// Tick ...
func (a *task) Tick(ctx *core.Context) core.Status {
	if a.cbFunc != nil {
		if !a.cbFunc() {
			return core.StatusRunning
		}
	}
	return core.StatusSuccess
}

// Leave ...
func (a *task) Leave(ctx *core.Context) {}
