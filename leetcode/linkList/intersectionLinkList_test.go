package linkList

import (
	"fmt"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	listC := ListNode{4, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}}

	listA := ListNode{3, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  2,
			Next: &listC,
		},
	}}

	listB := ListNode{5, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  2,
			Next: &listC,
		},
	}}

	x := getIntersectionNode(&listA, &listB)
	fmt.Println(x)
}
