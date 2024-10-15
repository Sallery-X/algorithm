package bytedance

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	helper(root, targetSum, &path, &res)
	return res
}

func helper(root *TreeNode, targetSum int, path *[]int, res *[][]int) {
	if root == nil {
		return
	}

	*path = append(*path, root.Val)
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		temp := make([]int, len(*path))
		copy(temp, *path)
		*res = append(*res, temp)
	}
	helper(root.Left, targetSum-root.Val, path, res)
	helper(root.Right, targetSum-root.Val, path, res)
	*path = (*path)[:len(*path)-1]
}
