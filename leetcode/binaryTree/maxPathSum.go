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
// 最大路径和
var maxSum = math.MinInt32

func maxPathSum(root *TreeNode) int {
	containCurNodeMax(root)
	return maxSum
}

func containCurNodeMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 左边最大值
	leftMax := math.Max(float64(containCurNodeMax(root.Left)), 0)

	//右边最大值
	rightMax := math.Max(float64(containCurNodeMax(root.Right)), 0)

	//当前最大值
	curMax := math.Max(float64(int(leftMax)+root.Val), float64(int(rightMax)+root.Val))

	// 总最大值
	maxSum = int(math.Max(float64(root.Val+int(rightMax)+int(leftMax)), float64(maxSum)))

	return int(curMax)

}
