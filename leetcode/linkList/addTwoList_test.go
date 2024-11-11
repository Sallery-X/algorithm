package linkList

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	listA := ListNode{3, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}}

	listB := ListNode{5, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}}

	fmt.Println(addTwoNumbers(&listA, &listB))
}
