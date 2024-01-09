package lc_1

import "bfw/wheel/adt"

func sortList(head *ListNode) *ListNode {
	h := adt.MakeHeap[*ListNode](func(a *ListNode, b *ListNode) bool {
		return a.Val < b.Val
	})
	for head != nil {
		next := head.Next
		head.Next = nil
		h.Push(head)
		head = next
	}
	i := makeIter(&ListNode{})
	for !h.Empty() {
		i.append(h.Pop(1)[0])
	}
	return i.head()
}

// A+B = 90
// B/A = 17
// B = 17A
// 18A = 90
// A = 5
