package la

import (
	"bfw/common/lang"
	"errors"
	"fmt"
	"math"
)

const (
	vectorNoSize          int = 0
	vectorIndexOutOfBound int = -1
	vectorSizeOne         int = 1
)

var (
	vectorNotSameShapeError      = errors.New("two vectors are not same shape")
	vectorCanNotMultiplyError    = errors.New("two vectors cannot cross multiply")
	vectorCanNotDotMultiplyError = errors.New("two vectors cannot dot multiply")
	vectorCanNotBeMatrixError    = errors.New("vector cannot be matrix")
	vectorIndexOutOfBoundError   = errors.New("vector index is out of bound")
	vectorInValidError           = errors.New("vector is invalid")
	NullVector                   = &Vector{}
)

// Vector
// shape: 1*m -> false, m*1 -> true
type Vector struct {
	slice       []float64
	size        int
	shape       bool
	coefficient float64
}

func ConstructVector(real1DArray []float64, shape ...bool) *Vector {
	vector := &Vector{}
	return vector.Construct(real1DArray, shape...)
}

func (vector *Vector) Construct(real1DArray []float64, shape ...bool) *Vector {
	if real1DArray == nil ||
		len(real1DArray) == 0 {
		return vector
	}
	var (
		vShape = true
	)
	if len(shape) > 0 && !shape[0] {
		vShape = false
	}
	vector.setValues(real1DArray, len(real1DArray), vShape)
	return vector
}

func (vector *Vector) validate() bool {
	if vector.size == vectorNoSize ||
		vector.slice == nil ||
		len(vector.slice) != vector.size {
		return false
	}
	return true
}

func (vector *Vector) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if index[indexIdx] < 0 ||
				index[indexIdx] >= vector.size {
				return false
			}
		}
	}
	return true
}

func (vector *Vector) isNull() bool {
	return !vector.validate()
}

func (vector *Vector) null() *Vector {
	return &Vector{}
}

func (vector *Vector) setNull() {
	vector.setValues(nil, vectorNoSize, false)
}

func (vector *Vector) makeCopy() *Vector {
	vCopy := &Vector{}
	vCopy.setValues(make([]float64, vector.size), vector.size, vector.shape)
	copy(vCopy.slice, vector.slice)
	return vCopy
}

func (vector *Vector) Equal(v *Vector) bool {
	if vector.sameShape(v) {
		for idx := 0; idx < vector.size; idx++ {
			if !lang.EqualFloat64ByAccuracy(vector.get(idx, idx), v.get(idx, idx)) {
				return false
			}
		}
		return true
	}
	return false
}

// assign
// function:
//
//	1.
//
// character:
//  1. change self
func (vector *Vector) assign(shape bool, size int) {
	vector.shape = shape
	vector.size = size
	vector.slice = make([]float64, size)
}

func (vector *Vector) swap(v *Vector) {
	vTemp := &Vector{}
	vTemp.setSelf(vector)
	vector.setSelf(v)
	v.setSelf(vTemp)
}

func (vector *Vector) get(index, indexRedundancy int) float64 {
	indexRedundancy = index
	if index < vector.size {
		return vector.slice[index]
	}
	panic(vectorIndexOutOfBoundError)
}

func (vector *Vector) setOpt(index, indexRedundancy int, value float64, opt ...rune) {
	if !vector.validateIndex(index, indexRedundancy) {
		panic(vectorIndexOutOfBound)
	}
	if len(opt) > 0 {
		switch opt[0] {
		case '=':
			{
				vector.slice[index] = value
			}
		case '+':
			{
				vector.slice[index] += value
			}
		case '-':
			{
				vector.slice[index] -= value
			}
		case '*':
			{
				vector.slice[index] *= value
			}
		case '/':
			{
				vector.slice[index] /= value
			}
		default:
			{

			}
		}
	}
}

// set
// change self
func (vector *Vector) set(index, indexRedundancy int, value float64) {
	vector.setOpt(index, indexRedundancy, value, '=')
}

func (vector *Vector) getSelf() *Vector {
	return vector
}

func (vector *Vector) setSelf(v *Vector) {
	vector.setValues(v.slice, v.size, v.shape)
}

