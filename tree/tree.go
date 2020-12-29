package tree

import "github.com/rrylee/go-algorithm/container"

type Node interface {
	container.Container
	InsertAsLeftChild(val interface{}) Node
	InsertAsRightChild(val interface{}) Node
	SetParent(parent Node)
	GetParent() Node
	IsRightChild() bool
	IsLeftChild() bool
	GetLeftChild() Node
	GetRightChild() Node
	SetLeftChild(Node)
	SetRightChild(Node)
	Succ() Node
	Traverse(func(tree Node))
	SetValue(interface{})
	Value() interface{}
	UpdateHeight()
	GetHeight() int
	GetTallerChild() Node
	LevelRet() []interface{}
}

type AVL interface {
	Insert(val container.Compare) Node
	Search(val container.Compare) Node
	Delete(val container.Compare) Node
	Root() Node
	Size() int
}
