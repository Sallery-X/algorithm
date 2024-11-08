package bytedance

import (
	"fmt"
	"testing"
)

func Test_diameterOfBinaryTree(t *testing.T) {
	// 构建一个示例二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	// 计算并打印二叉树的直径
	fmt.Println("Diameter of the binary tree:", diameterOfBinaryTree(root))
}
