package la

import (
	"bfw/common/lang"
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	matrixNotPhalanx       int     = -1
	matrixNoSize           int     = 0
	matrixIndexOutOfBound  int     = -1
	matrixDeterminantZero  float64 = 0
	matrixNoRank           int     = -1
	matrixToVectorRowSize  int     = 1
	matrixToVectorLineSize int     = 1
	simplePhalanxSizeOne   int     = 1
	simplePhalanxSizeTwo   int     = 2
	simplePhalanxSizeThree int     = 3
)

var (
	matrixNotSameShapeError           = errors.New("two matrices are not same shape")
	matrixCanNotMultiplyError         = errors.New("two matrices cannot multiply")
	matrixCanNotHadamardMultiplyError = errors.New("two matrices cannot hadamard multiply")
	matrixRowLineDiffer               = errors.New("matrix row and line differ")
	matrixCanNotBeVectorError         = errors.New("matrix cannot be vector")
	matrixIndexOutOfBoundError        = errors.New("matrix index is out of bound")
	matrixNotPhalanxError             = errors.New("matrix is not phalanx")
	matrixCanNotBeAdjoinError         = errors.New("matrix cannot be adjoin")
	matrixCanNotBeInverseError        = errors.New("matrix cannot be inverse")
	matrixCanNotBeIdentityError       = errors.New("matrix cannot be identity")
	matrixCanNotBeReShapedError       = errors.New("matrix cannot be reshaped")
	matrixInValidError                = errors.New("matrix is invalid")
	NullMatrix                        = &Matrix{}
)

type Matrix struct {
	slice       [][]float64
	rowSize     int
	lineSize    int
	coefficient float64
}

func (matrix *Matrix) validate() bool {
	if matrix.rowSize == matrixNoSize ||
		matrix.lineSize == matrixNoSize ||
		matrix.slice == nil ||
		len(matrix.slice) != matrix.rowSize {
		return false
	} else {
		for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
			if idxRow := matrix.getRow(rowIdx); idxRow == nil ||
				len(idxRow) != matrix.lineSize {
				return false
			}
		}
	}
	return true
}

func (matrix *Matrix) validateOneIndex(index int) bool {
	if index < 0 ||
		index >= matrix.rowSize ||
		index >= matrix.lineSize {
		return false
	}
	return true
}

func (matrix *Matrix) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		if indexLen >= 1 {
			if index[0] < 0 || index[0] >= matrix.rowSize {
				return false
			}
		}
		if indexLen >= 2 {
			if index[1] < 0 || index[1] >= matrix.lineSize {
				return false
			}
		}
		if indexLen >= 3 {
			for indexIdx := 2; indexIdx < indexLen; indexIdx++ {
				if !matrix.validateOneIndex(index[indexIdx]) {
					return false
				}
			}
		}
	}
	return true
}

// assign
// 1. character
// change self
// 2.function
// initial the location
func (matrix *Matrix) assign(rowSize, lineSize int) {
	matrix.setValues(make([][]float64, rowSize), rowSize, lineSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		matrix.setRow(rowIdx, make([]float64, lineSize))
	}
}

