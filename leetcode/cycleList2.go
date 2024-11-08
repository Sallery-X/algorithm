package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	slow, fast := head, head

	//快慢指针相遇
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 没环
	if fast == nil || fast.Next == nil {
		return nil
	}

	//有环 找起点
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow

}
