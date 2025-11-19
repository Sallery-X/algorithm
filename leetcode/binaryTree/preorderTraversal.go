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
		last := stack[0]
		stack = stack[1:]

		res = append(res, last.Val)

		if last.Left != nil {
			stack = append(stack, last.Left)
		}

		if last.Right != nil {
			stack = append(stack, last.Right)
		}

	}
	return res

}
