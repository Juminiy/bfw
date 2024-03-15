package lc_4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		return p.Val == q.Val &&
			isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

func swapTreeNode(p, q *TreeNode) {
	p.Val, p.Left, p.Right, q.Val, q.Left, q.Right = q.Val, q.Left, q.Right, p.Val, p.Left, p.Right
}
