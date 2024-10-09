package bytedance

import (
	"fmt"
	"testing"
)

func Test_reorderList(t *testing.T) {
	node0 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	reorderList(node0)
	for node0 != nil {
		fmt.Println(node0.Val)
		node0 = node0.Next
	}
}
