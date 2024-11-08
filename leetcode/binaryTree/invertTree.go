package bytedance

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 翻转二叉树/ 镜像二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 只考虑一层，交换左右
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	//返回跟节点
	return root
}
