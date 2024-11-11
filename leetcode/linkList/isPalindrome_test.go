package linkList

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
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

	fmt.Println(isPalindrome(&node0))

}
