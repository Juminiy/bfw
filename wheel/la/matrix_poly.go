package la

import (
	"bfw/wheel/poly"
)

type MatrixPoly struct {
	matrix *Matrix
	poly   *poly.Poly
}

func (mp *MatrixPoly) Matrix() *Matrix {
	return &Matrix{}
}

// Hamilton Cayley
// f(A)*g(A) = 0

type ComplexMatrixPoly struct {
	matrix *ComplexMatrix
	poly   *poly.Poly
}

func (cmp *ComplexMatrixPoly) ComplexMatrix() *ComplexMatrix {
	return &ComplexMatrix{}
}
