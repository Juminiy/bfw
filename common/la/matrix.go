package la

import (
	"bfw/common/lang"
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	matrixNotPhalanx         int     = -1
	matrixNoSize             int     = 0
	matrixIndexOutOfBound    int     = -1
	matrixDeterminantZero    float64 = 0
	matrixNoRank             int     = -1
	matrixToVectorRowSize    int     = 1
	matrixToVectorColumnSize int     = 1
	simplePhalanxSizeOne     int     = 1
	simplePhalanxSizeTwo     int     = 2
	simplePhalanxSizeThree   int     = 3
)

var (
	matrixNotSameShapeError           = errors.New("two matrices are not same shape")
	matrixCanNotMultiplyError         = errors.New("two matrices cannot multiply")
	matrixCanNotHadamardMultiplyError = errors.New("two matrices cannot hadamard multiply")
	matrixRowColumnDiffer             = errors.New("matrix row and column differ")
	matrixCanNotBeVectorError         = errors.New("matrix cannot be vector")
	matrixIndexOutOfBoundError        = errors.New("matrix index is out of bound")
	matrixNotPhalanxError             = errors.New("matrix is not phalanx")
	matrixCanNotBeAdjoinError         = errors.New("matrix cannot be adjoin")
	matrixCanNotBeInverseError        = errors.New("matrix cannot be inverse")
	matrixCanNotBeIdentityError       = errors.New("matrix cannot be identity")
	matrixCanNotBeReShapedError       = errors.New("matrix cannot be reshaped")
	matrixInValidError                = errors.New("matrix is invalid")
	rangeBoundError                   = errors.New("the num range bound is error")
	NullMatrix                        = &Matrix{}
)

type Matrix struct {
	slice       [][]float64
	rowSize     int
	columnSize  int
	coefficient float64
}

func ConstructMatrix(real2DArray [][]float64) *Matrix {
	matrix := &Matrix{}
	return matrix.Construct(real2DArray)
}

func (matrix *Matrix) Construct(real2DArray [][]float64) *Matrix {
	if real2DArray == nil ||
		len(real2DArray) == matrixNoSize {
		return matrix
	}
	var (
		maxColumnSize int = 1
		rowSize       int = len(real2DArray)
	)
	matrix.setValues(real2DArray, rowSize, maxColumnSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		maxColumnSize = lang.MaxInt(maxColumnSize, len(real2DArray[rowIdx]))
	}
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		matrix.setRowZeroPadding(maxColumnSize, rowIdx)
	}
	matrix.setSize(rowSize, maxColumnSize)
	return matrix
}

func (matrix *Matrix) validate() bool {
	if matrix.rowSize == matrixNoSize ||
		matrix.columnSize == matrixNoSize ||
		matrix.slice == nil ||
		len(matrix.slice) != matrix.rowSize {
		return false
	} else {
		for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
			if idxRow := matrix.getRow(rowIdx); idxRow == nil ||
				len(idxRow) != matrix.columnSize {
				return false
			}
		}
	}
	return true
}

func (matrix *Matrix) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		if indexLen >= 1 {
			if !matrix.validateRowIndex(index[0]) {
				return false
			}
		}
		if indexLen >= 2 {
			if !matrix.validateColumnIndex(index[1]) {
				return false
			}
		}
	}
	return true
}

func (matrix *Matrix) validateRowIndex(rowIndex int) bool {
	return rowIndex >= 0 &&
		rowIndex < matrix.rowSize
}

func (matrix *Matrix) validateColumnIndex(columnIndex int) bool {
	return columnIndex >= 0 &&
		columnIndex < matrix.columnSize
}

func (matrix *Matrix) null() *Matrix {
	return &Matrix{}
}

