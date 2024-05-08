package bytedance

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	node0 := LinkNode{
		Val: 0,
		Next: &LinkNode{
			Val: 1,
			Next: &LinkNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	head := ReverseLinkedList2(&node0)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
