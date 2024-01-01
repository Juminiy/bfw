package lc_1

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	puppet := &ListNode{0, head}
	h1, h2, k := puppet, puppet, n
	for k > 0 {
		k--
		h1 = h1.Next
	}
	for h1 != nil && h1.Next != nil {
		h2 = h2.Next
		h1 = h1.Next
	}
	if h2.Next != nil {
		h2.Next = h2.Next.Next
	}
	return puppet.Next
}
