package linkList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := first.Next

		first.Next = second.Next
		second.Next = first
		cur.Next = second

		cur = first

	}

	return dummy.Next

}
