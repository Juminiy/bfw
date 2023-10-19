package adt

import (
	"bfw/wheel/lang"
)

// GenericStack Template Type of stack
type GenericStack[T any] struct {
	slice []T
}

func (stack *GenericStack[T]) GetSlice() []T {
	return stack.slice
}

func (stack *GenericStack[T]) SetSlice(ts []T) {
	stack.slice = ts
}

func (stack *GenericStack[T]) Empty() bool {
	return !lang.ValidateInterfaceArrayOrSlice(stack.slice)
}

func (stack *GenericStack[T]) Len() int {
	if !stack.Empty() {
		return len(stack.slice)
	}
	return 0
}

func (stack *GenericStack[T]) Top() T {
	if !stack.Empty() {
		return stack.slice[stack.Len()-1]
	}
	var t T
	return t
}

func (stack *GenericStack[T]) Push(t T) {
	if stack.Empty() {
		stack.slice = make([]T, 0)
	}
	stack.slice = append(stack.slice, t)
}

func (stack *GenericStack[T]) Pop() {
	if !stack.Empty() {
		stack.slice = stack.slice[:stack.Len()-1]
	}
}
