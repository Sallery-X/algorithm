package bytedance

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	node := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}

	p := node.Left.Right
	q := node.Right.Right
	l := lowestCommonAncestor(&node, p, q)
	fmt.Println(l.Val)

}
