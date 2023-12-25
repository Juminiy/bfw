package leetcode

<<<<<<< HEAD
import (
	"fmt"
)
=======
import "fmt"
>>>>>>> ba4054c (1. all update)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenList(v ...int) *ListNode {
	i := makeIter(&ListNode{})
	for _, vV := range v {
		i.append(&ListNode{Val: vV})
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

func (l *ListNode) Len() int {
	return makeIter(l).len()
}

func (l *ListNode) NextK(k int) *ListNode {
	i := makeIter(l)
	i.next(k)
	return i.tail()
}

func (l *ListNode) Tail() *ListNode {
	return makeIter(l).end()
}

func (l *ListNode) Print() {
	i := makeIter(l)
	for i.tail() != nil {
		fmt.Printf("%d ", i.tail().Val)
		i.next()
	}
	fmt.Println()
}

type iter struct {
	pup, cur *ListNode
}

func makeIter(l *ListNode) *iter {
	return &iter{l, l}
}

func (i *iter) len() int {
	curCur, curLen := i.cur, 0
	for curCur != nil {
		curCur = curCur.Next
		curLen++
	}
	return curLen
}

func (i *iter) append(l ...*ListNode) {
	for _, lL := range l {
		i.cur.Next = lL
		i.cur = lL
	}
}

func (i *iter) next(n ...int) {
	nN := 1
	if len(n) > 0 && n[0] > 0 {
		nN = n[0]
	}
	for i.cur != nil && nN > 0 {
		nN--
		i.cur = i.cur.Next
	}
}

func (i *iter) head() *ListNode {
	return i.pup.Next
}

func (i *iter) tail() *ListNode {
	return i.cur
}

func (i *iter) end() *ListNode {
	cur := i.cur
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	return cur
}
