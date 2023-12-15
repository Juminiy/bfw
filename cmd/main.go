package main

import (
	"bfw/wheel/algo/classical"
	"fmt"
	"time"
)

func Spinner(sleepTime time.Duration) {
	for {
		for _, ch := range `-\|/` {
			fmt.Printf("\r%c", ch)
			time.Sleep(sleepTime)
		}
	}
}

func main() {
	go Spinner(100 * time.Millisecond)
	time0 := time.Now()
	n := 45
	fn := classical.Fibonacci(n)
	fmt.Println("\r\n", classical.DisplayFibonacci(n, fn))
	fmt.Printf("fibonacci CPU time: %v", time.Since(time0))
}
