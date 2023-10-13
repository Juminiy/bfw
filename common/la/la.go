// la
// 1. makeCopy() is deep copy
// 2. ConstructXXX() is for global use

// stand for Linear Algebra

package la

type RealAlgebraContainer interface {
	validate() bool
	get(int, int) float64
	set(int, int, float64)
}
