package adt

import (
	"errors"
	"fmt"
)

var (
	HeapUnknownError = errors.New("heap UnKnown Error")
	HeapIndexError   = errors.New("heap Index Error")
	HeapEmptyError   = errors.New("heap Empty Error")
)

// GenericPriorityQueue
// control assign to function pointer, for example: less,deepCopy,...
type GenericPriorityQueue[T any] struct {
	slice []T
	less  func(T, T) bool
	print func(T)
}

func MakeHeap[T any](less func(T, T) bool) *GenericPriorityQueue[T] {
	h := &GenericPriorityQueue[T]{}
	h.make(less)
	return h
}

func (h *GenericPriorityQueue[T]) make(less func(T, T) bool) {
	h.less = less
}

func (h *GenericPriorityQueue[T]) Empty() bool {
	return h.Len() == 0
}

func (h *GenericPriorityQueue[T]) Len() int {
	return len(h.slice)
}

func (h *GenericPriorityQueue[T]) Less(i, j int) bool {
	return !h.less(h.slice[i], h.slice[j])
}

func (h *GenericPriorityQueue[T]) Swap(i, j int) {
	h.slice[i], h.slice[j] =
		h.slice[j], h.slice[i]
}

func (h *GenericPriorityQueue[T]) Push(t ...T) {
	for _, tT := range t {
		h.push(tT)
	}
}

func (h *GenericPriorityQueue[T]) Pop(n int) []T {
	t := make([]T, 0)
	for n > 0 && !h.Empty() {
		t = append(t, h.pop())
		n--
	}
	return t
}

func (h *GenericPriorityQueue[T]) push(t T) {
	h.slice = append(h.slice, t)
	h.adjust(true)
}

func (h *GenericPriorityQueue[T]) pop() T {
	if h.Empty() {
		panic(HeapEmptyError)
	}
	si, ei := 0, h.Len()-1
	h.Swap(si, ei)
	t := h.slice[ei]
	h.slice = h.slice[:ei]
	h.adjust(false)
	return t
}

func (h *GenericPriorityQueue[T]) adjust(down2top bool) {
	if down2top {
		h.down2top()
	} else {
		h.top2down()
	}
}

func (h *GenericPriorityQueue[T]) down2top() {
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

func (h *GenericPriorityQueue[T]) top2down() {
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

func (h *GenericPriorityQueue[T]) iBySort(i ...int) int {
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

func (h *GenericPriorityQueue[T]) valI(i ...int) bool {
	for _, iI := range i {
		if !h.valII(iI) {
			return false
		}
	}
	return true
}

func (h *GenericPriorityQueue[T]) valII(i int) bool {
	return i >= 0 && i < h.Len()
}

func (h *GenericPriorityQueue[T]) Print() {
	for _, t := range h.slice {
		h.print(t)
		fmt.Print(" ")
	}
	fmt.Println()
}

func parents(i int) int {
	return (i - 1) >> 1
}

func leftChild(i int) int {
	return i<<1 + 1
}

func rightChild(i int) int {
	return i<<1 + 2
}
