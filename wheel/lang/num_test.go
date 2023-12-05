package lang

import (
	"crypto/sha256"
	"fmt"
	"math"
	"testing"
)

func TestFloat64ToString(t *testing.T) {
	DisplayOneComplex128(complex(0, 0))
	fmt.Println()
}

func TestMath_Pow(t *testing.T) {
	// NaN
	fmt.Println(math.Pow(-27, 1.0/3))
	fmt.Println(math.Pow(27, 1.0/3))
	fmt.Println(math.Pow(-27, 1.0/3))
}

func TestGetRandFloat64ByFloat64Range(t *testing.T) {
	fmt.Println(GetRandFloat64ByInt32Range(10, 20))
	fmt.Println(GetRandFloat64ByInt32Range(10, 20))
	fmt.Println(GetRandFloat64ByInt32Range(10, 20))
	fmt.Println(GetRandFloat64ByInt32Range(10, 20))
	fmt.Println(GetRandFloat64ByInt32Range(10, 20))
	fmt.Println(GetRandFloat64ByFloat64Range(10, 20))
	fmt.Println(GetRandFloat64ByFloat64Range(10, 20))
	fmt.Println(GetRandFloat64ByFloat64Range(10, 20))
	fmt.Println(GetRandFloat64ByFloat64Range(10, 20))
	fmt.Println(GetRandFloat64ByFloat64Range(10, 20))
	fmt.Println(GetRandFloat64ByFloat64Range(10, 20))
}

func TestFloat64ToString2(t *testing.T) {
	//fmt.Println(Float64ToString(-10.22, 2))
	fmt.Println(math.Log2(5))
}

func TestAbsInt(t *testing.T) {
	fmt.Println(AbsInt(11))
	fmt.Println(AbsInt(0))
	fmt.Println(AbsInt(+0))
	fmt.Println(AbsInt(-0))
	fmt.Println(AbsInt(-100))
}

func TestDisplayInt1DArrayInPythonFormat(t *testing.T) {
	DisplayInt1DArrayInPythonFormat([]int{1, 2, 3})
	DisplayInt1DArrayInPythonFormat([]int{})
	DisplayInt1DArrayInPythonFormat(nil)
}

func TestGetRandFloat64ArrayByRange(t *testing.T) {
	for bitCnt := 27; bitCnt < 48; bitCnt++ {
		if 1<<bitCnt > int64(4e9) {
			fmt.Printf("bit = %d, val = %d", bitCnt, 1<<bitCnt)
			break
		}
	}
}

func TestCountSHA256DiffBits(t *testing.T) {
	a, b := []byte{'x'}, []byte{'X'}
	fmt.Println(sha256.Sum256(a), "\n", sha256.Sum256(b))
	fmt.Println(CountSHA256DiffBits(a, b))
}
