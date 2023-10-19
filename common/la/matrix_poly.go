package la

import "bfw/common/cal"

type MatrixPoly struct {
	matrix *Matrix
	poly   *cal.Poly
}

func (mp *MatrixPoly) Matrix() *Matrix {
	return &Matrix{}
}

type ComplexMatrixPoly struct {
	matrix *ComplexMatrix
	poly   *cal.Poly
}

func (cmp *ComplexMatrixPoly) ComplexMatrix() *ComplexMatrix {
	return &ComplexMatrix{}
}