func (matrix *Matrix) get(rowIndex, lineIndex int) float64 {
	if !matrix.validateIndex(rowIndex, lineIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return matrix.slice[rowIndex][lineIndex]
}

func (matrix *Matrix) setOpt(rowIndex, lineIndex int, value float64, opt ...rune) {
	if !matrix.validateIndex(rowIndex, lineIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	if len(opt) > 0 {
		switch opt[0] {
		case '=':
			{
				matrix.slice[rowIndex][lineIndex] = value
			}
		case '+':
			{
				matrix.slice[rowIndex][lineIndex] += value
			}
		case '-':
			{
				matrix.slice[rowIndex][lineIndex] -= value
			}
		case '*':
			{
				matrix.slice[rowIndex][lineIndex] *= value
			}
		case '/':
			{
				matrix.slice[rowIndex][lineIndex] /= value
			}
		case '^':
			{
				matrix.slice[rowIndex][lineIndex] = math.Pow(
					matrix.get(rowIndex, lineIndex), value)
			}
		default:
			{

			}
		}
	}
}

// set
// change self
func (matrix *Matrix) set(rowIndex, lineIndex int, value float64) {
	matrix.setOpt(rowIndex, lineIndex, value, '=')
}

func (matrix *Matrix) setRow(rowIndex int, rowSlice []float64) {
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	matrix.slice[rowIndex] = rowSlice
}

func (matrix *Matrix) getRow(rowIndex int) []float64 {
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return matrix.slice[rowIndex]
}

func (matrix *Matrix) getRowCopy(rowIndex int) []float64 {
	indexRow := matrix.getRow(rowIndex)
	sliceCopy := make([]float64, len(indexRow))
	copy(sliceCopy, indexRow)
	return sliceCopy
}

func (matrix *Matrix) setLine(lineIndex int, lineSlice []float64) {
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		matrix.set(rowIdx, lineIndex, lineSlice[rowIdx])
	}
}

func (matrix *Matrix) getLine(lineIndex int) []float64 {
	return lang.GetReal2DArrayLine(matrix.GetSlice(), lineIndex)
}

func (matrix *Matrix) getLineCopy(lineIndex int) []float64 {
	return matrix.getLine(lineIndex)
}

func (matrix *Matrix) getRowIndexLen(rowIndex int) int {
	return len(matrix.getRow(rowIndex))
}

func (matrix *Matrix) setValues(slice [][]float64, size ...int) {
	matrix.setSlice(slice)
	matrix.setSize(size...)
}

func (matrix *Matrix) setSelf(m *Matrix) {
	matrix.setValues(m.slice, m.rowSize, m.lineSize)
}

func (matrix *Matrix) getSelf() *Matrix {
	return matrix
}

// setSlice
// change self
func (matrix *Matrix) setSlice(slice [][]float64) {
	matrix.slice = slice
}

func (matrix *Matrix) GetSlice() [][]float64 {
	if !matrix.validate() {
		panic(matrixInValidError)
	}
	return matrix.slice
}

func (matrix *Matrix) rowPadding(rowIndex, lineSize int, value float64) {
	if lineSize <= matrix.lineSize {
		return
	}
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	for lineIdx := matrix.getRowIndexLen(rowIndex); lineIdx < lineSize; lineIdx++ {
		matrix.slice[rowIndex] = append(matrix.slice[rowIndex], value)
	}
}

func (matrix *Matrix) rowZeroPadding(rowIndex, lineSize int) {
	matrix.rowPadding(rowIndex, lineSize, 0.0)
}

func (matrix *Matrix) setRowZeroPadding(lineSize int, rowIndex ...int) {
	for rowIndexIdx := 0; rowIndexIdx < len(rowIndex); rowIndexIdx++ {
		matrix.rowZeroPadding(rowIndex[rowIndexIdx], lineSize)
	}
}

// setSize
// change self
func (matrix *Matrix) setSize(size ...int) {
	if sizeLen := len(size); sizeLen > 0 {
		if sizeLen >= 1 {
			matrix.rowSize = size[0]
			matrix.lineSize = size[0]
		}
		if sizeLen >= 2 {
			matrix.lineSize = size[1]
		}
	}
}

func (matrix *Matrix) GetRowSize() int {
	return matrix.rowSize
}

func (matrix *Matrix) GetLineSize() int {
	return matrix.lineSize
}

func (matrix *Matrix) isNull() bool {
	return !matrix.validate()
}

func (matrix *Matrix) null() *Matrix {
	return &Matrix{}
}

func (matrix *Matrix) setNull() {
	matrix.setValues(nil, matrixNoSize, matrixNoSize)
}

func (matrix *Matrix) remake(m ...*Matrix) {
	if len(m) > 0 {
		matrix.setSelf(m[0])
	} else {
		matrix.setValues(nil, matrixNoSize)
	}
}

func (matrix *Matrix) sameShape(m *Matrix) bool {
	if m == nil ||
		matrix.rowSize != m.rowSize ||
		matrix.lineSize != m.lineSize {
		return false
	}
	return true
}

func (matrix *Matrix) canMultiply(m *Matrix) bool {
	if m == nil ||
		matrix.lineSize != m.rowSize {
		return false
	}
	return true
}

func (matrix *Matrix) isPhalanx() bool {
	return matrix.rowSize == matrix.lineSize
}

func (matrix *Matrix) getPhalanxSize() int {
	if matrix.isPhalanx() {
		return matrix.rowSize
	}
	return matrixNotPhalanx
}

// getPhalanxByIgnoreCentral
// chained option
func (matrix *Matrix) getPhalanxByIgnoreCentral(rowIndex, lineIndex int) *Matrix {
	phalanxSize := matrix.getPhalanxSize()
	remainedPhalanxSize := phalanxSize - 1
	remainedCnt, remainedRowIdx, remainedLineIdx := -1, 0, 0
	remainedPhalanx := &Matrix{}
	remainedPhalanx.assign(remainedPhalanxSize, remainedPhalanxSize)
	for rowIdx := 0; rowIdx < phalanxSize; rowIdx++ {
		for lineIdx := 0; lineIdx < phalanxSize; lineIdx++ {
			if rowIdx != rowIndex && lineIdx != lineIndex {
				remainedCnt++
				remainedRowIdx = remainedCnt / remainedPhalanxSize
				remainedLineIdx = remainedCnt % remainedPhalanxSize
				remainedPhalanx.set(remainedRowIdx, remainedLineIdx, matrix.get(rowIdx, lineIdx))
			}
		}
	}
	return remainedPhalanx
}

// getPhalanxByIgnoreKCentral
// chained option
func (matrix *Matrix) getPhalanxByIgnoreKCentral(rowIndexMap, lineIndexMap map[int]bool) (*Matrix, *Matrix) {
	if rowIndexMap == nil || lineIndexMap == nil {
		return &Matrix{}, &Matrix{}
	}
	phalanxSize := matrix.getPhalanxSize()
	rowsK, linesK := len(rowIndexMap), len(lineIndexMap)
	if rowsK != linesK ||
		phalanxSize < rowsK {
		return &Matrix{}, &Matrix{}
	}
	selected, remained := &Matrix{}, &Matrix{}
	selectedSize, remainedSize := rowsK, phalanxSize-rowsK
	selected.assign(selectedSize, selectedSize)
	remained.assign(remainedSize, remainedSize)
	slCnt, slRowIdx, slLineIdx := -1, 0, 0
	rmCnt, rmRowIdx, rmLineIdx := -1, 0, 0
	for rowIdx := 0; rowIdx < phalanxSize; rowIdx++ {
		for lineIdx := 0; lineIdx < phalanxSize; lineIdx++ {
			if !rowIndexMap[rowIdx] && !lineIndexMap[lineIdx] {
				rmCnt++
				rmRowIdx = rmCnt / remainedSize
				rmLineIdx = rmCnt % remainedSize
				remained.set(rmRowIdx, rmLineIdx, matrix.get(rowIdx, lineIdx))
			} else if rowIndexMap[rowIdx] && lineIndexMap[lineIdx] {
				slCnt++
				slRowIdx = slCnt / selectedSize
				slLineIdx = slCnt % selectedSize
				selected.set(slRowIdx, slLineIdx, matrix.get(rowIdx, lineIdx))
			}
		}
	}
	return selected, remained
}

func (matrix *Matrix) getPhalanxMinusOnePower(rowIndexMap, lineIndexMap map[int]bool) int {
	totalPower := 0
	for row, _ := range rowIndexMap {
		totalPower += row
	}
	for line, _ := range lineIndexMap {
		totalPower += line
	}
	return totalPower
}

func (matrix *Matrix) makeCopy() *Matrix {
	mCopy := &Matrix{}
	mCopy.setValues(make([][]float64, matrix.rowSize), matrix.rowSize, matrix.lineSize)
	for rowIdx := 0; rowIdx < mCopy.rowSize; rowIdx++ {
		mCopy.setRow(rowIdx, make([]float64, matrix.lineSize))
		copy(mCopy.slice[rowIdx], matrix.getRow(rowIdx))
	}
	return mCopy
}

func (matrix *Matrix) Construct(real2DArray [][]float64) *Matrix {
	if real2DArray == nil ||
		len(real2DArray) == matrixNoSize {
		return matrix
	}
	var (
		maxLineSize int = 1
		rowSize     int = len(real2DArray)
	)
	matrix.setValues(real2DArray, rowSize, maxLineSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		maxLineSize = lang.MaxInt(maxLineSize, len(real2DArray[rowIdx]))
	}
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		matrix.setRowZeroPadding(maxLineSize, rowIdx)
	}
	matrix.setSize(rowSize, maxLineSize)
	return matrix
}

