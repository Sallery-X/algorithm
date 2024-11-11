package linkList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 两个链表相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	yu := 0
	res := &ListNode{}
	head := res
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + yu
		res.Next = &ListNode{Val: sum % 10}
		yu = sum / 10

		res = res.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {

		sum := l1.Val + yu
		res.Next = &ListNode{Val: sum % 10}
		yu = sum / 10
		res = res.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := l2.Val + yu
		res.Next = &ListNode{Val: sum % 10}
		yu = sum / 10
		res = res.Next
		l2 = l2.Next
	}

	if yu > 0 {
		res.Next = &ListNode{Val: yu}
	}

	return head.Next

}
