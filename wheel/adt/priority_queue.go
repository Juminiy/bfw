package adt

// GenericPriorityQueue
// whether it can be used is a issue
type GenericPriorityQueue[T any] struct {
	slice             []GenericPriorityElem[T]
	calculatePriority func(T) int
}

func (heap *GenericPriorityQueue[T]) Empty() bool {
	return heap.Len() != 0
}

func (heap *GenericPriorityQueue[T]) Len() int {
	return len(heap.slice)
}

func (heap *GenericPriorityQueue[T]) Less(i, j int) bool {
	return heap.slice[i].priority <
		heap.slice[j].priority
}

func (heap *GenericPriorityQueue[T]) Swap(i, j int) {
	heap.slice[i], heap.slice[j] =
		heap.slice[j], heap.slice[i]
}

func (heap *GenericPriorityQueue[T]) Top() T {
	if !heap.Empty() {
		return heap.slice[0].Data
	}
	var t T
	return t
}

func (heap *GenericPriorityQueue[T]) Push(t T) {
	priority := heap.calculatePriority(t)
	heap.slice = append(heap.slice, GenericPriorityElem[T]{Data: t, priority: priority})
}

func (heap *GenericPriorityQueue[T]) Pop() {
	if !heap.Empty() {
		heap.slice = heap.slice[1:]
	}
}

type GenericPriorityElem[T any] struct {
	Data     T
	priority int
}