func (matrix *Matrix) isNull() bool {
	return !matrix.validate()
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

func (matrix *Matrix) makeCopy() *Matrix {
	mCopy := &Matrix{}
	mCopy.setValues(make([][]float64, matrix.rowSize), matrix.rowSize, matrix.columnSize)
	for rowIdx := 0; rowIdx < mCopy.rowSize; rowIdx++ {
		mCopy.setRow(rowIdx, make([]float64, matrix.columnSize))
		copy(mCopy.slice[rowIdx], matrix.getRow(rowIdx))
	}
	return mCopy
}

// assign
// 1. character
// change self
// 2.function
// initial the location
func (matrix *Matrix) assign(rowSize, columnSize int) {
	matrix.setValues(make([][]float64, rowSize), rowSize, columnSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		matrix.setRow(rowIdx, make([]float64, columnSize))
	}
}

func (matrix *Matrix) getSelf() *Matrix {
	return matrix
}

func (matrix *Matrix) setSelf(m *Matrix) {
	matrix.setValues(m.slice, m.rowSize, m.columnSize)
}

func (matrix *Matrix) setValues(slice [][]float64, size ...int) {
	matrix.setSlice(slice)
	matrix.setSize(size...)
}

// setSlice
// change self
func (matrix *Matrix) setSlice(slice [][]float64) {
	matrix.slice = slice
}

// setSize
// change self
func (matrix *Matrix) setSize(size ...int) {
	if sizeLen := len(size); sizeLen > 0 {
		if sizeLen >= 1 {
			matrix.setRowSize(size[0])
			matrix.setColumnSize(size[0])
		}
		if sizeLen >= 2 {
			matrix.setColumnSize(size[1])
		}
	}
}

func (matrix *Matrix) setRowSize(rowSize int) {
	matrix.rowSize = rowSize
}

func (matrix *Matrix) setColumnSize(columnSize int) {
	matrix.columnSize = columnSize
}

func (matrix *Matrix) swap(m *Matrix) {
	mTemp := &Matrix{}
	mTemp.setSelf(matrix)
	matrix.setSelf(m)
	m.setSelf(mTemp)
}

func (matrix *Matrix) get(rowIndex, columnIndex int) float64 {
	if !matrix.validateIndex(rowIndex, columnIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return matrix.slice[rowIndex][columnIndex]
}

// set
// change self
func (matrix *Matrix) set(rowIndex, columnIndex int, value float64) {
	matrix.setOpt(rowIndex, columnIndex, value, '=')
}

func (matrix *Matrix) setOpt(rowIndex, columnIndex int, value float64, opt ...rune) {
	if !matrix.validateIndex(rowIndex, columnIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	if len(opt) > 0 {
		switch opt[0] {
		case '=':
			{
				matrix.slice[rowIndex][columnIndex] = value
			}
		case '+':
			{
				matrix.slice[rowIndex][columnIndex] += value
			}
		case '-':
			{
				matrix.slice[rowIndex][columnIndex] -= value
			}
		case '*':
			{
				matrix.slice[rowIndex][columnIndex] *= value
			}
		case '/':
			{
				matrix.slice[rowIndex][columnIndex] /= value
			}
		case '^':
			{
				matrix.slice[rowIndex][columnIndex] = math.Pow(
					matrix.get(rowIndex, columnIndex), value)
			}
		default:
			{

			}
		}
	}
}

func (matrix *Matrix) getRowCopy(rowIndex int) []float64 {
	indexRow := matrix.getRow(rowIndex)
	sliceCopy := make([]float64, len(indexRow))
	copy(sliceCopy, indexRow)
	return sliceCopy
}

func (matrix *Matrix) getRow(rowIndex int) []float64 {
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return matrix.slice[rowIndex]
}

func (matrix *Matrix) setRow(rowIndex int, rowSlice []float64) {
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	matrix.slice[rowIndex] = rowSlice
}

func (matrix *Matrix) getColumnCopy(columnIndex int) []float64 {
	return matrix.getColumn(columnIndex)
}

func (matrix *Matrix) getColumn(columnIndex int) []float64 {
	if !matrix.validateColumnIndex(columnIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return lang.GetReal2DArrayColumn(matrix.GetSlice(), columnIndex)
}

func (matrix *Matrix) setColumn(columnIndex int, columnSlice []float64) {
	if !matrix.validateColumnIndex(columnIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		matrix.set(rowIdx, columnIndex, columnSlice[rowIdx])
	}
}

func (matrix *Matrix) setRowSwap(rowIIndex, rowJIndex int) {
	if !matrix.validateIndex(rowIIndex) ||
		!matrix.validateIndex(rowJIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	matrix.slice[rowIIndex], matrix.slice[rowJIndex] =
		matrix.slice[rowJIndex], matrix.slice[rowIIndex]
}

func (matrix *Matrix) setColumnSwap(columnIIndex, columnJIndex int) {
	if !matrix.validateColumnIndex(columnIIndex) ||
		!matrix.validateColumnIndex(columnJIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	columnICopy := matrix.getColumnCopy(columnIIndex)
	matrix.setColumn(columnIIndex, matrix.getColumn(columnJIndex))
	matrix.setColumn(columnJIndex, columnICopy)
}

func (matrix *Matrix) setRowElemOpt(rowIndex int, opt rune, value float64) {
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
		matrix.setOpt(rowIndex, columnIdx, value, opt)
	}
}

func (matrix *Matrix) setRowElemSwap(rowIndex int) {
	for columnIdx := 0; columnIdx < (matrix.columnSize >> 1); columnIdx++ {
		matrix.setElemSwap(rowIndex, columnIdx, rowIndex, matrix.columnSize-1-columnIdx)
	}
}

func (matrix *Matrix) setColumnElemOpt(columnIndex int, opt rune, value float64) {
	if !matrix.validateColumnIndex(columnIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		matrix.setOpt(rowIdx, columnIndex, value, opt)
	}
}

func (matrix *Matrix) setColumnElemSwap(columnIndex int) {
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		matrix.setElemSwap(rowIdx, columnIndex, matrix.rowSize-1-rowIdx, columnIndex)
	}
}

func (matrix *Matrix) getRowIndexLen(rowIndex int) int {
	return len(matrix.getRow(rowIndex))
}

func (matrix *Matrix) setElemSwap(rowIndexI, columnIndexI, rowIndexJ, columnIndexJ int) {
	matrix.slice[rowIndexI][columnIndexI], matrix.slice[rowIndexJ][columnIndexJ] =
		matrix.slice[rowIndexJ][columnIndexJ], matrix.slice[rowIndexI][columnIndexI]
}

func (matrix *Matrix) rowPadding(rowIndex, columnSize int, value float64) {
	if columnSize <= matrix.columnSize {
		return
	}
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	for columnIdx := matrix.getRowIndexLen(rowIndex); columnIdx < columnSize; columnIdx++ {
		matrix.slice[rowIndex] = append(matrix.slice[rowIndex], value)
	}
}

func (matrix *Matrix) rowZeroPadding(rowIndex, columnSize int) {
	matrix.rowPadding(rowIndex, columnSize, 0.0)
}

func (matrix *Matrix) setRowZeroPadding(columnSize int, rowIndex ...int) {
	for rowIndexIdx := 0; rowIndexIdx < len(rowIndex); rowIndexIdx++ {
		matrix.rowZeroPadding(rowIndex[rowIndexIdx], columnSize)
	}
}

func (matrix *Matrix) sameShape(m *Matrix) bool {
	if m == nil ||
		matrix.rowSize != m.rowSize ||
		matrix.columnSize != m.columnSize {
		return false
	}
	return true
}

func (matrix *Matrix) Equal(m *Matrix) bool {
	if matrix.sameShape(m) {
		for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
			for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
				if !lang.EqualFloat64ByAccuracy(matrix.get(rowIdx, columnIdx), m.get(rowIdx, columnIdx)) {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (matrix *Matrix) canMultiply(m *Matrix) bool {
	if m == nil ||
		matrix.columnSize != m.rowSize {
		return false
	}
	return true
}

func (matrix *Matrix) isPhalanx() bool {
	return matrix.rowSize == matrix.columnSize
}

func (matrix *Matrix) getPhalanxSize() int {
	if matrix.isPhalanx() {
		return matrix.rowSize
	}
	return matrixNotPhalanx
}

// getPhalanxByIgnoreCentral
// chained option
func (matrix *Matrix) getPhalanxByIgnoreCentral(rowIndex, columnIndex int) *Matrix {
	phalanxSize := matrix.getPhalanxSize()
	remainedPhalanxSize := phalanxSize - 1
	remainedCnt, remainedRowIdx, remainedColumnIdx := -1, 0, 0
	remainedPhalanx := &Matrix{}
	remainedPhalanx.assign(remainedPhalanxSize, remainedPhalanxSize)
	for rowIdx := 0; rowIdx < phalanxSize; rowIdx++ {
		for columnIdx := 0; columnIdx < phalanxSize; columnIdx++ {
			if rowIdx != rowIndex && columnIdx != columnIndex {
				remainedCnt++
				remainedRowIdx = remainedCnt / remainedPhalanxSize
				remainedColumnIdx = remainedCnt % remainedPhalanxSize
				remainedPhalanx.set(remainedRowIdx, remainedColumnIdx, matrix.get(rowIdx, columnIdx))
			}
		}
	}
	return remainedPhalanx
}

// getPhalanxByIgnoreKCentral
// chained option
func (matrix *Matrix) getPhalanxByIgnoreKCentral(rowIndexMap, columnIndexMap map[int]bool) (*Matrix, *Matrix) {
	if rowIndexMap == nil || columnIndexMap == nil {
		return &Matrix{}, &Matrix{}
	}
	phalanxSize := matrix.getPhalanxSize()
	rowsK, columnsK := len(rowIndexMap), len(columnIndexMap)
	if rowsK != columnsK ||
		phalanxSize < rowsK {
		return &Matrix{}, &Matrix{}
	}
	selected, remained := &Matrix{}, &Matrix{}
	selectedSize, remainedSize := rowsK, phalanxSize-rowsK
	selected.assign(selectedSize, selectedSize)
	remained.assign(remainedSize, remainedSize)
	slCnt, slRowIdx, slColumnIdx := -1, 0, 0
	rmCnt, rmRowIdx, rmColumnIdx := -1, 0, 0
	for rowIdx := 0; rowIdx < phalanxSize; rowIdx++ {
		for columnIdx := 0; columnIdx < phalanxSize; columnIdx++ {
			if !rowIndexMap[rowIdx] && !columnIndexMap[columnIdx] {
				rmCnt++
				rmRowIdx = rmCnt / remainedSize
				rmColumnIdx = rmCnt % remainedSize
				remained.set(rmRowIdx, rmColumnIdx, matrix.get(rowIdx, columnIdx))
			} else if rowIndexMap[rowIdx] && columnIndexMap[columnIdx] {
				slCnt++
				slRowIdx = slCnt / selectedSize
				slColumnIdx = slCnt % selectedSize
				selected.set(slRowIdx, slColumnIdx, matrix.get(rowIdx, columnIdx))
			}
		}
	}
	return selected, remained
}

func (matrix *Matrix) getPhalanxMinusOnePower(rowIndexMap, columnIndexMap map[int]bool) int {
	totalPower := 0
	for row, _ := range rowIndexMap {
		totalPower += row
	}
	for column, _ := range columnIndexMap {
		totalPower += column
	}
	return totalPower
}

func (matrix *Matrix) GetSlice() [][]float64 {
	if !matrix.validate() {
		panic(matrixInValidError)
	}
	return matrix.slice
}

func (matrix *Matrix) GetRowSize() int {
	return matrix.rowSize
}

func (matrix *Matrix) GetColumnSize() int {
	return matrix.columnSize
}

func (matrix *Matrix) GetElemSize() int {
	return matrix.rowSize * matrix.columnSize
}

// one2OneOpt
// change self
// chained option
func (matrix *Matrix) one2OneOpt(opt rune, m *Matrix) *Matrix {
	if !matrix.sameShape(m) {
		panic(matrixNotSameShapeError)
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
			switch opt {
			case '+':
				{
					matrix.setOpt(rowIdx, columnIdx, m.get(rowIdx, columnIdx), '+')
				}
			case '-':
				{
					matrix.setOpt(rowIdx, columnIdx, m.get(rowIdx, columnIdx), '-')
				}
			case '*':
				{
					matrix.setOpt(rowIdx, columnIdx, m.get(rowIdx, columnIdx), '*')
				}
			case '/':
				{
					matrix.setOpt(rowIdx, columnIdx, m.get(rowIdx, columnIdx), '/')
				}
			case '^':
				{
					matrix.setOpt(rowIdx, columnIdx, m.get(rowIdx, columnIdx), '^')
				}
			default:
				{

				}
			}
		}
	}
	return matrix
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
		panic(matrixRowColumnDiffer)
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
	for columnIdx := 0; columnIdx < totalN; columnIdx++ {
		sumByRandomRow += matrix.getPhalanxByIgnoreCentral(randomRow, columnIdx).det() *
			matrix.get(randomRow, columnIdx) *
			float64(lang.MinusOnePower(randomRow+columnIdx))
	}
	return sumByRandomRow
}

func (matrix *Matrix) laplaceDet(totalN int) float64 {
	var (
		randomKRow      int            = matrix.getPhalanxSize() / 2
		randomKRows     map[int]bool   = lang.GetRandomMapValue(totalN, randomKRow)
		kColumnsMap     []map[int]bool = lang.GetCombinationSliceMap(totalN, randomKRow)
		sumByRandomKRow float64        = 0.0
	)
	for _, columnsMap := range kColumnsMap {
		sel, rem := matrix.getPhalanxByIgnoreKCentral(randomKRows, columnsMap)
		power := matrix.getPhalanxMinusOnePower(randomKRows, columnsMap)
		sumByRandomKRow += sel.det() * rem.det() * float64(lang.MinusOnePower(power))
	}
	return sumByRandomKRow
}

// Matrix Opt Matrix self

// GetTranspose
// chained option
func (matrix *Matrix) GetTranspose() *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.transpose()
}

// transpose
// change self
// chained option
// update version: the space complexity is: O(1), impossible
func (matrix *Matrix) transpose() *Matrix {
	newRowSize, newColumnSize := matrix.columnSize, matrix.rowSize
	newSlice := make([][]float64, newRowSize)
	for newRowIdx := 0; newRowIdx < newRowSize; newRowIdx++ {
		newSlice[newRowIdx] = make([]float64, newColumnSize)
		for newColumnIdx := 0; newColumnIdx < newColumnSize; newColumnIdx++ {
			newSlice[newRowIdx][newColumnIdx] = matrix.get(newColumnIdx, newRowIdx)
		}
	}
	matrix.setValues(newSlice, newRowSize, newColumnSize)
	return matrix
}

// GetInverse
// chained option
func (matrix *Matrix) GetInverse() *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.inverse()
}

// Inverse Matrix by Adjoin Matrix
// change self
// chained option
// |M| != 0, M^{-1} = M^{*}/|M|
func (matrix *Matrix) inverse() *Matrix {
	if !matrix.isPhalanx() {
		panic(matrixRowColumnDiffer)
	}
	det := matrix.det()
	if lang.EqualFloat64ByAccuracy(det, matrixDeterminantZero) {
		panic(matrixCanNotBeInverseError)
	}
	matrix.adjoin()
	matrix.MulLambda(1 / det)
	return matrix
}

// GetAdjoin
// chained option
func (matrix *Matrix) GetAdjoin() *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.adjoin()
}

// Adjoin Matrix by remained Matrix
// change self
// chained option
func (matrix *Matrix) adjoin() *Matrix {
	if !matrix.isPhalanx() {
		panic(matrixRowColumnDiffer)
	}
	size := matrix.getPhalanxSize()
	adjoinMatrix := &Matrix{}
	adjoinMatrix.assign(size, size)
	// 1. each element calculate remainder value by (getPhalanxByIgnoreCentral)
	// 2. each element calculate algebraic remainder value by (MinusOnePower)
	for rowIdx := 0; rowIdx < size; rowIdx++ {
		for columnIdx := 0; columnIdx < size; columnIdx++ {
			algebraicRemainderValue := matrix.getPhalanxByIgnoreCentral(rowIdx, columnIdx).det() *
				float64(lang.MinusOnePower(rowIdx+columnIdx))
			adjoinMatrix.set(rowIdx, columnIdx, algebraicRemainderValue)
		}
	}
	// 3. matrix transpose
	adjoinMatrix.transpose()
	matrix.setSlice(adjoinMatrix.slice)
	return matrix
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

func (matrix *Matrix) Sum() float64 {
	return matrix.Accu()
}

func (matrix *Matrix) Accu() float64 {
	return matrix.accumulate()
}

func (matrix *Matrix) accumulate() float64 {
	var (
		accuRes float64 = 0.0
	)
	matrix.traverse(func(elemValue float64) {
		accuRes += elemValue
	})
	return accuRes
}

func (matrix *Matrix) All(conditionFunc func(float64) bool) bool {
	var (
		allFlag = true
	)
	matrix.traverse(func(elemValue float64) {
		if !conditionFunc(elemValue) {
			allFlag = false
		}
	})
	return allFlag
}

func (matrix *Matrix) Any(conditionFunc func(float64) bool) bool {
	var (
		anyFlag = false
	)
	matrix.traverse(func(elemValue float64) {
		if conditionFunc(elemValue) {
			anyFlag = true
		}
	})
	return anyFlag
}

func (matrix *Matrix) Find(conditionFunc func(float64) bool) []float64 {
	var (
		conditionValues []float64 = make([]float64, 0)
	)
	matrix.traverse(func(elemValue float64) {
		if conditionFunc(elemValue) {
			conditionValues = append(conditionValues, elemValue)
		}
	})
	return conditionValues
}

func (matrix *Matrix) FindOne(conditionFunc func(float64) bool) float64 {
	return matrix.Find(conditionFunc)[0]
}

func (matrix *Matrix) Approx(accBits ...int) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.approximation(accBits...)
}

func (matrix *Matrix) approximation(accAfterDotBits ...int) *Matrix {
	accBits := 5
	if len(accAfterDotBits) > 0 {
		accBits = accAfterDotBits[0]
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
			newVal := 0.0
			valStr := fmt.Sprintf("%."+strconv.Itoa(accBits)+"f", matrix.get(rowIdx, columnIdx))
			valF, err := strconv.ParseFloat(valStr, 64)
			if err != nil {
				panic(err)
			}
			if lang.IsStringValueIntZero(valStr) {
				newVal = 0
			} else {
				newVal = valF
			}
			matrix.set(rowIdx, columnIdx, newVal)
		}
	}
	return matrix
}

func (matrix *Matrix) ReShape(rowSize, columnSize int) *Matrix {
	if matrix.rowSize*matrix.columnSize !=
		rowSize*columnSize {
		panic(matrixCanNotBeReShapedError)
	}
	reShapedMatrix := &Matrix{}
	reShapedMatrix.assign(rowSize, columnSize)
	mCnt, mRowIdx, mColumnIdx := -1, 0, 0
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
			mCnt++
			mRowIdx = mCnt / columnSize
			mColumnIdx = mCnt % columnSize
			reShapedMatrix.set(mRowIdx, mColumnIdx, matrix.get(rowIdx, columnIdx))
		}
	}
	matrix.remake(reShapedMatrix)
	return matrix
}

