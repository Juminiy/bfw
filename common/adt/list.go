package adt

import "errors"

var (
	genericListHeadAndTailError = errors.New("GenericList head and tail error")
	genericListHeadError        = errors.New("GenericList head error")
	genericListTailError        = errors.New("GenericList tail error")
	genericListSizeError        = errors.New("GenericList size error")
)

// GenericList Template Type of forward_list
type GenericList[T any] struct {
	head *GenericListNode[T]
	tail *GenericListNode[T]
	size int
}

func (list *GenericList[T]) validate() error {
	if list.size < 0 {
		return genericListSizeError
	} else if list.size == 0 {
		if list.head != nil || list.tail != nil {
			return genericListHeadAndTailError
		}
	} else {
		if list.head == nil ||
			list.tail == nil ||
			list.head.Prev != nil ||
			list.tail.Next != nil {
			return genericListHeadAndTailError
		}
		if list.size == 1 {
			if list.head != list.tail ||
				list.head.Next == list.tail ||
				list.tail.Prev == list.head {
				return genericListHeadAndTailError
			}
		}
	}
	return nil
}

// GetHead
// O(1)
func (list *GenericList[T]) GetHead() interface{} {
	return list.head
}

// Empty
// O(1)
func (list *GenericList[T]) Empty() bool {
	err := list.validate()
	if err != nil {
		panic(err)
	}
	return list.size == 0
}

// Len
// O(1)
func (list *GenericList[T]) Len() int {
	if !list.Empty() {
		return list.size
	}
	return 0
}

type GenericListNode[T any] struct {
	Data T
	Prev *GenericListNode[T]
	Next *GenericListNode[T]
}

func (node *GenericListNode[T]) construct(t T, prev, next *GenericListNode[T]) {
	node.Data = t
	node.Prev = prev
	node.Next = next
}

func (node *GenericListNode[T]) Assign(nd *GenericListNode[T]) {
	if nd != nil {
		node.construct(nd.Data, nd.Prev, nd.Next)
	} else {
		node.SetZero()
	}
}

func (node *GenericListNode[T]) SetZero() {
	var t T
	node.construct(t, nil, nil)
}

// addHead
// O(1)
// head    tail
// |       |
// 2 < - > 1
func (list *GenericList[T]) addHead(node *GenericListNode[T]) error {
	if node != nil {
		node.Prev = nil
		node.Next = nil
		if list.head == nil {
			list.head = node
			if list.tail == nil {
				if list.size == 0 {
					list.tail = node
				} else {
					return genericListSizeError
				}
			} else {
				return genericListTailError
			}
		} else {
			node.Next = list.head
			list.head.Prev = node
			list.head = node
		}

		list.size++
	}

	return nil
}

// delHead
// O(1)
// head tail
// |     |
// 1  -  1
func (list *GenericList[T]) delHead() (*GenericListNode[T], error) {
	var (
		node *GenericListNode[T]
	)
	if list.head == nil {
		if list.tail == nil {
			if list.size == 0 {
				return nil, nil
			} else {
				return nil, genericListSizeError
			}
		} else {
			return nil, genericListTailError
		}
	} else {
		if list.size == 1 {
			if list.head == list.tail {
				node = list.head
				list.head = nil
				list.tail = nil
			} else {
				return nil, genericListHeadAndTailError
			}
		} else {
			node = list.head
			list.head = list.head.Next
			list.head.Prev = nil
		}
		node.Prev = nil
		node.Next = nil
		list.size--
	}

	return node, nil
}

// addTail
// O(1)
func (list *GenericList[T]) addTail(node *GenericListNode[T]) error {
	if node != nil {
		node.Prev = nil
		node.Next = nil
		if list.tail == nil {
			list.tail = node
			if list.head == nil {
				if list.size == 0 {
					list.head = node
				} else {
					return genericListSizeError
				}
			} else {
				return genericListHeadError
			}
		} else {
			node.Prev = list.tail
			list.tail.Next = node
			list.tail = node
		}

		list.size++
	}

	return nil
}

// delTail
// O(1)
func (list *GenericList[T]) delTail() (*GenericListNode[T], error) {
	var (
		node *GenericListNode[T]
	)
	if list.tail == nil {
		if list.head == nil {
			if list.size == 0 {
				return nil, nil
			} else {
				return nil, genericListSizeError
			}
		} else {
			return nil, genericListHeadError
		}
	} else {
		if list.size == 1 {
			if list.tail == list.head {
				node = list.tail
				list.tail = nil
				list.head = nil
			} else {
				return nil, genericListHeadAndTailError
			}
		} else {
			node = list.tail
			list.tail = list.tail.Prev
			list.tail.Next = nil
		}
		node.Prev = nil
		node.Next = nil
		list.size--
	}

	return node, nil
}

