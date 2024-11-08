package bytedance

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历二叉树
func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorderTraversalHelper(root, &res)
	return res
}

func inorderTraversalHelper(root *TreeNode, res *[]int) {
	if root != nil {
		inorderTraversalHelper(root.Left, res)
		*res = append(*res, root.Val)
		inorderTraversalHelper(root.Right, res)
	}
}
