package bytedance

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	// pre 的左子树是 in的root的索引前面的数； in的左边是索引左边的数
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	// 右边和左边相反
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
