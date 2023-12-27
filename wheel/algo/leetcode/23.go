package leetcode

import "bfw/wheel/adt"

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

func mergeKListsV2(lists []*ListNode) *ListNode {
	h := adt.MakeHeap[*ListNode](func(a, b *ListNode) bool {
		return a.Val < b.Val
	})
	for _, n := range lists {
		if n != nil {
			h.Push(n)
		}
	}
	i := makeIter(&ListNode{})
	for !h.Empty() {
		t := h.Pop(1)[0]
		i.append(t)
		if t.Next != nil {
			h.Push(t.Next)
		}
	}
	return i.head()
}
