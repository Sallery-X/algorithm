package bytedance

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	maxD = 1
	depth(root)
	return maxD

}

var maxD int

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)
	maxD = max(maxD, left+right)

	return max(left, right) + 1
}
