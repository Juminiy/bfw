package adt

import (
	"errors"
	"fmt"
)

var (
	intHeapUnknownError = errors.New("IntHeap UnKnown Error")
	intHeapIndexError   = errors.New("IntHeap Index Error")
	intHeapEmptyError   = errors.New("IntHeap Empty Error")
)

type IntHeap struct {
	slice []int
	sort  bool
}

func MakeIntHeap(sort bool) *IntHeap {
	h := &IntHeap{}
	h.make(sort)
	return h
}

func (h *IntHeap) make(sort bool) {
	h.sort = sort
}

func (h *IntHeap) Len() int {
	return len(h.slice)
}

func (h *IntHeap) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func (h *IntHeap) Less(i, j int) bool {
	r := h.slice[i] < h.slice[j]
	if h.sort {
		return r
	} else {
		return !r
	}
}

func (h *IntHeap) Empty() bool {
	return h.Len() == 0
}

func (h *IntHeap) Push(e ...int) {
	for _, eE := range e {
		h.push(eE)
	}
}

func (h *IntHeap) Pop(n int) []int {
	e := make([]int, 0)
	for n > 0 && !h.Empty() {
		e = append(e, h.pop())
		n--
	}
	return e
}

func (h *IntHeap) push(e int) {
	h.slice = append(h.slice, e)
	h.adjust(true)
}

func (h *IntHeap) pop() int {
	if h.Empty() {
		panic(intHeapEmptyError)
	}
	si, ei := 0, h.Len()-1
	h.Swap(si, ei)
	e := h.slice[ei]
	h.slice = h.slice[:ei]
	h.adjust(false)
	return e
}

func (h *IntHeap) adjust(down2top bool) {
	if down2top {
		h.down2top()
	} else {
		h.top2down()
	}
}

// 1. down to top
// 2. sort by Less
// 3. O(log(N))
func (h *IntHeap) down2top() {
	if h.Len() <= 1 {
		return
	}
	curI := h.Len() - 1
	parI := h.up(curI)
	for curI > 0 {
		if h.right(parI) == curI {
			sI := h.iBySort(parI, curI-1, curI)
			h.Swap(sI, parI)
		} else if h.left(parI) == curI {
			sI := h.iBySort(parI, curI)
			h.Swap(sI, parI)
		} else {
			// never can occur case
			panic(intHeapUnknownError)
		}
		curI, parI = parI, h.up(parI)
	}
}

// 1. top to down
// 2. sort by Less
// 3. O(log(N))
func (h *IntHeap) top2down() {
	curI := 0
	for curI < h.Len() {
		sI := curI
		if rI := h.right(curI); rI < h.Len() {
			sI = h.iBySort(curI, rI-1, rI)
		} else if lI := h.left(curI); lI < h.Len() {
			sI = h.iBySort(curI, lI)
		} else {
			break
		}
		if sI != curI {
			h.Swap(sI, curI)
			curI = sI
		} else {
			break
		}
	}
}

func (h *IntHeap) up(i int) int {
	return (i - 1) >> 1
}

func (h *IntHeap) left(i int) int {
	return i<<1 + 1
}

func (h *IntHeap) right(i int) int {
	return i<<1 + 2
}

func (h *IntHeap) iBySort(i ...int) int {
	if len(i) == 0 {
		panic(intHeapIndexError)
	}
	iE := i[0]
	for _, iI := range i {
		if !h.valI(iI) {
			panic(intHeapIndexError)
		}
		if h.Less(iE, iI) {
			iE = iI
		}
	}
	return iE
}

func (h *IntHeap) valI(i ...int) bool {
	for _, iI := range i {
		if !h.valII(iI) {
			return false
		}
	}
	return true
}

func (h *IntHeap) valII(i int) bool {
	return i >= 0 && i < h.Len()
}

func (h *IntHeap) Print() {
	fmt.Println(h.slice)
}
