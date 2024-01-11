package condition

import (
	"github.com/hnlxhzw/go-behave/core"
)

// Condition creates a new condition node.
func Condition(name string, checkFunc func() bool) core.Node {
	return &condition{
		Leaf:      core.NewLeaf(name),
		checkFunc: checkFunc,
	}
}

// Condition is a simple condition node.
type condition struct {
	*core.Leaf
	checkFunc func() bool
}

// Enter doesn't do anything special for a condition.
func (c *condition) Enter(ctx *core.Context) {}

// Tick checks the condition using the provided check function.
func (c *condition) Tick(ctx *core.Context) core.Status {
	if c.checkFunc() {
		return core.StatusSuccess
	}
	return core.StatusFailure
}

// Leave doesn't do anything special for a condition.
func (c *condition) Leave(ctx *core.Context) {}