func ConstructMatrix(real2DArray [][]float64) *Matrix {
	matrix := &Matrix{}
	return matrix.Construct(real2DArray)
}

// Rank
// uncompleted
func (matrix *Matrix) Rank() int {
	if lang.EqualFloat64ByAccuracy(matrix.Det(), matrixDeterminantZero) {
		return matrix.getPhalanxSize()
	}
	return 0
}

func (matrix *Matrix) Det() float64 {
	if !matrix.isPhalanx() {
		panic(matrixRowLineDiffer)
	}
	return matrix.det()
}

func (matrix *Matrix) det() float64 {
	if n := matrix.getPhalanxSize(); n == simplePhalanxSizeOne {
		return matrix.get(0, 0)
	} else if n == simplePhalanxSizeTwo {
		return matrix.get(0, 0)*matrix.get(1, 1) -
			matrix.get(0, 1)*matrix.get(1, 0)
	} else if n == simplePhalanxSizeThree {
		return matrix.get(0, 0)*matrix.get(1, 1)*matrix.get(2, 2) +
			matrix.get(1, 0)*matrix.get(2, 1)*matrix.get(0, 2) +
			matrix.get(0, 1)*matrix.get(1, 2)*matrix.get(2, 0) -
			matrix.get(0, 2)*matrix.get(1, 1)*matrix.get(2, 0) -
			matrix.get(0, 1)*matrix.get(1, 0)*matrix.get(2, 2) -
			matrix.get(1, 2)*matrix.get(2, 1)*matrix.get(0, 0)
	} else {
		// 1. simple calculation
		//return matrix.simpleDet(n)
		// 2. laplace calculation
		//return matrix.laplaceDet(n)
		// 3. mixture calculation
		if lang.Odd(n) {
			return matrix.simpleDet(n)
		} else {
			return matrix.laplaceDet(n)
		}
	}
}

