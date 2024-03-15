package lc_4

import (
	"fmt"
	"testing"
)

func TestLC1(t *testing.T) {
	for i := 1; i <= 9; i++ {
		fmt.Println(totalNQueens(i))
	}
}

func TestLC2(t *testing.T) {
	fmt.Printf("res := []int{")
	for i := 1; i <= 8; i++ {
		fmt.Printf("%d,", numTrees(i))
	}
	fmt.Printf("}\n")
}

func TestLC4(t *testing.T) {
	p, q := &TreeNode{Val: 1}, &TreeNode{Val: 2}
	swapTreeNode(p, q)
	fmt.Println(p, q)
}
