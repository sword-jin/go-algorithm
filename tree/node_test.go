package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_postorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Right: &TreeNode{
			Val:  2,
			Left:  &TreeNode{
				Val: 3,
			},
		},
	}
	ret := postorderTraversal(root)
	if ret[0] != 3 && ret[1] != 2 && ret[2] != 1 {
		t.FailNow()
	}
	assert.Equal(t, root.Size(), 3)
}

func Test_Something(t *testing.T) {
	root := &TreeNode{}
	pointer := getLeftPointer(root)
	*pointer = &TreeNode{Val: 2}
	fmt.Println(root.Left)
}

func getLeftPointer(node *TreeNode) **TreeNode {
	return &(node.Left)
}
