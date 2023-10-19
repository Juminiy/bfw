package cal

import (
	"bfw/common/la"
	"bfw/common/lang"
	"fmt"
)

const (
	eigenPolyMatrixDefaultAES rune = 'λ'
)

var (
	NullPolyMatrix = &PolyMatrix{}
)

type PolyMatrix struct {
	slice       [][]*Poly
	rowSize     int
	columnSize  int
	coefficient float64
}

func (pm *PolyMatrix) validate() bool {
	if pm.rowSize == la.matrixNoSize ||
		pm.columnSize == la.matrixNoSize ||
		pm.slice == nil ||
		len(pm.slice) != pm.rowSize {
		return false
	} else {
		for rowIdx := 0; rowIdx < pm.rowSize; rowIdx++ {
			if idxRow := pm.getRow(rowIdx); idxRow == nil ||
				len(idxRow) != pm.columnSize {
				return false
			}
		}
	}
	return true
}

func (pm *PolyMatrix) validateOneIndex(index int) bool {
	if index < 0 ||
		index >= pm.rowSize ||
		index >= pm.columnSize {
		return false
	}
	return true
}

func (pm *PolyMatrix) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		if indexLen >= 1 {
			if index[0] < 0 || index[0] >= pm.rowSize {
				return false
			}
		}
		if indexLen >= 2 {
			if index[1] < 0 || index[1] >= pm.columnSize {
				return false
			}
		}
		if indexLen >= 3 {
			for indexIdx := 2; indexIdx < indexLen; indexIdx++ {
				if !pm.validateOneIndex(index[indexIdx]) {
					return false
				}
			}
		}
	}
	return true
}

func (pm *PolyMatrix) isNull() bool {
	return !pm.validate()
}

func (pm *PolyMatrix) null() *PolyMatrix {
	return &PolyMatrix{}
}

func (pm *PolyMatrix) assign(rowSize, columnSize int) {
	pm.setValues(make([][]*Poly, rowSize), rowSize, columnSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		pm.setRow(rowIdx, make([]*Poly, columnSize))
	}
}

func (pm *PolyMatrix) swap(m *PolyMatrix) {
	pmTemp := &PolyMatrix{}
	pmTemp.setSelf(pm)
	pm.setSelf(m)
	m.setSelf(pmTemp)
}

func (pm *PolyMatrix) getSelf() *PolyMatrix {
	return pm
}

func (pm *PolyMatrix) setSelf(pmt *PolyMatrix) {
	pm.setValues(pmt.slice, pmt.rowSize, pmt.columnSize)
}

func (pm *PolyMatrix) setValues(slice [][]*Poly, rowSize, columnSize int) {
	pm.setSlice(slice)
	pm.setRowSize(rowSize)
	pm.setColumnSize(columnSize)
}

func (pm *PolyMatrix) setSlice(slice [][]*Poly) {
	pm.slice = slice
}

func (pm *PolyMatrix) setRowSize(rowSize int) {
	pm.rowSize = rowSize
}

func (pm *PolyMatrix) setColumnSize(columnSize int) {
	pm.columnSize = columnSize
}

func (pm *PolyMatrix) setElemByOne2OneOpt(rowIndex, columnIndex int, opt rune, pmt *PolyMatrix) {
	if !pm.validateIndex(rowIndex, columnIndex) {
		panic(la.matrixIndexOutOfBoundError)
	}
	pm.get(rowIndex, columnIndex).one2OneOpt(opt, pmt.get(rowIndex, columnIndex))
}

func (pm *PolyMatrix) setElemByOptElem(rowIndex, columnIndex int, opt rune, p *Poly) {
	if !pm.validateIndex(rowIndex, columnIndex) {
		panic(la.matrixIndexOutOfBoundError)
	}
	pm.get(rowIndex, columnIndex).one2OneOpt(opt, p)
}

func (pm *PolyMatrix) get(rowIndex, columnIndex int) *Poly {
	if !pm.validateIndex(rowIndex, columnIndex) {
		panic(la.matrixIndexOutOfBoundError)
	}
	return pm.slice[rowIndex][columnIndex]
}

func (pm *PolyMatrix) set(rowIndex, columnIndex int, value *Poly) {
	pm.slice[rowIndex][columnIndex] = value
}

func (pm *PolyMatrix) getRow(rowIndex int) []*Poly {
	if !pm.validateIndex(rowIndex) {
		panic(la.matrixIndexOutOfBoundError)
	}
	return pm.slice[rowIndex]
}

