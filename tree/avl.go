package tree

import "github.com/rrylee/go-algorithm/container"

type avltree struct {
	root Node
	hot  Node //搜索目标节点的父节点
	size int
}

func NewAvlTree() AVL {
	return &avltree{}
}

func (j avltree) Root() Node {
	return j.root
}

func (j avltree) Size() int {
	return j.size
}

func (j *avltree) Insert(val container.Compare) Node {
	x := j.Search(val)
	if !nodeIsNil(x) {
		return x
	}
	x = NewNode(val)
	if j.root == nil {
		j.root = x
	} else {
		// 设置插入节点的父节点
		x.SetParent(j.hot)
		// 设置父节点的left或者right为插入节点
		if val.Compare(j.hot.Value()) == container.CompareGt {
			j.hot.SetRightChild(x)
		} else {
			j.hot.SetLeftChild(x)
		}
	}
	j.size += 1

	node := j.hot
	for !nodeIsNil(node) {
		if !avlBalanced(node) {
			p := node.GetParent()
			isLeft := node.IsLeftChild()
			rotateNode := rotate(node.GetTallerChild().GetTallerChild())

			if nodeIsNil(p) {
				node = rotateNode
				j.root = node
			} else if isLeft {
				p.SetLeftChild(rotateNode)
			} else {
				p.SetRightChild(rotateNode)
			}
			break
		} else {
			node.UpdateHeight()
		}
		node = node.GetParent()
	}

	return x
}

func getParentPointer(node Node) **TreeNode {
	if nodeIsNil(node.GetParent()) {
		n := node.(*TreeNode)
		return &n
	}
	p := node.GetParent().(*TreeNode)
	if node.IsLeftChild() {
		return &(p.Left)
	} else {
		return &(p.Right)
	}
}

func avlBalanced(node Node) bool {
	fac := node.GetLeftChild().GetHeight() - node.GetRightChild().GetHeight()
	return -2 < fac && fac < 2
}

func rotate(node Node) *TreeNode {
	p := node.GetParent()
	g := p.GetParent()
	if p.IsLeftChild() {
		if node.IsLeftChild() {
			p.SetParent(g.GetParent())
			return connect34(node, p, g, node.GetLeftChild(), node.GetRightChild(), p.GetRightChild(), g.GetRightChild())
		} else {
			node.SetParent(g.GetParent())
			return connect34(p, node, g, p.GetLeftChild(), node.GetLeftChild(), node.GetRightChild(), g.GetRightChild())
		}
	} else {
		if node.IsLeftChild() {
			node.SetParent(g.GetParent())
			return connect34(g, node, p, g.GetLeftChild(), node.GetLeftChild(), node.GetRightChild(), p.GetRightChild())
		} else {
			p.SetParent(g.GetParent())
			return connect34(g, p, node, g.GetLeftChild(), p.GetLeftChild(), node.GetLeftChild(), node.GetRightChild())
		}
	}
}

func (j *avltree) Search(val container.Compare) Node {
	j.hot = nil
	if j.root == nil {
		return nil
	}
	return j.SearchIn(j.root, val)
}

func (j *avltree) SearchIn(node Node, val container.Compare) Node {
	if nodeIsNil(node) {
		return nil
	}
	switch node.Value().(container.Compare).Compare(val) {
	case container.CompareEqual:
		return node
	case container.CompareGt:
		j.hot = node
		return j.SearchIn(node.GetLeftChild(), val)
	default: /* CompareLt */
		j.hot = node
		return j.SearchIn(node.GetRightChild(), val)
	}
}

func (j *avltree) Delete(val container.Compare) Node {
	x := j.Search(val)
	if nodeIsNil(x) {
		return nil
	}
	j.RemoveAt(x, j.hot)
	j.size -= 1

	g := j.hot
	for !nodeIsNil(g) {
		if !avlBalanced(g) {
			p := g.GetParent()
			isLeft := g.IsLeftChild()
			rotatedNode := rotate(g.GetTallerChild().GetTallerChild())

			if nodeIsNil(p) {
				g = rotatedNode
				j.root = g
			} else {
				if isLeft {
					p.SetLeftChild(rotatedNode)
				} else {
					p.SetRightChild(rotatedNode)
				}
			}
		}
		g.UpdateHeight()
		g = g.GetParent()
	}

	return x
}

// RemoveAt 返回被删节点的接替者
func (j *avltree) RemoveAt(x Node, parent Node) Node {
	var succ Node
	a := nodeIsNil(x.GetLeftChild())
	b := nodeIsNil(x.GetRightChild())
	if a || b {
		if a {
			succ = x.GetRightChild()
		} else {
			succ = x.GetLeftChild()
		}
		if !nodeIsNil(parent) {
			if x.IsLeftChild() {
				parent.SetLeftChild(succ)
			} else {
				parent.SetRightChild(succ)
			}
		}
	} else {
		// 左右节点都存在
		succ := x.Succ()
		x.SetValue(succ.Value())
		succParent := succ.GetParent()
		if succ.IsLeftChild() {
			succParent.SetLeftChild(succ.GetRightChild())
		} else {
			succParent.SetRightChild(succ.GetRightChild())
		}
		if !nodeIsNil(succ.GetRightChild()) {
			succ.GetRightChild().SetParent(succParent)
		}
		// 重新指向 hot 为删除节点的父节点
		j.hot = succ.GetParent()
		succ = x
	}
	return succ
}

//connect34 构造节点树
/*
		  b
	 a         c
  t1   t2   t3   t4
*/
func connect34(a, b, c Node, t1, t2, t3, t4 Node) *TreeNode {
	a.SetLeftChild(t1)
	if !nodeIsNil(t1) {
		t1.SetParent(a)
	}
	a.SetRightChild(t2)
	if !nodeIsNil(t2) {
		t2.SetParent(a)
	}
	a.UpdateHeight()

	c.SetLeftChild(t3)
	if !nodeIsNil(t3) {
		t3.SetParent(c)
	}
	c.SetRightChild(t4)
	if !nodeIsNil(t4) {
		t4.SetParent(c)
	}
	c.UpdateHeight()

	b.SetLeftChild(a)
	b.SetRightChild(c)
	a.SetParent(b)
	c.SetParent(b)
	b.UpdateHeight()
	return b.(*TreeNode)
}

func nodeIsNil(node Node) bool {
	return node == nil || node.(*TreeNode) == nil
}
