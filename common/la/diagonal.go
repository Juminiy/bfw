package la

import (
	"bfw/common/lang"
	"errors"
)

const (
	diagonalNoSize          int = 0
	diagonalIndexOutOfBound int = -1
)

var (
	diagonalInValidError         = errors.New("diagonal is invalid")
	diagonalIndexOutOfBoundError = errors.New("diagonal index is out of bound")
	NullDiagonal                 = &Diagonal{}
)

type Diagonal struct {
	slice       []float64
	size        int
	shape       bool
	coefficient float64
}

func (diagonal *Diagonal) validate() bool {
	if diagonal.size == diagonalNoSize ||
		diagonal == nil ||
		len(diagonal.slice) != diagonal.size {
		return false
	}
	return true
}

func (diagonal *Diagonal) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if index[indexIdx] < 0 ||
				index[indexIdx] >= diagonal.size {
				return false
			}
		}
	}
	return true
}

func (diagonal *Diagonal) sameShape(d *Diagonal) bool {
	if d == nil ||
		diagonal.size != d.size {
		return false
	}
	return true
}

func (diagonal *Diagonal) makeCopy() *Diagonal {
	dCopy := &Diagonal{}
	dCopy.setValues(make([]float64, diagonal.size), diagonal.size)
	copy(dCopy.slice, diagonal.slice)
	return dCopy
}

func (diagonal *Diagonal) set(index int, value float64) {
	diagonal.slice[index] = value
}

func (diagonal *Diagonal) get(index int) float64 {
	return diagonal.slice[index]
}

func (diagonal *Diagonal) setSlice(slice []float64) {
	diagonal.slice = slice
}

func (diagonal *Diagonal) setSize(size int) {
	diagonal.size = size
}

func (diagonal *Diagonal) setValues(slice []float64, size int) {
	diagonal.setSize(size)
	diagonal.setSlice(slice)
}

func (diagonal *Diagonal) setSelf(d *Diagonal) {
	diagonal.setValues(d.slice, d.size)
}

func (diagonal *Diagonal) null() *Diagonal {
	return &Diagonal{}
}

// one2OneOpt
// change self
// chained option
func (diagonal *Diagonal) one2OneOpt(opt rune, d ...*Diagonal) *Diagonal {
	if len(d) > 0 {
		if !diagonal.sameShape(d[0]) {
			panic(matrixNotSameShapeError)
		}
	}
	for idx := 0; idx < diagonal.size; idx++ {
		switch opt {
		case '+':
			{
				diagonal.slice[idx] += d[0].slice[idx]
			}
		case '-':
			{
				diagonal.slice[idx] -= d[0].slice[idx]
			}
		case '*':
			{
				diagonal.slice[idx] *= d[0].slice[idx]
			}
		case 'i':
			{
				diagonal.slice[idx] = 1.0 / diagonal.slice[idx]
			}
		default:
			{

			}
		}
	}
	return diagonal
}

// Add
// change self
// chained option
func (diagonal *Diagonal) Add(d *Diagonal) *Diagonal {
	return diagonal.one2OneOpt('+', d)
}

// Sub
// change self
// chained option
func (diagonal *Diagonal) Sub(d *Diagonal) *Diagonal {
	return diagonal.one2OneOpt('-', d)
}

// Mul
// change self
// chained option
func (diagonal *Diagonal) Mul(d *Diagonal) *Diagonal {
	return diagonal.one2OneOpt('*', d)
}

func (diagonal *Diagonal) inverse() *Diagonal {
	return diagonal.one2OneOpt('i')
}

func (diagonal *Diagonal) GetInverse() *Diagonal {
	dCopy := diagonal.makeCopy()
	return dCopy.inverse()
}

func (diagonal *Diagonal) IsIdentity() bool {
	if size := diagonal.size; size != matrixNoSize {
		for idx := 0; idx < size; idx++ {
			if !lang.EqualFloat64ByAccuracy(identityMatrixValue, diagonal.get(idx)) {
				return false
			}
		}
		return true
	}
	return false
}

func (diagonal *Diagonal) Identity() *Identity {
	if !diagonal.IsIdentity() {
		panic(matrixCanNotBeIdentityError)
	}
	return &Identity{diagonal.size}
}

func (diagonal *Diagonal) Matrix() *Matrix {
	size := diagonal.size
	matrix := &Matrix{}
	matrix.assign(size, size)
	for rowIdx := 0; rowIdx < size; rowIdx++ {
		for lineIdx := 0; lineIdx < size; lineIdx++ {
			if rowIdx == lineIdx {
				matrix.set(rowIdx, lineIdx, diagonal.slice[rowIdx])
			}
		}
	}
	return matrix
}
