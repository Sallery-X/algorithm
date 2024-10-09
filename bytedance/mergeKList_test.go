package bytedance

import (
	"fmt"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	// 初始化第一个链表
	head1 := &ListNode{Val: 0}
	current := head1
	for i := 1; i < 5; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
	}

	// 初始化第二个链表
	head2 := &ListNode{Val: 5}
	current = head2
	for i := 6; i < 10; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
	}

	// 初始化第三个链表
	head3 := &ListNode{Val: 10}
	current = head3
	for i := 11; i < 15; i++ {
		current.Next = &ListNode{Val: i}
		current = current.Next
	}

	n := []*ListNode{head1, head2, head3}
	r := mergeKLists(n)
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}

}
