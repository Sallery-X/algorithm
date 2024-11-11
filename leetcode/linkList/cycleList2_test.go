package linkList

import (
	"fmt"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	lis1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	lis2 := &ListNode{
		Val:  2,
		Next: lis1,
	}
	lis3 := &ListNode{
		Val:  3,
		Next: lis2,
	}
	lis4 := &ListNode{
		Val:  4,
		Next: lis3,
	}
	lis1.Next = lis4
	fmt.Println(detectCycle(lis4).Val)
}
