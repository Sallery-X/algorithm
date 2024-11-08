package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeLinkList(t *testing.T) {
	ListNode1 := &ListNode{
		Val: 0,
		Next: &ListNode{Val: 1, Next: &ListNode{
			Val:  8,
			Next: nil,
		}},
	}

	ListNode2 := &ListNode{
		Val: 3,
		Next: &ListNode{Val: 4, Next: &ListNode{
			Val:  5,
			Next: nil,
		}},
	}
	l := MergeLinkList(ListNode1, ListNode2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