func (matrix *Matrix) simpleDet(totalN int) float64 {
	var (
		randomRow      int     = lang.GetRandomIntValue(matrix.getPhalanxSize())
		sumByRandomRow float64 = 0.0
	)
	for lineIdx := 0; lineIdx < totalN; lineIdx++ {
		sumByRandomRow += matrix.getPhalanxByIgnoreCentral(randomRow, lineIdx).det() *
			matrix.get(randomRow, lineIdx) *
			float64(lang.MinusOnePower(randomRow+lineIdx))
	}
	return sumByRandomRow
}

func (matrix *Matrix) laplaceDet(totalN int) float64 {
	var (
		randomKRow      int            = matrix.getPhalanxSize() / 2
		randomKRows     map[int]bool   = lang.GetRandomMapValue(totalN, randomKRow)
		kLinesMap       []map[int]bool = lang.GetCombinationSliceMap(totalN, randomKRow)
		sumByRandomKRow float64        = 0.0
	)
	for _, linesMap := range kLinesMap {
		sel, rem := matrix.getPhalanxByIgnoreKCentral(randomKRows, linesMap)
		power := matrix.getPhalanxMinusOnePower(randomKRows, linesMap)
		sumByRandomKRow += sel.det() * rem.det() * float64(lang.MinusOnePower(power))
	}
	return sumByRandomKRow
}

func (matrix *Matrix) GetSameSizeIdentity() *Identity {
	if size := matrix.getPhalanxSize(); size != matrixNoSize {
		return &Identity{
			size: size,
		}
	}
	return &Identity{}
}

func (matrix *Matrix) isDiagonalMatrix(value ...interface{}) bool {
	if matrix.isPhalanx() {
		var (
			compValue float64   = math.Inf(0)
			compSlice []float64 = nil
		)
		if len(value) > 0 {
			switch value[0].(type) {
			case float32, float64:
				{
					compValue = value[0].(float64)
				}
			case []float32, []float64:
				{
					compSlice = value[0].([]float64)
				}
			default:
				{

				}
			}
		}
		for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
			for lineIdx := 0; lineIdx < matrix.lineSize; lineIdx++ {
				val := matrix.get(rowIdx, lineIdx)
				if rowIdx == lineIdx {
					if compValue != math.Inf(0) {
						if !lang.EqualFloat64ByAccuracy(1.0, val) {
							return false
						}
					}
					if compSlice != nil && len(compSlice) > rowIdx {
						if !lang.EqualFloat64ByAccuracy(compSlice[rowIdx], val) {
							return false
						}
					}
				} else {
					if !lang.EqualFloat64ByAccuracy(0.0, val) {
						return false
					}
				}
			}
		}
		return true
	}
	return false
}

func (matrix *Matrix) IsIdentity() bool {
	return matrix.isDiagonalMatrix(1.0)
}

func (matrix *Matrix) Identity() *Identity {
	if matrix.IsIdentity() {
		return &Identity{size: matrix.getPhalanxSize()}
	}
	panic(matrixCanNotBeIdentityError)
}

// Inverse Matrix by Adjoin Matrix
// change self
// chained option
// |M| != 0, M^{-1} = M^{*}/|M|
func (matrix *Matrix) inverse() *Matrix {
	if !matrix.isPhalanx() {
		panic(matrixRowLineDiffer)
	}
	det := matrix.det()
	if lang.EqualFloat64ByAccuracy(det, matrixDeterminantZero) {
		panic(matrixCanNotBeInverseError)
	}
	matrix.adjoin()
	matrix.MulLambda(1 / det)
	return matrix
}

// GetInverse
// chained option
func (matrix *Matrix) GetInverse() *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.inverse()
}

