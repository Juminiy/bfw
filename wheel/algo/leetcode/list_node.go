package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenList(v ...int) *ListNode {
	i := makeIterator(&ListNode{})
	for _, vV := range v {
		i.splice(&ListNode{vV, nil})
	}
	return i.head()
}

func GenLists(v [][]int) []*ListNode {
	ls := make([]*ListNode, len(v))
	for i, vV := range v {
		ls[i] = GenList(vV...)
	}
	return ls
}

// s,t is order by asc
// l = puppet->merge(s,t)
func (l *ListNode) merge(s, t *ListNode) *ListNode {
	i := makeIterator(l)
	for s != nil && t != nil {
		if s.Val < t.Val {
			i.splice(s)
			s = s.Next
		} else {
			i.splice(t)
			t = t.Next
		}
	}
	if s != nil {
		i.splice(s)
	}
	if t != nil {
		i.splice(t)
	}
	return l.Next
}

func (l *ListNode) Print() {
	h := l
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
	fmt.Println()
}

type iterator struct {
	puppet *ListNode
	cur    *ListNode
}

func makeIterator(head *ListNode) *iterator {
	i := &iterator{}
	i.puppet = head
	i.cur = head
	return i
}

func (i *iterator) head() *ListNode {
	return i.puppet.Next
}

func (i *iterator) walk() {
	i.cur = i.cur.Next
}

func (i *iterator) next() *ListNode {
	return i.cur.Next
}

func (i *iterator) hasNext() bool {
	return i.next() == nil
}

func (i *iterator) splice(n ...*ListNode) {
	for _, nN := range n {
		i.cur.Next = nN
		i.cur = nN
	}
}
