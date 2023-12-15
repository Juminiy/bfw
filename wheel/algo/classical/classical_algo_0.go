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

func DisplayFibonacci(n, fibN int) string {
	return "Fibonacci(" + strconv.Itoa(n) + ") = " + strconv.Itoa(fibN)
}
