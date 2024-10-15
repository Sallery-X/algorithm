package bytedance

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

1-2-3-4
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	l1, slow := findMid(head)

	l2 := reverse(slow)
	// 都为nil
	for l1 != l2 {
		n1 := l1.Next
		n2 := l2.Next
		l1.Next = l2
		if n1 == nil {
			break
		}
		l2.Next = n1
		l1 = n1
		l2 = n2
	}
}

func findMid(head *ListNode) (preMid, afMid *ListNode) {

	slow := head
	fast := head
	temp := head
	for fast != nil && fast.Next != nil {
		temp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	temp.Next = nil
	return head, slow
}

func reverse(head *ListNode) *ListNode {
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
