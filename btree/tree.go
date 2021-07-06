package btree

import (
	"github.com/rrylee/go-algorithm/container"
)

type BTree interface {
	Search(item container.Item) *Node
	Insert(item container.Item) bool
	Remove(item container.Item)
	Size() int
}

type Key int

var nilNode *Node

func (k Key) Compare(i interface{}) container.CompareRet {
	if i.(Key) == k {
		return container.CompareEqual
	} else if i.(Key) < k {
		return container.CompareGt
	} else {
		return container.CompareLt
	}
}

type bTree struct {
	root  *Node
	hot   *Node
	size  int
	order int //阶数
	midRank int //发生分裂的中间索引位置
}

func NewBTree(order int) BTree {
	return &bTree{
		root:  NewEmptyNode(),
		hot:   nil,
		size:  0,
		order: order,
		midRank: order >> 1,
	}
}

func (b *bTree) Search(item container.Item) *Node {
	v := b.root
	b.hot = nil
	for v != nil {
		// 找到最后一个不大于 item 的位置
		i := v.Keys.Search(item)
		// i == -1，没有找到key
		if i >= 0 && v.Keys.Get(i).(container.Item).Compare(item) == container.CompareEqual {
			return v
		}
		b.hot = v
		v = getNode(v.Children.Get(i + 1))
	}
	return nil
}

func (b *bTree) Size() int {
	return b.size
}

func (b *bTree) Insert(item container.Item) bool {
	v := b.Search(item)
	if v != nil {
		return false
	}
	r := b.hot.Keys.Search(item)
	b.hot.Keys.Insert(r+1, item)
	b.hot.Children.Insert(r+2, nilNode)
	b.size++
	b.solveOverflow(b.hot)
	return true
}

func (b *bTree) solveOverflow(hot *Node) {
	if b.order >= hot.Children.Size() {
		return
	}

	newNode := NewEmptyNode()
	for j := 0; j < b.order - b.midRank - 1; j ++ {
		newNode.Children.Insert(j, hot.Children.Remove(b.midRank + 1))
		newNode.Keys.Insert(j, hot.Keys.Remove(b.midRank + 1))
	}
	newNode.Children.Set(b.order - b.midRank - 1, hot.Children.Remove(b.midRank + 1))
}

func (b bTree) Remove(item container.Item) {
	panic("implement me")
}

func getNode(node interface{}) *Node {
	if v, ok := node.(*Node); ok {
		return v
	}
	return nil
}


  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

type BSTIterator struct {
	stack *[]*TreeNode
	next  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	leftBranch(root, &stack)
	return BSTIterator{
		stack: &stack,
		next:  stackPop(&stack),
	}
}

func stackPop(stack *[]*TreeNode) *TreeNode {
	if len(*stack) > 0 {
		v := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		return v
	}
	return nil
}

func leftBranch(node *TreeNode, stack *[]*TreeNode) {
	for node != nil {
		*stack = append(*stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	v := this.next.Val

	leftBranch(this.next.Right, this.stack)
	this.next = stackPop(this.stack)

	return v
}

func (this *BSTIterator) HasNext() bool {
	return this.next != nil
}