func (matrix *Matrix) GetReShape(rowSize, columnSize int) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.ReShape(rowSize, columnSize)
}

func (matrix *Matrix) GetMirror() *Matrix {
	return matrix
}

func (matrix *Matrix) mirror() *Matrix {
	return matrix
}

func (matrix *Matrix) GetFlip() *Matrix {
	m := matrix.makeCopy()
	return m.flip()
}

func (matrix *Matrix) flip() *Matrix {
	size := matrix.getPhalanxSize()
	for rowIdx := 0; rowIdx < (size >> 1); rowIdx++ {
		for columnIdx := 0; columnIdx < size; columnIdx++ {
			matrix.setElemSwap(rowIdx, columnIdx, size-1-rowIdx, size-1-columnIdx)
		}
	}
	if lang.Odd(size) {
		matrix.setRowElemSwap(size >> 1)
	}
	return matrix
}

func (matrix *Matrix) GetRotate(clockWise bool, rotateCount int) *Matrix {
	return matrix.rotate(clockWise, rotateCount)
}

func (matrix *Matrix) rotate(clockWise bool, rotateCount int) *Matrix {
	rotateCount %= 4
	switch rotateCount {
	case 1:
		{
			matrix.rotate90(clockWise)
		}
	case 2:
		{
			matrix.flip()
		}
	case 3:
		{
			matrix.rotate90(!clockWise)
		}
	default:
		{
			// do nothing
		}
	}
	return matrix
}

