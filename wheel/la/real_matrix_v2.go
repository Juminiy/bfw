package la

import "C"
import (
	"bfw/wheel/lang"
)

const (
	maxRecursiveDepth = 22
	minAtomMatrixSize = 2
)

var ()

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

func (matrix *Matrix) setPower2ZeroPadding(size ...int) *Matrix {
	destSize := lang.CeilBin(lang.MaxInt(matrix.rowSize, matrix.columnSize))
	if sizeLen := len(size); sizeLen > 0 {
		if sizeLen == 1 && size[0] > destSize {
			destSize = lang.CeilBin(size[0])
		}
		if sizeLen == 2 && (size[0] > destSize || size[1] > destSize) {
			destSize = lang.CeilBin(lang.MaxInt(size[0], size[1]))
		}
	}
	return matrix.zeroPadding(destSize, destSize)
}

func (matrix *Matrix) zeroPadding(rowSize, columnSize int) *Matrix {
	destRowSize := matrix.rowSize
	destColumnSize := matrix.columnSize
	if destColumnSize < columnSize {
		for rowIdx := 0; rowIdx < destRowSize; rowIdx++ {
			matrix.setRowElemAppend(rowIdx, make([]float64, columnSize-destColumnSize))
		}
		destColumnSize = columnSize
		matrix.setColumnSize(destColumnSize)
	}
	if destRowSize < rowSize {
		matrix.setRowAppend(make([][]float64, rowSize-destRowSize))
		matrix.setRowSize(rowSize)
		for rowIdx := destRowSize; rowIdx < rowSize; rowIdx++ {
			matrix.setRow(rowIdx, make([]float64, destColumnSize))
		}
	}
	return matrix
}

func (matrix *Matrix) zeroTruncate(keepRowSize, keepColumnSize int) *Matrix {
	return matrix
}

