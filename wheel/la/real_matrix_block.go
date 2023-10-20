package la

// BlockMatrix
// r*c [r*c][r*c]
// we need to divide the whole matrix into regular phalanx if possible and spare no effort
type BlockMatrix struct {
	block      [][]*Matrix
	rowSize    int
	columnSize int
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

func (bm *BlockMatrix) setSelf(bmt *BlockMatrix) {
	if !validateBM(bmt) {
		return
	}
	bm.setValues(bmt.block, bmt.rowSize, bmt.columnSize)
}

func (bm *BlockMatrix) setValues(block [][]*Matrix, size ...int) {
	bm.setBlock(block)
	bm.setSize(size...)
}

func (bm *BlockMatrix) setBlock(block [][]*Matrix) {
	bm.block = nil
	bm.block = block
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
