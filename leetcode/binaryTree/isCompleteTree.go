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

//给你一棵二叉树的根节点 root ，请你判断这棵树是否是一棵 完全二叉树 。
//
//在一棵 完全二叉树 中，除了最后一层外，所有层都被完全填满，并且最后一层中的所有节点都尽可能靠左。最后一层（第 h 层）中可以包含 1 到 2h 个节点。

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

func isCompleteTree2(root *TreeNode) bool {
	isNil := false
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top == nil {
			isNil = true
		} else {
			if isNil {
				return false
			}
			queue = append(queue, top.Left)
			queue = append(queue, top.Right)
		}
	}

	return true

}
