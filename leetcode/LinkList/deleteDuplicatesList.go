package LinkList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head

	for cur != nil && cur.Next != nil {

		if cur.Val == cur.Next.Val {
			temp := cur.Next.Next
			cur.Next = temp
		} else {
			cur = cur.Next
		}
	}
	return head
}
