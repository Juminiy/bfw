package num

import (
	"bfw/wheel/lang"
	"fmt"
	"testing"
	"time"
)

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

func TestQ_rsqrt(t *testing.T) {
	fmt.Println(QInverseSqrt(0.15625))
	fmt.Println(QInverseSqrt(0.15625))
	//var a float32 = 3.9999999
	//fmt.Println(int32(a))
}

func TestBigNumberMultiplication(t *testing.T) {
	A, B := "123458987", "67890"
	// 8381630627430
	fmt.Println(NaiveBigNumberMultiplication(A, B))
	fmt.Println(KaratsubaBigNumberMultiplication(A, B))
	fmt.Println(FFTBigNumberMultiplication(A, B))
}

func TestBigNumberMultiplication2(t *testing.T) {
	sizeN, naiveT, karatsubaT, fftT := GetTimeOfBigNumberMultiplyByBit(14)
	fmt.Println(sizeN, lang.GetMS(naiveT), lang.GetMS(karatsubaT), lang.GetMS(fftT))
}

func TestBigNumberMultiply(t *testing.T) {
	time0 := time.Now()
	maxBit, eachBitLoop := 13, 1
	sizeArray := make([]int, 0)
	naiveArray := make([]int64, 0)
	karatsubaArray := make([]int64, 0)
	fftArray := make([]int64, 0)
	for bit := 0; bit < maxBit; bit++ {
		for lp := 0; lp < eachBitLoop; lp++ {
			sizeN, naiveT, karatsubaT, fftT := GetTimeOfBigNumberMultiplyByBit(bit)
			sizeArray = append(sizeArray, sizeN)
			naiveArray = append(naiveArray, lang.GetMS(naiveT))
			karatsubaArray = append(karatsubaArray, lang.GetMS(karatsubaT))
			fftArray = append(fftArray, lang.GetMS(fftT))
		}
	}
	lang.DisplayInt1DArrayInPythonFormat(sizeArray)
	lang.DisplayInt641DArrayInPythonFormat(naiveArray)
	lang.DisplayInt641DArrayInPythonFormat(karatsubaArray)
	lang.DisplayInt641DArrayInPythonFormat(fftArray)
	fmt.Println("total time:", time.Since(time0))
}

func TestDfftBigNumberMultiplication(t *testing.T) {
	fmt.Println(len("fft 10840 length multiply 10840 length number string time"))
}

func TestDfftBigNumberPower(t *testing.T) {
	//fmt.Println(BigNumberPower(2, 10))
	var a uint64 = 0xffffffffffffffff
	fmt.Println(a)
}

func TestBigNumberPower(t *testing.T) {
	for i := 0; i <= 13; i++ {
		fmt.Printf("%v\n%s\n\n", uint64(lang.Power2MulByBitCalculation(2, i)), BigNumberPower(2, i))
	}
}

func TestFraction_Add(t *testing.T) {
	frac := MakeFraction(-1, 2)
	frac.Display(true).
		setND(-1, -2).Display().
		setND(1, 2).Display().
		setND(1, -2).Display()
}

func TestFraction_Display(t *testing.T) {
	// 1/2
	f1 := MakeFraction(-3, -6).Display()
	// -2/1
	f2 := MakeFraction(-2, 1).Display()
	// 5/2
	f1.Sub(f2).Display()
	// -3/2
	f1.Add(f2).Display()
	// -1/1
	f1.Mul(f2).Display()
	// -1/4
	f1.Div(f2).Display()
}

func TestFraction_Display1(t *testing.T) {
	// -1/5
	f1 := MakeFraction(2, -10).Display()
	// -9/1
	f2 := MakeFraction(9, -1).Display()
	// 44/5
	f1.Sub(f2).Display()
	// -46/5
	f1.Add(f2).Display()
	// 9/5
	f1.Mul(f2).Display()
	// 1/45
	f1.Div(f2).Display()
}

func TestFraction_Display2(t *testing.T) {
	// 1/5
	f1 := MakeFraction(-2, -10).Display()
	// 9/1
	f2 := MakeFraction(-9, -1).Display()
	// -44/5
	f1.Sub(f2).Display()
	// 46/5
	f1.Add(f2).Display()
	// 9/5
	f1.Mul(f2).Display()
	// 1/45
	f1.Div(f2).Display()
}
