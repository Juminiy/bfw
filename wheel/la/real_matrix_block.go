package la

import (
	"bfw/wheel/lang"
	"errors"
)

var (
	notSupportedBlockMatrixError = errors.New("not supported block matrix shape")
	canNotGetMatrixBlockError    = errors.New("cannot get the matrix block")
)

// BlockMatrix
// r*c [r*c][r*c]
// we need to divide the whole matrix into regular phalanx if possible and spare no effort
type BlockMatrix struct {
	block      [][]*Matrix
	rowSize    int
	columnSize int
}

func GenBlockMatrixBlocks(rowSize, columnSize int, dataType string, bRowSize, bColumnSize int, dataRange ...float64) [][]*Matrix {
	bMatrix := make([][]*Matrix, rowSize)
	for bRowIdx := 0; bRowIdx < rowSize; bRowIdx++ {
		bMatrix[bRowIdx] = make([]*Matrix, columnSize)
		for bColumnIdx := 0; bColumnIdx < columnSize; bColumnIdx++ {
			bMatrix[bRowIdx][bColumnIdx] = GenMatrix(bRowSize, bColumnSize, dataType, dataRange...)
		}
	}
	return bMatrix
}

func GenBlockMatrix(rowSize, columnSize int, dataType string, bRowSize, bColumnSize int, dataRange ...float64) *BlockMatrix {
	bm := &BlockMatrix{}
	bm.setValues(GenBlockMatrixBlocks(rowSize, columnSize, dataType, bRowSize, bColumnSize, dataRange...), rowSize, columnSize)
	return bm
}

func ValidateBlockMatrix(bm ...*BlockMatrix) bool {
	return validateBM(bm...)
}

func validateBM(bm ...*BlockMatrix) bool {
	if bmLen := len(bm); bmLen > 0 {
		for bmIdx := 0; bmIdx < bmLen; bmIdx++ {
			if bm[bmIdx] == nil ||
				!bm[bmIdx].validate() {
				return false
			}
		}
	}
	return true
}

func (bm *BlockMatrix) validate() bool {
	return bm.rowSize != matrixNoSize &&
		bm.columnSize != matrixNoSize &&
		len(bm.block) == bm.rowSize
	// columns and elems will not be checked
}

func (bm *BlockMatrix) null() *BlockMatrix {
	return &BlockMatrix{}
}

func (bm *BlockMatrix) isNull() bool {
	return !bm.validate()
}

func (bm *BlockMatrix) setNull() {
	bm.block = nil
	bm.rowSize = matrixNoSize
	bm.columnSize = matrixNoSize
}

func (bm *BlockMatrix) assign(rowSize, columnSize int, blockRowSize, blockColumnSize int) {
	bm.setValues(make([][]*Matrix, rowSize), rowSize, columnSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		bm.block[rowIdx] = make([]*Matrix, columnSize)
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			matrix := &Matrix{}
			matrix.assign(blockRowSize, blockColumnSize)
			bm.block[rowIdx][columnIdx] = matrix
		}
	}
}

func (bm *BlockMatrix) setSelf(bmt *BlockMatrix) {
	if !validateBM(bmt) {
		return
	}
	bm.setValues(bmt.block, bmt.rowSize, bmt.columnSize)
}

func (bm *BlockMatrix) setValues(block [][]*Matrix, size ...int) {
	bm.setSize(size...)
	bm.setBlock(block)
}

func (bm *BlockMatrix) setBlock(block [][]*Matrix) {
	bm.block = nil
	bm.block = block
	// assume the size
	//var (
	//	destRowSize    = bm.rowSize
	//	destColumnSize = bm.columnSize
	//)
	//destRowSize = lang.MinInt(destRowSize, len(block))
	//if len(block) > 0 {
	//	destColumnSize = lang.MinInt(destColumnSize, len(block[0]))
	//}
	//bm.setSize(destRowSize, destColumnSize)
}

func (bm *BlockMatrix) setSize(size ...int) {
	if sizeLen := len(size); sizeLen > 0 {
		if sizeLen >= 1 {
			bm.setRowSize(size[0])
			bm.setColumnSize(size[0])
		}
		if sizeLen >= 2 {
			bm.setColumnSize(size[1])
		}
	}
}

func (bm *BlockMatrix) setRowSize(rowSize int) {
	bm.rowSize = rowSize
}

func (bm *BlockMatrix) setColumnSize(columnSize int) {
	bm.columnSize = columnSize
}

func (bm *BlockMatrix) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		if indexLen >= 1 {
			if !bm.validateRowIndex(index[0]) {
				return false
			}
		}
		if indexLen >= 2 {
			if !bm.validateColumnIndex(index[1]) {
				return false
			}
		}
	}
	return true
}

func (bm *BlockMatrix) validateRowIndex(rowIndex int) bool {
	return rowIndex >= 0 &&
		rowIndex < bm.rowSize
}

func (bm *BlockMatrix) validateColumnIndex(columnIndex int) bool {
	return columnIndex >= 0 &&
		columnIndex < bm.columnSize
}

