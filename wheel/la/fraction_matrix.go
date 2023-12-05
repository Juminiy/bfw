package la

import (
	"bfw/wheel/lang"
	"bfw/wheel/num"
	"fmt"
)

type FractionMatrix struct {
	slice [][]*num.Fraction
}

func MakeFractionMatrix(slice [][]*num.Fraction) *FractionMatrix {
	fm := &FractionMatrix{}
	fm.setSlice(slice)
	return fm
}

func (fm *FractionMatrix) make(rowSize, columnSize int) {
	slice := make([][]*num.Fraction, rowSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		slice[rowIdx] = make([]*num.Fraction, columnSize)
	}
	fm.setSlice(slice)
}

func (fm *FractionMatrix) copy(fma *FractionMatrix) {
	if !fm.sameSize(fma) {
		fm.make(fma.size())
	}
	for rowIdx := 0; rowIdx < fm.rowSize(); rowIdx++ {
		copy(fm.slice[rowIdx], fma.slice[rowIdx])
	}
}

// deep copy
func (fm *FractionMatrix) makeCopy() *FractionMatrix {
	fma := &FractionMatrix{}
	fma.make(fm.size())
	fma.copy(fm)
	return fma
}

func (fm *FractionMatrix) clear() {
	fm.slice = nil
}

func (fm *FractionMatrix) setSlice(slice [][]*num.Fraction) {
	fm.clear()
	fm.slice = slice
}

func (fm *FractionMatrix) validateIndex(rowIndex, columnIndex int) bool {
	return rowIndex >= 0 && rowIndex < fm.rowSize() &&
		columnIndex >= 0 && columnIndex < fm.columnSize()
}

func (fm *FractionMatrix) set(rowIndex, columnIndex int, fraction *num.Fraction) {
	if fm.validateIndex(rowIndex, columnIndex) {
		fm.slice[rowIndex][columnIndex] = fraction
	}
}

func (fm *FractionMatrix) get(rowIndex, columnIndex int) *num.Fraction {
	if fm.validateIndex(rowIndex, columnIndex) {
		return fm.slice[rowIndex][columnIndex]
	}
	return nil
}

func (fm *FractionMatrix) isZero() bool {
	return fm.slice == nil ||
		len(fm.slice) == matrixNoSize ||
		len(fm.slice[0]) == matrixNoSize
}

func (fm *FractionMatrix) size() (int, int) {
	return fm.rowSize(), fm.columnSize()
}

func (fm *FractionMatrix) rowSize() int {
	if fm.isZero() {
		return matrixNoSize
	}
	return len(fm.slice)
}

func (fm *FractionMatrix) columnSize() int {
	if fm.isZero() {
		return matrixNoSize
	}
	return len(fm.slice[0])
}

func (fm *FractionMatrix) sameSize(fma *FractionMatrix) bool {
	fmRowSize, fmColumnSize := fm.size()
	fmaRowSize, fmaColumnSize := fma.size()
	return lang.PairIntEqual(fmRowSize, fmColumnSize, fmaRowSize, fmaColumnSize)
}

func (fm *FractionMatrix) canLeftMultiply(fma *FractionMatrix) bool {
	return fm.columnSize() == fma.rowSize()
}

func (fm *FractionMatrix) canRightMultiply(fma *FractionMatrix) bool {
	return fm.rowSize() == fma.columnSize()
}

// one2OneOpt
func (fm *FractionMatrix) ternaryArithmetic(fma, fmb *FractionMatrix, operator rune) {
	if !fma.sameSize(fmb) {
		panic(matrixNotSameShapeError)
	}
	if !fm.sameSize(fma) {
		fm.make(fma.size())
	}
	for rowIdx := 0; rowIdx < fm.rowSize(); rowIdx++ {
		for columnIdx := 0; columnIdx < fm.columnSize(); columnIdx++ {
			elemA := fma.get(rowIdx, columnIdx)
			elemB := fmb.get(rowIdx, columnIdx)
			switch operator {
			case '+':
				{
					fm.slice[rowIdx][columnIdx] = elemA.Add(elemB)
				}
			case '-':
				{
					fm.slice[rowIdx][columnIdx] = elemA.Sub(elemB)
				}
			case '*':
				{
					fm.slice[rowIdx][columnIdx] = elemA.Mul(elemB)
				}
			case '/':
				{
					fm.slice[rowIdx][columnIdx] = elemA.Div(elemB)
				}
			case '^':
				{

				}
			}
		}
	}
}

func (fm *FractionMatrix) add(fma, fmb *FractionMatrix) {
	fm.ternaryArithmetic(fma, fmb, '+')
}

