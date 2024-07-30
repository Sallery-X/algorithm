package bytedance

import "fmt"

// https://leetcode.cn/problems/reverse-linked-list/
type LinkNode struct {
	Val  int
	Next *LinkNode
}

// 迭代法： 每次只反转一个节点

func ReverseLinkedList(head *LinkNode) *LinkNode {
	var pre *LinkNode
	cur := head
	for cur != nil {
		//保存next节点
		temp := cur.Next
		//反转当前节点
		cur.Next = pre

		// pre后移
		pre = cur

		//cur后移
		cur = temp
	}

	fmt.Println(pre.Val, head.Val)
	return pre

}

func ReverseLinkedList2(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseLinkedList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