// setValues
// change self
func (vector *Vector) setValues(slice []float64, size int, shape bool) {
	vector.setSlice(slice)
	vector.setSize(size)
	vector.setShape(shape)
}

func (vector *Vector) setSlice(slice []float64) {
	vector.slice = slice
}

func (vector *Vector) setSize(size int) {
	vector.size = size
}

func (vector *Vector) setShape(shape bool) {
	vector.shape = shape
}

func (vector *Vector) setElemSwap(indexI, indexJ int) {
	if !vector.validateIndex(indexI, indexJ) {
		panic(vectorIndexOutOfBound)
	}
	vector.slice[indexI], vector.slice[indexJ] =
		vector.slice[indexJ], vector.slice[indexI]
}

func (vector *Vector) sameShape(v *Vector) bool {
	if v == nil ||
		vector.shape != v.shape ||
		vector.size != v.size {
		return false
	}
	return true
}

func (vector *Vector) GetSlice() []float64 {
	return vector.slice
}

// Vector Opt Vector Res Vector

// canMultiply
// m*1 * 1*n
// m*1 * n*1
// 1*m * 1*n
// 1*m * n*1
func (vector *Vector) canMultiply(v *Vector) bool {
	if vector == nil ||
		v == nil {
		return false
	}
	if vector.shape {
		if !v.shape ||
			(v.shape && v.size == vectorSizeOne) {
			return true
		} else {
			return false
		}
	} else {
		if (!v.shape && vector.size == vectorSizeOne) ||
			(v.shape && vector.size == v.size) {
			return true
		} else {
			return false
		}
	}
}

// one2OneOpt
// change self
// chained option
func (vector *Vector) one2OneOpt(opt rune, v *Vector) *Vector {
	if !vector.sameShape(v) {
		panic(vectorNotSameShapeError)
	}
	for idx := 0; idx < vector.size; idx++ {
		switch opt {
		case '+':
			{
				vector.setOpt(idx, idx, v.get(idx, idx), '+')
			}
		case '-':
			{
				vector.setOpt(idx, idx, v.get(idx, idx), '-')
			}
		case '*':
			{
				vector.setOpt(idx, idx, v.get(idx, idx), '*')
			}
		default:
			{

			}
		}
	}
	return vector
}

// Transpose
// change self
// chained option
func (vector *Vector) Transpose() *Vector {
	vector.shape = !vector.shape
	return vector
}

func (vector *Vector) GetPlus(v *Vector) *Vector {
	vCopy := vector.makeCopy()
	return vCopy.Add(v)
}

// Add
// change self
// chained option
func (vector *Vector) Add(v *Vector) *Vector {
	return vector.one2OneOpt('+', v)
}

func (vector *Vector) GetMinus(v *Vector) *Vector {
	vCopy := vector.makeCopy()
	return vCopy.Sub(v)
}

// Sub
// change self
// chained option
func (vector *Vector) Sub(v *Vector) *Vector {
	return vector.one2OneOpt('-', v)
}

func (vector *Vector) GetMTimes(v *Vector) *Matrix {
	vCopy := vector.makeCopy()
	return vCopy.Mul(v)
}

// Mul
// change self
// chained option
func (vector *Vector) Mul(v *Vector) *Matrix {
	if !vector.canMultiply(v) {
		panic(vectorCanNotMultiplyError)
	} else {
		return vector.convertToMatrix().Mul(v.convertToMatrix())
	}
}

// MulLambda
// change self
// chained option
func (vector *Vector) MulLambda(lambda float64) *Vector {
	for idx := 0; idx < vector.size; idx++ {
		vector.setOpt(idx, idx, lambda, '*')
	}
	return vector
}

// dotMulByRange
// [startIndex, endIndex]
func (vector *Vector) dotMulByRange(v *Vector, startIndex, endIndex int, gap ...int) float64 {
	var (
		gapValue int     = 0
		dotSum   float64 = 0.0
	)
	if len(gap) > 0 {
		gapValue = gap[0]
	}
	if startIndex < 0 ||
		endIndex >= vector.size ||
		startIndex+gapValue < 0 ||
		endIndex+gapValue >= v.size ||
		startIndex > endIndex {
		panic(vectorIndexOutOfBoundError)
	}

	for idx := startIndex; idx <= endIndex; idx++ {
		dotSum += vector.get(idx, idx) * v.get(idx+gapValue, idx+gapValue)
	}
	return dotSum
}

