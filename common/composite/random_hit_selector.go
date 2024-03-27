package composite

import (
	"github.com/hnlxhzw/go-behave/core"
	"math/rand"
	"time"
)

func RandomHitSelector(children ...core.Node) core.Node {
	base := core.NewComposite("RandomHitSelector", children)
	return &randomHitSelector{Composite: base}
}

// randomHitSelector ...
type randomHitSelector struct {
	*core.Composite
	count    int32
	maxCount int32
}

// Enter ...
func (s *randomHitSelector) Enter(ctx *core.Context) {
	s.maxCount = 100
}

// Tick ...
func (s *randomHitSelector) Tick(ctx *core.Context) core.Status {
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
func (s *randomHitSelector) Leave(ctx *core.Context) {}
