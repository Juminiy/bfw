package lc2

import (
	"bfw/wheel/adt"
	"math/bits"
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
	   2  -> 3 ->
	  / \   / \
	 4   5 6   7 ->
*/
// 1 3 7 15
func connect(root *Node) *Node {
	cnt, q := 0, adt.GenericQueue[*Node]{}
	var prev *Node
	if root != nil {
		q.Push(root)
	}
	lastInLevel := func() bool {
		return bits.OnesCount(uint(cnt+1)) == 1
	}
	for !q.Empty() {
		e := q.Front()
		q.Pop()
		cnt++
		if !lastInLevel() && prev != nil {
			prev.Next = e
			prev = e
		} else {
			prev = nil
		}
		if e.Left != nil {
			q.Push(e.Left)
		}
		if e.Right != nil {
			q.Push(e.Right)
		}
	}
	return root
}