// Adjoin Matrix by remained Matrix
// change self
// chained option
func (matrix *Matrix) adjoin() *Matrix {
	if !matrix.isPhalanx() {
		panic(matrixRowLineDiffer)
	}
	size := matrix.getPhalanxSize()
	adjoinMatrix := &Matrix{}
	adjoinMatrix.assign(size, size)
	// 1. each element calculate remainder value by (getPhalanxByIgnoreCentral)
	// 2. each element calculate algebraic remainder value by (MinusOnePower)
	for rowIdx := 0; rowIdx < size; rowIdx++ {
		for lineIdx := 0; lineIdx < size; lineIdx++ {
			algebraicRemainderValue := matrix.getPhalanxByIgnoreCentral(rowIdx, lineIdx).det() *
				float64(lang.MinusOnePower(rowIdx+lineIdx))
			adjoinMatrix.set(rowIdx, lineIdx, algebraicRemainderValue)
		}
	}
	// 3. matrix transpose
	adjoinMatrix.transpose()
	matrix.setSlice(adjoinMatrix.slice)
	return matrix
}

// GetAdjoin
// chained option
func (matrix *Matrix) GetAdjoin() *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.adjoin()
}

func (matrix *Matrix) RankAdjoin() int {
	if size := matrix.getPhalanxSize(); size != matrixNoSize {
		if mRank := matrix.Rank(); mRank == size {
			return size
		} else if mRank == size-1 {
			return 1
		} else {
			return 0
		}
	}
	return matrixNoRank
}

func (matrix *Matrix) Trace() float64 {
	if n := matrix.getPhalanxSize(); n > 0 {
		trace := 0.0
		for idx := 0; idx < n; idx++ {
			trace += matrix.get(idx, idx)
		}
		return trace
	}
	panic(matrixNotPhalanx)
}

func (matrix *Matrix) EigenValues() *EigenValue {
	if !matrix.isPhalanx() {
		panic(matrixRowLineDiffer)
	}
	return &EigenValue{}
}

func (matrix *Matrix) EigenVectors() *EigenVector {
	if !matrix.isPhalanx() {
		panic(matrixRowLineDiffer)
	}
	return &EigenVector{}
}

func (matrix *Matrix) EigenMatrix() *PolyMatrix {
	return &PolyMatrix{}
}

func (matrix *Matrix) SmithStandard() *PolyDiagonal {
	return &PolyDiagonal{}
}

func (matrix *Matrix) JordanStandard() *JordanMatrix {
	return &JordanMatrix{}
}

func (matrix *Matrix) dkLambda(k int) *Poly {
	return matrix.constantFactors()[k]
}

func (matrix *Matrix) constantFactors() []*Poly {
	return nil
}

func (matrix *Matrix) DkLambda(k int) *Poly {
	return matrix.determinantFactors()[k]
}

func (matrix *Matrix) determinantFactors() []*Poly {
	return nil
}

// GetTranspose
// chained option
func (matrix *Matrix) GetTranspose() *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.transpose()
}

// transpose
// change self
// chained option
func (matrix *Matrix) transpose() *Matrix {
	newRowSize, newLineSize := matrix.lineSize, matrix.rowSize
	newSlice := make([][]float64, newRowSize)
	for newRowIdx := 0; newRowIdx < newRowSize; newRowIdx++ {
		newSlice[newRowIdx] = make([]float64, newLineSize)
		for newLineIdx := 0; newLineIdx < newLineSize; newLineIdx++ {
			newSlice[newRowIdx][newLineIdx] = matrix.get(newLineIdx, newRowIdx)
		}
	}
	matrix.setValues(newSlice, newRowSize, newLineSize)
	return matrix
}

// one2OneOpt
// change self
// chained option
func (matrix *Matrix) one2OneOpt(opt rune, m *Matrix) *Matrix {
	if !matrix.sameShape(m) {
		panic(matrixNotSameShapeError)
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for lineIdx := 0; lineIdx < matrix.lineSize; lineIdx++ {
			switch opt {
			case '+':
				{
					matrix.setOpt(rowIdx, lineIdx, m.get(rowIdx, lineIdx), '+')
				}
			case '-':
				{
					matrix.setOpt(rowIdx, lineIdx, m.get(rowIdx, lineIdx), '-')
				}
			case '*':
				{
					matrix.setOpt(rowIdx, lineIdx, m.get(rowIdx, lineIdx), '*')
				}
			case '/':
				{
					matrix.setOpt(rowIdx, lineIdx, m.get(rowIdx, lineIdx), '/')
				}
			case '^':
				{
					matrix.setOpt(rowIdx, lineIdx, m.get(rowIdx, lineIdx), '^')
				}
			default:
				{

				}
			}
		}
	}
	return matrix
}

