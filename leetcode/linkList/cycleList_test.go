package linkList

import (
	"fmt"
	"testing"
)

func Test_hasCycle(t *testing.T) {

}

func Test_hasCycle2(t *testing.T) {
	lis1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	lis2 := &ListNode{
		Val:  0,
		Next: lis1,
	}
	lis3 := &ListNode{
		Val:  0,
		Next: lis2,
	}
	lis4 := &ListNode{
		Val:  0,
		Next: lis3,
	}
	//lis1.Next = lis4
	fmt.Println(hasCycle2(lis4))
}
