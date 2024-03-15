package lc_1

// 1->2->3->4
func reverseList(head *ListNode) *ListNode {
	var p, q *ListNode
	l := head
	for l != nil {
		p = l.Next
		l.Next = q
		q = l
		l = p
	}
	return q
}

// 1->2->3->4->5

// 1->2->nil
// 1
func reverseListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := reverseListV2(head.Next)
	head.Next.Next = n
	head.Next = nil
	return n
}
