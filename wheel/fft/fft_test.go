package fft

import (
	"bfw/wheel/lang"
	"fmt"
	"testing"
	"time"
)

func TestFFT(t *testing.T) {
	size := lang.GetRandomIntValue(1 << 16)
	time0 := time.Now()
	A, B := lang.GenerateNumberString(size), lang.GenerateNumberString(size)
	fmt.Printf("generate two %d length number string time: %v\n", size, time.Since(time0))
	time1 := time.Now()
	lang.NaiveBigNumberMultiplication(A, B)
	fmt.Printf("naive length %d multiply %d number string time: %v\n", size, size, time.Since(time1))
	time2 := time.Now()
	lang.KaratsubaBigNumberMultiplication(A, B)
	fmt.Printf("karatsuba length %d multiply %d number string time: %v\n", size, size, time.Since(time2))
	time3 := time.Now()
	DfftBigNumberMultiplication(A, B)
	fmt.Printf("fft %d length multiply %d length number string time: %v\n", size, size, time.Since(time3))
}

func TestDfftBigNumberMultiplication(t *testing.T) {
	fmt.Println(len("fft 10840 length multiply 10840 length number string time"))
}
