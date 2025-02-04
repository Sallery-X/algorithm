package bytedance

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 二叉搜索树，第k小的值是多少
func kthSmallest(root *TreeNode, k int) int {
	var count int
	res := -1

	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil || count > k {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			res = node.Val
			return
		}
		inorder(node.Right)
	}

	inorder(root)
	return res
}
