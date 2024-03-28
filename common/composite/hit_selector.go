package composite

import (
	"github.com/hnlxhzw/go-behave/core"
	"math/rand"
	"time"
)

func HitSelector(children ...core.Node) core.Node {
	base := core.NewComposite("HitSelector", children)
	return &hitSelector{Composite: base}
}

// hitSelector ...
type hitSelector struct {
	*core.Composite
	count    int32
	maxCount int32
}

// Enter ...
func (s *hitSelector) Enter(ctx *core.Context) {
	s.maxCount = 100
}

// Tick ...
func (s *hitSelector) Tick(ctx *core.Context) core.Status {
	for {
		if s.count > s.maxCount {
			break
		}
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(s.Children))
		status := core.Update(s.Children[index], ctx)
		if status != core.StatusFailure {
			return status
		}
		s.count++
	}
	return core.StatusFailure
}

// Leave ...
func (s *hitSelector) Leave(ctx *core.Context) {}