func (matrix *Matrix) getBlock(rowIIndex, columnIIndex, rowJIndex, columnJIndex int) *Matrix {
	if !matrix.validateIndex(rowIIndex, columnIIndex) ||
		!matrix.validateIndex(rowJIndex, columnJIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	if rowIIndex > rowJIndex ||
		columnIIndex > columnJIndex {
		panic(canNotGetMatrixBlockError)
	}
	bRowSize, bColumnSize := lang.AbsInt(rowJIndex-rowIIndex+1), lang.AbsInt(columnJIndex-columnIIndex+1)
	bMatrix := &Matrix{}
	bMatrix.assignRow(bRowSize, bColumnSize)
	for rowIdx := rowIIndex; rowIdx <= rowJIndex; rowIdx++ {
		bMatrix.setRow(rowIdx-rowIIndex, matrix.getRow(rowIdx)[columnIIndex:columnJIndex+1])
	}
	return bMatrix
}

// speed up the cal, ignore the check func
func (matrix *Matrix) setBlock(rowIIndex, columnIIndex, rowJIndex, columnJIndex int, bMatrix *Matrix) {
	for rowIdx := rowIIndex; rowIdx <= rowJIndex; rowIdx++ {
		for columnIdx := columnIIndex; columnIdx <= columnJIndex; columnIdx++ {
			matrix.slice[rowIdx][columnIdx] = bMatrix.slice[rowIdx-rowIIndex][columnIdx-columnIIndex]
		}
	}
}

func (matrix *Matrix) divBlock(strategy string) *Matrix {
	switch strategy {
	case "2", "b", "bin", "binary":
		{
			return matrix.binDivBlock()
		}
	default:
		{

		}
	}
	return &Matrix{}
}

func (matrix *Matrix) binDivBlock() *Matrix {
	return &Matrix{}
}

func (matrix *Matrix) mulByDivBlock(m *Matrix) *Matrix {
	if !matrix.canMultiply(m) {
		panic(matrixCanNotMultiplyError)
	}
	matrixRowSize, matrixColumnSize := matrix.rowSize, matrix.columnSize
	mRowSize, mColumnSize := m.rowSize, m.columnSize
	destSize := lang.MaxIntCeilBin(matrixRowSize, matrixColumnSize, mRowSize, mColumnSize)
	matrix.zeroPadding(destSize, destSize)
	m.zeroPadding(destSize, destSize)
	resMatrix := MakeZeroMatrix(destSize, destSize)

	simplePhalanxBlockMul(matrix, m, 0, 0, destSize-1, destSize-1, 0, 0, destSize-1, destSize-1, destSize)

	resMatrix.zeroTruncate(matrixRowSize, mColumnSize)
	// whether to recover
	//matrix.zeroTruncate(matrixRowSize, matrixColumnSize)
	//m.zeroTruncate(mRowSize, mColumnSize)
	return resMatrix
}

func phalanxBlockPatch(C11, C12, C21, C22 *Matrix, sz int) *Matrix {
	return &Matrix{}
}

// speed up, none use more check api
// cannot write back to C, error
func simplePhalanx22Mul22(A, B *Matrix, ari, aci, arj, acj, bri, bci, brj, bcj int) *Matrix {
	res := MakeZeroMatrix(minAtomMatrixSize, minAtomMatrixSize)
	//c[0,0] = a[0,0]*b[0,0]+a[0,1]*b[1,0]; c[0,1] = a[0,0]*b[0,1]+a[0,1]*b[1,1]
	//c[1,0] = a[1,0]*b[0,0]+a[1,1]*b[1,0]; c[1,1] = a[1,0]*b[0,1]+a[1,1]*b[1,1]
	res.slice[0][0] = A.slice[ari][aci]*B.slice[bri][bci] + A.slice[ari][acj]*B.slice[arj][aci]
	res.slice[0][1] = A.slice[ari][aci]*B.slice[ari][acj] + A.slice[ari][acj]*B.slice[arj][acj]
	res.slice[1][0] = A.slice[arj][aci]*B.slice[ari][aci] + A.slice[arj][acj]*B.slice[arj][aci]
	res.slice[1][1] = A.slice[acj][aci]*B.slice[ari][acj] + A.slice[arj][acj]*B.slice[arj][acj]
	return res
}

// consider speed up in go application level
// 1. multiple threads
// 2. application level virtual stack to put temp matrix instead gmp itself
// 3. algorithm change
// 4. multiple algorithm calculate and assume the loss cost.
// 5. binary div convert to three-parts div
// 6. 1-5 multiple exams
// 7. cgo to call C lang pure code.
func simplePhalanxBlockMul(A, B *Matrix, ari, aci, arj, acj, bri, bci, brj, bcj, sz int) *Matrix {
	//recursiveDepth := int(math.Log2(float64(sz)))
	if sz == minAtomMatrixSize {
		return simplePhalanx22Mul22(A, B, ari, aci, arj, acj, bri, bci, brj, bcj)
	}
	sz >>= 1
	aris, acis := ari+sz, aci+sz
	aris1, acis1 := aris-1, acis-1
	bris, bcis := bri+sz, bci+sz
	bris1, bcis1 := bris-1, bcis-1
	//A11, A12 := A.getBlock(ari, aci, aris1, acis1), A.getBlock(ari, acis, aris1, acj)
	//A21, A22 := A.getBlock(aris, aci, arj, acis1), A.getBlock(aris, acis, arj, acj)
	//B11, B12 := B.getBlock(bri, bci, bris1, bcis1), B.getBlock(bri, bcis, bris1, bcj)
	//B21, B22 := B.getBlock(bris, bci, brj, bcis1), B.getBlock(bris, bcis, brj, bcj)
	//C11 = A11*B11+A12*B21; C12 = A11*B12+A12+B22
	//C21 = A21*B11+A22*B21; C22 = A21*B12+A22*B22
	// calculate along with put data
	// 1. calculate then put data, the space complexity enlarge * 2
	// 2. use Aij.mul(), the original address will be changed, because A11,A12...,B11,B12,... is weak reference
	// 3. case 1 will not be considered
	P1 := simplePhalanxBlockMul(A, B, ari, aci, aris1, acis1, bri, bci, bris1, bcis1, sz)
	P2 := simplePhalanxBlockMul(A, B, ari, acis, aris1, acj, bris, bci, brj, bcis1, sz)
	P3 := simplePhalanxBlockMul(A, B, ari, aci, aris1, acis1, bri, bcis, bris1, bcj, sz)
	P4 := simplePhalanxBlockMul(A, B, ari, acis, aris1, acj, bris, bcis, brj, bcj, sz)
	P5 := simplePhalanxBlockMul(A, B, aris, aci, arj, acis1, bri, bci, bris1, bcis1, sz)
	P6 := simplePhalanxBlockMul(A, B, aris, acis, arj, acj, bris, bci, brj, bcis1, sz)
	P7 := simplePhalanxBlockMul(A, B, aris, aci, arj, acis1, bri, bcis, bris1, bcj, sz)
	P8 := simplePhalanxBlockMul(A, B, aris, acis, arj, acj, bris, bcis, brj, bcj, sz)
	//C.setBlock(ari, aci, aris1, acis1)
	return phalanxBlockPatch(P1.add(P2), P3.add(P4), P5.add(P6), P7.add(P8), sz)
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
