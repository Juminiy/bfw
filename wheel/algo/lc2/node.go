package lc2

import (
	"bfw/wheel/adt"
	"errors"
	"fmt"
	"math/bits"
	"strconv"
)

const (
	nodeNilString = "#"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
		  1
	    /   \
	   #     3
	  / \   / \
	 #   # #   7
*/

func BuildTreeNodeByBFS(v ...string) *Node {
	m := make(map[int]*Node)
	assertV := func(vV string) *Node {
		if vVI, err := strconv.Atoi(vV); err != nil {
			return nil
		} else {
			return &Node{Val: vVI}
		}
	}
	pushV := func(i int) {
		if i < len(v) {
			if vV := assertV(v[i]); vV != nil {
				m[i] = vV
			}
		} else {
			fmt.Println("index out of bound")
		}
	}
	for i := range v {
		pushV(i)
	}
	for i := 0; i < len(v)>>1; i++ {
		if m[i] != nil {
			m[i].Left = m[i<<1+1]
			m[i].Right = m[i<<1+2]
		}
	}

	return m[0]
}

func (n *Node) visitByBFS() []int {
	q, sl := adt.GenericQueue[*Node]{}, make([]int, 0)
	if n != nil {
		q.Push(n)
	}
	for !q.Empty() {
		e := q.Front()
		sl = append(sl, e.Val)
		q.Pop()
		if e.Left != nil {
			q.Push(e.Left)
		}
		if e.Right != nil {
			q.Push(e.Right)
		}
	}
	return sl
}

func (n *Node) visitByDFS(mode rune) []int {
	switch mode {
	case 'E', 'e':
		{
			return n.visitByPreDFS()
		}
	case 'M', 'm':
		{
			return n.visitByMidDFS()
		}
	case 'T', 't':
		{
			return n.visitByPostDFS()
		}
	default:
		{
			panic(errors.New("unsupported dfs mode: " + string(mode)))
		}
	}
}

func (n *Node) visitByPreDFS() []int {
	return nil
}

func (n *Node) visitByMidDFS() []int {
	return nil
}

func (n *Node) visitByPostDFS() []int {
	return nil
}
func (n *Node) Print(mode ...rune) {
	m := 'B'
	if len(mode) > 0 {
		m = mode[0]
	}
	switch m {
	case 'B', 'b':
		{
			fmt.Println(n.visitByBFS())
		}
	default:
		{
			fmt.Println(n.visitByDFS(m))
		}
	}
}

func (n *Node) PrintNext() {
	nt := n
	ntl := n.Left
	for nt != nil {
		fmt.Printf("%d ", nt.Val)
		nt = nt.Next
		if nt == nil && ntl != nil {
			nt = ntl
			ntl = nt.Left
		}
	}
}

func firstInLevel(cnt int) bool {
	return cnt >= 0 && bits.OnesCount(uint(cnt)) == 1
}
func lastInLevel(cnt int) bool {
	return cnt >= 0 && bits.OnesCount(uint(cnt+1)) == 1
}
