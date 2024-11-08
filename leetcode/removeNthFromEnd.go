package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//1->2>3
	dummy := &ListNode{0, head}
	cur := dummy
	l := length(head)

	for i := 0; i < l-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next

	return dummy.Next

}

func length(head *ListNode) int {
	i := 0
	for head != nil {
		i++
		head = head.Next
	}
	return i
}
