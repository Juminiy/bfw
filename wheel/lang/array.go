package lang

import (
	"errors"
	"fmt"
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
	return (r2da[i][1] == r2da[j][1] && r2da[i][0] < r2da[j][0]) || r2da[i][1] < r2da[j][1]
}
func (r2da Real2DArray) Swap(i, j int) {
	r2da[i], r2da[j] = r2da[j], r2da[i]
}

func ConstructReal2DArrayBySize(size int) Real2DArray {
	return ConstructReal2DArray(make([][]float64, size))
}

func ConstructReal2DArray(real2DArray [][]float64) Real2DArray {
	for idx := 0; idx < len(real2DArray); idx++ {
		real2DArray[idx] = make([]float64, 2)
	}
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

func GetReal2DArrayColumn(real2DArray [][]float64, columnIndex int) []float64 {
	if real2DArray == nil ||
		len(real2DArray) == 0 {
		panic(realArrayInValidError)
	}
	columnSlice := make([]float64, 0)
	for rowIdx := 0; rowIdx < len(real2DArray); rowIdx++ {
		if columnIndex < 0 ||
			columnIndex >= len(real2DArray[rowIdx]) {
			panic(realArrayIndexOutOfBoundError)
		}
		columnSlice = append(columnSlice, real2DArray[rowIdx][columnIndex])
	}
	return columnSlice
}

func GetInitialReal2DArray(rowSize, columnSize int) [][]float64 {
	if rowSize == realArrayNoSize ||
		columnSize == realArrayNoSize {
		return nil
	}
	slice := make([][]float64, rowSize)
	for rowIdx := 0; rowIdx < rowSize; rowIdx++ {
		slice[rowIdx] = make([]float64, columnSize)
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

func ConvertReal2DArrayToInt2DArray(real2DArray [][]float64) [][]int {
	int2DArray := make([][]int, len(real2DArray))
	for rowIdx := 0; rowIdx < len(real2DArray); rowIdx++ {
		int2DArray[rowIdx] = make([]int, len(real2DArray[rowIdx]))
		for columnIdx := 0; columnIdx < len(real2DArray[rowIdx]); columnIdx++ {
			int2DArray[rowIdx][columnIdx] = int(real2DArray[rowIdx][columnIdx])
		}
	}
	return int2DArray
}

func ConvertInt2DArrayToReal2DArray(int2DArray [][]int) [][]float64 {
	real2DArray := make([][]float64, len(int2DArray))
	for rowIdx := 0; rowIdx < len(int2DArray); rowIdx++ {
		real2DArray[rowIdx] = make([]float64, len(int2DArray[rowIdx]))
		for columnIdx := 0; columnIdx < len(int2DArray[rowIdx]); columnIdx++ {
			real2DArray[rowIdx][columnIdx] = float64(int2DArray[rowIdx][columnIdx])
		}
	}
	return real2DArray
}

func Int1DArrayZeroPadding(int1DArray []int, zeroCnt int, direction ...bool) []int {
	if len(direction) > 0 && direction[0] {
		return append(make([]int, zeroCnt), int1DArray...)
	}
	return append(int1DArray, make([]int, zeroCnt)...)
}

func Divide1DArrayEvenOddPart(int1DArray []int) ([]int, []int) {
	if int1DArray == nil ||
		len(int1DArray) == 0 {
		return nil, nil
	}
	size := len(int1DArray)
	evenPart, oddPart := make([]int, 0), make([]int, 0)
	for idx := 0; idx < size; idx++ {
		if IsOdd(idx) {
			oddPart = append(oddPart, int1DArray[idx])
		} else {
			evenPart = append(evenPart, int1DArray[idx])
		}
	}
	return evenPart, oddPart
}

func Int1DArrayContribute(a []int, inverse bool) []int {
	if inverse {
		for idx := len(a) - 1; idx > 0; idx-- {
			a[idx-1] += a[idx] / 10
			a[idx] %= 10
		}
	} else {
		for idx := 0; idx < len(a)-1; idx++ {
			a[idx+1] += a[idx] / 10
			a[idx] %= 10
		}
	}
	return a
}

func Int1DArrayMulLambda(a []int, lambda int) []int {
	for idx := 0; idx < len(a); idx++ {
		a[idx] *= lambda
	}
	return a
}

func Int1DArrayDivLambda(a []int, lambda int) []int {
	for idx := 0; idx < len(a); idx++ {
		a[idx] /= lambda
	}
	return a
}

func Int1DArrayShiftBit(a []int, bit int) []int {
	for idx := 0; idx < len(a); idx++ {
		a[idx] >>= bit
	}
	return a
}

func Int1DArrayReverse(a []int) []int {
	for idx := 0; idx < (len(a) >> 1); idx++ {
		a[idx], a[len(a)-1-idx] = a[len(a)-1-idx], a[idx]
	}
	return a
}

func Int1DArrayTruncateLeadingZero(a []int) []int {
	sIndex := 0
	for ; sIndex < len(a) && a[sIndex] == 0; sIndex++ {
	}
	return a[sIndex:]
}

func Int1DArrayTruncateTrailingZero(a []int) []int {
	sIndex := len(a) - 1
	for ; sIndex >= 0 && a[sIndex] == 0; sIndex-- {
	}
	return a[:sIndex+1]
}

func DisplayInt1DArrayInPythonFormat(a []int) {
	aLen := len(a)
	fmt.Printf("[")
	if aLen > 0 {
		fmt.Printf("%d", a[0])
		for idx := 1; idx < aLen; idx++ {
			fmt.Printf(", %d", a[idx])
		}
	} else {
		fmt.Printf("null")
	}
	fmt.Println("]")
}

func DisplayInt641DArrayInPythonFormat(a []int64) {
	aLen := len(a)
	fmt.Printf("[")
	if aLen > 0 {
		fmt.Printf("%d", a[0])
		for idx := 1; idx < aLen; idx++ {
			fmt.Printf(", %d", a[idx])
		}
	} else {
		fmt.Printf("null")
	}
	fmt.Println("]")
}

func DisplayFloat641DArrayInPythonFormat(a []float64) {
	aLen := len(a)
	fmt.Printf("[")
	if aLen > 0 {
		fmt.Printf("%f", a[0])
		for idx := 1; idx < aLen; idx++ {
			fmt.Printf(", %f", a[idx])
		}
	} else {
		fmt.Printf("null")
	}
	fmt.Println("]")
}

func GenerateInt641DArray(size int, maxValAbs int64) []int64 {
	i641da := make([]int64, size)
	for idx, _ := range i641da {
		i641da[idx] = GetRandomInt64Value(int64(AbsInt(int(maxValAbs))), true)
	}
	return i641da
}
func GenerateUInt641DArray(size int, maxValAbs int64) []int64 {
	i641da := make([]int64, size)
	for idx, _ := range i641da {
		i641da[idx] = GetRandomInt64Value(int64(AbsInt(int(maxValAbs))))
	}
	return i641da
}

func EqualInt641DArray(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := 0; idx < len(a); idx++ {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func PairIntEqual(a1, b1, a2, b2 int) bool {
	return a1 == a2 && b1 == b2
}
