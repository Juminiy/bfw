package fft

import (
	"fmt"
	"testing"
)

func TestFFT(t *testing.T) {
	fmt.Println(polyMul([]int{5, 4, 3, 2, 1}, []int{0, 9, 8, 7, 6}))
}
