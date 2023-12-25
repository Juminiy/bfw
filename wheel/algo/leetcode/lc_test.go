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

func TestLC3(t *testing.T) {
	//fmt.Println(numOfBurgers(16, 7))
	//fmt.Println(numOfBurgers(17, 4))
	//fmt.Println(numOfBurgers(4, 17))
	//fmt.Println(numOfBurgers(0, 0))
	//fmt.Println(numOfBurgers(2, 1))
	//fmt.Println(isPalindrome(31013))
	//fmt.Println(countDigitOne(13))
}

func TestLC4(t *testing.T) {
	l := GenList(1, 2, 3, 4, 5)
	//reverseList(l).Print()
	//swapPairs(l).Print()
	//i := makeIter(l)
	//i.next(2)
	//fmt.Println(i.tail().Val)
	//reverseKGroup(l, 3).Print()
	rotateRight(l, 1).Print()
}
