package lang

import (
	"math"
	"math/rand"
	"strconv"
)

const (
	defaultFloat32Accuracy      = 1e-5
	defaultFloat64Accuracy      = 1e-7
	defaultFloat32Precision int = 5
	defaultFloat64Precision int = 5
)

func EqualFloat32ByAccuracy(a, b float32, acc ...float32) bool {
	var rAcc float32 = 0.0
	if len(acc) == 0 {
		rAcc = defaultFloat32Accuracy
	} else {
		rAcc = acc[0]
	}
	if math.Abs(float64(a-b)) < float64(rAcc) {
		return true
	}
	return false
}

func EqualFloat64ByAccuracy(a, b float64, acc ...float64) bool {
	var rAcc float64 = 0.0
	if len(acc) == 0 {
		rAcc = defaultFloat64Accuracy
	} else {
		rAcc = acc[0]
	}
	if math.Abs(a-b) < rAcc {
		return true
	}
	return false
}

// GetRandomIntValue -> [0,n)
func GetRandomIntValue(n int) int {
	return rand.Intn(n)
}

// GetRandomIntValue2 -> [1,n)
func GetRandomIntValue2(n int) int {
	return rand.Intn(n-1) + 1
}

func GetRandomMapValue(n, k int) map[int]bool {
	if k >= n {
		return nil
	}
	genMap := make(map[int]bool)
	for idx := 0; idx < k; idx++ {
		genVal := 0
		for ; genMap[genVal]; genVal = GetRandomIntValue(n) {
		}
		genMap[genVal] = true
	}
	return genMap
}

func GetCombinationSliceMap(n, k int) []map[int]bool {
	sliceMap := make([]map[int]bool, 0)
	for bin := 0; bin < (1 << n); bin++ {
		if CountBits(bin) == k {
			intMap := make(map[int]bool)
			for bit := 0; bit < n; bit++ {
				if (bin & (1 << bit)) > 0 {
					intMap[bit] = true
				}
			}
			sliceMap = append(sliceMap, intMap)
		}
	}
	return sliceMap
}

func CountBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func Odd(n int) bool {
	return n%2 == 1
}

func MinusOnePower(n int) int {
	if Odd(n) {
		return -1
	}
	return 1
}

func IsStringValueIntZero(valStr string) bool {
	strIdx, strLen := 0, len(valStr)
	for strIdx < strLen {
		if valStr[strIdx] == ' ' ||
			valStr[strIdx] == '+' ||
			valStr[strIdx] == '-' ||
			valStr[strIdx] == '.' ||
			valStr[strIdx] == '0' {
			strIdx++
		} else {
			break
		}
	}
	return strIdx == strLen
}

func MaxInt(a ...int) int {
	if len(a) > 0 {
		maxInt := a[0]
		for _, val := range a {
			if val > maxInt {
				maxInt = val
			}
		}
		return maxInt
	}
	return 0
}

func MinInt(a ...int) int {
	if len(a) > 0 {
		minInt := a[0]
		for _, val := range a {
			if val < minInt {
				minInt = val
			}
		}
		return minInt
	}
	return 0
}

func Square(x float64) float64 {
	return x * x
}

func Float64ToString(a float64, precision ...int) string {
	precisionBit := defaultFloat64Precision
	if len(precision) > 0 {
		precisionBit = precision[0]
	}
	return strconv.FormatFloat(a, byte('f'), precisionBit, 64)
}
