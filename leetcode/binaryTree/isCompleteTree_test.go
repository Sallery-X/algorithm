package bytedance

import (
	"fmt"
	"testing"
)

func Test_isCompleteTree(t *testing.T) {
	node := TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Right: nil,
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
		},
	}
	fmt.Println(isCompleteTree2(&node))
}
