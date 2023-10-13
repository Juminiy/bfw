package la

import (
	"bfw/common/lang"
	"errors"
	"fmt"
)

const (
	polyFactorsNoSize int = 0
)

var (
	polyCanNotFactoring = errors.New("poly can not be factored")
)

type PolyFactors struct {
	slice []*Poly
	size  int
}

func ConstructPolyFactors(slice []complex128) *PolyFactors {
	pf := &PolyFactors{}
	return pf.Construct(slice)
}

func (pf *PolyFactors) Construct(slice []complex128) *PolyFactors {
	pf.assign(len(slice))
	for idx := 0; idx < len(slice); idx++ {
		idxSolution := slice[idx]
		if !lang.IsComplex128PureReal(idxSolution) {
			panic(polyCanNotFactoring)
		}
		pf.setElem(idx, ConstructPoly([][]float64{{-real(idxSolution), 0}, {1.0, 1}}))
	}
	return pf
}

func (pf *PolyFactors) validate() bool {
	if pf.slice == nil ||
		pf.size == polyFactorsNoSize ||
		len(pf.slice) != pf.size {
		return false
	}
	return true
}

func (pf *PolyFactors) makeCopy() *PolyFactors {
	pfCopy := &PolyFactors{}
	pfSize := pf.size
	pfCopy.setValues(make([]*Poly, pfSize), pfSize)
	for idx := 0; idx < pfSize; idx++ {
		pfCopy.setElemValue(idx, pf.getElem(idx))
	}
	return pfCopy
}

func (pf *PolyFactors) assign(size int) {
	pf.setValues(make([]*Poly, size), size)
}

func (pf *PolyFactors) getSelf() *PolyFactors {
	return pf
}

func (pf *PolyFactors) setSelf(pft *PolyFactors) {
	pf.setValues(pft.slice, pft.size)
}

func (pf *PolyFactors) setValues(slice []*Poly, size int) {
	pf.setSlice(slice)
	pf.setSize(size)
}

func (pf *PolyFactors) setSlice(slice []*Poly) {
	pf.slice = slice
}

func (pf *PolyFactors) setSize(size int) {
	pf.size = size
}

func (pf *PolyFactors) getElem(index int) *Poly {
	return pf.slice[index]
}

func (pf *PolyFactors) setElem(index int, poly *Poly) {
	pf.slice[index] = poly
}

func (pf *PolyFactors) setElemValue(index int, poly *Poly) {
	pf.slice[index] = poly.makeCopy()
}

func (pf *PolyFactors) Display(precisionBits ...int) *PolyFactors {
	for idx := 0; idx < pf.size; idx++ {
		fmt.Printf("(")
		pf.getElem(idx).DisplayV2(false, true, precisionBits...)
		fmt.Printf(")")
	}
	return pf
}
