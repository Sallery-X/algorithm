package bytedance

import "container/list"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New()
	res := make([]int, 0)
	stack.PushBack(root)

	for stack.Len() != 0 {
		levelSize := stack.Len()
		for i := 0; i < levelSize; i++ {
			node := stack.Front()
			if i == levelSize-1 {
				res = append(res, node.Value.(*TreeNode).Val)
			}
			stack.Remove(node)
			if node.Value.(*TreeNode).Left != nil {
				stack.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil {
				stack.PushBack(node.Value.(*TreeNode).Right)
			}

		}
	}
	return res
}
