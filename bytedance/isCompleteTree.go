package bytedance

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root)
	encounteredNil := false
	for queue.Len() > 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		if node == nil {
			encounteredNil = true
		} else {
			if encounteredNil {
				return false
			}
			queue.PushBack(node.Left)
			queue.PushBack(node.Right)
		}
	}
	return true

}
