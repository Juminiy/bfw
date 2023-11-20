package lang

import (
	"fmt"
	"testing"
	"time"
)

func TestInt1DArrayContribute(t *testing.T) {
	size, maxAbsVal := 1<<32, 1<<48
	time0 := time.Now()
	i641da := GenerateInt641DArrayInfo(size, int64(maxAbsVal))
	fmt.Printf("generate %d size int64 array time:%v\n", size, time.Since(time0))

	time0 = time.Now()
	i641da.QSort()
	fmt.Printf("my quick sort %d time:%v\n", size, time.Since(time0))
	i641da.Display()
	qSortRes := i641da.getResult()

	time0 = time.Now()
	i641da.LSort()
	fmt.Printf("lang mix sort %d time:%v\n", size, time.Since(time0))
	i641da.Display()
	lSortRes := i641da.getResult()

	//time0 = time.Now()
	//i641da.MSort()
	//fmt.Printf("my merge sort %d time:%v\n", size, time.Since(time0))
	//i641da.Display()
	//mSortRes := i641da.getResult()
	//EqualInt641DArray(lSortRes, mSortRes)

	time0 = time.Now()
	fmt.Println(EqualInt641DArray(lSortRes, qSortRes))
	fmt.Printf("traverse %d size int64 array time:%v\n", size, time.Since(time0))

}
