package la

import (
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
	lineSize    int
	coefficient float64
}

func (pm *PolyMatrix) validate() bool {
	if pm.rowSize == matrixNoSize ||
		pm.lineSize == matrixNoSize ||
		pm.slice == nil ||
		len(pm.slice) != pm.rowSize {
		return false
	} else {
		for rowIdx := 0; rowIdx < pm.rowSize; rowIdx++ {
			if idxRow := pm.getRow(rowIdx); idxRow == nil ||
				len(idxRow) != pm.lineSize {
				return false
			}
		}
	}
	return true
}

func (pm *PolyMatrix) validateOneIndex(index int) bool {
	if index < 0 ||
		index >= pm.rowSize ||
		index >= pm.lineSize {
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
			if index[1] < 0 || index[1] >= pm.lineSize {
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

func (pm *PolyMatrix) assign(rowSize, lineSize int) {
	pm.setValues(make([][]*Poly, rowSize), rowSize, lineSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		pm.setRow(rowIdx, make([]*Poly, lineSize))
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
	pm.setValues(pmt.slice, pmt.rowSize, pmt.lineSize)
}

func (pm *PolyMatrix) setValues(slice [][]*Poly, rowSize, lineSize int) {
	pm.setSlice(slice)
	pm.setRowSize(rowSize)
	pm.setLineSize(lineSize)
}

func (pm *PolyMatrix) setSlice(slice [][]*Poly) {
	pm.slice = slice
}

func (pm *PolyMatrix) setRowSize(rowSize int) {
	pm.rowSize = rowSize
}

func (pm *PolyMatrix) setLineSize(lineSize int) {
	pm.lineSize = lineSize
}

func (pm *PolyMatrix) setElemByOne2OneOpt(rowIndex, lineIndex int, opt rune, pmt *PolyMatrix) {
	if !pm.validateIndex(rowIndex, lineIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	pm.get(rowIndex, lineIndex).one2OneOpt(opt, pmt.get(rowIndex, lineIndex))
}

func (pm *PolyMatrix) setElemByOptElem(rowIndex, lineIndex int, opt rune, p *Poly) {
	if !pm.validateIndex(rowIndex, lineIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	pm.get(rowIndex, lineIndex).one2OneOpt(opt, p)
}

func (pm *PolyMatrix) get(rowIndex, lineIndex int) *Poly {
	if !pm.validateIndex(rowIndex, lineIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return pm.slice[rowIndex][lineIndex]
}

func (pm *PolyMatrix) set(rowIndex, lineIndex int, value *Poly) {
	pm.slice[rowIndex][lineIndex] = value
}

func (pm *PolyMatrix) getRow(rowIndex int) []*Poly {
	if !pm.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return pm.slice[rowIndex]
}

func (pm *PolyMatrix) setRow(rowIndex int, rowSlice []*Poly) {
	if !pm.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	pm.slice[rowIndex] = rowSlice
}

func (pm *PolyMatrix) Equal(pmt *PolyMatrix) bool {
	if pm.sameShape(pmt) {
		for rowIdx := 0; rowIdx < pm.rowSize; rowIdx++ {
			for lineIdx := 0; lineIdx < pm.lineSize; lineIdx++ {
				if !pm.get(rowIdx, lineIdx).Equal(pmt.get(rowIdx, lineIdx)) {
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
		pm.lineSize != pmt.lineSize {
		return false
	}
	return true
}

func (pm *PolyMatrix) isPhalanx() bool {
	return pm.rowSize == pm.lineSize
}

func (pm *PolyMatrix) getPhalanxSize() int {
	if pm.isPhalanx() {
		return pm.rowSize
	}
	return matrixNotPhalanx
}

func (pm *PolyMatrix) Det() *Poly {
	if !pm.isPhalanx() {
		panic(matrixRowLineDiffer)
	}
	return pm.det()
}

// det
// TODO: to complete the PolyMatrix det algorithm
// simpleDet
// laplaceDet
func (pm *PolyMatrix) det() *Poly {
	if n := pm.getPhalanxSize(); n == simplePhalanxSizeOne {
		return pm.get(0, 0)
	} else if n == simplePhalanxSizeTwo {
		return pm.get(0, 0).GetTimes(pm.get(1, 1)).
			GetMinus(pm.get(0, 1).GetTimes(pm.get(1, 0)))
	} else if n == simplePhalanxSizeThree {
		res1 := pm.get(0, 0).GetTimes(pm.get(1, 1)).GetTimes(pm.get(2, 2))
		res2 := pm.get(1, 0).GetTimes(pm.get(2, 1)).GetTimes(pm.get(0, 2))
		res3 := pm.get(0, 1).GetTimes(pm.get(1, 2)).GetTimes(pm.get(2, 0))
		res4 := pm.get(0, 2).GetTimes(pm.get(1, 1)).GetTimes(pm.get(2, 0))
		res5 := pm.get(0, 1).GetTimes(pm.get(1, 0)).GetTimes(pm.get(2, 2))
		res6 := pm.get(1, 2).GetTimes(pm.get(2, 1)).GetTimes(pm.get(0, 0))
		return res1.GetPlus(res2).GetPlus(res3).GetMinus(res4).GetMinus(res5).GetMinus(res6)
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

func (pm *PolyMatrix) eigenMatrixLineExchangeET() {

}

func (pm *PolyMatrix) eigenMatrixRowMulLambdaET() {

}

func (pm *PolyMatrix) eigenMatrixLineMulLambdaET() {

}

func (pm *PolyMatrix) eigenMatrixRow1MulPolyAddRow2ET() {

}

func (pm *PolyMatrix) eigenMatrixLine1MulPolyAddLine2ET() {

}

func (pm *PolyMatrix) eigenMatrixElementaryTransformation() {

}

func (pm *PolyMatrix) Display(precisionBits ...int) *PolyMatrix {
	if pm.isNull() {
		fmt.Println("[null]")
		return pm
	}
	for rowIdx := 0; rowIdx < pm.rowSize; rowIdx++ {
		for lineIdx := 0; lineIdx < pm.lineSize; lineIdx++ {
			poly := pm.get(rowIdx, lineIdx)
			poly.Display(false, precisionBits...)
			fmt.Printf(" ")
		}
		fmt.Println()
	}
	return pm
}
