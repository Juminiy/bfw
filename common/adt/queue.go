package adt

import (
	"bfw/common/lang"
)

// GenericQueue Template Type of queue
type GenericQueue[T any] struct {
	slice []T
}

func (queue *GenericQueue[T]) GetSlice() []T {
	return queue.slice
}

func (queue *GenericQueue[T]) SetSlice(ts []T) {
	queue.slice = ts
}

func (queue *GenericQueue[T]) Empty() bool {
	return !lang.ValidateInterfaceArrayOrSlice(queue.slice)
}

func (queue *GenericQueue[T]) Len() int {
	if !queue.Empty() {
		return len(queue.slice)
	}
	return 0
}

func (queue *GenericQueue[T]) Front() T {
	if !queue.Empty() {
		return queue.slice[0]
	}
	var t T
	return t
}

func (queue *GenericQueue[T]) Back() T {
	if !queue.Empty() {
		return queue.slice[queue.Len()-1]
	}
	var t T
	return t
}

func (queue *GenericQueue[T]) Push(t T) {
	if queue.Empty() {
		queue.slice = make([]T, 0)
	}
	queue.slice = append(queue.slice, t)
}

func (queue *GenericQueue[T]) Pop() {
	if !queue.Empty() {
		queue.slice = queue.slice[1:]
	}
}
