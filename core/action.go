package core

import (
	"fmt"
)

// Leaf is the base type for any specific leaf node (domain-specific).
// Each leaf node has Params: data keys that the implementation imports
// and Returns: data keys that the implementation exports.
type Leaf struct {
	*BaseNode
}

// NewLeaf creates a new leaf base node.
func NewLeaf(name string) *Leaf {
	return &Leaf{
		BaseNode: newBaseNode(CategoryLeaf, name),
	}
}

// GetChildren returns an empty list of Node, since a leaf has no children.
// This method is required for Leaf in order to implement Node.
func (a *Leaf) GetChildren() []Node {
	return []Node{}
}

// String returns a string representation of the leaf node.
func (a *Leaf) String() string {
	return fmt.Sprintf("! %s", a.name)
}