func (matrix *Matrix) convertToVector() *Vector {
	if matrix.rowSize != matrixToVectorRowSize &&
		matrix.lineSize != matrixToVectorLineSize {
		panic(matrixCanNotBeVectorError)
	}
	if matrix.rowSize == matrixToVectorRowSize {
		vector := &Vector{}
		vector.setValues(make([]float64, matrix.lineSize), matrix.lineSize, false)
		copy(vector.slice, matrix.getRow(0))
		return vector
	} else {
		vector := &Vector{}
		vector.setValues(make([]float64, matrix.rowSize), matrix.rowSize, true)
		for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
			vector.set(rowIdx, rowIdx, matrix.get(rowIdx, 0))
		}
		return vector
	}
}

func (matrix *Matrix) approximation(accAfterDotBits ...int) *Matrix {
	accBits := 5
	if len(accAfterDotBits) > 0 {
		accBits = accAfterDotBits[0]
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for lineIdx := 0; lineIdx < matrix.lineSize; lineIdx++ {
			newVal := 0.0
			valStr := fmt.Sprintf("%."+strconv.Itoa(accBits)+"f", matrix.get(rowIdx, lineIdx))
			valF, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				panic(err)
			}
			if lang.IsStringValueIntZero(valStr) {
				newVal = 0
			} else {
				newVal = valF
			}
			matrix.set(rowIdx, lineIdx, newVal)
		}
	}
	return matrix
}

func (matrix *Matrix) Approx(accBits ...int) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.approximation(accBits...)
}

