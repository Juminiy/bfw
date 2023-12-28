package lc2

import "testing"

func TestLC1(t *testing.T) {
	n := BuildTreeNodeByBFS(
		"1",
		"2", "3",
		"4", "#", "#", "6",
		"7", "#", "#", "#", "#", "#", "#", "8")
	//n := BuildTreeNodeByBFS(
	//	"1", "2", "3", "#", "#", "4", "5")
	connectV3(n)
	n.PrintNext()
}
