package cc

import (
	"fmt"
	"math/rand"
	"time"
)

func SendSign(ch chan int, s int) {
	ch <- s
}

func ReceiveSign(ch chan int) {
	s := <-ch
	fmt.Println(s)
}

func PipelineSquare() {
	numChan := make(chan int)
	nsqChan := make(chan int)

	go genNum(numChan)

	go calNum(nsqChan, numChan)

	for nsq := range nsqChan {
		fmt.Printf("%d,", nsq)
	}
}

func genNum(out chan<- int) {
	for num := 0; num < 100; num++ {
		out <- num
	}
	close(out)
}

func calNum(out chan<- int, in <-chan int) {
	for num := range in {
		out <- num * num
	}
	close(out)
}

type ConnResp struct {
	Data string
	Time time.Duration
}

func GetRapidestConnByMultipleConns(urls []string) ConnResp {
	resps := make(chan ConnResp, 3)
	for i, _ := range urls {
		url := urls[i]
		go func() {
			resps <- requestByURL(url)
		}()
	}
	return <-resps
}

func requestByURL(url string) ConnResp {
	var (
		resp string
		dura time.Duration
	)
	switch url {
	case "host1":
		{
			dura = time.Duration(rand.Intn(5)+1) * time.Second
			time.Sleep(dura)
			fmt.Println(dura)
			resp = "resp1"
		}
	case "host2":
		{
			dura = time.Duration(rand.Intn(5)+1) * time.Second
			time.Sleep(dura)
			fmt.Println(dura)
			resp = "resp2"
		}
	case "host3":
		{
			dura = time.Duration(rand.Intn(5)+1) * time.Second
			time.Sleep(dura)
			fmt.Println(dura)
			resp = "resp3"
		}
	default:
		{
			dura = time.Duration(rand.Intn(5)+1) * time.Second
			time.Sleep(dura)
			fmt.Println(dura)
			resp = "error conn"
		}
	}
	return ConnResp{
		resp,
		dura,
	}
}

func RocketLaunch() {
	tChan := time.Tick(5 * time.Second)
	<-tChan
}

func PrintEven(n int) {
	ch := make(chan int, 1)
	for i := 0; i < (n<<1)+1; i++ {
		select {
		case x := <-ch:
			{
				fmt.Printf("%d ", x)
			}
		case ch <- i:
			{

			}
		}
	}
}
