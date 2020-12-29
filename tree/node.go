package tree

import (
	"github.com/rrylee/go-algorithm/queue"
	stack2 "github.com/rrylee/go-algorithm/stack"
)

type TreeNode struct {
	Val    interface{}
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
	height int
}

func NewNode(val interface{}) Node {
	return &TreeNode{Val: val, height: 1}
}

func (t *TreeNode) SetLeftChild(node Node) {
	t.Left = node.(*TreeNode)
}

func (t *TreeNode) SetRightChild(node Node) {
	t.Right = node.(*TreeNode)
}

func (t *TreeNode) Value() interface{} {
	return t.Val
}

func (t *TreeNode) UpdateHeight() {
	t.height = 1 + max(t.Left.GetHeight(), t.Right.GetHeight())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *TreeNode) LevelRet() []interface{} {
	ret := []interface{}{}
	q := queue.NewListQueue()
	q.Enqueue(t)
	for !q.Empty() {
		i, _ := q.Dequeue()
		d := i.(*TreeNode)
		ret = append(ret, d.Val)
		if d.Left != nil {
			q.Enqueue(d.Left)
		}
		if d.Right != nil {
			q.Enqueue(d.Right)
		}
	}
	return ret
}

func (t *TreeNode) GetHeight() int {
	if t == nil {
		return 0
	}
	return t.height
}

func (t *TreeNode) GetTallerChild() Node {
	if t.GetLeftChild().GetHeight() > t.GetRightChild().GetHeight() {
		return t.GetLeftChild()
	}

	if t.GetLeftChild().GetHeight() < t.GetRightChild().GetHeight() {
		return t.GetRightChild()
	}

	if t.IsLeftChild() {
		return t.GetLeftChild()
	} else {
		return t.GetRightChild()
	}
}

func (t *TreeNode) IsRightChild() bool {
	return t.Parent != nil && t.Parent.Right == t
}

func (t *TreeNode) IsLeftChild() bool {
	return t.Parent != nil && t.Parent.Left == t
}

func (t *TreeNode) SetValue(value interface{}) {
	t.Val = value
}

func (t *TreeNode) Succ() Node {
	if t == nil || t.Right == nil {
		return nil
	}
	succ := t.Right
	for succ.Left != nil {
		succ = succ.Left
	}
	return succ
}

func (t TreeNode) Size() int {
	count := 0
	t.Traverse(func(_ Node) {
		count += 1
	})
	return count
}

func (t TreeNode) Empty() bool {
	return t.Size() == 0
}

func (t *TreeNode) InsertAsLeftChild(val interface{}) Node {
	if t.Size() == 0 {
		return &TreeNode{Val: val}
	}

	node := &TreeNode{Val: val}

	return node
}

func (t *TreeNode) InsertAsRightChild(val interface{}) Node {
	if t.Size() == 0 {
		return &TreeNode{Val: val}
	}

	node := &TreeNode{Val: val}

	return node
}

func (t *TreeNode) SetParent(parent Node) {
	t.Parent = parent.(*TreeNode)
}

func (t *TreeNode) GetParent() Node {
	return t.Parent
}

func (t TreeNode) GetLeftChild() Node {
	return t.Left
}

func (t TreeNode) GetRightChild() Node {
	return t.Right
}

func (t *TreeNode) Traverse(f func(tree Node)) {
	q := queue.NewListQueue()
	q.Enqueue(t)

	for !q.Empty() {
		nodeI, _ := q.Dequeue()
		node := nodeI.(*TreeNode)
		f(node)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}
}

func postorderTraversal(root *TreeNode) []int {
	ret := []int{}
	stack := stack2.NewArrayStack()
	for {
		leftBranch(root, stack)
		if stack.Empty() {
			break
		}

		//pop
		nodeI, _ := stack.Pop()
		node := nodeI.(*TreeNode)

		if next, ok := stack.Peek(); node.Right != nil && ok && node.Right == next.(*TreeNode) {
			stack.Pop()
			stack.Push(node)
			root = node.Right
		} else {
			ret = append(ret, node.Val.(int))
			root = nil
		}
	}
	return ret
}

func leftBranch(node *TreeNode, stack stack2.Stack) {
	for node != nil {
		if node.Right != nil {
			stack.Push(node.Right)
		}
		stack.Push(node)
		node = node.Left
	}
}
