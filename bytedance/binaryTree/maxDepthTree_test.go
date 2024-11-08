package bytedance

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	fmt.Println(maxDepth(root))
}
