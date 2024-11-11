package linkList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 是否是回文链表
func isPalindrome(head *ListNode) bool {
	temp := make([]int, 0)
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}

	for i := 0; i < len(temp)/2; i++ {
		if temp[i] != temp[len(temp)-i-1] {
			return false
		}
	}

	return true

}