func (matrix *Matrix) rotate90(clockWise bool) *Matrix {
	matrix.flipByMainDiagonal()
	if clockWise {
		matrix.flipByMiddleColumn()
	} else {
		matrix.flipByMiddleRow()
	}
	return matrix
}

func (matrix *Matrix) flipByMiddleRow() *Matrix {
	size := matrix.getPhalanxSize()
	for rowIdx := 0; rowIdx < (size >> 1); rowIdx++ {
		matrix.setRowSwap(rowIdx, size-1-rowIdx)
	}
	return matrix
}

func (matrix *Matrix) flipByMiddleColumn() *Matrix {
	size := matrix.getPhalanxSize()
	for rowIdx := 0; rowIdx < (size >> 1); rowIdx++ {
		matrix.setColumnSwap(rowIdx, size-1-rowIdx)
	}
	return matrix
}

func (matrix *Matrix) flipByMainDiagonal() *Matrix {
	size := matrix.getPhalanxSize()
	for rowIdx := 0; rowIdx < size; rowIdx++ {
		for columnIdx := 0; columnIdx < rowIdx; columnIdx++ {
			matrix.setElemSwap(rowIdx, columnIdx, columnIdx, rowIdx)
		}
	}
	return matrix
}

func (matrix *Matrix) flipBySubDiagonal() *Matrix {
	size := matrix.getPhalanxSize()
	for rowIdx := 0; rowIdx < size; rowIdx++ {
		for columnIdx := rowIdx; columnIdx < size; columnIdx++ {
			matrix.setElemSwap(rowIdx, columnIdx, columnIdx, rowIdx)
		}
	}
	return matrix
}

