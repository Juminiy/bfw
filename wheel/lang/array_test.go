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

func TestReal2DArray_Len(t *testing.T) {
	arr := [32]byte{'0', '1', '2', '3', '4', '5'}
	sl1 := arr[2:4]
	sl2 := arr[3:7]
	fmt.Println(arr, "\n", sl1, "\n", sl2)
	sl1[0] = 'x'
	sl2[1] = 'X'
	fmt.Println(arr, "\n", sl1, "\n", sl2)
	revByteSlice := func(ba []byte) {
		baL := len(ba)
		for i := 0; i < (baL >> 1); i++ {
			ba[i], ba[baL-i-1] = ba[baL-i-1], ba[i]
		}
	}
	fmt.Println(sl1)
	revByteSlice(sl1)
	fmt.Println(sl1)
}

func TestRunChangeSlice(t *testing.T) {
	var arr []int
	fmt.Printf("nih:(%p(%p))\n", &arr, arr)
	RunChangeSlice(&arr)
	fmt.Printf("nih:(%p(%p))\n", &arr, arr)
	fmt.Println(arr)
	RunNoChangeSlice(arr)
	fmt.Printf("nih:(%p(%p))\n", &arr, arr)
	fmt.Println(arr)
}

//nih:(0x1400012c0a8(0x0))
//param address: 0x14000106048
//param value address: 0x1400012c0a8
//param value pointer to address: 0x1400011c060
//nih:(0x1400012c0a8(0x1400011c060))
//param address: 0x1400012c0f0
//param value address: 0x1400012e060
//param value pointer to address: 0
//nih:(0x1400012c0a8(0x1400011c060))
