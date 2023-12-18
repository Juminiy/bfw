package main

import (
	"bfw/wheel/cc"
	"fmt"
	"time"
)

func main() {
	time0 := time.Now()
	concount := cc.ConcurrentCount(1<<20, 1<<10)
	fmt.Println(time.Since(time0), concount)
	time0 = time.Now()
	concount = cc.ConcurrentCount(1<<20, 1)
	fmt.Println(time.Since(time0), concount)
}
