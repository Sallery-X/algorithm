package bytedance

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return helper2(root, math.MinInt, math.MaxInt)
}

func helper2(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val >= max || root.Val <= min {
		return false
	}

	return helper2(root.Left, min, root.Val) && helper2(root.Right, root.Val, max)

}