func (vector *Vector) DotMul(v *Vector) float64 {
	return vector.dotMulByRange(v, 0, vector.size-1)
}

func (vector *Vector) CrossMul(v *Vector) *Matrix {
	return vector.Mul(v)
}

func (vector *Vector) InnerMul() float64 {
	return vector.DotMul(vector)
}

func (vector *Vector) OuterMul(v *Vector) *Matrix {
	return vector.CrossMul(v)
}

// MulMatrix
// change self
// chained option
func (vector *Vector) MulMatrix(m *Matrix) *Vector {
	if m != nil &&
		!vector.shape &&
		vector.size == m.rowSize {
		vCopy := vector.makeCopy()
		for columnIdx := 0; columnIdx < m.columnSize; columnIdx++ {
			var columnSum float64 = 0.0
			for rowIdx := 0; rowIdx < m.rowSize; rowIdx++ {
				columnSum += vCopy.slice[rowIdx] * m.get(rowIdx, columnIdx)
			}
			vector.set(columnIdx, columnIdx, columnSum)
		}
		return vector
	}
	return &Vector{}
}

// PowerMatrix
// change self
// chained option
func (vector *Vector) PowerMatrix(m *Matrix, n int) *Vector {
	for t := 0; t < n; t++ {
		vector.MulMatrix(m)
	}
	return vector
}

// Vector Scaling flexible

func (vector *Vector) padding(size int, value float64) *Vector {
	if vector.size < size {
		vector.setSize(size)
		for idx := vector.size; idx < size; idx++ {
			vector.slice = append(vector.slice, value)
		}
	}
	return vector
}

func (vector *Vector) ZeroPadding(size int) *Vector {
	return vector.padding(size, 0.0)
}

func (vector *Vector) GetUnit() *Vector {
	vCopy := vector.makeCopy()
	return vCopy.unit()
}

func (vector *Vector) unit() *Vector {
	euNorm := vector.euclideanNorm()
	for idx := 0; idx < vector.size; idx++ {
		vector.setOpt(idx, idx, euNorm, '/')
	}
	return vector
}

func (vector *Vector) validateUnit() bool {
	return lang.EqualFloat64ByAccuracy(1.0, vector.euclideanNorm())
}

func (vector *Vector) validateOrthogonal(v *Vector) bool {
	return lang.EqualFloat64Zero(vector.DotMul(v))
}

func (vector *Vector) Reverse() *Vector {
	vSize := vector.size
	for idx := 0; idx < vSize>>1; idx++ {
		vector.setElemSwap(idx, vSize-idx-1)
	}
	return vector
}

func (vector *Vector) Convolution(v *Vector) *Vector {
	v1 := vector.makeCopy()
	v2 := v.makeCopy()
	v1Size, v2Size, destSize := v1.size, v2.size, 0
	if v1Size > v2Size {
		v2.ZeroPadding(v1Size)
		destSize = v1Size
	} else if v2Size > v1Size {
		v1.ZeroPadding(v2Size)
		destSize = v2Size
	} else {
		destSize = v1Size
	}
	v2.Reverse()
	vRes := &Vector{}
	vRes.setValues(make([]float64, destSize<<1-1), destSize<<1-1, vector.shape)
	gap, mIdx := 1, destSize-1
	lIdx, rIdx := mIdx, mIdx
	for gap < destSize {
		lIdx--
		rIdx++
		vRes.set(lIdx, lIdx, v1.dotMulByRange(v2, 0, lIdx, gap))
		vRes.set(rIdx, lIdx, v1.dotMulByRange(v2, gap, mIdx, -gap))
		gap++
	}
	vRes.set(mIdx, mIdx, v1.DotMul(v2))
	return vRes
}

// Vector Norm Theory

