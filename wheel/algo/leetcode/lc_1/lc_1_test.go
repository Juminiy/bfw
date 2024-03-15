package lc_1

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

func TestLC3(t *testing.T) {
	//fmt.Println(numOfBurgers(16, 7))
	//fmt.Println(numOfBurgers(17, 4))
	//fmt.Println(numOfBurgers(4, 17))
	//fmt.Println(numOfBurgers(0, 0))
	//fmt.Println(numOfBurgers(2, 1))
	//fmt.Println(isPalindrome(31013))
	//fmt.Println(countDigitOne(13))
}

func TestMerge(t *testing.T) {
	lists := GenLists([][]int{{1, 2, 3}, {1, 3, 4}, {1, 5, 6}, {1, 7, 8}, {1, 10, 25}})
	//lists := GenLists([][]int{})
	//lists := GenLists([][]int{{}, {}, {}})
	//mergeKLists(lists).Print()
	mergeKListsV2(lists).Print()
}

func TestLC4(t *testing.T) {
	list := GenList(1, 2, 3, 4, 5)
	//reverseList(l).Print()
	//swapPairs(l).Print()
	//i := makeIter(l)
	//i.next(2)
	//fmt.Println(i.tail().Val)
	//reverseKGroup(l, 3).Print()
	//rotateRight(i.head(), 1).Print()
	//rotateRight(i.head(), 2).Print()
	//rotateRight(i.head(), 3).Print()
	//rotateRight(i.head(), 4).Print()
	//rotateRight(i.head(), 5).Print()
	//rotateRight(i.head(), 6).Print()
	//deleteDuplicates(list).Print()
	//list, nl := NextEquals(list)
	//list.Print()
	//nl.Print()
	//deleteAllDuplicates(list).Print()
	reverseBetween(list, 2, 4).Print()
}

func TestLC5(t *testing.T) {
	//seats := [][]string{
	//	{"#", ".", "#", "#", ".", "#"},
	//	{".", "#", "#", "#", "#", "."},
	//	{"#", ".", "#", "#", ".", "#"},
	//}
	// dp[i][j] =
}

func TestLC6(t *testing.T) {
	node := GenNode(1, 2, 3, 4, 5)
	copyRandomList(node).Print()
}

func TestLRUCache_Get(t *testing.T) {
	//case 1
	//c := Constructor(2)
	//c.Put(1, 1)
	//c.Put(2, 2)
	//fmt.Println(c.Get(1))
	//c.Put(3, 3)
	//fmt.Println(c.Get(2))
	//c.Put(4, 4)
	//fmt.Println(c.Get(1))
	//fmt.Println(c.Get(3))
	//fmt.Println(c.Get(4))
	//1
	//-1
	//-1
	//3
	//4

	//case 2
	c := Constructor(1)
	c.Put(2, 1)
	fmt.Println(c.Get(2))
	c.Put(3, 2)
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
}

func TestSortList(t *testing.T) {
	sortList(GenList(6, 8, 7, 3, 1, 3, 8, 3, 4)).Print()
}

func TestLongs(t *testing.T) {
	fmt.Println(longestConsecutive([]int{0}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
