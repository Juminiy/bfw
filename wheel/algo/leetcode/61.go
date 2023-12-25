package leetcode

// 1->2->3->4->5
func rotateRight(head *ListNode, k int) *ListNode {

	Len := func(l *ListNode) int {
		return makeIter(l).len()
	}

	NextK := func(l *ListNode, kK int) *ListNode {
		i := makeIter(l)
		i.next(k)
		return i.tail()
	}

	Tail := func(l *ListNode) *ListNode {
		return makeIter(l).end()
	}

	hLen := Len(head)
	k %= hLen
	if k == 0 {
		return head
	}
	cur := &ListNode{0, head}
	hkPrev := NextK(cur, hLen-k)
	hK := hkPrev.Next
	hKTail := Tail(hK)
	hKTail.Next = head
	hkPrev.Next = nil
	return hK
}
