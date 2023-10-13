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

// 999000
// 000999
// 000001

// 0100
// 0010

// 02222111
// 01231231

// 0002000
// 0000442
// 0004000
// 0080000
// 0800000
// 0884000
func TestBigNum_Construct(t *testing.T) {
	num1 := ConstructBigNum("2000")
	num2 := ConstructBigNum("442")
	//num1.Add(num2).Display(false, false)
	//num1.Sub(num2).Display(false, false)
	num1.Mul(num2).Display(false, false)
}