func (matrix *Matrix) IsX() bool {
	size := matrix.getPhalanxSize()
	flag := true
	matrix.traverseV2(func(rowIdx int, columnIdx int, elemValue ...float64) {
		if rowIdx == columnIdx || rowIdx+columnIdx == size-1 {
			if lang.EqualFloat64Zero(elemValue[0]) {
				flag = false
			}
		} else {
			if !lang.EqualFloat64Zero(elemValue[0]) {
				flag = false
			}
		}
	})
	return flag
}

// Matrix Opt Matrix Res Matrix

func (matrix *Matrix) GetPlus(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Add(m)
}

// Add
// change self
// chained option
func (matrix *Matrix) Add(m *Matrix) *Matrix {
	return matrix.one2OneOpt('+', m)
}

func (matrix *Matrix) GetMinus(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Sub(m)
}

// Sub
// change self
// chained option
func (matrix *Matrix) Sub(m *Matrix) *Matrix {
	return matrix.one2OneOpt('-', m)
}

func (matrix *Matrix) GetMTimes(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Mul(m)
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
		newSlice[rowIdx] = make([]float64, m.columnSize)
		for columnIdx := 0; columnIdx < m.columnSize; columnIdx++ {
			newSlice[rowIdx][columnIdx] = matrix.GetRowAsVector(rowIdx).DotMul(m.GetColumnAsVector(columnIdx))
		}
	}
	matrix.setValues(newSlice, matrix.rowSize, m.columnSize)
	return matrix
}

func (matrix *Matrix) GetMRDivide(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Div(m)
}

// Div
// change self
// chained option
func (matrix *Matrix) Div(m *Matrix) *Matrix {
	inverseM := m.GetInverse()
	return matrix.Mul(inverseM)
}

func (matrix *Matrix) GetMPower(n int) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.Power(n)
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

// Matrix Opt Matrix Res Matrix

// MulLambda
// change self
// chained option
func (matrix *Matrix) MulLambda(lambda float64) *Matrix {
	for rowIdx, _ := range matrix.GetSlice() {
		for columnIdx, _ := range matrix.getRow(rowIdx) {
			matrix.setOpt(rowIdx, columnIdx, lambda, '*')
		}
	}
	return matrix
}

func (matrix *Matrix) MulVector(v *Vector) *Vector {
	if v != nil &&
		v.shape &&
		matrix.columnSize == v.size {
		vCopy := v.makeCopy()
		for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
			var rowSum float64 = 0.0
			for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
				rowSum += vCopy.get(columnIdx, columnIdx) * matrix.get(rowIdx, columnIdx)
			}
			v.slice[rowIdx] = rowSum
		}
		return v
	}
	return &Vector{}
}

func (matrix *Matrix) GetTimes(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.DotMul(m)
}

