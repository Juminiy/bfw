package leetcode

// 1->2->3->4
func reverseList(head *ListNode) *ListNode {
	var p, q *ListNode
	for head != nil {
		p = head.Next
		head.Next = q
		q = head
		head = p
	}
	return q
}
