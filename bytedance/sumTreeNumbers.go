package bytedance

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return dfs3(root, 0)
}

func dfs3(root *TreeNode, currentSum int) int {
	if root == nil {
		return 0
	}

	currentSum = currentSum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return currentSum
	}

	leftSum := dfs3(root.Left, currentSum)
	rightSum := dfs3(root.Right, currentSum)

	return leftSum + rightSum

}
