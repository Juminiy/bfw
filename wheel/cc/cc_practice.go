package cc

import (
	"bfw/wheel/lang"
	"sync"
)

func sumOfFloat64Array(a []float64, wg *sync.WaitGroup, sumChan chan float64) {
	var (
		aLen = len(a)
		aIdx = 0
		sum  = 0.0
	)
	if lang.Odd(aLen) {
		sum += a[aIdx]
		aIdx++
	}
	for ; aIdx < aLen; aIdx += 2 {
		sum += a[aIdx] + a[aIdx+1]
	}
	sumChan <- sum
	wg.Done()
}

func CCSumOfFloat64Array(a []float64, routineCnt int) float64 {
	routineCnt = lang.CeilBin(routineCnt)
	totalSum := 0.0
	preIdx := 0
	aLen := len(a)
	aDestLen := lang.CeilBin(aLen)
	a = append(a, make([]float64, aDestLen-aLen)...)
	segLen := aDestLen / routineCnt

	fChan := make(chan float64)
	wg := new(sync.WaitGroup)
	wg.Add(routineCnt)

	for tIdx := 0; tIdx < routineCnt; tIdx++ {
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
