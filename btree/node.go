package btree

import (
	"github.com/rrylee/go-algorithm/container"
	"github.com/rrylee/go-algorithm/vector"
)

type Node struct {
	Parent   *Node
	Keys     vector.Vector
	Children vector.Vector
}

func NewEmptyNode() *Node {
	node := &Node{
		Parent:   nil,
		Keys:     vector.NewVector(0),
		Children: vector.NewVector(0),
	}
	var nilNode *Node
	node.Children.Insert(0, nilNode)
	return node
}

func NewNodeWithTwoChild(key container.Item, c1, c2 *Node) *Node {
	node := &Node{
		Parent:   nil,
		Keys:     vector.NewVector(2),
		Children: vector.NewVector(2),
	}

	node.Keys.Insert(0, key)
	node.Children.Insert(0, c1)
	node.Children.Insert(1, c2)

	if c1 != nil {
		c1.Parent = node
	}
	if c2 != nil {
		c2.Parent = node
	}
	return node
}
