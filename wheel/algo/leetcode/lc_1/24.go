package lc_1

// cur->1->2->3->4->5

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := &ListNode{0, head}
	var newHead *ListNode
	// l->ln->lnn->p
	// l->lnn->ln->p
	swapPair := func(l *ListNode) *ListNode {
		ln, lnn := l.Next, l.Next.Next
		l.Next = lnn
		ln.Next = lnn.Next
		lnn.Next = ln
		return l
	}

	for cur.Next != nil && cur.Next.Next != nil {
		cur = swapPair(cur)
		if newHead == nil {
			newHead = cur
		}
		cur = cur.Next.Next
	}
	return newHead.Next
}
