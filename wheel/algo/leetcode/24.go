package leetcode

// p->1->2->3->4->5
func swapPairs(head *ListNode) *ListNode {
	next2 := func(l *ListNode) (*ListNode, *ListNode, bool) {
		if l1, l2 := l.Next, l.Next.Next; l1 != nil && l2 != nil {
			return l1, l2, true
		}
		return nil, nil, false
	}
	puppet := &ListNode{0, head}
	cur := puppet
	for l1, l2, ok := next2(cur); ok; {
		l1.Next = l2.Next
		l2.Next = l1
		cur.Next = l2
		cur = l1
	}

	return puppet.Next
}
