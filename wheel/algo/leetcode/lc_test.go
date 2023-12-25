package leetcode

import (
	"fmt"
	"testing"
)

func TestLC(t *testing.T) {
	firstMissingPositive([]int{3, 4, -1, 1})
}

func TestLC1(t *testing.T) {
	fmt.Println(getUglyNumberSlice(10))
	fmt.Println(nthUglyNumber(10))
}

func TestLC2(t *testing.T) {
	lists := GenLists([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}})
	//lists := GenLists([][]int{{}, {}})
	//mergeKListsV2(lists).Print()
	mergeKLists(lists).Print()
}

func TestLC3(t *testing.T) {
	l := GenList(1, 2, 3, 4, 5)
	//removeNthFromEnd(l, 2).Print()
	swapPairs(l).Print()
}
