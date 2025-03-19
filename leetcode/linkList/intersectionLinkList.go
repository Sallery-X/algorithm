package linkList

//https://leetcode.cn/problems/intersection-of-two-linked-lists/description/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB
	lenA, lenB := 0, 0
	for a != nil {
		lenA++
		a = a.Next
	}
	for b != nil {
		lenB++
		b = b.Next
	}

	a = headA
	b = headB

	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			a = a.Next
		}
	} else if lenB > lenA {
		for i := 0; i < lenB-lenB; i++ {
			b = b.Next
		}
	}

	for a != nil && b != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil

}
