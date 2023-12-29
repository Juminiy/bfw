package adt

// adt Stand for abstract data type
const (
	defaultPrintLen = 1 << 5
)

// GenericContainer
// Sequential Container include: queue, stack, deque, list, vector(formed from array or slice), forward_list...
// Association Container include: set(formed from map)...
type (
	GenericContainer[T any] interface {
		Empty() bool
		Len() int
	}

	GenericSequentialContainer[T any] interface {
		GetSlice() []T
		SetSlice([]T)
		Empty() bool
		Len() int
	}

	GenericLinkedContainer[T any] interface {
		GetHead() interface{}
		Empty() bool
		Len() int
		Clear()
	}
)

// GenericIterated
// Type Container And Function Pointer that can be iterated
type (
	GenericIterated[T any] interface {
		Next(n ...int)
		Append(...*T)
		Reset()
		Dummy() *T
		Head() *T
		Tail() *T
		End() *T
		Nil() bool
		Len() int
		Swap(int, int)
		Less(int, int) bool
		Equal(int, int) bool
	}
)

// GenericValue
// Generci Value foramt and regular
type (
	GenericZero[T any] interface {
		Zero() T
	}
)