func (fm *FractionMatrix) sub(fma, fmb *FractionMatrix) {
	fm.ternaryArithmetic(fma, fmb, '-')
}

func (fm *FractionMatrix) dotMul(fma, fmb *FractionMatrix) {
	fm.ternaryArithmetic(fma, fmb, '*')
}

func (fm *FractionMatrix) dotDiv(fma, fmb *FractionMatrix) {
	fm.ternaryArithmetic(fma, fmb, '/')
}

func (fm *FractionMatrix) mul(fma, fmb *FractionMatrix) {
	if !fma.canLeftMultiply(fmb) {
		panic(matrixCanNotMultiplyError)
	}
	fm.one(fma.rowSize(), fmb.columnSize())
	for kIdx := 0; kIdx < fma.columnSize(); kIdx++ {
		for iIdx := 0; iIdx < fma.rowSize(); iIdx++ {
			valIK := fma.slice[iIdx][kIdx]
			for jIdx := 0; jIdx < fmb.columnSize(); jIdx++ {
				fm.slice[iIdx][jIdx].Add(valIK.Mul(fmb.slice[kIdx][jIdx]))
			}
		}
	}
}

func (fm *FractionMatrix) power(fma *FractionMatrix, n int) {
	fm.identity(fm.size())
	for n > 0 {
		if n&1 != 0 {
			fm.mul(fm.makeCopy(), fma)
		}
		fma.mul(fma.makeCopy(), fma.makeCopy())
		n >>= 1
	}
}

func (fm *FractionMatrix) setAllVal(rowSize, columnSize int, val rune) {
	fm.make(rowSize, columnSize)
	fm.traverse(func(rowIdx int, columnIdx int) {
		switch val {
		case '0':
			{
				fm.slice[rowIdx][columnIdx] = num.MakeZero()
			}
		case '1':
			{
				fm.slice[rowIdx][columnIdx] = num.MakeOne()
			}
		}
	})
}

func (fm *FractionMatrix) isPhalanx() bool {
	return fm.rowSize() == fm.columnSize()
}

func MakeFractionIdentity(rowSize, columnSize int) *FractionMatrix {
	fm := &FractionMatrix{}
	fm.identity(rowSize, columnSize)
	return fm
}

func (fm *FractionMatrix) identity(rowSize, columnSize int) {
	if !fm.isPhalanx() {
		panic(matrixNotPhalanxError)
	}
	fm.make(rowSize, columnSize)
	fm.traverse(func(rowIdx int, columnIdx int) {
		if rowIdx == columnIdx {
			fm.slice[rowIdx][columnIdx] = num.MakeOne()
		} else {
			fm.slice[rowIdx][columnIdx] = num.MakeZero()
		}
	})
}

func (fm *FractionMatrix) zero(rowSize, columnSize int) {
	fm.setAllVal(rowSize, columnSize, '0')
}

func (fm *FractionMatrix) one(rowSize, columnSize int) {
	fm.setAllVal(rowSize, columnSize, '1')
}

func (fm *FractionMatrix) Display() {
	curRow := 0
	fm.traverse(func(rowIdx int, columnIdx int) {
		if rowIdx > curRow {
			fmt.Println()
			curRow = rowIdx
		}
		fm.slice[rowIdx][columnIdx].Display(false)
		fmt.Printf(" ")
	})
}

func (fm *FractionMatrix) traverse(funcPtr func(int, int)) {
	rowSize, columnSize := fm.size()
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			funcPtr(rowIdx, columnIdx)
		}
	}
}

func (fm *FractionMatrix) traverseV2(funcPtr func(int, int, *num.Fraction)) {
	rowSize, columnSize := fm.size()
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			funcPtr(rowIdx, columnIdx, fm.slice[rowIdx][columnIdx])
		}
	}
}

func (fm *FractionMatrix) traverseV3(funcPtr func(int, int, *num.Fraction, bool), stop bool) {
	rowSize, columnSize := fm.size()
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			funcPtr(rowIdx, columnIdx, fm.slice[rowIdx][columnIdx], stop)
		}
	}
}

func (fm *FractionMatrix) rowTraverse() {

}

func (fm *FractionMatrix) columnTraverse() {

}

func (fm *FractionMatrix) traverseRow(rowIndex int, funcPtr func(int, *num.Fraction)) {
	columnSize := fm.columnSize()
	for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
		funcPtr(columnIdx, fm.slice[rowIndex][columnIdx])
	}
}

func (fm *FractionMatrix) traverseColumn(columnIndex int, funcPtr func(int, *num.Fraction)) {

}
