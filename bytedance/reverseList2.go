package bytedance

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}

	pre := dummy

	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	mL := pre.Next
	end := mL

	for i := 0; i < right-left; i++ {
		end = end.Next
	}

	next := end.Next
	end.Next = nil

	pre.Next = reverse2(mL)

	mL.Next = next

	return dummy.Next

}

func reverse2(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}
