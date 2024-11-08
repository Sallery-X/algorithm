package leetcode

import (
	"fmt"
	"testing"
)

func Test_sortList(t *testing.T) {
	node0 := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	x := sortList(&node0)
	for x != nil {
		fmt.Println(x.Val)
		x = x.Next
	}
}