func (matrix *Matrix) DotMul(m *Matrix) *Matrix {
	return matrix.hadamardMul(m)
}

// HadamardMul
// change self
// chained option
func (matrix *Matrix) hadamardMul(m *Matrix) *Matrix {
	return matrix.one2OneOpt('*', m)
}

func (matrix *Matrix) GetRDivide(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.DotDivide(m)
}

func (matrix *Matrix) DotDivide(m *Matrix) *Matrix {
	return matrix.one2OneOpt('/', m)
}

func (matrix *Matrix) GetPower(m *Matrix) *Matrix {
	mCopy := matrix.makeCopy()
	return mCopy.DotPower(m)
}

func (matrix *Matrix) DotPower(m *Matrix) *Matrix {
	return matrix.one2OneOpt('^', m)
}

func (matrix *Matrix) Convolution(m *Matrix) float64 {
	if !matrix.sameShape(m) {
		panic(matrixNotSameShapeError)
	}
	return matrix.GetTimes(m.GetFlip()).Sum()
}

func (matrix *Matrix) Rand() *Matrix {
	return matrix.SetRandom()
}

func (matrix *Matrix) SetRandom() *Matrix {
	matrix.setSelf(matrix.generateRandomFloat64(matrix.rowSize, matrix.columnSize, 0.0, 1.0))
	return matrix
}

// GenMatrix
// Matrix generate Matrix
// len(dataRange) = 0, [0.0,1.0)
// len(dataRange) = 1, [0.0,dataRange[0])
// len(dataRange) = 2, [dataRange[0],dataRange[1])
func GenMatrix(rowSize, columnSize int, dataType string, dataRange ...float64) *Matrix {
	var (
		rangeStart   float64 = 0.0
		rangeEnd     float64 = 1.0
		puppetMatrix         = &Matrix{}
	)
	if rangeLen := len(dataRange); rangeLen > 0 {
		if rangeLen == 1 {
			rangeEnd = dataRange[0]
		} else if rangeLen == 2 {
			rangeStart = dataRange[0]
			rangeEnd = dataRange[1]
		}
	}
	switch dataType {
	case "i", "int":
		{
			return puppetMatrix.generateRandomInt(rowSize, columnSize, int(rangeStart), int(rangeEnd))
		}
	case "i32", "int32":
		{
			return puppetMatrix.generateRandomInt32(rowSize, columnSize, int32(rangeStart), int32(rangeEnd))
		}
	case "i64", "int64":
		{
			return puppetMatrix.generateRandomInt64(rowSize, columnSize, int64(rangeStart), int64(rangeEnd))
		}
	case "f", "f64", "float64":
		{
			return puppetMatrix.generateRandomFloat64(rowSize, columnSize, rangeStart, rangeEnd)
		}
	default:
		{

		}
	}
	return &Matrix{}
}

func (matrix *Matrix) generateMatrix() *Matrix {
	return matrix
}

func (matrix *Matrix) generateRandomInt(rowSize, columnSize int, a, b int) *Matrix {
	if a > b {
		panic(rangeBoundError)
	}
	mt := &Matrix{}
	mt.assign(rowSize, columnSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			mt.set(rowIdx, columnIdx, lang.GetRandFloat64ByIntRange(a, b))
		}
	}
	return mt
}

func (matrix *Matrix) generateRandomInt32(rowSize, columnSize int, a, b int32) *Matrix {
	if a > b {
		panic(rangeBoundError)
	}
	mt := &Matrix{}
	mt.assign(rowSize, columnSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			mt.set(rowIdx, columnIdx, lang.GetRandFloat64ByInt32Range(a, b))
		}
	}
	return mt
}

func (matrix *Matrix) generateRandomInt64(rowSize, columnSize int, a, b int64) *Matrix {
	if a > b {
		panic(rangeBoundError)
	}
	mt := &Matrix{}
	mt.assign(rowSize, columnSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			mt.set(rowIdx, columnIdx, lang.GetRandFloat64ByInt64Range(a, b))
		}
	}
	return mt
}

func (matrix *Matrix) generateRandomFloat64(rowSize, columnSize int, a, b float64) *Matrix {
	if a > b {
		panic(rangeBoundError)
	}
	mt := &Matrix{}
	mt.assign(rowSize, columnSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < columnSize; columnIdx++ {
			mt.set(rowIdx, columnIdx, lang.GetRandFloat64ByFloat64Range(a, b))
		}
	}
	return mt
}

