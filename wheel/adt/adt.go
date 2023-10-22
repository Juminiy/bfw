package adt

// stand for Abstract Data Type

// GenericContainer
// Sequential Container include: queue, stack, deque, list, vector(formed from array or slice), forward_list...
// Association Container include: set(formed from map)...
type GenericContainer[T any] interface {
	Empty() bool
	Len() int
}

type GenericSequentialContainer[T any] interface {
	GetSlice() []T
	SetSlice([]T)
	Empty() bool
	Len() int
}

type GenericLinkedContainer[T any] interface {
	GetHead() interface{}
	Empty() bool
	Len() int
	Clear()
}
