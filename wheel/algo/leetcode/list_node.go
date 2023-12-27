package leetcode

import (
	"fmt"
)

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

func Len(l *ListNode) int {
	return makeIter(l).len()
}

func NextK(l *ListNode, k int) *ListNode {
	i := makeIter(l)
	i.next(k)
	return i.tail()
}

func Tail(l *ListNode) *ListNode {
	return makeIter(l).end()
}

func Merge(s, t *ListNode) *ListNode {
	li, si, ti := makeIter(&ListNode{}), makeIter(s), makeIter(t)
	for !si.nil() && !ti.nil() {
		if li.less(si, ti) {
			li.appendI(si)
			si.next()
		} else {
			li.appendI(ti)
			ti.next()
		}
	}
	if !si.nil() {
		li.appendI(si)
	}
	if !ti.nil() {
		li.appendI(ti)
	}
	return li.head()
}

func NextEquals(l *ListNode) *ListNode {
	i, p := makeIter(l), l
	for !i.nil() && i.pup.Val == i.cur.Val {
		p = i.cur
		i.next()
	}
	return p
}

func GetRange(l *ListNode, s, t int) (*ListNode, *ListNode) {
	i := makeIter(&ListNode{0, l})
	i.next(s - 1)
	sL := i.tail()
	i.next(t - s + 1)
	tL := i.tail()
	return sL, tL
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

func (i *iter) appendI(it ...*iter) {
	for _, itT := range it {
		i.append(itT.cur)
	}
}

func (i *iter) next(n ...int) {
	nN := 1
	if len(n) > 0 && n[0] >= 0 {
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

func (i *iter) nil() bool {
	return i.cur == nil
}

func (i *iter) less(p, q *iter) bool {
	if !validateI(p, q) {
		return false
	}
	return p.cur.Val < q.cur.Val
}

func (i *iter) equal(p, q *iter) bool {
	if !validateI(p, q) {
		return false
	}
	return p.cur.Val == q.cur.Val
}

func (i *iter) end() *ListNode {
	cur := i.cur
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

func (i *iter) reset() {
	i.cur = i.pup
}

func validateI(i ...*iter) bool {
	for _, iI := range i {
		if iI == nil || iI.nil() {
			return false
		}
	}
	return true
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func GenNode(v ...int) *Node {
	i := makeNIter(&Node{})
	for _, vV := range v {
		i.append(&Node{Val: vV})
	}
	return i.head()
}

func (n *Node) print() {
	fmt.Printf("(%d", n.Val)
	if n.Next != nil {
		fmt.Printf(",%d", n.Next.Val)
	}
	if n.Random != nil {
		fmt.Printf(",%d", n.Random.Val)
	}
	fmt.Printf(")")
}

func (n *Node) Print() {
	cur := n
	for cur != nil {
		cur.print()
		cur = cur.Next
	}
	fmt.Println()
}

type niter struct {
	pup, cur *Node
}

func makeNIter(n *Node) *niter {
	return &niter{n, n}
}

func (i *niter) appendNI(ni ...*niter) {
	for _, niI := range ni {
		i.cur.Next = niI.cur

	}
}

func (i *niter) append(n ...*Node) {
	for _, nN := range n {
		i.cur.Next = nN
		i.cur = nN
	}
}

func (i *niter) appendCopy(n ...*Node) {
	for _, nN := range n {
		newNN := &Node{Val: nN.Val}
		i.cur.Next = newNN
		i.cur = newNN
	}
}

func (i *niter) next(n ...int) {
	nN := 1
	if len(n) > 0 && n[0] >= 0 {
		nN = n[0]
	}
	for i.cur != nil && nN > 0 {
		nN--
		i.cur = i.cur.Next
	}
}

func (i *niter) head() *Node {
	return i.pup.Next
}

func (i *niter) tail() *Node {
	return i.cur
}

func (i *niter) nil() bool {
	return i.cur == nil
}

func (i *niter) reset() {
	i.cur = i.pup
}