func (bm *BlockMatrix) get(rowIndex, columnIndex int) *Matrix {
	if !bm.validateIndex(rowIndex, columnIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return bm.block[rowIndex][columnIndex]
}

func (bm *BlockMatrix) setElemSwap(rowIndexI, columnIndexI, rowIndexJ, columnIndexJ int) {
	bm.block[rowIndexI][columnIndexI], bm.block[rowIndexJ][columnIndexJ] =
		bm.block[rowIndexJ][columnIndexJ], bm.block[rowIndexI][columnIndexI]
}

func (bm *BlockMatrix) Transpose() *BlockMatrix {
	return bm.transpose()
}

func (bm *BlockMatrix) transpose() *BlockMatrix {
	if bm.isPhalanx() {
		bm.phalanxTranspose()
		bm.elemsTranspose()
		return bm
	}
	newRowSize, newColumnSize := bm.columnSize, bm.rowSize
	newBlock := make([][]*Matrix, newRowSize)
	for newRowIdx := 0; newRowIdx < newRowSize; newRowIdx++ {
		newBlock[newRowIdx] = make([]*Matrix, newColumnSize)
		for newColumnIdx := 0; newColumnIdx < newColumnSize; newColumnIdx++ {
			newBlock[newRowIdx][newColumnIdx] = bm.get(newColumnIdx, newRowIdx).phalanxTranspose()
		}
	}
	bm.setValues(newBlock, newRowSize, newColumnSize)
	return bm
}

func (bm *BlockMatrix) elemsTranspose() *BlockMatrix {
	for rowIdx := 0; rowIdx < bm.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < bm.columnSize; columnIdx++ {
			bm.block[rowIdx][columnIdx].transpose()
		}
	}
	return bm
}

func (bm *BlockMatrix) blocksTranspose() *BlockMatrix {
	newRowSize, newColumnSize := bm.columnSize, bm.rowSize
	newBlock := make([][]*Matrix, newRowSize)
	for newRowIdx := 0; newRowIdx < newRowSize; newRowIdx++ {
		newBlock[newRowIdx] = make([]*Matrix, newColumnSize)
		for newColumnIdx := 0; newColumnIdx < newColumnSize; newColumnIdx++ {
			newBlock[newRowIdx][newColumnIdx] = bm.get(newColumnIdx, newRowIdx)
		}
	}
	bm.setValues(newBlock, newRowSize, newColumnSize)
	return bm
}

func (bm *BlockMatrix) getPhalanxSize() int {
	if bm.isPhalanx() {
		return bm.rowSize
	}
	return matrixNoSize
}

func (bm *BlockMatrix) getPhalanxSizePanic() int {
	if !bm.isPhalanx() {
		panic(matrixNotPhalanxError)
	}
	return bm.rowSize
}

// need to be reconsidered
func (bm *BlockMatrix) isPhalanx() bool {
	return bm.rowSize == bm.columnSize
}

func (bm *BlockMatrix) phalanxTranspose() *BlockMatrix {
	for columnIdx := 0; columnIdx < bm.columnSize; columnIdx++ {
		for rowIdx := columnIdx + 1; rowIdx < bm.rowSize; rowIdx++ {
			bm.setElemSwap(rowIdx, columnIdx, columnIdx, rowIdx)
		}
	}
	return bm
}

func (bm *BlockMatrix) getElemPhalanxSize() int {
	return bm.get(lang.GetTwoRandomIntValue(bm.rowSize)).getPhalanxSize()
}

func (bm *BlockMatrix) isElemPhalanx() bool {
	return bm.get(lang.GetTwoRandomIntValue(bm.rowSize)).isPhalanx()
}

func (bm *BlockMatrix) Mul(bmt *BlockMatrix) *BlockMatrix {
	return bm.mul(bmt)
}

func (bm *BlockMatrix) mul(bmt *BlockMatrix) *BlockMatrix {
	if !bm.canMul(bmt) {
		panic(matrixCanNotMultiplyError)
	}
	btSize, bSize := bm.getPhalanxSize(), bmt.getElemPhalanxSize()
	resMatrix := &BlockMatrix{}
	resMatrix.assign(bSize, bSize, btSize, btSize)
	for kIdx := 0; kIdx < bm.columnSize; kIdx++ {
		for iIdx := 0; iIdx < bm.rowSize; iIdx++ {
			valIK := bm.block[iIdx][kIdx]
			for jIdx := 0; jIdx < bmt.columnSize; jIdx++ {
				resMatrix.block[iIdx][jIdx].add(valIK.MTimes(bmt.block[kIdx][jIdx]))
			}
		}
	}
	return resMatrix
}

func (bm *BlockMatrix) checkPhalanx() bool {
	return bm.isPhalanx() &&
		bm.isElemPhalanx()
}

// canMul
// let random choose the calculation destiny and lifecycle
func (bm *BlockMatrix) canMul(bmt *BlockMatrix) bool {
	bmBTSize, bmtBTSize := bm.getPhalanxSize(), bmt.getPhalanxSize()
	return bmBTSize == bmtBTSize &&
		bm.getElemPhalanxSize() == bmt.getElemPhalanxSize()
}

// Matrix
// the matrix keep the previous block address pointer
// none change them, none copy them
// assume each block is same size phalanx
// assume the blocks is phalanx
// 1 2 | 3 4
// 5 6 | 7 8
// - - - - -
// 9 8 | 7 6
// 5 4 | 3 2
func (bm *BlockMatrix) Matrix() *Matrix {
	if !bm.checkPhalanx() {
		panic(notSupportedBlockMatrixError)
	}
	btSize := bm.getPhalanxSize()
	bSize := bm.getElemPhalanxSize()
	matrix := &Matrix{}
	matrix.assignRow(btSize*bSize, btSize*bSize)
	for btRowIdx := 0; btRowIdx < btSize; btRowIdx++ {
		for bRowIdx := 0; bRowIdx < bSize; bRowIdx++ {
			for btColumnIdx := 0; btColumnIdx < btSize; btColumnIdx++ {
				matrix.setRowElemAppend(btRowIdx*bSize+bRowIdx, bm.get(btRowIdx, btColumnIdx).getRow(bRowIdx))
			}
		}
	}
	return matrix
}

// Display
// to Matrix Display
func (bm *BlockMatrix) Display() *BlockMatrix {
	bm.Matrix().Display()
	return bm
}
