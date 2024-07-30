package bytedance

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = mReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func mReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre = tail.Next
	cur := head

	for pre != tail {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre, head
}
