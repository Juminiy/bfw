package adt

import (
	"fmt"
)

// Deprecated: Use adt.GenericPriorityQueue instead.
type IntHeap struct {
	slice []int
	asc   bool
}

// Deprecated: Use adt.GenericPriorityQueue instead.
func MakeIntHeap(asc bool) *IntHeap {
	h := &IntHeap{}
	h.make(asc)
	return h
}

func (h *IntHeap) make(sort bool) {
	h.asc = sort
}

func (h *IntHeap) Len() int {
	return len(h.slice)
}

func (h *IntHeap) Swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

func (h *IntHeap) Less(i, j int) bool {
	r := h.slice[i] < h.slice[j]
	if h.asc {
		return !r
	} else {
		return r
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
		panic(HeapEmptyError)
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
// 2. asc by Less
// 3. O(log(N))
func (h *IntHeap) down2top() {
	if h.Len() <= 1 {
		return
	}
	curI := h.Len() - 1
	parI := parents(curI)
	for curI > 0 {
		if rightChild(parI) == curI {
			sI := h.iBySort(parI, curI-1, curI)
			h.Swap(sI, parI)
		} else if leftChild(parI) == curI {
			sI := h.iBySort(parI, curI)
			h.Swap(sI, parI)
		} else {
			// never can occur case
			panic(HeapUnknownError)
		}
		curI, parI = parI, parents(parI)
	}
}

// 1. top to down
// 2. asc by Less
// 3. O(log(N))
func (h *IntHeap) top2down() {
	curI := 0
	for curI < h.Len() {
		sI := curI
		if rI := rightChild(curI); rI < h.Len() {
			sI = h.iBySort(curI, rI-1, rI)
		} else if lI := leftChild(curI); lI < h.Len() {
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

func (h *IntHeap) iBySort(i ...int) int {
	if len(i) == 0 {
		panic(HeapIndexError)
	}
	iE := i[0]
	for _, iI := range i {
		if !h.valI(iI) {
			panic(HeapIndexError)
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
