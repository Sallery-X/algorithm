package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

归并逻辑，先拆分成最小的，链表，然后每个进行合并


*/

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	mid := findMidNode(head)
	left := head
	right := mid.Next
	mid.Next = nil

	left = sortList(left)
	right = sortList(right)
	//递归，递去，拆分成最小的，每个节点一个

	res := mergeList(left, right)
	//归来，每个节点，进行两两合并
	return res

}

func findMidNode(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow

}

func mergeList(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}

		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}

	return dummy.Next

}
