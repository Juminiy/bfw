package leetcode

// 1->2->3->4->5
// k = 1
// 5->1->2->3->4
// k = 2
// 4->5->1->2->3
func rotateRight(head *ListNode, k int) *ListNode {
	hL := Len(head)
	if hL == 0 {
		return head
	}
	k %= hL
	if k == 0 {
		return head
	}
	newTail := NextK(head, hL-k-1)
	newHead := newTail.Next
	newTail.Next = nil
	Tail(newHead).Next = head
	return newHead
}
