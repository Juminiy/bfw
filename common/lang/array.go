package lang

import (
	"errors"
	"sort"
)

const (
	realArrayNoSize     int = 0
	realArrayStartIndex int = 0
)

var (
	realArrayInValidError         = errors.New("real array is invalid")
	realArrayIndexOutOfBoundError = errors.New("real array index is out of bound")
)

type Real2DArray [][]float64

func (r2da Real2DArray) Len() int { return len(r2da) }
func (r2da Real2DArray) Less(i, j int) bool {
	return r2da[i][0] < r2da[j][0] || (r2da[i][0] == r2da[j][0] && r2da[i][1] < r2da[j][1])
}
func (r2da Real2DArray) Swap(i, j int) {
	r2da[i], r2da[j] = r2da[j], r2da[i]
}

func ConstructReal2DArray(real2DArray [][]float64) Real2DArray {
	return real2DArray
}

func GetReal2DArrayRow(real2DArray [][]float64, rowIndex int) []float64 {
	if real2DArray == nil ||
		len(real2DArray) == realArrayNoSize {
		panic(realArrayInValidError)
	}
	if rowIndex < realArrayStartIndex ||
		rowIndex >= len(real2DArray) {
		panic(realArrayIndexOutOfBoundError)
	}
	return real2DArray[rowIndex]
}

func GetReal2DArrayLine(real2DArray [][]float64, lineIndex int) []float64 {
	if real2DArray == nil ||
		len(real2DArray) == 0 {
		panic(realArrayInValidError)
	}
	lineSlice := make([]float64, 0)
	for rowIdx := 0; rowIdx < len(real2DArray); rowIdx++ {
		if lineIndex < 0 ||
			lineIndex >= len(real2DArray[rowIdx]) {
			panic(realArrayIndexOutOfBoundError)
		}
		lineSlice = append(lineSlice, real2DArray[rowIdx][lineIndex])
	}
	return lineSlice
}

func GetInitialReal2DArray(rowSize, lineSize int) [][]float64 {
	if rowSize == realArrayNoSize ||
		lineSize == realArrayNoSize {
		return nil
	}
	slice := make([][]float64, rowSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		slice[rowIdx] = make([]float64, lineSize)
	}
	return slice
}

func GetInitialReal1DArray(size int) []float64 {
	if size == realArrayNoSize {
		return nil
	}
	return make([]float64, size)
}

func SortReal2DArrayBySecondFactor(real2DArray [][]float64) {
	r2da := ConstructReal2DArray(real2DArray)
	sort.Sort(r2da)
}
