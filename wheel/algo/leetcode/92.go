package leetcode

// pup->1->2->3->4->5
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	pup := &ListNode{0, head}
	cur, prev := pup, pup
	for k := 1; k <= right; k++ {
		cur = cur.Next
		if k == left-1 {
			prev = cur
		}
	}
	newHead, newTail, newNext := prev.Next, prev.Next, cur.Next
	prev.Next, cur.Next = nil, nil
	newHead = reverseList(newHead)
	prev.Next = newHead
	newTail.Next = newNext
	return pup.Next
}
