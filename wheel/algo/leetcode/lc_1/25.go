package lc_1

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := &ListNode{0, head}
	var newHead, nextGroup *ListNode

	for cur != nil {
		t := NextK(cur, k)
		if t == nil {
			break
		}
		nextGroup = t.Next
		t.Next = nil
		cur.Next = reverseList(cur.Next)
		if newHead == nil {
			newHead = cur
		}
		cur = NextK(cur, k)
		cur.Next = nextGroup
	}
	if newHead == nil {
		return head
	}
	return newHead.Next
}
