package poly

import (
	"bfw/wheel/lang"
	"errors"
	"fmt"
)

const (
	polyFactorsNoSize int = 0
)

var (
	polyCanNotFactoring = errors.New("poly can not be factored")
)

type Factors struct {
	slice []*Poly
	size  int
}

func ConstructPolyFactors(slice []complex128) *Factors {
	pf := &Factors{}
	return pf.Construct(slice)
}

func (f *Factors) Construct(slice []complex128) *Factors {
	f.assign(len(slice))
	for idx := 0; idx < len(slice); idx++ {
		idxSolution := slice[idx]
		if !lang.IsComplex128PureReal(idxSolution) {
			panic(polyCanNotFactoring)
		}
		f.setElem(idx, ConstructPoly([][]float64{{-real(idxSolution), 0}, {1.0, 1}}))
	}
	return f
}

func (f *Factors) validate() bool {
	if f.slice == nil ||
		f.size == polyFactorsNoSize ||
		len(f.slice) != f.size {
		return false
	}
	return true
}

func (f *Factors) makeCopy() *Factors {
	pfCopy := &Factors{}
	pfSize := f.size
	pfCopy.setValues(make([]*Poly, pfSize), pfSize)
	for idx := 0; idx < pfSize; idx++ {
		pfCopy.setElemValue(idx, f.getElem(idx))
	}
	return pfCopy
}

func (f *Factors) assign(size int) {
	f.setValues(make([]*Poly, size), size)
}

func (f *Factors) getSelf() *Factors {
	return f
}

func (f *Factors) setSelf(pft *Factors) {
	f.setValues(pft.slice, pft.size)
}

func (f *Factors) setValues(slice []*Poly, size int) {
	f.setSlice(slice)
	f.setSize(size)
}

func (f *Factors) setSlice(slice []*Poly) {
	f.slice = slice
}

func (f *Factors) setSize(size int) {
	f.size = size
}

func (f *Factors) getElem(index int) *Poly {
	return f.slice[index]
}

func (f *Factors) setElem(index int, poly *Poly) {
	f.slice[index] = poly
}

func (f *Factors) setElemValue(index int, poly *Poly) {
	f.slice[index] = poly.makeCopy()
}

func (f *Factors) Display(precisionBits ...int) *Factors {
	for idx := 0; idx < f.size; idx++ {
		fmt.Printf("(")
		f.getElem(idx).DisplayV2(false, true, precisionBits...)
		fmt.Printf(")")
	}
	return f
}
