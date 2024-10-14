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

var maxSum = math.MinInt32

func maxPathSum(root *TreeNode) int {
	containCurNodeMax(root)
	return maxSum
}

func containCurNodeMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := math.Max(float64(containCurNodeMax(root.Left)), 0)
	rightMax := math.Max(float64(containCurNodeMax(root.Right)), 0)

	curMax := math.Max(float64(int(leftMax)+root.Val), float64(int(rightMax)+root.Val))

	maxSum = int(math.Max(float64(root.Val+int(rightMax)+int(leftMax)), float64(maxSum)))

	return int(curMax)

}
