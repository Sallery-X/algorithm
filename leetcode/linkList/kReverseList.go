package linkList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse3(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next

	cur := head

	// 前一个不是最后一个节点继续，当前节点是最后一个时，代表已经反转完毕
	for pre != tail {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre, head

}

func realKr(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		temp := tail.Next

		head, tail = reverse3(head, tail)

		pre.Next = head
		tail.Next = temp

		pre = tail
		head = tail.Next

	}
	return dummy.Next

}