func (matrix *Matrix) Equal(m *Matrix) bool {
	if matrix.sameShape(m) {
		for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
			for lineIdx := 0; lineIdx < matrix.lineSize; lineIdx++ {
				if !lang.EqualFloat64ByAccuracy(matrix.get(rowIdx, lineIdx), m.get(rowIdx, lineIdx)) {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (matrix *Matrix) Vector() *Vector {
	if matrix.rowSize == matrixToVectorRowSize ||
		matrix.lineSize == matrixToVectorLineSize {
		return matrix.convertToVector()
	}
	panic(matrixCanNotBeVectorError)
}

// Add
// change self
// chained option
func (matrix *Matrix) Add(m *Matrix) *Matrix {
	return matrix.one2OneOpt('+', m)
}

func (matrix *Matrix) GetPlus(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Add(m)
}

// Sub
// change self
// chained option
func (matrix *Matrix) Sub(m *Matrix) *Matrix {
	return matrix.one2OneOpt('-', m)
}

func (matrix *Matrix) GetMinus(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Sub(m)
}

func (matrix *Matrix) GetRowAsVector(rowIndex int) *Vector {
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return ConstructVector(lang.GetReal2DArrayRow(matrix.GetSlice(), rowIndex), false)
}

func (matrix *Matrix) GetLineAsVector(lineIndex int) *Vector {
	if !matrix.validateIndex(0, lineIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return ConstructVector(lang.GetReal2DArrayLine(matrix.GetSlice(), lineIndex), true)
}

// Mul
// change self
// chained option
func (matrix *Matrix) Mul(m *Matrix) *Matrix {
	if !matrix.canMultiply(m) {
		panic(matrixCanNotMultiplyError)
	}
	newSlice := make([][]float64, matrix.rowSize)
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		newSlice[rowIdx] = make([]float64, m.lineSize)
		for lineIdx := 0; lineIdx < m.lineSize; lineIdx++ {
			newSlice[rowIdx][lineIdx] = matrix.GetRowAsVector(rowIdx).DotMul(m.GetLineAsVector(lineIdx))
		}
	}
	matrix.setValues(newSlice, matrix.rowSize, m.lineSize)
	return matrix
}

func (matrix *Matrix) GetMTimes(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Mul(m)
}

// Div
// change self
// chained option
func (matrix *Matrix) Div(m *Matrix) *Matrix {
	inverseM := m.GetInverse()
	return matrix.Mul(inverseM)
}

func (matrix *Matrix) GetMRDivide(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Div(m)
}

// MulLambda
// change self
// chained option
func (matrix *Matrix) MulLambda(lambda float64) *Matrix {
	for rowIdx, _ := range matrix.GetSlice() {
		for lineIdx, _ := range matrix.getRow(rowIdx) {
			matrix.setOpt(rowIdx, lineIdx, lambda, '*')
		}
	}
	return matrix
}

func (matrix *Matrix) MulVector(v *Vector) *Vector {
	if v != nil &&
		v.shape &&
		matrix.lineSize == v.size {
		vCopy := v.makeCopy()
		for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
			var rowSum float64 = 0.0
			for lineIdx := 0; lineIdx < matrix.lineSize; lineIdx++ {
				rowSum += vCopy.get(lineIdx, lineIdx) * matrix.get(rowIdx, lineIdx)
			}
			v.slice[rowIdx] = rowSum
		}
		return v
	}
	return &Vector{}
}

// Power
// change self
// chained option
func (matrix *Matrix) Power(n int) *Matrix {
	mCopy := matrix.makeCopy()
	for t := 0; t < n; t++ {
		matrix.Mul(mCopy)
	}
	return matrix
}

func (matrix *Matrix) GetMPower(n int) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Power(n)
}

// HadamardMul
// change self
// chained option
func (matrix *Matrix) hadamardMul(m *Matrix) *Matrix {
	return matrix.one2OneOpt('*', m)
}

func (matrix *Matrix) DotMul(m *Matrix) *Matrix {
	return matrix.hadamardMul(m)
}

func (matrix *Matrix) GetTimes(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.DotMul(m)
}

func (matrix *Matrix) DotDivide(m *Matrix) *Matrix {
	return matrix.one2OneOpt('/', m)
}

func (matrix *Matrix) GetRDivide(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.DotDivide(m)
}

func (matrix *Matrix) DotPower(m *Matrix) *Matrix {
	return matrix.one2OneOpt('^', m)
}

func (matrix *Matrix) GetPower(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.DotPower(m)
}

// Display
// chained option
func (matrix *Matrix) Display() *Matrix {
	if matrix.isNull() {
		fmt.Println("[null]")
		return matrix
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for lineIdx := 0; lineIdx < matrix.lineSize; lineIdx++ {
			val := matrix.get(rowIdx, lineIdx)
			if lang.EqualFloat64ByAccuracy(0.0, val) {
				val = 0.0
			}
			fmt.Printf(" %.5v", val)
		}
		fmt.Println()
	}
	return matrix
}

func (matrix *Matrix) DisplayDelimiter() *Matrix {
	fmt.Println("--------------------------")
	return matrix
}

// Convergence
// change self
// chained option
func (matrix *Matrix) Convergence() (*Matrix, int) {
	convergenceMatrix := matrix.makeCopy()
	convergenceIteratorTime := 0
	for {
		prev := convergenceMatrix.makeCopy()
		convergenceIteratorTime++
		convergenceMatrix.Power(1)
		if convergenceMatrix.Equal(prev) {
			break
		}
	}
	return convergenceMatrix, convergenceIteratorTime
}

func (matrix *Matrix) ReShape(rowSize, lineSize int) *Matrix {
	if matrix.rowSize*matrix.lineSize !=
		rowSize*lineSize {
		panic(matrixCanNotBeReShapedError)
	}
	reShapedMatrix := &Matrix{}
	reShapedMatrix.assign(rowSize, lineSize)
	mCnt, mRowIdx, mLineIdx := -1, 0, 0
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for lineIdx := 0; lineIdx < matrix.lineSize; lineIdx++ {
			mCnt++
			mRowIdx = mCnt / lineSize
			mLineIdx = mCnt % lineSize
			reShapedMatrix.set(mRowIdx, mLineIdx, matrix.get(rowIdx, lineIdx))
		}
	}
	matrix.remake(reShapedMatrix)
	return matrix
}

func (matrix *Matrix) GetReShape(rowSize, lineSize int) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.ReShape(rowSize, lineSize)
}

func (matrix *Matrix) generate(value interface{}, size ...int) *Matrix {
	generatedMatrix := &Matrix{}
	gRowSize, gLineSize := matrix.rowSize, matrix.lineSize
	if sizeLen := len(size); sizeLen > 0 {
		if sizeLen >= 1 {
			gRowSize = size[0]
			gLineSize = gRowSize
		}
		if sizeLen >= 2 {
			gLineSize = size[1]
		}
	}
	switch value.(type) {
	case float32, float64:
		{
			generatedMatrix.assign(gRowSize, gLineSize)
			for rowIdx := 0; rowIdx < gRowSize; rowIdx++ {
				for lineIdx := 0; lineIdx < gLineSize; lineIdx++ {
					generatedMatrix.set(rowIdx, lineIdx, value.(float64))
				}
			}
		}
	case [][]float32, [][]float64:
		{
			generatedMatrix.setValues(value.([][]float64), size...)
		}
	default:
		{

		}
	}
	return generatedMatrix
}

