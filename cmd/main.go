package main

import (
	"bfw/cmd/web"
	"bfw/wheel/la"
	"fmt"
	"time"
)

func runWebApi() {
	web.ServeRun()
}

func runMatrixTest() {
	genBTSize, genBSize := 50, 50
	time0 := time.Now()
	bm1 := la.GenBlockMatrix(genBTSize, genBTSize, "f", genBSize, genBSize, 100)
	bm2 := la.GenBlockMatrix(genBTSize, genBTSize, "f", genBSize, genBSize, 100)
	fmt.Printf("2 * %d*%d Matrix Generate time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time0))

	time1 := time.Now()
	bm1.Mul(bm2)
	fmt.Printf("%d*%d Matrix Multiply After Speed & DivBlock time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time1))

	time2 := time.Now()
	bm1.Matrix().MTimes(bm2.Matrix())
	fmt.Printf("%d*%d Matrix Multiply After speed time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time2))

	time2dot5 := time.Now()
	bm1.Matrix().Mul(bm2.Matrix())
	fmt.Printf("%d*%d Matrix Multiply None speed time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time2dot5))

	time3 := time.Now()
	bm1.Matrix().Equal(bm2.Matrix())
	fmt.Printf("%d*%d Matrix Traverse time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time3))
}

func main() {
	runMatrixTest()
}