func (pm *PolyMatrix) setRow(rowIndex int, rowSlice []*Poly) {
	if !pm.validateIndex(rowIndex) {
		panic(la.matrixIndexOutOfBoundError)
	}
	pm.slice[rowIndex] = rowSlice
}

func (pm *PolyMatrix) Equal(pmt *PolyMatrix) bool {
	if pm.sameShape(pmt) {
		for rowIdx := 0; rowIdx < pm.rowSize; rowIdx++ {
			for columnIdx := 0; columnIdx < pm.columnSize; columnIdx++ {
				if !pm.get(rowIdx, columnIdx).Equal(pmt.get(rowIdx, columnIdx)) {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (pm *PolyMatrix) sameShape(pmt *PolyMatrix) bool {
	if pmt == nil ||
		pm.rowSize != pmt.rowSize ||
		pm.columnSize != pmt.columnSize {
		return false
	}
	return true
}

func (pm *PolyMatrix) isPhalanx() bool {
	return pm.rowSize == pm.columnSize
}

func (pm *PolyMatrix) getPhalanxSize() int {
	if pm.isPhalanx() {
		return pm.rowSize
	}
	return la.matrixNotPhalanx
}

func (pm *PolyMatrix) Det() *Poly {
	if !pm.isPhalanx() {
		panic(la.matrixRowColumnDiffer)
	}
	return pm.det()
}

// det
// TODO: to complete the PolyMatrix det algorithm
// simpleDet
// laplaceDet
func (pm *PolyMatrix) det() *Poly {
	if n := pm.getPhalanxSize(); n == la.simplePhalanxSizeOne {
		return pm.get(0, 0)
	} else if n == la.simplePhalanxSizeTwo {
		res1 := PolyChainedMulV2(pm.get(0, 0), pm.get(1, 1))
		res2 := PolyChainedMulV2(pm.get(0, 1), pm.get(1, 1))
		return PolyChainedSub(res1, res2)
	} else if n == la.simplePhalanxSizeThree {
		res1 := PolyChainedMulV2(pm.get(0, 0), pm.get(1, 1), pm.get(2, 2))
		res2 := PolyChainedMulV2(pm.get(1, 0), pm.get(2, 1), pm.get(0, 2))
		res3 := PolyChainedMulV2(pm.get(0, 1), pm.get(1, 2), pm.get(2, 0))
		res4 := PolyChainedMulV2(pm.get(0, 2), pm.get(1, 1), pm.get(2, 0))
		res5 := PolyChainedMulV2(pm.get(0, 1), pm.get(1, 0), pm.get(2, 2))
		res6 := PolyChainedMulV2(pm.get(1, 2), pm.get(2, 1), pm.get(0, 0))
		res1 = PolyChainedAdd(res1, res2, res3)
		return PolyChainedSub(res1, res4, res5, res6)
	} else {
		// 1. simple calculation
		//return pm.simpleDet(n)
		// 2. laplace calculation
		//return pm.laplaceDet(n)
		// 3. mixture calculation
		if lang.Odd(n) {
			return pm.simpleDet(n)
		} else {
			return pm.laplaceDet(n)
		}
	}
}

func (pm *PolyMatrix) simpleDet(totalN int) *Poly {
	return &Poly{}
}

func (pm *PolyMatrix) laplaceDet(totalN int) *Poly {
	return &Poly{}
}

func (pm *PolyMatrix) Smith() {

}

func (pm *PolyMatrix) DD(k int) {

}

func (pm *PolyMatrix) Dd(k int) {

}

func (pm *PolyMatrix) greatestCommonFactor() {

}

func (pm *PolyMatrix) eigenMatrixRowExchangeET() {

}

func (pm *PolyMatrix) eigenMatrixColumnExchangeET() {

}

func (pm *PolyMatrix) eigenMatrixRowMulLambdaET() {

}

func (pm *PolyMatrix) eigenMatrixColumnMulLambdaET() {

}

func (pm *PolyMatrix) eigenMatrixRow1MulPolyAddRow2ET() {

}

func (pm *PolyMatrix) eigenMatrixColumn1MulPolyAddColumn2ET() {

}

func (pm *PolyMatrix) eigenMatrixElementaryTransformation() {

}

func (pm *PolyMatrix) Display(precisionBits ...int) *PolyMatrix {
	if pm.isNull() {
		fmt.Println("[null]")
		return pm
	}
	for rowIdx := 0; rowIdx < pm.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < pm.columnSize; columnIdx++ {
			poly := pm.get(rowIdx, columnIdx)
			poly.Display(false, precisionBits...)
			fmt.Printf(" ")
		}
		fmt.Println()
	}
	return pm
}