func (matrix *Matrix) SetZero(size ...int) *Matrix {
	matrix.setSelf(matrix.GetZero(size...))
	return matrix
}

func (matrix *Matrix) GetZero(size ...int) *Matrix {
	return matrix.generate(0.0, size...)
}

func (matrix *Matrix) SetOne(size ...int) *Matrix {
	matrix.setSelf(matrix.GetOne(size...))
	return matrix
}

func (matrix *Matrix) GetOne(size ...int) *Matrix {
	return matrix.generate(1.0, size...)
}

func (matrix *Matrix) SetEye(size ...int) *Matrix {
	matrix.setSelf(matrix.GetEye(size...))
	return matrix
}

func (matrix *Matrix) GetEye(size ...int) *Matrix {
	if sizeLen := len(size); sizeLen > 0 {
		eyeIdentity := &Identity{size[0]}
		return eyeIdentity.Matrix()
	}
	return &Matrix{}
}

func (matrix *Matrix) IsDiagonal() bool {
	return matrix.isDiagonalMatrix()
}

func (matrix *Matrix) Diagonal() *Diagonal {
	if matrix.IsDiagonal() {
		d := &Diagonal{}
		d.setValues(make([]float64, matrix.getPhalanxSize()), matrix.getPhalanxSize())
		for idx := 0; idx < matrix.rowSize; idx++ {
			d.set(idx, matrix.get(idx, idx))
		}
		return d
	}
	return &Diagonal{}
}

func (matrix *Matrix) CanDiagonalizing() bool {
	return false
}

func (matrix *Matrix) Diagonalizing() *Diagonal {
	return &Diagonal{}
}

func (matrix *Matrix) IsSimilar(m *Matrix) bool {
	return false
}

func (matrix *Matrix) Similar(m *Matrix) *Matrix {
	return m.GetInverse().GetMTimes(matrix).GetMTimes(m)
}

// schmidtOrthogonality
// current version do not support vector expand
func (matrix *Matrix) schmidtOrthogonality() *Matrix {
	if !matrix.isPhalanx() {
		return &Matrix{}
	}
	return matrix.VectorGroup(true).GetSchmidt().Matrix()
}

func (matrix *Matrix) Schmidt() *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.schmidtOrthogonality()
}

func (matrix *Matrix) unitization() *Matrix {
	if !matrix.isPhalanx() {
		return &Matrix{}
	}
	return matrix.VectorGroup(true).GetUnit().Matrix()
}

func (matrix *Matrix) Unit() *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.unitization()
}

func (matrix *Matrix) convertToVectorGroup(shape bool) *VectorGroup {
	vg := &VectorGroup{}
	return vg.Construct(matrix.GetSlice())
}

func (matrix *Matrix) VectorGroup(shape bool) *VectorGroup {
	return matrix.convertToVectorGroup(shape)
}

func (matrix *Matrix) ZeroSpace() {}
func (matrix *Matrix) OrthonormalBasis() *Matrix {
	return matrix.Schmidt().Unit()
}

func (matrix *Matrix) RowSimplest() {}

// ET
// ElementaryTransformation
func (matrix *Matrix) ET() {}

func (matrix *Matrix) RowET() {}

func (matrix *Matrix) LineET() {}

func (matrix *Matrix) rowExchangeET() {}

func (matrix *Matrix) lineExchangeET() {}

func (matrix *Matrix) rowMulLambdaET() {}

func (matrix *Matrix) lineMulLambdaET() {}

func (matrix *Matrix) rowIMulLambdaAddRowJET() {}

func (matrix *Matrix) lineIMulLambdaAddLineJET() {}

// Matrix Decomposition

func (matrix *Matrix) LU()         {}
func (matrix *Matrix) QR()         {}
func (matrix *Matrix) Cholesky()   {}
func (matrix *Matrix) SVD()        {}
func (matrix *Matrix) Schur()      {}
func (matrix *Matrix) Hessenberg() {}

func (matrix *Matrix) norm() float64         { return 0 }
func (matrix *Matrix) spectralNorm() float64 { return 0 }
func (matrix *Matrix) L1Norm() float64       { return 0 }
func (matrix *Matrix) L2Norm() float64 {
	return matrix.spectralNorm()
}
func (matrix *Matrix) InfiniteNorm() float64  { return 0 }
func (matrix *Matrix) FrobeniusNorm() float64 { return 0 }

func (matrix *Matrix) IsPositiveDefine() bool { return false }
func (matrix *Matrix) IsNegativeDefine() bool { return false }

func (matrix *Matrix) IsSymmetric() bool     { return false }
func (matrix *Matrix) IsAntiSymmetric() bool { return false }