func (matrix *Matrix) generateSameSizeMatrix(value interface{}, size ...int) *Matrix {
	generatedMatrix := &Matrix{}
	gRowSize, gColumnSize := matrix.rowSize, matrix.columnSize
	if sizeLen := len(size); sizeLen > 0 {
		if sizeLen >= 1 {
			gRowSize = size[0]
			gColumnSize = gRowSize
		}
		if sizeLen >= 2 {
			gColumnSize = size[1]
		}
	}
	switch value.(type) {
	case float32, float64:
		{
			generatedMatrix.assign(gRowSize, gColumnSize)
			for rowIdx := 0; rowIdx < gRowSize; rowIdx++ {
				for columnIdx := 0; columnIdx < gColumnSize; columnIdx++ {
					generatedMatrix.set(rowIdx, columnIdx, value.(float64))
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

func (matrix *Matrix) Zeros() *Matrix {
	return matrix.SetZero()
}

func (matrix *Matrix) SetZero(size ...int) *Matrix {
	matrix.setSelf(matrix.GetZero(size...))
	return matrix
}

func (matrix *Matrix) GetZero(size ...int) *Matrix {
	return matrix.generateSameSizeMatrix(0.0, size...)
}

func (matrix *Matrix) Ones() *Matrix {
	return matrix.SetOne()
}

func (matrix *Matrix) SetOne(size ...int) *Matrix {
	matrix.setSelf(matrix.GetOne(size...))
	return matrix
}

func (matrix *Matrix) GetOne(size ...int) *Matrix {
	return matrix.generateSameSizeMatrix(1.0, size...)
}

func (matrix *Matrix) SetEye(size ...int) *Matrix {
	matrix.setSelf(matrix.GetEye(size...))
	return matrix
}

func (matrix *Matrix) GetEye(size ...int) *Matrix {
	if sizeLen := len(size); sizeLen > 0 {
		eyeIdentity := ConstructIdentity(size[0])
		return eyeIdentity.Matrix()
	}
	return &Matrix{}
}

// Matrix Similarity Theory

func (matrix *Matrix) CanDiagonalizing() bool {
	return false
}

func (matrix *Matrix) Diagonalizing() *Diagonal {
	return &Diagonal{}
}

func (matrix *Matrix) IsSimilar(m *Matrix) bool {
	return false
}

func (matrix *Matrix) SimilarityTransformation(m *Matrix) *Matrix {
	return m.GetInverse().GetMTimes(matrix).GetMTimes(m)
}

func (matrix *Matrix) OrthonormalBasis() *Matrix {
	return matrix.Schmidt().Unit()
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

func (matrix *Matrix) EigenValues() *EigenValues {
	if !matrix.isPhalanx() {
		panic(matrixRowColumnDiffer)
	}
	eigenMatrix := matrix.EigenMatrix()
	eigenPoly := eigenMatrix.Det()
	solution := eigenPoly.Solve()
	return solution.EigenValues()
}

// TODO: 1.
func (matrix *Matrix) EigenVectors() *EigenVectors {
	if !matrix.isPhalanx() {
		panic(matrixRowColumnDiffer)
	}
	return &EigenVectors{}
}

// EigenMatrix
// λI-A <-> A-λI
func (matrix *Matrix) EigenMatrix() *PolyMatrix {
	phalanxSize := matrix.getPhalanxSize()
	polyMatrix := matrix.PolyMatrix()
	for phaIdx := 0; phaIdx < phalanxSize; phaIdx++ {
		lambdaPoly := ConstructPolyNode(1.0, 1).Poly(eigenPolyMatrixDefaultAES)
		polyMatrix.setElemByOptElem(phaIdx, phaIdx, '-', lambdaPoly)
	}
	return polyMatrix
}

func (matrix *Matrix) SmithStandard() *PolyDiagonal {
	return &PolyDiagonal{}
}

func (matrix *Matrix) JordanStandard() *JordanMatrix {
	return &JordanMatrix{}
}

func (matrix *Matrix) DdLambda(k int) *Poly {
	return matrix.constantFactors()[k]
}

func (matrix *Matrix) constantFactors() []*Poly {
	return nil
}

func (matrix *Matrix) DDLambda(k int) *Poly {
	return matrix.determinantFactors()[k]
}

func (matrix *Matrix) determinantFactors() []*Poly {
	return nil
}

func (matrix *Matrix) ZeroSpace() {}

// RowSimplest
// Matrix ElementaryTransformation
func (matrix *Matrix) RowSimplest() *Matrix {
	return matrix
}

func (matrix *Matrix) rowExchangeET(rowIIndex, rowJIndex int) *Matrix {
	matrix.setRowSwap(rowIIndex, rowJIndex)
	return matrix
}

func (matrix *Matrix) columnExchangeET(columnIIndex, columnJIndex int) *Matrix {
	matrix.setColumnSwap(columnIIndex, columnJIndex)
	return matrix
}

func (matrix *Matrix) rowMulLambdaET(rowIndex int, lambda float64) *Matrix {
	matrix.setRowElemOpt(rowIndex, '*', lambda)
	return matrix
}

func (matrix *Matrix) columnMulLambdaET(columnIndex int, lambda float64) *Matrix {
	matrix.setColumnElemOpt(columnIndex, '*', lambda)
	return matrix
}

func (matrix *Matrix) rowIMulLambdaAddRowJET(rowIIndex, rowJIndex int, lambda float64) *Matrix {
	if !matrix.validateIndex(rowIIndex) ||
		!matrix.validateIndex(rowJIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
		matrix.setOpt(rowJIndex, columnIdx, matrix.get(rowIIndex, columnIdx)*lambda, '+')
	}
	return matrix
}

func (matrix *Matrix) columnIMulLambdaAddColumnJET(columnIIndex, columnJIndex int, lambda float64) *Matrix {
	if !matrix.validateColumnIndex(columnIIndex) ||
		!matrix.validateColumnIndex(columnJIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		matrix.setOpt(rowIdx, columnJIndex, matrix.get(rowIdx, columnIIndex)*lambda, '+')
	}
	return matrix
}

// Matrix Decomposition

func (matrix *Matrix) LU()         {}
func (matrix *Matrix) QR()         {}
func (matrix *Matrix) Cholesky()   {}
func (matrix *Matrix) SVD()        {}
func (matrix *Matrix) Schur()      {}
func (matrix *Matrix) Hessenberg() {}

// Matrix Norm

func (matrix *Matrix) norm() float64          { return 0 }
func (matrix *Matrix) spectralNorm() float64  { return 0 }
func (matrix *Matrix) L1Norm() float64        { return 0 }
func (matrix *Matrix) L2Norm() float64        { return matrix.spectralNorm() }
func (matrix *Matrix) InfiniteNorm() float64  { return 0 }
func (matrix *Matrix) FrobeniusNorm() float64 { return 0 }

// Matrix Positive Define

func (matrix *Matrix) IsPositiveDefine() bool { return false }
func (matrix *Matrix) IsNegativeDefine() bool { return false }

// Matrix Symmetric

func (matrix *Matrix) IsSymmetric() bool {
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := rowIdx + 1; columnIdx < matrix.columnSize; columnIdx++ {
			if !lang.EqualFloat64ByAccuracy(matrix.get(rowIdx, columnIdx), matrix.get(columnIdx, rowIdx)) {
				return false
			}
		}
	}
	return true
}
func (matrix *Matrix) IsAntiSymmetric() bool {
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := rowIdx + 1; columnIdx < matrix.columnSize; columnIdx++ {
			if !lang.EqualFloat64Zero(matrix.get(rowIdx, columnIdx) + matrix.get(columnIdx, rowIdx)) {
				return false
			}
		}
	}
	return true
}

// AI algorithm

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

// Matrix to other shape

func (matrix *Matrix) IsIdentity() bool {
	return matrix.isDiagonalMatrix(1.0)
}

func (matrix *Matrix) Identity() *Identity {
	if matrix.IsIdentity() {
		return ConstructIdentity(matrix.getPhalanxSize())
	}
	panic(matrixCanNotBeIdentityError)
}

func (matrix *Matrix) GetSameSizeIdentity() *Identity {
	if size := matrix.getPhalanxSize(); size != matrixNoSize {
		return ConstructIdentity(size)
	}
	return &Identity{}
}

func (matrix *Matrix) IsDiagonal() bool {
	return matrix.isDiagonalMatrix()
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
			for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
				val := matrix.get(rowIdx, columnIdx)
				if rowIdx == columnIdx {
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
					if !lang.EqualFloat64Zero(val) {
						return false
					}
				}
			}
		}
		return true
	}
	return false
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

func (matrix *Matrix) Vector() *Vector {
	if matrix.rowSize == matrixToVectorRowSize ||
		matrix.columnSize == matrixToVectorColumnSize {
		return matrix.convertToVector()
	}
	panic(matrixCanNotBeVectorError)
}

func (matrix *Matrix) convertToVector() *Vector {
	if matrix.rowSize != matrixToVectorRowSize &&
		matrix.columnSize != matrixToVectorColumnSize {
		panic(matrixCanNotBeVectorError)
	}
	if matrix.rowSize == matrixToVectorRowSize {
		vector := &Vector{}
		vector.setValues(make([]float64, matrix.columnSize), matrix.columnSize, false)
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

func (matrix *Matrix) GetRowAsVector(rowIndex int) *Vector {
	if !matrix.validateIndex(rowIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return ConstructVector(lang.GetReal2DArrayRow(matrix.GetSlice(), rowIndex), false)
}

func (matrix *Matrix) GetColumnAsVector(columnIndex int) *Vector {
	if !matrix.validateIndex(0, columnIndex) {
		panic(matrixIndexOutOfBoundError)
	}
	return ConstructVector(lang.GetReal2DArrayColumn(matrix.GetSlice(), columnIndex), true)
}

func (matrix *Matrix) VectorGroup(shape bool) *VectorGroup {
	return matrix.convertToVectorGroup(shape)
}

func (matrix *Matrix) convertToVectorGroup(shape bool) *VectorGroup {
	vg := &VectorGroup{}
	return vg.Construct(matrix.GetSlice())
}

func (matrix *Matrix) PolyMatrix() *PolyMatrix {
	return matrix.convertToPolyMatrix()
}

func (matrix *Matrix) convertToPolyMatrix() *PolyMatrix {
	pm := &PolyMatrix{}
	pm.assign(matrix.rowSize, matrix.columnSize)
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
			pm.set(rowIdx, columnIdx, ConstructPolyNode(matrix.get(rowIdx, columnIdx), 0).Poly(eigenPolyMatrixDefaultAES))
		}
	}
	return pm
}

func (matrix *Matrix) getSpiralOrder() []float64 {
	var (
		elemSize                                  = matrix.GetElemSize()
		spiralSlice                               = make([]float64, elemSize)
		cnt, turn, x, y                           = 0, 0, 0, 0
		rightBound, downBound, leftBound, upBound = matrix.GetColumnSize() - 1, matrix.GetRowSize() - 1, 0, 1
	)
	for cnt < elemSize {
		switch turn {
		// right
		case 0:
			{
				if y > rightBound {
					turn = 1
					y--
					x++
					rightBound--
				} else {
					spiralSlice[cnt] = matrix.get(x, y)
					y++
					cnt++
				}

			}
			// down
		case 1:
			{
				if x > downBound {
					turn = 2
					x--
					y--
					downBound--
				} else {
					spiralSlice[cnt] = matrix.get(x, y)
					x++
					cnt++
				}
			}
			// left
		case 2:
			{
				if y < leftBound {
					turn = 3
					y++
					x--
					leftBound++
				} else {
					spiralSlice[cnt] = matrix.get(x, y)
					y--
					cnt++
				}
			}
			// up
		case 3:
			{
				if x < upBound {
					turn = 0
					x++
					y++
					upBound++
				} else {
					spiralSlice[cnt] = matrix.get(x, y)
					x--
					cnt++
				}
			}
		default:
			{

			}
		}
	}
	return spiralSlice
}

func (matrix *Matrix) traverse(funcPtr func(float64)) *Matrix {
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
			funcPtr(matrix.get(rowIdx, columnIdx))
		}
	}
	return matrix
}

func (matrix *Matrix) traverseV2(funcPtr func(int, int, ...float64)) *Matrix {
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
			funcPtr(rowIdx, columnIdx, matrix.get(rowIdx, columnIdx))
		}
	}
	return matrix
}

// Display
// chained option
func (matrix *Matrix) Display() *Matrix {
	if matrix.isNull() {
		fmt.Println("[null]")
		return matrix
	}
	for rowIdx := 0; rowIdx < matrix.rowSize; rowIdx++ {
		for columnIdx := 0; columnIdx < matrix.columnSize; columnIdx++ {
			val := matrix.get(rowIdx, columnIdx)
			if lang.EqualFloat64Zero(val) {
				val = 0.0
			}
			fmt.Printf(" %.5v", val)
		}
		fmt.Println()
	}
	return matrix
}
