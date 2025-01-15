package bytedance

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	x := make([]int, 0)
	helper4(root, &x)
	return x
}

func helper4(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	helper4(root.Left, res)
	helper4(root.Right, res)
}
func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, last.Val)

		if last.Right != nil {
			stack = append(stack, last.Right)
		}

		if last.Left != nil {
			stack = append(stack, last.Left)
		}

	}
	return res

}
