package fft

import (
	"fmt"
	"testing"
)

func TestFFT(t *testing.T) {
	// after hadamard and div (1 << bitCnt)
	//[0 45 76 94 100 70 40 19 6 0 0 0 0 0 0 0]
	// destValue
	//[8 3 8 1 0 2 0 5 0]
	fmt.Println(polyIntMul([]int{5, 4, 3, 2, 1}, []int{0, 9, 8, 7, 6}))
}
