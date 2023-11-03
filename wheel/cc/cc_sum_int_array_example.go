package cc

import (
	"fmt"
	"sync"
)

func sumOfIntArray(arr []int, wg *sync.WaitGroup, ch chan int) {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	ch <- sum
	wg.Done()
}

func runConcurrencySumOfArray() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	mid := len(arr) / 2

	wg.Add(2)
	go sumOfIntArray(arr[:mid], wg, ch)
	go sumOfIntArray(arr[mid:], wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	total := 0
	for sum := range ch {
		total += sum
	}

	fmt.Println("Total:", total)
}
