package lang

import (
	"fmt"
	"sort"
)

const (
	maxRecursiveDepth int = 14
	insertSortSize    int = 1 << 4
	canDisplaySize    int = 1 << 4
)

type Int641DArrayInfo struct {
	slice   []int64
	result  []int64
	initial bool
	sorted  bool
}

func GenerateInt641DArrayInfo(size int, maxAbsVal int64) *Int641DArrayInfo {
	i641da := &Int641DArrayInfo{}
	i641da.slice = GenerateInt641DArray(size, maxAbsVal)
	return i641da
}

func MakeInt641DArray(slice []int64, direct ...bool) *Int641DArrayInfo {
	i641da := &Int641DArrayInfo{}
	if len(direct) > 0 && direct[0] {
		i641da.result = slice
		i641da.initial = true
	} else {
		i641da.slice = slice
	}
	return i641da
}

func (i641da *Int641DArrayInfo) makeCopy() *Int641DArrayInfo {
	i641daCopy := &Int641DArrayInfo{slice: make([]int64, i641da.Len())}
	copy(i641daCopy.slice, i641da.slice)
	return i641daCopy
}

func (i641da *Int641DArrayInfo) resetResult() {
	if !i641da.initial {
		i641da.result = make([]int64, i641da.Len())
		i641da.initial = true
	}
	if i641da.result == nil {
		copy(i641da.result, i641da.slice)
	}
}

func (i641da *Int641DArrayInfo) clearResult() {
	i641da.result = nil
	i641da.initial = false
	i641da.sorted = false
}

func (i641da *Int641DArrayInfo) getResult() []int64 {
	if i641da.sorted {
		return i641da.result
	}
	return nil
}

func (i641da *Int641DArrayInfo) writeResult() {
	if i641da.sorted {
		copy(i641da.slice, i641da.result)
	}
}

func (i641da *Int641DArrayInfo) Len() int {
	destLen := 0
	if i641da.slice != nil {
		destLen = len(i641da.slice)
	}
	if i641da.result != nil {
		destLen = max(destLen, len(i641da.result))
	}
	return destLen
}

func (i641da *Int641DArrayInfo) Less(i, j int) bool {
	return i641da.result[i] < i641da.result[j]
}

func (i641da *Int641DArrayInfo) Swap(i, j int) {
	i641da.result[i], i641da.result[j] = i641da.result[j], i641da.result[i]
}

func (i641da *Int641DArrayInfo) LSort() *Int641DArrayInfo {
	i641da.resetResult()
	sort.Sort(i641da)
	i641da.sorted = true
	return i641da
}

func (i641da *Int641DArrayInfo) QSort() *Int641DArrayInfo {
	i641da.resetResult()
	i641da.quickSort(0, i641da.Len()-1)
	i641da.sorted = true
	return i641da
}

// quickSort faster than LSort
func (i641da *Int641DArrayInfo) quickSort(si, ei int) {
	if si > ei {
		return
	}
	li, ri := si, ei
	pivotIndex := getRandomIntValue(ei, si)
	pivotVal := i641da.result[pivotIndex]
	for li <= ri {
		for ; si < ri && pivotVal < i641da.result[ri]; ri-- {
		}
		for ; li < ei && i641da.result[li] < pivotVal; li++ {
		}
		if li <= ri {
			i641da.Swap(ri, li)
			ri--
			li++
		}
	}
	if li < ei {
		i641da.quickSort(li, ei)
	}
	if si < ri {
		i641da.quickSort(si, ri)
	}
}

func (i641da *Int641DArrayInfo) MSort() *Int641DArrayInfo {
	i641da.resetResult()
	i641da.mergeSort(0, i641da.Len()-1, 0)
	i641da.sorted = true
	return i641da
}

// should use mix sort method, recursive depth too much
// recursive div and conquer cost too much
func (i641da *Int641DArrayInfo) mergeSort(si, ei, depth int) {
	if depth > maxRecursiveDepth ||
		ei-si < insertSortSize {
		i641da.insertSort(si, ei)
	}
	if si >= ei {
		return
	}
	mi := (ei-si)>>1 + si
	ta := make([]int64, i641da.Len())
	i641da.mergeSort(si, mi, depth+1)
	i641da.mergeSort(mi+1, ei, depth+1)
	i641da.mergeSortMerge(ta, si, mi, mi+1, ei)
}

// 0, 3, 4, 7  -> destLen = 8
func (i641da *Int641DArrayInfo) mergeSortMerge(ta []int64, si0, ei0, si1, ei1 int) {
	si00, si01 := si0, si0
	for ; si0 <= ei0 && si1 <= ei1; si00++ {
		if i641da.Less(si0, si1) {
			ta[si00] = i641da.result[si0]
			si0++
		} else {
			ta[si00] = i641da.result[si1]
			si1++
		}
	}
	if si0 <= ei0 {
		copy(ta[si00:ei1+1], i641da.result[si0:ei0+1])
	}
	if si1 <= ei1 {
		copy(ta[si00:ei1+1], i641da.result[si1:ei1+1])
	}
	copy(i641da.result[si01:ei1+1], ta[si01:ei1+1])
}

func (i641da *Int641DArrayInfo) HSort() *Int641DArrayInfo {
	return i641da
}

func (i641da *Int641DArrayInfo) heapSort() {

}

func (i641da *Int641DArrayInfo) insertSort(si, ei int) {
	for i := si + 1; i <= ei; i++ {
		for j := i; j > si && i641da.Less(j, j-1); j-- {
			i641da.Swap(j, j-1)
		}
	}
}

func (i641da *Int641DArrayInfo) Display() *Int641DArrayInfo {
	if i641da.sorted {
		if i641da.Len() <= canDisplaySize {
			fmt.Println(i641da.result)
		} else {
			fmt.Println("too long to can not display")
		}
	} else {
		fmt.Println("has not been sorted")
	}
	return i641da
}

func (i641da *Int641DArrayInfo) CmdDisplay() *Int641DArrayInfo {
	for idx := 0; idx < i641da.Len(); idx++ {
		fmt.Printf("%d ", i641da.result[idx])
	}
	return i641da
}
