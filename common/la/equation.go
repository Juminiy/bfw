package la

import (
	"bfw/common/lang"
	"errors"
	"fmt"
)

const (
	solutionNoSize int = 0
)

var (
	solutionIndexOutOfBoundError = errors.New("solution index is out of bound")
)

type Solution struct {
	slice []complex128
	size  int
}

func ConstructSolution(slice []complex128) *Solution {
	s := &Solution{}
	return s.Construct(slice)
}

func (s *Solution) Construct(slice []complex128) *Solution {

	return s
}

func (s *Solution) validate() bool {
	if s.size == solutionNoSize ||
		len(s.slice) != s.size {
		return false
	}
	return true
}

func (s *Solution) validateOneIndex(index int) bool {
	if index < 0 ||
		index >= s.size {
		return false
	}
	return true
}

func (s *Solution) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if !s.validateOneIndex(index[indexIdx]) {
				return false
			}
		}
	}
	return true
}

func (s *Solution) null() *Solution {
	return &Solution{}
}

func (s *Solution) isNull() bool {
	return !s.validate()
}

func (s *Solution) makeCopy() *Solution {
	sCopy := &Solution{}
	sCopy.setValues(make([]complex128, s.size), s.size)
	return sCopy
}

func (s *Solution) setSelf(st *Solution) {
	s.setValues(st.slice, st.size)
}

func (s *Solution) setValues(slice []complex128, size int) {
	s.setSlice(slice)
	s.setSize(size)
}

func (s *Solution) setSlice(slice []complex128) {
	s.slice = slice
}

func (s *Solution) setSize(size int) {
	s.size = size
}

func (s *Solution) getElem(index int) complex128 {
	if !s.validateIndex(index) {
		panic(solutionIndexOutOfBoundError)
	}
	return s.slice[index]
}

func (s *Solution) setElem(index int, value complex128) {
	if !s.validateIndex(index) {
		panic(solutionIndexOutOfBoundError)
	}
	s.slice[index] = value
}

func (s *Solution) Display(realPrecision, imagPrecision int) *Solution {
	if s.size ==
		solutionNoSize {
		fmt.Printf("[null]")
	}
	fmt.Printf("[")
	lang.DisplayComplex128(realPrecision, imagPrecision, s.getElem(0))
	for idx := 1; idx < s.size; idx++ {
		fmt.Printf(", ")
		lang.DisplayComplex128(realPrecision, imagPrecision, s.getElem(idx))
	}
	fmt.Printf("]")
	return s
}

type PolyEquation struct {
	poly     *Poly
	value    complex128
	solution *Solution
}

func (pe *PolyEquation) Solve() *Solution {
	return pe.solution
}
