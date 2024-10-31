package bytedance

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates2(t *testing.T) {
	node0 := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
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

	x := deleteDuplicates2(&node0)
	for x != nil {
		fmt.Println(x.Val)
		x = x.Next
	}
}
