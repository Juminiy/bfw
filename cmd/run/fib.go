package run

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

func ProgressBar(s chan bool, firstLine ...string) {
	if len(firstLine) > 0 {
		fmt.Printf(firstLine[0])
	}
	t0 := time.Now()
	for {
		select {
		case <-s:
			{
				fmt.Printf("\u001B[2J\r[==========100.0%%=========]")
				break
			}
		default:
			{
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("\u001B[2J\r[==========%.1f%%=========]", time.Since(t0).Seconds()/6.0682377*100.0)
			}
		}
	}
}

func RunFibonacci() {
	s := make(chan bool)
	n := 45
	go ProgressBar(s)
	time0 := time.Now()
	fn := classical.Fibonacci(n)
	s <- true
	fmt.Println("\n", classical.DisplayFibonacci(n, fn))
	fmt.Printf("fibonacci CPU time: %v", time.Since(time0))
}
