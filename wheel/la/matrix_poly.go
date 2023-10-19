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

type ComplexMatrixPoly struct {
	matrix *ComplexMatrix
	poly   *poly.Poly
}

func (cmp *ComplexMatrixPoly) ComplexMatrix() *ComplexMatrix {
	return &ComplexMatrix{}
}