func (vector *Vector) norm(opt rune, extra ...interface{}) float64 {
	if !vector.validate() {
		panic(vectorInValidError)
	}
	var (
		resNum float64 = 0.0
		powerP float64 = 0.0
	)
	if len(extra) > 0 {
		powerP = extra[0].(float64)
		if powerP == math.Inf(1) {
			opt = '+'
		} else if powerP == math.Inf(-1) {
			opt = '-'
		}
	}
	for idx := 0; idx < vector.size; idx++ {
		val := vector.get(idx, idx)
		switch opt {
		case '0':
			{
				if !lang.EqualFloat64Zero(val) {
					resNum += 1.0
				}
			}
		case '1', 'm':
			{
				resNum += math.Abs(val)
			}
		case '2', 'e':
			{
				resNum += lang.Square(val)
			}
		case 'i', '+':
			{
				resNum = math.Max(resNum, math.Abs(val))
			}
		case '-':
			{
				resNum = math.Min(resNum, math.Abs(val))
			}
		case 'p':
			{
				resNum += math.Pow(math.Abs(val), powerP)
			}
		default:
			{

			}
		}

	}
	if opt == '2' || opt == 'm' {
		resNum = math.Sqrt(resNum)
	} else if opt == 'p' {
		resNum = math.Pow(resNum, 1.0/powerP)
	}
	return resNum
}

func (vector *Vector) manhattanDistance() float64 {
	return vector.norm('m')
}

// euclideanNorm
// 2-norm
// ||x||_{2}
func (vector *Vector) euclideanNorm() float64 {
	return vector.norm('e')
}

func (vector *Vector) L0Norm() float64 {
	return vector.norm('0')
}

func (vector *Vector) L1Norm() float64 {
	return vector.manhattanDistance()
}

func (vector *Vector) L2Norm() float64 {
	return vector.euclideanNorm()
}

func (vector *Vector) PositiveInfiniteNorm() float64 {
	return vector.norm('+')
}

func (vector *Vector) NegativeInfiniteNorm() float64 {
	return vector.norm('-')
}

func (vector *Vector) InfiniteNorm() float64 {
	return vector.PositiveInfiniteNorm()
}
func (vector *Vector) PNorm(p float64) float64 {
	return vector.norm('p', p)
}

func (vector *Vector) GetProjection(v *Vector) *Vector {
	vCopy := vector.makeCopy()
	return vCopy.projection(v)
}

func (vector *Vector) projection(v *Vector) *Vector {
	return vector
}

func (vector *Vector) IsLinearCorrelation(v *Vector) bool {
	return false
}

func (vector *Vector) IsUnitary() bool {
	return false
}

func (vector *Vector) IsNormal() bool {
	return false
}

func (vector *Vector) IsHermite() bool {
	return false
}

// AI algorithm

// Convergence
// change self
// chained option
func (vector *Vector) Convergence(matrix *Matrix) (*Vector, int) {
	convergenceVector := vector.makeCopy()
	convergenceIteratorTime := 0
	for {
		preVector := convergenceVector.makeCopy()
		convergenceIteratorTime++
		convergenceVector.PowerMatrix(matrix, 1)
		if convergenceVector.Equal(preVector) {
			break
		}
	}
	return convergenceVector, convergenceIteratorTime
}

// Display
// chained option
func (vector *Vector) Display() *Vector {
	if vector.isNull() {
		fmt.Println("[null]")
		return vector
	}
	fmt.Printf("%c", '[')
	for idx := 0; idx < vector.size; idx++ {
		val := vector.get(idx, idx)
		if lang.EqualFloat64Zero(val) {
			val = 0.0
		}
		fmt.Printf(" %.5v", val)
	}
	fmt.Printf("%c", ']')
	if vector.shape {
		fmt.Printf("%c\n", 'ᐪ')
	}
	return vector
}

func (vector *Vector) Matrix() *Matrix {
	return vector.convertToMatrix()
}

func (vector *Vector) convertToMatrix() *Matrix {
	matrix := &Matrix{}
	if vector.shape {
		matrix.setValues(make([][]float64, vector.size), vector.size, 1)
		for idx := 0; idx < vector.size; idx++ {
			matrix.setRow(idx, make([]float64, 1))
			matrix.set(idx, 0, vector.get(idx, idx))
		}
	} else {
		matrix.setValues(make([][]float64, 1), 1, vector.size)
		copy(matrix.slice[0], vector.slice)
	}
	return matrix
}
