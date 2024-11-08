package bytedance

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 是否是平衡二叉树
func isBalanced(root *TreeNode) bool {
	// 出口条件
	if root == nil {
		return true
	}
	// 左右子树差距不超过1
	// 做是否平衡，右是否平衡
	return math.Abs(float64(height(root.Left))-float64(height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// 计算树的高度
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(height(root.Left)), float64(height(root.Right))) + 1)
}
