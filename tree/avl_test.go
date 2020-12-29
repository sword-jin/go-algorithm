package tree

import (
	"testing"

	"github.com/rrylee/go-algorithm/container"
	"github.com/stretchr/testify/assert"
)

type IntValue int

func (i IntValue) Compare(v interface{}) container.CompareRet {
	if i == v.(IntValue) {
		return container.CompareEqual
	} else if i < v.(IntValue) {
		return container.CompareLt
	} else {
		return container.CompareGt
	}
}

func Test_avltree_Search(t *testing.T) {
	tree := &avltree{}
	t.Run("search empty tree", func(t *testing.T) {
		node := tree.Search(IntValue(10))
		assert.Nil(t, node)
	})

	t.Run("search root node", func(t *testing.T) {
		tree.root = &TreeNode{
			Val:    IntValue(10),
			height: 1,
		}

		node := tree.Search(IntValue(10))
		assert.NotNil(t, node)
		assert.Equal(t, node.Value().(IntValue), IntValue(10))
	})

	t.Run("search left child", func(t *testing.T) {
		tree.root = &TreeNode{
			Val:    IntValue(10),
			height: 1,
			Left: &TreeNode{
				Val:    IntValue(8),
				Right:  &TreeNode{Val: IntValue(9)},
				height: 2,
			},
		}

		node := tree.Search(IntValue(9))
		assert.NotNil(t, node)
		assert.Equal(t, node.Value().(IntValue), IntValue(9))

		node = tree.Search(IntValue(7))
		assert.Nil(t, node)
		assert.Equal(t, tree.hot.Value(), IntValue(8))
	})
}

func Test_ASLInsert(t *testing.T) {
	t.Run("insert a root node", func(t *testing.T) {
		avl := NewAvlTree()
		root := avl.Insert(IntValue(10))
		assert.Equal(t, root.Value(), IntValue(10))
	})

	t.Run("insert zig-zag nodes", func(t *testing.T) {
		avl := NewAvlTree()
		root := avl.Insert(IntValue(100))
		node2 := avl.Insert(IntValue(20))
		assert.Equal(t, node2.Value(), IntValue(20))
		assert.Equal(t, root.GetLeftChild().Value(), IntValue(20))
		assert.Equal(t, avl.Size(), 2)

		//再插入一个，发生旋转
		avl.Insert(IntValue(34))
		assert.Equal(t, avl.Size(), 3)
		assert.Equal(t, IntValue(34), avl.Root().Value())
	})

	t.Run("insert zig-zig nodes", func(t *testing.T) {
		avl := NewAvlTree()
		root := avl.Insert(IntValue(100))
		node2 := avl.Insert(IntValue(20))
		assert.Equal(t, node2.Value(), IntValue(20))
		assert.Equal(t, root.GetLeftChild().Value(), IntValue(20))
		assert.Equal(t, avl.Size(), 2)

		//再插入一个，发生旋转
		avl.Insert(IntValue(10))
		assert.Equal(t, avl.Size(), 3)
		assert.Equal(t, IntValue(20), avl.Root().Value())
	})

	t.Run("insert zag-zag nodes", func(t *testing.T) {
		avl := NewAvlTree()
		root := avl.Insert(IntValue(100))
		node2 := avl.Insert(IntValue(200))
		assert.Equal(t, node2.Value(), IntValue(200))
		assert.Equal(t, root.GetRightChild().Value(), IntValue(200))
		assert.Equal(t, avl.Size(), 2)

		//再插入一个，发生旋转
		avl.Insert(IntValue(300))
		assert.Equal(t, avl.Size(), 3)
		assert.Equal(t, IntValue(200), avl.Root().Value())
	})

	t.Run("insert zag-zig nodes", func(t *testing.T) {
		avl := NewAvlTree()
		root := avl.Insert(IntValue(100))
		node2 := avl.Insert(IntValue(200))
		assert.Equal(t, node2.Value(), IntValue(200))
		assert.Equal(t, root.GetRightChild().Value(), IntValue(200))
		assert.Equal(t, avl.Size(), 2)

		//再插入一个，发生旋转
		avl.Insert(IntValue(150))
		assert.Equal(t, avl.Size(), 3)
		assert.Equal(t, IntValue(150), avl.Root().Value())
	})

	t.Run("insert zag-zig nodes", func(t *testing.T) {
		avl := NewAvlTree()
		root := avl.Insert(IntValue(100))
		node2 := avl.Insert(IntValue(200))
		assert.Equal(t, node2.Value(), IntValue(200))
		assert.Equal(t, root.GetRightChild().Value(), IntValue(200))
		assert.Equal(t, avl.Size(), 2)

		//再插入一个，发生旋转
		avl.Insert(IntValue(150))
		assert.Equal(t, avl.Size(), 3)
		assert.Equal(t, IntValue(150), avl.Root().Value())

		avl.Insert(IntValue(201))
		assert.Equal(t, avl.Root().GetHeight(), 3)
		avl.Insert(IntValue(202))
		assert.Equal(t, avl.Root().GetHeight(), 3)
		avl.Insert(IntValue(203))
		assert.Equal(t, avl.Root().GetHeight(), 3)
		assert.Equal(t, avl.Root().Value(), IntValue(201))
		avl.Insert(IntValue(90))
		assert.Equal(t, avl.Root().GetHeight(), 4)
		avl.Insert(IntValue(162))
		assert.Equal(t, avl.Root().GetHeight(), 4)
		assert.Equal(t, avl.Size(), 8)

		assert.Equal(t, avl.Root().LevelRet(), []interface{}{IntValue(201), IntValue(150), IntValue(202), IntValue(100), IntValue(200), IntValue(203), IntValue(90), IntValue(162)})

		// 开始删除
		avl.Delete(IntValue(90))
		assert.Equal(t, avl.Root().GetHeight(), 4)
		assert.Equal(t, avl.Root().LevelRet(), []interface{}{IntValue(201), IntValue(150), IntValue(202), IntValue(100), IntValue(200), IntValue(203), IntValue(162)})
		assert.Equal(t, avl.Size(), 7)
		avl.Delete(IntValue(100))
		assert.Equal(t, []interface{}{IntValue(201), IntValue(162), IntValue(202), IntValue(150), IntValue(200), IntValue(203)}, avl.Root().LevelRet())
		assert.Equal(t, avl.Root().GetHeight(), 3)
		avl.Delete(IntValue(201))
		assert.Equal(t, []interface{}{IntValue(202), IntValue(162), IntValue(203), IntValue(150), IntValue(200)}, avl.Root().LevelRet())

		avl.Delete(IntValue(203))
		assert.Equal(t, avl.Root().Value(), IntValue(162))
		assert.Equal(t, []interface{}{IntValue(162), IntValue(150), IntValue(202), IntValue(200)}, avl.Root().LevelRet())
	})
}
