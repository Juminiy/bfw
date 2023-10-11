// 1. all makeCopy() is deep copy

package la

type RealAlgebraContainer interface {
	validate() bool
	get(int, int) float64
	set(int, int, float64)
}
