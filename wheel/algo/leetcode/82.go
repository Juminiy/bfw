package leetcode

// pup->1->1->2->2->3->3->3->4->4->5->5->6
// pup->6
func deleteAllDuplicates(head *ListNode) *ListNode {
	pup := &ListNode{-999, head}
	prev, cur := pup, head
	for cur != nil {
		cn := NextEquals(cur)
		if cur == cn {
			prev = cur
			cur = cur.Next
		} else {
			prev.Next = cn.Next
			cur = cn.Next
		}
	}
	return pup.Next
}
