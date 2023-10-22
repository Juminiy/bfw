package la

import (
	"bfw/wheel/lang"
	"errors"
)

const (
	vandermondeNoSize int = 0
)

var (
	vandermondeIndexOutOfBound = errors.New("vanermonde index is out of bound")
	vandermondeInvalidError    = errors.New("vandermonde is invalid")
)

type Vandermonde struct {
	values []float64
}

func ConstructVandermonde(values []float64) *Vandermonde {
	v := &Vandermonde{}
	return v.Construct(values)
}

func (v *Vandermonde) Construct(values []float64) *Vandermonde {
	v.setValues(values)
	return v
}

func (v *Vandermonde) setSelf(vt *Vandermonde) {
	v.setValues(vt.values)
}

func (v *Vandermonde) setValues(values []float64) {
	v.setValuesValues(values)
}

func (v *Vandermonde) setValuesValues(values []float64) {
	v.values = nil
	v.values = values
}

func (v *Vandermonde) size() int {
	return len(v.values)
}

func (v *Vandermonde) validateIndex(index int) bool {
	return index >= 0 &&
		index < v.size()
}

func (v *Vandermonde) validate() bool {
	return v.size() == vandermondeNoSize
}

func (v *Vandermonde) get(index, indexRedundancy int) float64 {
	if !v.validateIndex(index) {
		panic(vandermondeIndexOutOfBound)
	}
	return v.values[index]
}

func (v *Vandermonde) set(index, indexRedundancy int, value float64) {
	// do nothing
}

func (v *Vandermonde) Det() float64 {
	return v.det()
}

func (v *Vandermonde) det() float64 {
	if !v.validate() {
		panic(vandermondeInvalidError)
	}
	var (
		realResult float64 = 1.0
	)
	for iIdx := 1; iIdx < v.size(); iIdx++ {
		for jIdx := 0; jIdx < iIdx; jIdx++ {
			tRes := v.values[iIdx] - v.values[jIdx]
			if lang.EqualFloat64Zero(tRes) {
				return 0.0
			}
			realResult *= tRes
		}
	}
	return realResult
}

func (v *Vandermonde) Matrix() *Matrix {
	var (
		matrix = &Matrix{}
		vSize  = v.size()
		setVal float64
	)
	matrix.assign(vSize, vSize)
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
			if rowIdx == 0 {
				setVal = 1.0
			} else {
				setVal = matrix.get(rowIdx-1, columnIdx) * v.get(columnIdx, columnIdx)
			}
			matrix.set(rowIdx, columnIdx, setVal)
		}
	}
	return matrix
}
