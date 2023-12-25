package algo

import (
	"bfw/wheel/algo/beyond"
	"bfw/wheel/algo/classical"
	"bfw/wheel/lang"
	"fmt"
	"testing"
)

func TestGenIntLog2(t *testing.T) {
	arr := beyond.GenIntLog2(100)
	lang.DisplayInt1DArrayInPythonFormat(arr)
}

func TestMakeST(t *testing.T) {
	st := beyond.MakeST([]int{9, 3, 1, 7, 5, 6, 0, 8}, func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	})
	querySlice := [][]int{
		{1, 6},
		{1, 5},
		{2, 7},
		{2, 6},
		{1, 8},
		{4, 8},
		{3, 7},
		{1, 8}}
	//9
	//9
	//7
	//7
	//9
	//8
	//7
	//9
	for _, q := range querySlice {
		fmt.Println(st.Query(q[0], q[1]))
	}
}

func TestFib(t *testing.T) {
	fmt.Println(classical.FibonacciV2(45), "\n\r",
		classical.FibonacciV3(45))
}
