package bytedance

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	node := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	xx := LevelOrder(&node)
	fmt.Println(xx)

}
