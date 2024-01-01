package lc_1

// pup->1->2->3->3->3->4->4->5
func deleteDuplicates(head *ListNode) *ListNode {
	puppet := &ListNode{-111, head}
	prev, cur := puppet, head
	for cur != nil {
		if prev.Val == cur.Val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return puppet.Next
}
