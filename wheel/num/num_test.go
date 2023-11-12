package num

import (
	"bfw/wheel/lang"
	"errors"
	"fmt"
	"testing"
	"time"
)

func GetTimeBigNumberMultiplyByBit(bit int) (int, time.Duration, time.Duration, time.Duration) {
	size := lang.GetRandomIntValue(1 << bit)
	if size == 0 {
		return size, 0, 0, 0
	}
	var (
		time0         time.Time
		du1, du2, du3 time.Duration
	)
	//time0 := time.Now()
	A, B := GenerateNumberString(size), GenerateNumberString(size)
	//fmt.Printf("generate two %d length number string time: %v\n", size, time.Since(time0))
	time0 = time.Now()
	res1 := NaiveBigNumberMultiplication(A, B)
	du1 = time.Since(time0)
	//fmt.Printf("naive length %d multiply %d number string time: %v\n", size, size, time.Since(time1))
	time0 = time.Now()
	res2 := KaratsubaBigNumberMultiplication(A, B)
	du2 = time.Since(time0)
	//fmt.Printf("karatsuba length %d multiply %d number string time: %v\n", size, size, time.Since(time2))
	time0 = time.Now()
	res3 := FFTBigNumberMultiplication(A, B)
	du3 = time.Since(time0)
	//fmt.Printf("fft %d length multiply %d length number string time: %v\n", size, size, time.Since(time3))
	if res1 != res2 || res1 != res3 {
		panic(errors.New("algorithm is something wrong"))
	}
	return size, du1, du2, du3
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
	sizeN, naiveT, karatsubaT, fftT := GetTimeBigNumberMultiplyByBit(14)
	fmt.Println(sizeN, lang.GetMS(naiveT), lang.GetMS(karatsubaT), lang.GetMS(fftT))
}

func TestBigNumberMultiply(t *testing.T) {
	maxBit, eachBitLoop := 20, 4
	sizeArray := make([]int, 0)
	naiveArray := make([]float64, 0)
	karatsubaArray := make([]float64, 0)
	fftArray := make([]float64, 0)
	for bit := 0; bit < maxBit; bit++ {
		for lp := 0; lp < eachBitLoop; lp++ {
			sizeN, naiveT, karatsubaT, fftT := GetTimeBigNumberMultiplyByBit(bit)
			sizeArray = append(sizeArray, sizeN)
			naiveArray = append(naiveArray, lang.GetMS(naiveT))
			karatsubaArray = append(karatsubaArray, lang.GetMS(karatsubaT))
			fftArray = append(fftArray, lang.GetMS(fftT))
		}
	}
	fmt.Println(sizeArray)
	fmt.Println(naiveArray)
	fmt.Println(karatsubaArray)
	fmt.Println(fftArray)
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
	var i uint64 = 0
	for ; i <= 63; i++ {
		fmt.Printf("%v\n%s\n\n", uint64(1<<i), BigNumberPower(2, int(i)))
	}
}
