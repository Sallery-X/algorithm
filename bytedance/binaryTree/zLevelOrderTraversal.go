package bytedance

import (
	"container/list"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := list.New()
	res := make([][]int, 0)
	stack.PushBack(root)
	director := true
	for stack.Len() != 0 {
		size := stack.Len()
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := stack.Front()
			level = append(level, node.Value.(*TreeNode).Val)
			stack.Remove(node)
			if !director {
				if node.Value.(*TreeNode).Left != nil {
					stack.PushBack(node.Value.(*TreeNode).Left)
				}
				if node.Value.(*TreeNode).Right != nil {
					stack.PushBack(node.Value.(*TreeNode).Right)
				}
			}
			if director {
				if node.Value.(*TreeNode).Right != nil {
					stack.PushBack(node.Value.(*TreeNode).Right)
				}
				if node.Value.(*TreeNode).Left != nil {
					stack.PushBack(node.Value.(*TreeNode).Left)
				}
			}
		}
		director = !director
		res = append(res, level)
	}

	return res

}
