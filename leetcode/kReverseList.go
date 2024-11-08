package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func kRev(head, tail *ListNode) (*ListNode, *ListNode) {

	pre := tail.Next
	cur := head

	for pre != tail {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre, head
}

func realKr(head *ListNode, k int) *ListNode {
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
		temp := tail.Next
		head, tail = kRev(head, tail)

		pre.Next = head
		tail.Next = temp
		pre = tail
		head = tail.Next
	}

	return hair.Next

}
