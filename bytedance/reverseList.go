package bytedance

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func ReverseLinkedList(head *LinkNode) *LinkNode {
	var pre *LinkNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre

}

func ReverseLinkedList2(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
