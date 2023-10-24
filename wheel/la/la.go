package la

// la
// 1. makeCopy() is deep copy
// 2. ConstructXXX() is for global use

// la Stand for linear algebra

type RealAlgebraContainer interface {
	validate() bool
	get(int, int) float64
	set(int, int, float64)
}
