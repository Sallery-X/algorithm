package bytedance

import (
	"fmt"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	node0 := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	head := reverseBetween(&node0, 2, 3)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
