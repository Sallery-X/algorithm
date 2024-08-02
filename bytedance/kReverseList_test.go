package bytedance

import (
	"fmt"
	"testing"
)

func TestKReverseLinkedList(t *testing.T) {
	node0 := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	head := realKr(&node0, 2)
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
}
