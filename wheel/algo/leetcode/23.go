package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tNode *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	for l1 != nil && l2 != nil {
		xNode := l1
		if l1.Val < l2.Val {
			xNode = l1
			l1 = l1.Next
		} else {
			xNode = l2
			l2 = l2.Next
		}
		if head == nil {
			tNode = xNode
			head = tNode
			head.Next = nil
		} else {
			tNode.Next = xNode
			tNode = xNode
			tNode.Next = nil
		}
	}
	if l1 == nil {
		tNode.Next = l2
	}
	if l2 == nil {
		tNode.Next = l1
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	} else if l == 1 {
		return lists[0]
	} else {
		mid := l >> 1
		p1 := mergeKLists(lists[:mid])
		p2 := mergeKLists(lists[mid:])
		return mergeTwoLists(p1, p2)
	}
}
