package bytedance

import (
	"container/list"
	"fmt"
)

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	var result [][]int

	for stack.Len() != 0 {
		//levelSize := stack.Len()
		level := make([]int, 0)

		//for i := 0; i < levelSize; i++ {
		node := stack.Front()
		fmt.Print(node.Value.(*TreeNode).Val)
		level = append(level, node.Value.(*TreeNode).Val)
		stack.Remove(node)
		if node.Value.(*TreeNode).Left != nil {
			stack.PushBack(node.Value.(*TreeNode).Left)
		}
		if node.Value.(*TreeNode).Right != nil {
			stack.PushBack(node.Value.(*TreeNode).Right)
		}
		//}

		result = append(result, level)

	}
	return result
}

func zigZagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	var result [][]int
	director := true

	for stack.Len() != 0 {
		levelSize := stack.Len()
		level := make([]int, 0)

		for i := 0; i < levelSize; i++ {
			node := stack.Front()
			fmt.Print(node.Value.(*TreeNode).Val)
			level = append(level, node.Value.(*TreeNode).Val)
			stack.Remove(node)
			if !director {
				if node.Value.(*TreeNode).Left != nil {
					stack.PushBack(node.Value.(*TreeNode).Left)
				}
				if node.Value.(*TreeNode).Right != nil {
					stack.PushBack(node.Value.(*TreeNode).Right)
				}
			} else {
				if node.Value.(*TreeNode).Right != nil {
					stack.PushBack(node.Value.(*TreeNode).Right)
				}
				if node.Value.(*TreeNode).Left != nil {
					stack.PushBack(node.Value.(*TreeNode).Left)
				}
			}

		}
		director = !director
		result = append(result, level)
	}
	return result
}

func LevelOrder2(root *TreeNode) [][]int {
	var res [][]int
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			if top != nil {
				level = append(level, top.Val)
				if top.Left != nil {
					queue = append(queue, top.Left)
				}
				if top.Right != nil {
					queue = append(queue, top.Right)
				}
			}
		}
		res = append(res, level)
	}

	return res

}
