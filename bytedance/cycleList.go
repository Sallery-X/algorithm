package bytedance

func hasCycle(head *ListNode) bool {
	sets := make(map[*ListNode]int, 0)

	for head != nil {
		if _, ok := sets[head]; ok {
			return true
		}
		head = head.Next

	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
