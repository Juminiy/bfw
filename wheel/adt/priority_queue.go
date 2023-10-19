package adt

type GenericPriorityQueue[T, Compare any] struct {
}

func (heap *GenericPriorityQueue[T, Compare]) Empty() bool {
	return false
}

func (heap *GenericPriorityQueue[T, Compare]) Len() int {
	return 0
}

func (heap *GenericPriorityQueue[T, Compare]) Top() T {
	var t T
	return t
}

func (heap *GenericPriorityQueue[T, Compare]) Push(t T) {

}

func (heap *GenericPriorityQueue[T, Compare]) Pop() {

}
