package cc

import (
	"bfw/wheel/lang"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

func sumOfFloat64Array(a []float64, wg *sync.WaitGroup, sumChan chan float64) {
	var (
		aLen = len(a)
		aIdx = 0
		sum  = 0.0
	)
	if lang.IsOdd(aLen) {
		sum += a[aIdx]
		aIdx++
	}
	for ; aIdx < aLen; aIdx += 2 {
		sum += a[aIdx] + a[aIdx+1]
	}
	sumChan <- sum
	wg.Done()
}

func SumOfFloat64Array(a []float64, rCnt int) float64 {
	rCnt = lang.CeilBin(rCnt)
	totalSum := 0.0
	preIdx := 0
	aLen := len(a)
	aDestLen := lang.CeilBin(aLen)
	a = append(a, make([]float64, aDestLen-aLen)...)
	segLen := aDestLen / rCnt

	fChan := make(chan float64)
	wg := new(sync.WaitGroup)
	wg.Add(rCnt)

	for tIdx := 0; tIdx < rCnt; tIdx++ {
		go sumOfFloat64Array(a[preIdx:segLen+preIdx], wg, fChan)
		preIdx += segLen
	}

	go func() {
		wg.Wait()
		close(fChan)
	}()

	for fSum := range fChan {
		totalSum += fSum
	}

	return totalSum
}

func runCCSumOfFloat64BenchMark(routines int) float64 {
	//time0 := time.Now()
	f64Arr := lang.GetRandFloat64ArrayByRange(1<<22, 0.0, 1.0)
	//fmt.Println("Generate 2"+cal.GetExponent("22"), "float64 num, time:", time.Since(time0))
	//time1 := time.Now()
	//res0 := func(f []float64) float64 {
	//	sum0 := 0.0
	//	for idx := 0; idx < len(f); idx++ {
	//		sum0 += f[idx]
	//	}
	//	return sum0
	//}(f64Arr)
	//fmt.Println("single routine calculate sum of 2"+cal.GetExponent("22"), "float64 num, time:", time.Since(time1))
	time2 := time.Now()
	SumOfFloat64Array(f64Arr, routines)
	du := time.Since(time2)
	//fmt.Printf("%d routines calculate sum of 2"+cal.GetExponent("22")+" float64 num, time:%v\n", routines, du)
	//fmt.Println("concurrent calculate result is:", lang.EqualFloat64Zero(res0-res1), res0, res1)
	//res0, res1 = res1, res0
	return du.Seconds()
}

func ParallelF64BenchMark() {
	time0 := time.Now()
	timeSlice := lang.ConstructReal2DArrayBySize(1 << 10)
	minSec, maxSec := 1e10, -1e10
	for routineCnt := 1; routineCnt < (1 << 12); routineCnt += 4 {
		duSec := runCCSumOfFloat64BenchMark(routineCnt)
		timeSlice[routineCnt>>2][0] = float64(routineCnt)
		timeSlice[routineCnt>>2][1] = duSec
		minSec = math.Min(minSec, duSec)
		maxSec = math.Max(maxSec, duSec)
	}
	fmt.Printf("min time: %fs, max time: %fs\n", minSec, maxSec)
	sort.Sort(timeSlice)
	fmt.Println("top 5 min routines:", timeSlice[1], timeSlice[2], timeSlice[3], timeSlice[4], timeSlice[5])
	fmt.Println("1024 rounds calculate, total time:", time.Since(time0))
}