// Front
// O(1)
func (list *GenericList[T]) Front() T {
	if list.Empty() {
		var t T
		return t
	}

	return list.head.Data
}

// Back
// O(1)
func (list *GenericList[T]) Back() T {
	if list.Empty() {
		var t T
		return t
	}

	return list.tail.Data
}

// PushFront
// O(1)
func (list *GenericList[T]) PushFront(t T) {
	err := list.validate()
	if err != nil {
		panic(err)
	}

	err = list.addHead(&GenericListNode[T]{Data: t})
	if err != nil {
		panic(err)
	}
}

// PopFront
// O(1)
func (list *GenericList[T]) PopFront() {
	err := list.validate()
	if err != nil {
		panic(err)
	}

	_, err = list.delTail()
	if err != nil {
		panic(err)
	}
}

// PushBack
// O(1)
func (list *GenericList[T]) PushBack(t T) {
	err := list.validate()
	if err != nil {
		panic(err)
	}

	err = list.addTail(&GenericListNode[T]{Data: t})
	if err != nil {
		panic(err)
	}
}

// PopBack
// O(1)
func (list *GenericList[T]) PopBack() {
	err := list.validate()
	if err != nil {
		panic(err)
	}

	_, err = list.delTail()
	if err != nil {
		panic(err)
	}
}

// ForwardTraverse
// O(N)
func (list *GenericList[T]) ForwardTraverse(funcPtr func(params ...any) (int, error)) {
	err := list.validate()
	if err != nil {
		panic(err)
	}
	var (
		node = list.head
	)
	if node != nil {
		for node != nil {
			funcPtr(node.Data)
			node = node.Next
		}
	}
	// head is nil, do nothing
}

// ReverseTraverse
// O(N)
func (list *GenericList[T]) ReverseTraverse(funcPtr func(params ...any) (int, error)) {
	err := list.validate()
	if err != nil {
		panic(err)
	}
	var (
		node = list.tail
	)
	if node != nil {
		for node != nil {
			funcPtr(node.Data)
			node = node.Prev
		}
	}
	// tail is nil, do nothing
}

func (list *GenericList[T]) construct(head, tail *GenericListNode[T], size int) {
	list.head = head
	list.tail = tail
	list.size = size
}

func (list *GenericList[T]) Assign(ts *GenericList[T]) {
	if ts != nil {
		err := ts.validate()
		if err != nil {
			panic(err)
		}
		list.construct(ts.head, ts.tail, ts.size)
	} else {
		list.Clear()
	}
}

// Swap
// O(1)
// can not be used, go do not permit common assertion in generic
func (list *GenericList[T]) Swap(ts *GenericList[T]) {
	var tsCopy *GenericList[T]
	tsCopy.Assign(ts)
	ts.Assign(list)
	list.Assign(tsCopy)
}

func (list *GenericList[T]) Clear() {
	list.construct(nil, nil, 0)
}

func (list *GenericList[T]) At(index int) T {
	var t T
	return t
}

// Merge
// O(1)
// can not be used, go do not permit common assertion in generic
func (list *GenericList[T]) Merge(ts *GenericList[T]) {
	err := list.validate()
	if err != nil {
		panic(err)
	}

	if ts != nil {
		err = ts.validate()
		if err != nil {
			panic(err)
		}
		if list.size == 0 {
			list.Assign(ts)
		} else {
			if ts.size != 0 {
				list.tail.Next = ts.head
				list.tail = ts.tail
				list.size += ts.size
			}
		}
	}

}

func (list *GenericList[T]) Insert(index int, t T) {

}

func (list *GenericList[T]) Erase(index int) {

}

func (list *GenericList[T]) Unique() {

}

func (list *GenericList[T]) Sort(func(comp, comped T) bool) {

}

func (list *GenericList[T]) ChainedPushFront(t T) *GenericList[T] {
	list.PushFront(t)
	return list
}

func (list *GenericList[T]) ChainedPopFront() *GenericList[T] {
	list.PopFront()
	return list
}

func (list *GenericList[T]) ChainedPushBack(t T) *GenericList[T] {
	list.PushBack(t)
	return list
}

func (list *GenericList[T]) ChainedPopBack() *GenericList[T] {
	list.PopBack()
	return list
}
