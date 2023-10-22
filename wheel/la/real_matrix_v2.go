package la

import (
	"bfw/wheel/lang"
	"fmt"
	"math"
)

const (
	// max rec depth follow MNN
	maxRecursiveDepth = 15
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

func (matrix *Matrix) mulV1(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.mul(m)
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

func (matrix *Matrix) mulV2Dot5(m *Matrix) *Matrix {
	return matrix.mulByRotate2(m)
}

func (matrix *Matrix) mulV3(m *Matrix) *Matrix {
	return matrix.mulByDivBlock(m, SimplePhalanxBlockMul)
}

func (matrix *Matrix) mulV4(m *Matrix) *Matrix {
	return matrix.mulByDivBlock(m, SimplePhalanxBlockMulV2)
}

func (matrix *Matrix) mulV5(m *Matrix) *Matrix {
	return matrix.mulByDivBlock(m, StraseenPhalanxBlockMul)
}

func (matrix *Matrix) getRowSum(rowIndex int) float64 {
	sum := 0.0
	for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
		sum += matrix.get(rowIndex, columnIdx)
	}
	return sum
}

func (matrix *Matrix) rowDot(mRowIndex int, mt *Matrix, mtRowIndex int) float64 {
	return ConstructVector(matrix.getRow(mRowIndex)).DotMul(ConstructVector(mt.getRow(mtRowIndex)))
}

// 1 2 | 5 6
// 3 4 | 7 8
// 1 2 | 5 7
// 3 4 | 6 8
func (matrix *Matrix) mulByRotate2(m *Matrix) *Matrix {
	if !matrix.canMultiply(m) {
		panic(matrixCanNotMultiplyError)
	}
	mCopy := m.makeCopy()
	mCopy.transpose()
	resMatrix := &Matrix{}
	resMatrix.assign(matrix.rowSize, m.columnSize)
	for rowIIdx := 0; rowIIdx < matrix.rowSize; rowIIdx++ {
		for rowJIdx := 0; rowJIdx < mCopy.rowSize; rowJIdx++ {
			resMatrix.set(rowIIdx, rowJIdx, matrix.rowDot(rowIIdx, mCopy, rowJIdx))
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
	matrix.setRowTruncate(keepRowSize)
	for rowIdx := 0; rowIdx < keepRowSize; rowIdx++ {
		matrix.setRowElemTruncate(rowIdx, keepColumnSize)
	}
	matrix.setColumnSize(keepColumnSize)
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

func (matrix *Matrix) mulByDivBlock(m *Matrix, matrixMultiplyAlgorithm ...func(*Matrix, *Matrix, int, int, int, int, int, int, int, int, int) *Matrix) *Matrix {
	if !matrix.canMultiply(m) {
		panic(matrixCanNotMultiplyError)
	}
	destMatrixMultiplyAlgorithm := SimplePhalanxBlockMul
	if mmaLen := len(matrixMultiplyAlgorithm); mmaLen > 0 {
		if mmaLen == 1 {
			destMatrixMultiplyAlgorithm = matrixMultiplyAlgorithm[0]
		}
	}
	matrixRowSize, matrixColumnSize := matrix.rowSize, matrix.columnSize
	mRowSize, mColumnSize := m.rowSize, m.columnSize
	destSize := lang.MaxIntCeilBin(matrixRowSize, matrixColumnSize, mRowSize, mColumnSize)
	matrix.zeroPadding(destSize, destSize)
	m.zeroPadding(destSize, destSize)

	resMatrix := destMatrixMultiplyAlgorithm(matrix, m, 0, 0, destSize-1, destSize-1, 0, 0, destSize-1, destSize-1, destSize)
	resMatrix.zeroTruncate(matrixRowSize, mColumnSize)

	// whether to recover
	//matrix.zeroTruncate(matrixRowSize, matrixColumnSize)
	//m.zeroTruncate(mRowSize, mColumnSize)
	return resMatrix
}

func (matrix *Matrix) phalanxBlockRowElemPatch(m *Matrix) *Matrix {
	size := matrix.getPhalanxSize()
	for rowIdx := 0; rowIdx < size; rowIdx++ {
		matrix.setRowElemAppend(rowIdx, m.getRow(rowIdx))
	}
	return matrix
}

func (matrix *Matrix) phalanxBlockRowPatch(m *Matrix) *Matrix {
	matrix.setRowAppend(m.getSlice())
	return matrix
}

// ----------|
// | C11 C12 |
// | C21 C22 |
// |----------
func phalanxBlock22Patch(C11, C12, C21, C22 *Matrix) *Matrix {
	return C11.phalanxBlockRowElemPatch(C12).
		phalanxBlockRowPatch(
			C21.phalanxBlockRowElemPatch(C22)).
		setRowSizeInc(C21.rowSize).
		setColumnSizeInc(C12.columnSize)
}

// speed up, none use more check api
// cannot write back to C, error
// Test Pass
func simplePhalanx22Mul22(A, B *Matrix, ari, aci, arj, acj, bri, bci, brj, bcj int) *Matrix {
	res := MakeZeroMatrix(minAtomMatrixSize, minAtomMatrixSize)
	//c[0,0] = a[0,0]*b[0,0]+a[0,1]*b[1,0]; c[0,1] = a[0,0]*b[0,1]+a[0,1]*b[1,1]
	//c[1,0] = a[1,0]*b[0,0]+a[1,1]*b[1,0]; c[1,1] = a[1,0]*b[0,1]+a[1,1]*b[1,1]
	res.slice[0][0] = A.slice[ari][aci]*B.slice[bri][bci] + A.slice[ari][acj]*B.slice[brj][bci]
	res.slice[0][1] = A.slice[ari][aci]*B.slice[bri][bcj] + A.slice[ari][acj]*B.slice[brj][bcj]
	res.slice[1][0] = A.slice[arj][aci]*B.slice[bri][bci] + A.slice[arj][acj]*B.slice[brj][bci]
	res.slice[1][1] = A.slice[arj][aci]*B.slice[bri][bcj] + A.slice[arj][acj]*B.slice[brj][bcj]
	return res
}

func debugPrintln(a ...any) {
	fmt.Println(a...)
}

// SimplePhalanxBlockMul
// consider speed up in go application level
// 1. multiple threads
// 2. application level virtual stack to put temp matrix instead gmp itself
// 3. algorithm change
// 4. multiple algorithm calculate and assume the loss cost.
// 5. binary div convert to three-parts div
// 6. 1-5 multiple exams
// 7. cgo to call C lang pure code.
func SimplePhalanxBlockMul(A, B *Matrix, ari, aci, arj, acj, bri, bci, brj, bcj, sz int) *Matrix {
	recursiveDepth := int(math.Log2(float64(sz)))
	if recursiveDepth >= maxRecursiveDepth {
		return A.getBlock(ari, aci, arj, acj).mulV2(B.getBlock(bri, bci, brj, bcj))
	}
	if sz == minAtomMatrixSize {
		return simplePhalanx22Mul22(A, B, ari, aci, arj, acj, bri, bci, brj, bcj)
	}
	sz >>= 1
	aris, acis := ari+sz, aci+sz
	aris1, acis1 := aris-1, acis-1
	bris, bcis := bri+sz, bci+sz
	bris1, bcis1 := bris-1, bcis-1

	// calculate along with put data
	// 1. calculate then put data, the space complexity enlarge * 2
	// 2. use Aij.mul(), the original address will be changed, because A11,A12...,B11,B12,... is weak reference
	// 3. case 1 will not be considered
	//debugPrintln("A11", ari, aci, aris1, acis1, "B11", bri, bci, bris1, bcis1)
	//debugPrintln("A12", ari, acis, aris1, acj, "B21", bris, bci, brj, bcis1)
	//debugPrintln("A11", ari, aci, aris1, acis1, "B12", bri, bcis, bris1, bcj)
	//debugPrintln("A12", ari, acis, aris1, acj, "B22", bris, bcis, brj, bcj)
	//debugPrintln("A21", aris, aci, arj, acis1, "B11", bri, bci, bris1, bcis1)
	//debugPrintln("A22", aris, acis, arj, acj, "B21", bris, bci, brj, bcis1)
	//debugPrintln("A21", aris, aci, arj, acis1, "B12", bri, bcis, bris1, bcj)
	//debugPrintln("A22", aris, acis, arj, acj, "B22", bris, bcis, brj, bcj)

	C11 := SimplePhalanxBlockMul(A, B, ari, aci, aris1, acis1, bri, bci, bris1, bcis1, sz).add(
		SimplePhalanxBlockMul(A, B, ari, acis, aris1, acj, bris, bci, brj, bcis1, sz))

	C12 := SimplePhalanxBlockMul(A, B, ari, aci, aris1, acis1, bri, bcis, bris1, bcj, sz).add(
		SimplePhalanxBlockMul(A, B, ari, acis, aris1, acj, bris, bcis, brj, bcj, sz))

	C21 := SimplePhalanxBlockMul(A, B, aris, aci, arj, acis1, bri, bci, bris1, bcis1, sz).add(
		SimplePhalanxBlockMul(A, B, aris, acis, arj, acj, bris, bci, brj, bcis1, sz))

	C22 := SimplePhalanxBlockMul(A, B, aris, aci, arj, acis1, bri, bcis, bris1, bcj, sz).add(
		SimplePhalanxBlockMul(A, B, aris, acis, arj, acj, bris, bcis, brj, bcj, sz))

	return phalanxBlock22Patch(C11, C12, C21, C22)
}

// SimplePhalanxBlockMulV2
// linear no recursive
func SimplePhalanxBlockMulV2(A, B *Matrix, ari, aci, arj, acj, bri, bci, brj, bcj, sz int) *Matrix {
	if sz == minAtomMatrixSize {
		return simplePhalanx22Mul22(A, B, ari, aci, arj, acj, bri, bci, brj, bcj)
	}
	sz >>= 1
	aris, acis := ari+sz, aci+sz
	aris1, acis1 := aris-1, acis-1
	bris, bcis := bri+sz, bci+sz
	bris1, bcis1 := bris-1, bcis-1
	//C11 = A11*B11+A12*B21; C12 = A11*B12+A12+B22
	//C21 = A21*B11+A22*B21; C22 = A21*B12+A22*B22
	A11, A12 := A.getBlock(ari, aci, aris1, acis1), A.getBlock(ari, acis, aris1, acj)
	A21, A22 := A.getBlock(aris, aci, arj, acis1), A.getBlock(aris, acis, arj, acj)
	B11, B12 := B.getBlock(bri, bci, bris1, bcis1), B.getBlock(bri, bcis, bris1, bcj)
	B21, B22 := B.getBlock(bris, bci, brj, bcis1), B.getBlock(bris, bcis, brj, bcj)

	P1 := A11.makeCopy().mul(B11)
	P2 := A12.makeCopy().mul(B21)
	P3 := A11.makeCopy().mul(B12)
	P4 := A12.makeCopy().mul(B22)
	P5 := A21.makeCopy().mul(B11)
	P6 := A22.makeCopy().mul(B21)
	P7 := A21.makeCopy().mul(B12)
	P8 := A22.makeCopy().mul(B22)

	//C.setBlock(ari, aci, aris1, acis1)
	return phalanxBlock22Patch(P1.add(P2), P3.add(P4), P5.add(P6), P7.add(P8))
}

// StraseenPhalanxBlockMul
// Strassen Matrix Multiply
func StraseenPhalanxBlockMul(A, B *Matrix, ari, aci, arj, acj, bri, bci, brj, bcj, sz int) *Matrix {
	if sz == minAtomMatrixSize {
		return simplePhalanx22Mul22(A, B, ari, aci, arj, acj, bri, bci, brj, bcj)
	}
	sz >>= 1
	aris, acis := ari+sz, aci+sz
	aris1, acis1 := aris-1, acis-1
	bris, bcis := bri+sz, bci+sz
	bris1, bcis1 := bris-1, bcis-1

	A11, A12 := A.getBlock(ari, aci, aris1, acis1), A.getBlock(ari, acis, aris1, acj)
	A21, A22 := A.getBlock(aris, aci, arj, acis1), A.getBlock(aris, acis, arj, acj)
	B11, B12 := B.getBlock(bri, bci, bris1, bcis1), B.getBlock(bri, bcis, bris1, bcj)
	B21, B22 := B.getBlock(bris, bci, brj, bcis1), B.getBlock(bris, bcis, brj, bcj)

	S1 := B12.makeCopy().sub(B22)
	S2 := A11.makeCopy().add(A12)
	S3 := A21.makeCopy().add(A22)
	S4 := B21.makeCopy().sub(B11)
	S5 := A11.makeCopy().add(A22)
	S6 := B11.makeCopy().add(B22)
	S7 := A12.makeCopy().sub(A22)
	S8 := B21.makeCopy().add(B22)
	S9 := A11.makeCopy().sub(A21)
	S10 := B11.makeCopy().add(B12)

	P1 := StraseenPhalanxBlockMul(A, S1, ari, aci, aris1, acis1, 0, 0, sz-1, sz-1, sz)
	P2 := StraseenPhalanxBlockMul(S2, B, 0, 0, sz-1, sz-1, bris, bcis, brj, bcj, sz)
	P3 := StraseenPhalanxBlockMul(S3, B, 0, 0, sz-1, sz-1, bri, bci, bris1, bcis1, sz)
	P4 := StraseenPhalanxBlockMul(A, S4, aris, acis, arj, acj, 0, 0, sz-1, sz-1, sz)
	P5 := StraseenPhalanxBlockMul(S5, S6, 0, 0, sz-1, sz-1, 0, 0, sz-1, sz-1, sz)
	P6 := StraseenPhalanxBlockMul(S7, S8, 0, 0, sz-1, sz-1, 0, 0, sz-1, sz-1, sz)
	P7 := StraseenPhalanxBlockMul(S9, S10, 0, 0, sz-1, sz-1, 0, 0, sz-1, sz-1, sz)

	C11 := P5.add(P4).sub(P2).add(P6)
	C12 := P1.add(P2)
	C21 := P3.add(P4)
	C22 := P5.add(P1).sub(P3).sub(P7)

	return phalanxBlock22Patch(C11, C12, C21, C22)
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
