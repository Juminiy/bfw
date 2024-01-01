package lc_2

import (
	"bfw/wheel/adt"
)

/* case 1
		 1->
        /  \
      2     3
	/	\  /  \
   #	# 4    5
*/

/* case 2
 			  1 ->
	    /     		   \
	   2  ->   		    3 ->
	  /    \     	  /    \
	 4->   #-> 		# ->    6 ->
	/ \   / \  	   / \     / \
   7->#->#->#->   #->#	->#-> 8->
*/

// Unfinished
func connectV3(root *Node) *Node {
	for r := root; r != nil && nextLevel(r) != nil; r = nextLevel(r) {
		for t := r; t != nil; t = t.Next {
			var ln, rn *Node
			if t.Next != nil {
				dN := nextNext(t.Next)
				ln, rn = dN, dN
			}
			if t.Left != nil {
				if t.Right != nil {
					ln = t.Right
				}
				t.Left.Next = ln
			}
			if t.Right != nil {
				t.Right.Next = rn
			}
		}
	}
	return root
}

func nextLevel(n *Node) *Node {
	nt := n
	for nt != nil {
		if nt.Left != nil {
			return nt.Left
		}
		if nt.Right != nil {
			return nt.Right
		}
		nt = nt.Next
	}
	return nil
}

func nextNext(n *Node) *Node {
	nt := n
	for nt != nil {
		if nt.Left != nil {
			return nt.Left
		}
		if nt.Right != nil {
			return nt.Right
		}
		nt = nt.Next
	}
	return nil
}

// first: 1 2 4 8 16
// last: 1 3 7 15
func connectV2(root *Node) *Node {
	for r := root; r != nil && r.Left != nil; r = r.Left {
		for t := r; t != nil; t = t.Next {
			t.Left.Next = t.Right
			if t.Next != nil {
				t.Right.Next = t.Next.Left
			}
		}
	}
	return root
}

// O(n)空间改不好了，呜呜呜，嘤嘤嘤
func connect(root *Node) *Node {
	type qNode struct {
		r   *Node
		cnt int
	}
	q := adt.GenericQueue[*qNode]{}
	var prev *Node
	if root != nil {
		q.Push(&qNode{root, 1})
	}
	for !q.Empty() {
		e := q.Front()
		q.Pop()
		if prev != nil && !firstInLevel(e.cnt) {
			prev.Next = e.r
		}
		prev = e.r
		if lastInLevel(e.cnt) {
			prev = nil
		}
		if e.r.Left != nil {
			q.Push(&qNode{e.r.Left, e.cnt << 1})
		}
		if e.r.Right != nil {
			q.Push(&qNode{e.r.Right, e.cnt<<1 + 1})
		}
	}
	return root
}
