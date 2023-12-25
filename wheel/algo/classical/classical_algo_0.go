package classical

import (
	"errors"
	"strconv"
)

const (
	fibonacciAssertValue = 3
)

var (
	fibonacciInputNError = errors.New("fibonacci input N Error")
)

func FibonacciAssert(n int) int {
	if n < 0 {
		panic(fibonacciInputNError)
	} else if n <= 1 {
		return n
	}
	return fibonacciAssertValue
}

func Fibonacci(n int) int {
	if fibAssert := FibonacciAssert(n); fibAssert < fibonacciAssertValue {
		return fibAssert
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

func FibonacciV2(n int) int {
	if n < 0 {
		panic(fibonacciInputNError)
	} else if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1
	for idx := 2; idx <= n; idx++ {
		fib[idx] = fib[idx-1] + fib[idx-2]
	}
	return fib[n]
}

func FibonacciV3(n uint) uint {
	if n == 0 {
		return 0
	}
	var n1, n2 uint = 0, 1
	for i := uint(1); i < n; i++ {
		n3 := n1 + n2
		n1 = n2
		n2 = n3
	}
	return n2
}

func DisplayFibonacci(n, fibN int) string {
	return "Fibonacci(" + strconv.Itoa(n) + ") = " + strconv.Itoa(fibN)
}
