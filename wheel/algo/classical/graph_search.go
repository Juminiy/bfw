package classical

type AStar[T any] struct {
	g    func(T) int
	h    func(T) int
	open []T
}

func (a *AStar[T]) f(s T) int {
	return a.g(s) + a.h(s)
}
