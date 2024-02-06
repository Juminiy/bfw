package generic

import "cmp"

type Int0 int

type Float0 float64

type LessFunc[T any] func(T, T) bool

type Ref[A comparable, B cmp.Ordered, C any, Less LessFunc[C]] struct{}

type Ref0 = Ref[int, int, Int0, LessFunc[Int0]]

type Ref1 = Ref[float64, float64, float64, LessFunc[float64]]

type Ref2 = Ref[Int0, Int0, Int0, LessFunc[Int0]]

type Ref3 Ref[int, int, Int0, LessFunc[Int0]]

func LessT[T cmp.Ordered](x, y T) bool {
	return x < y
}
