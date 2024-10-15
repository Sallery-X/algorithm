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

	slow, fast := head, head
	has := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			has = true
			break
		}
	}
	return has
}
