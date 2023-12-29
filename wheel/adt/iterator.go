package adt

type LinkedIterator[T any] struct {
	pup, cur *T
}

func MakeIterator[T any](t *T) *LinkedIterator[T] {
	return &LinkedIterator[T]{t, t}
}

func ValidateI[T any](i ...*LinkedIterator[T]) bool {
	for _, iI := range i {
		if iI == nil || iI.Nil() {
			return false
		}
	}
	return true
}

func (i *LinkedIterator[T]) Next(n ...int) {
	nN := 1
	if len(n) > 0 && n[0] >= 0 {
		nN = n[0]
	}
	for i.cur != nil && nN > 0 {
		nN--
		//i.cur = i.cur.Next
	}
}

func (i *LinkedIterator[T]) Append(t ...*T) {
	for _, tT := range t {
		//i.cur.Next = tT
		i.cur = tT
	}
}

func (i *LinkedIterator[T]) AppendI(it ...*LinkedIterator[T]) {
	for _, itT := range it {
		i.Append(itT.cur)
	}
}

func (i *LinkedIterator[T]) Reset() {
	i.cur = i.pup
}

func (i *LinkedIterator[T]) Dummy() *T {
	return i.pup
}
func (i *LinkedIterator[T]) Head() *T {
	//return i.pup.Next
	return i.pup
}
func (i *LinkedIterator[T]) Tail() *T {
	return i.cur
}
func (i *LinkedIterator[T]) End() *T {
	cur := i.cur
	for cur != nil { //&& cur.Next != nil
		//cur= cur.Next
	}
	return cur
}

func (i *LinkedIterator[T]) Nil() bool {
	return i.cur == nil
}

func (i *LinkedIterator[T]) Len() int {
	cur, curLen := i.cur, 0
	for cur != nil {
		//cur = cur.Next
		curLen++
	}
	return curLen
}

func (i *LinkedIterator[T]) Swap(xi, xj int) {
	//i.Next(xi)
	//tXi := i.cur
	//i.Next(xj - xi)
	//tXj := i.cur
	//
}

func (i *LinkedIterator[T]) Less(xi, xj int) bool {
	return xi < xj
}

func (i *LinkedIterator[T]) Equal(xi, xj int) bool {
	return xi == xj
}
