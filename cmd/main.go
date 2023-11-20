package main

import (
	"bfw/wheel/adt"
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int64, n)
	var maxVal int64 = 0
	for idx := 0; idx < n; idx++ {
		fmt.Scan(&arr[idx])
		if arr[idx] > maxVal {
			maxVal = arr[idx]
		}
	}
	// 100,000,000
	arr = adt.BitMapSort(arr, maxVal)
	for idx := 0; idx < n; idx++ {
		fmt.Printf("%d ", arr[idx])
	}
}
