package la

func (matrix *Matrix) calBlockSize() {

}

// transposeV2
// elemTranspose the Matrix by DivBlock

// A =
// 1 2 3 4 5
// 0 8 6 4 2

// A1 = |	A1T =
// 1 2	|	1 0
// 0 8	|	2 8
// A2 = |	A2T =
// 3 4 	|	3 6
// 6 4	|	4 4
// A3 = |	A3T =
// 5	|	5 2
// 2	|

// AT =
// A1T
// A2T
// A3T

func (matrix *Matrix) transposeV2() *Matrix {
	return matrix
}

// inverseV2
// inverse by LU Composition
func (matrix *Matrix) inverseV2() *Matrix {
	return matrix
}

// mulV2
// no change self
// Matrix Multiply is a fantastic technique, algorithms and papers flow much
// the easiest make effective is to CHANGE THE MULTIPLE ORDER
func (matrix *Matrix) mulV2(m *Matrix) *Matrix {
	if !matrix.canMultiply(m) {
		panic(matrixCanNotMultiplyError)
	}
	resMatrix := &Matrix{}
	resMatrix.assign(matrix.rowSize, m.columnSize)
	for kIdx := 0; kIdx < matrix.columnSize; kIdx++ {
		for iIdx := 0; iIdx < matrix.rowSize; iIdx++ {
			valIK := matrix.slice[iIdx][kIdx]
			for jIdx := 0; jIdx < m.columnSize; jIdx++ {
				resMatrix.slice[iIdx][jIdx] += valIK * m.slice[kIdx][jIdx]
			}
		}
	}
	return resMatrix
}

func (matrix *Matrix) mulByDivBlock(m *Matrix) *Matrix {
	return matrix
}

// mPowerV2
// quick Power
func (matrix *Matrix) mPowerV2(n int) *Matrix {
	resMatrix := matrix.GetIdentity().Matrix()
	for n > 0 {
		if n&1 != 0 {
			resMatrix.MTimes(matrix)
		}
		matrix.MTimes(matrix)
		n >>= 1
	}
	return resMatrix
}
