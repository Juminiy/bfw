package leetcode

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
