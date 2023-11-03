package la

import (
	"bfw/wheel/lang"
	"errors"
)

var (
	ComplexMatrixCanNotConvertToRealMatrix = errors.New("complex matrix can not convert to real matrix")
)

type ComplexMatrix struct {
	slice [][]complex128
}

func ConstructComplexMatrix(slice [][]complex128) *ComplexMatrix {
	cm := &ComplexMatrix{}
	cm.Construct(slice)
	return cm
}

func (cm *ComplexMatrix) Construct(slice [][]complex128) *ComplexMatrix {
	cm.setSlice(slice)
	return cm
}

func (cm *ComplexMatrix) validateRowIndex(index int) bool {
	return index >= 0 && index < cm.rowSize()
}
func (cm *ComplexMatrix) validateColumnIndex(index int) bool {
	return index >= 0 && index < cm.columnSize()
}

func (cm *ComplexMatrix) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		if indexLen == 1 {
			if !cm.validateRowIndex(index[0]) {
				return false
			}
		} else if indexLen == 2 {
			if !cm.validateColumnIndex(index[1]) {
				return false
			}
		}
	}
	return true
}

func (cm *ComplexMatrix) rowSize() int {
	return len(cm.slice)
}

func (cm *ComplexMatrix) columnSize() int {
	if cm.rowSize() == matrixNoSize {
		return matrixNoSize
	}
	return len(cm.slice[0])
}

func (cm *ComplexMatrix) get(rowIndex, columnIndex int) complex128 {
	return cm.slice[rowIndex][columnIndex]
}

func (cm *ComplexMatrix) set(rowIndex, columnIndex int, value complex128) {
	cm.slice[rowIndex][columnIndex] = value
}

func (cm *ComplexMatrix) setSlice(slice [][]complex128) {
	cm.slice = nil
	cm.slice = slice
}

func (cm *ComplexMatrix) getSlice() [][]complex128 {
	return cm.slice
}

func (cm *ComplexMatrix) setRow(rowIndex int, complex1DArray []complex128) {
	if !cm.validateRowIndex(rowIndex) ||
		len(complex1DArray) != cm.columnSize() {
		panic(matrixIndexOutOfBoundError)
	}
	cm.slice[rowIndex] = nil
	cm.slice[rowIndex] = complex1DArray
}

func (cm *ComplexMatrix) setColumn(columnIndex int, complex1DArray []complex128) {
	if !cm.validateColumnIndex(columnIndex) ||
		len(complex1DArray) != cm.rowSize() {
		panic(matrixIndexOutOfBoundError)
	}
	cm.traverseRowElem(columnIndex, func(elemValue complex128, index ...int) {
		cm.set(index[0], columnIndex, complex1DArray[index[0]])
	})
}

func (cm *ComplexMatrix) getRow(rowIndex int) []complex128 {
	if !cm.validateRowIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return cm.slice[rowIndex]
}

func (cm *ComplexMatrix) getColumn(columnIndex int) []complex128 {
	if !cm.validateColumnIndex(columnIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	column := make([]complex128, cm.rowSize())
	cm.traverseColumnElem(columnIndex, func(elemValue complex128, index ...int) {
		column[index[0]] = elemValue
	})
	return column
}

func (cm *ComplexMatrix) makeCopy() *ComplexMatrix {
	cmCopy := &ComplexMatrix{}
	cmCopy.setSlice(make([][]complex128, cm.rowSize()))
	cm.traverseByRow(func(rowArray []complex128, index ...int) {
		cmCopy.slice[index[0]] = make([]complex128, cm.columnSize())
		copy(cmCopy.slice[index[0]], rowArray)
	})
	return cmCopy
}

// Aᵀ
func (cm *ComplexMatrix) transpose() *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) conjugate() *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) CTranspose() *ComplexMatrix {
	return cm.transpose().conjugate()
}

func (cm *ComplexMatrix) phaseAngle() *Matrix {
	return &Matrix{}
}

func (cm *ComplexMatrix) null() *ComplexMatrix {
	return &ComplexMatrix{}
}

// UnitarySimilar
// return UᴴAU = U⁻¹AU (UᴴU = UUᴴ = I)
func (cm *ComplexMatrix) UnitarySimilar(U *UnitaryMatrix) *ComplexMatrix {
	return cm
}

func (cm *ComplexMatrix) IsUnitarySimilar(B *ComplexMatrix) bool {
	return false
}

func (cm *ComplexMatrix) convertToMatrix() *Matrix {
	matrix := &Matrix{}
	matrix.assign(cm.rowSize(), cm.columnSize())
	cm.traverse(func(elemValue complex128, index ...int) {
		if !lang.IsComplex128PureReal(elemValue) {
			panic(ComplexMatrixCanNotConvertToRealMatrix)
		}
		matrix.set(index[0], index[1], real(elemValue))
	})
	return matrix
}

func (cm *ComplexMatrix) Matrix() *Matrix {
	return cm.convertToMatrix()
}

func (cm *ComplexMatrix) traverse(funcPtr func(elemValue complex128, index ...int)) {
	rowSize, columnSize := cm.rowSize(), cm.columnSize()
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			funcPtr(cm.slice[rowIdx][columnIdx], rowIdx, columnIdx)
		}
	}
}

func (cm *ComplexMatrix) traverseByRow(funcPtr func(rowArray []complex128, index ...int)) {
	for rowIdx := 0; rowIdx < cm.rowSize(); rowIdx++ {
		funcPtr(cm.slice[rowIdx], rowIdx)
	}
}

func (cm *ComplexMatrix) traverseColumnElem(rowIndex int, funcPtr func(elemValue complex128, index ...int)) {
	for columnIdx := 0; columnIdx < cm.columnSize(); columnIdx++ {
		funcPtr(cm.slice[rowIndex][columnIdx], columnIdx)
	}
}

func (cm *ComplexMatrix) traverseRowElem(columnIndex int, funcPtr func(elemValue complex128, index ...int)) {
	for rowIdx := 0; rowIdx < cm.rowSize(); rowIdx++ {
		funcPtr(cm.slice[rowIdx][columnIndex], rowIdx)
	}
}
