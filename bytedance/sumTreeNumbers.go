package bytedance

/*
*
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }

输入：root = [1,2,3]
输出：25
解释：
从根到叶子节点路径 1->2 代表数字 12
从根到叶子节点路径 1->3 代表数字 13
因此，数字总和 = 12 + 13 = 25
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
