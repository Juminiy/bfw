package adt

type Iterator[T any] struct {
	pup, cur *T
}

func MakeIterator[T any](t *T) *Iterator[T] {
	return &Iterator[T]{t, t}
}

func ValidateI[T any](i ...*Iterator[T]) bool {
	for _, iI := range i {
		if iI == nil || iI.Nil() {
			return false
		}
	}
	return true
}

func (i *Iterator[T]) Next(n ...int) {
	nN := 1
	if len(n) > 0 && n[0] >= 0 {
		nN = n[0]
	}
	for i.cur != nil && nN > 0 {
		nN--
		//i.cur = i.cur.Next
	}
}

func (i *Iterator[T]) Append(t ...*T) {
	for _, tT := range t {
		//i.cur.Next = tT
		i.cur = tT
	}
}

func (i *Iterator[T]) AppendI(it ...*Iterator[T]) {
	for _, itT := range it {
		i.Append(itT.cur)
	}
}

func (i *Iterator[T]) Reset() {
	i.cur = i.pup
}

func (i *Iterator[T]) Dummy() *T {
	return i.pup
}
func (i *Iterator[T]) Head() *T {
	//return i.pup.Next
	return i.pup
}
func (i *Iterator[T]) Tail() *T {
	return i.cur
}
func (i *Iterator[T]) End() *T {
	cur := i.cur
	for cur != nil { //&& cur.Next != nil
		//cur= cur.Next
	}
	return cur
}

func (i *Iterator[T]) Nil() bool {
	return i.cur == nil
}

func (i *Iterator[T]) Len() int {
	cur, curLen := i.cur, 0
	for cur != nil {
		//cur = cur.Next
		curLen++
	}
	return curLen
}

func (i *Iterator[T]) Swap(xi, xj int) {
	//i.Next(xi)
	//tXi := i.cur
	//i.Next(xj - xi)
	//tXj := i.cur
	//
}

func (i *Iterator[T]) Less(xi, xj int) bool {
	return xi < xj
}

func (i *Iterator[T]) Equal(xi, xj int) bool {
	return xi == xj
}
