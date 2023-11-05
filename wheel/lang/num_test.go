package lang

import (
	"fmt"
	"math"
	"testing"
	"time"
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
	num1 := ConstructBigNum("-1234567877")
	num2 := ConstructBigNum("123456")
	//num1.add(num2).Display(false, false)
	//num1.sub(num2).Display(false, false)
	num1.Mul(num2).Display(false, false)
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

func TestKaratsubaBigNumberMultiplication(t *testing.T) {
	size := GetRandomIntValue(1 << 20)
	time0 := time.Now()
	A, B := GenerateNumberString(size), GenerateNumberString(size)
	fmt.Printf("generate two %d number string time: %v\n", size, time.Since(time0))
	time1 := time.Now()
	NaiveBigNumberMultiplication(A, B)
	fmt.Printf("naive %d multiply %d number string time: %v\n", size, size, time.Since(time1))
	time2 := time.Now()
	KaratsubaBigNumberMultiplication(A, B)
	fmt.Printf("karatsuba %d multiply %d number string time: %v\n", size, size, time.Since(time2))
}
