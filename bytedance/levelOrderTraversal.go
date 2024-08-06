package bytedance

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	var result [][]int

	for stack.Len() != 0 {
		levelSize := stack.Len()
		level := make([]int, 0)

		for i := 0; i < levelSize; i++ {
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
		}

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
