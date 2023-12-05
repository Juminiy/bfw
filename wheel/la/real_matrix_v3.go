package la

func (matrix *Matrix) addAB(A, B *Matrix) {
	matrix.setSelf(A.GetPlus(B))
}

func (matrix *Matrix) subAB(A, B *Matrix) {
	matrix.setSelf(A.GetMinus(B))
}

func (matrix *Matrix) mulAB(A, B *Matrix) {
	matrix.setSelf(A.mulV2(B))
}

func (matrix *Matrix) powerAn(A *Matrix, n int) {
	matrix.setSelf(A.mPowerV2(n))
}
