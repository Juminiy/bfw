package lang

import (
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
