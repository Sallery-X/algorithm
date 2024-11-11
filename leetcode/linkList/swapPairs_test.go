package linkList

import (
	"fmt"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	// 创建链表 1 -> 2 -> 3 -> 4
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	fmt.Println("Original list:")
	printList(head)

	// 交换节点
	swapped := swapPairs(head)

	fmt.Println("Swapped list:")
	printList(swapped)
}

// 打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
