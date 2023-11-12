package main

import (
	"bfw/wheel/lang"
	"bfw/wheel/num"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	var (
		bits int = 16
		err  error
	)
	if len(args) > 0 {
		bits, err = strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		if bits > 32 {
			panic(errors.New("bits too long"))
		}
	}
	size := lang.GetRandomIntValue(1 << bits)
	time0 := time.Now()
	A, B := num.GenerateNumberString(size), num.GenerateNumberString(size)
	fmt.Printf("generate two %d length number string time: %v\n", size, time.Since(time0))
	time1 := time.Now()
	num.NaiveBigNumberMultiplication(A, B)
	fmt.Printf("naive length %d multiply %d number string time: %v\n", size, size, time.Since(time1))
	time2 := time.Now()
	num.KaratsubaBigNumberMultiplication(A, B)
	fmt.Printf("karatsuba length %d multiply %d number string time: %v\n", size, size, time.Since(time2))
	time3 := time.Now()
	num.FFTBigNumberMultiplication(A, B)
	fmt.Printf("fft %d length multiply %d length number string time: %v\n", size, size, time.Since(time3))
}
