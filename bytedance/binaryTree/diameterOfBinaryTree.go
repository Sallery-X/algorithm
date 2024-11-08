package bytedance

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	maxD := 0
	depth(root, &maxD)
	return maxD

}

func depth(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}

	left := depth(root.Left, m)
	right := depth(root.Right, m)

	//二叉树最大直径
	*m = max(*m, left+right)

	//当前树高度
	return max(left, right) + 1
}
