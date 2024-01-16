package run_any

import (
	"bfw/wheel/cc"
	"fmt"
	"os"
	"time"
)

func RunChan0() {
	ch := make(chan bool)
	go cc.PipelineSquare(ch, 1<<20)
	time0 := time.Now()
	for {
		select {
		case ret := <-ch:
			{
				if ret {
					fmt.Println("end job")
				} else {
					fmt.Println("err job")
				}
				fmt.Printf("cost time: %v\n", time.Since(time0))
				os.Exit(0)
			}
		default:
			{
				time.Sleep(1 * time.Second)
			}
		}
	}
}
