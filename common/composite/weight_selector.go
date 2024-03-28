package composite

import (
	"math/rand"
	"time"

	"github.com/hnlxhzw/go-behave/core"
)

func WeightSelector(children []core.NodeWithWeight) core.Node {
	base := core.NewComposite("weightSelector", getNodes(children))
	return &weightSelector{Composite: base, children: children}
}

// weightSelector ...
type weightSelector struct {
	*core.Composite
	children []core.NodeWithWeight
}

// Enter ...
func (s *weightSelector) Enter(ctx *core.Context) {}

// Tick ...
func (s *weightSelector) Tick(ctx *core.Context) core.Status {
	rand.Seed(time.Now().UnixNano())

	// 计算总权重
	totalWeight := 0
	for _, child := range s.children {
		totalWeight += child.Weight
	}

	// 随机选择一个权重
	chosenWeight := rand.Intn(totalWeight)

	// 选择对应的子节点
	cumulativeWeight := 0
	for _, child := range s.children {
		cumulativeWeight += child.Weight
		if chosenWeight < cumulativeWeight {
			return core.Update(child.Node, ctx)
		}
	}

	// 如果所有权重都为零，则随机选择一个子节点
	index := rand.Intn(len(s.children))
	return core.Update(s.children[index].Node, ctx)
}

// Leave ...
func (s *weightSelector) Leave(ctx *core.Context) {}

// getNodes 从 ChildWithWeight 结构中提取节点
func getNodes(children []core.NodeWithWeight) []core.Node {
	nodes := make([]core.Node, len(children))
	for i, child := range children {
		nodes[i] = child.Node
	}
	return nodes
}
