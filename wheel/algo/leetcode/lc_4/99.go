package lc_4

func recoverTree(root *TreeNode) {

}

func visTreeNode(r *TreeNode) {
	if r.Left != nil {
		visTreeNode(r.Left)
		if r.Left.Val > r.Val {
			swapTreeNode(r.Left, r)
			return
		}
	}
	if r.Right != nil {
		visTreeNode(r.Right)
		if r.Right.Val < r.Val {
			swapTreeNode(r.Right, r)
			return
		}
	}
}
