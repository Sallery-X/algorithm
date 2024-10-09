package bytedance

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode

	for _, v := range lists {
		res = MergeLinkList2(res, v)
	}
	return res
}

func MergeLinkList2(list1, list2 *ListNode) *ListNode {
	//var res *ListNode
	res := &ListNode{}
	head := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			res.Next = list1
			list1 = list1.Next
		} else {
			res.Next = list2
			list2 = list2.Next
		}
		res = res.Next
	}

	if list1 != nil {
		res.Next = list1
	}

	if list2 != nil {
		res.Next = list2
	}
	return head.Next

}
