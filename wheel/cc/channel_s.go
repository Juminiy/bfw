package cc

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func SendSign(ch chan int, s int) {
	for {
		select {
		case ch <- s:
			{

			}
		default:
			{
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func ReceiveSign(ch chan int) {
	for {
		select {
		case s := <-ch:
			{
				fmt.Println(s)
			}
		default:
			{
				time.Sleep(20 * time.Second)
			}
		}
	}
}

func getFilePtr(dir string) *os.File {
	filePtr, err := os.OpenFile(dir, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	return filePtr
}

func PipelineSquare(fi chan bool, maxSize int) {
	numChan := make(chan int)
	nsqChan := make(chan int)

	go genNum(numChan, maxSize)

	go calNum(nsqChan, numChan)

	// write procedure
	fptr := getFilePtr("num_square.txt")
	defer func(fp *os.File) {
		err := fp.Close()
		if err != nil {
			panic(err)
		}
	}(fptr)
	cnt := 0
	for {
		for nsq := range nsqChan {
			_, err := fptr.WriteString(strconv.Itoa(nsq) + ",")
			if err != nil {
				panic(err)
			}
			cnt++
			if cnt%100 == 0 && cnt > 0 {
				_, err = fptr.WriteString("\n")
				if err != nil {
					panic(err)
				}
			}
		}
		if cnt == maxSize {
			fi <- true
		}
	}

}

func genNum(out chan<- int, maxSize int) {
	for num := 0; num < maxSize; num++ {
		out <- rand.Intn(num + 1)
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
