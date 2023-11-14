package num

import (
	"bfw/wheel/fft"
	"bfw/wheel/lang"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

// history:
// 1800BC math, algebra
// 1946AD ENIAC
// BigNum Multiplication simple: 6N² ~ O(N²)
// 1960AD, BigNum Multiplication karatsuba: 60N¹`⁶ ~ O(Nˡᵒᵍ²⁽³⁾)
// 1805AD -> 1965AD (fft theorem) -> 1994AD (FFT Realization), BigNum Multiplication fft: O(N*log(N)*log(log(N))) ~ O(N*log(N))

const (
	naiveLen  int = 1 << 5
	digitTen  int = 10
	digitsLen int = 10
)

var (
	InvalidNumberError = errors.New("invalid number found")
)

// BigNumberMultiply
// Exam to definite how to choose the mul
func BigNumberMultiply(A, B string) string {
	randInt := lang.GetRandomIntValue(2)
	switch randInt {
	case 0:
		{
			return KaratsubaBigNumberMultiplication(A, B)
		}
	case 1:
		{
			return FFTBigNumberMultiplication(A, B)
		}
	case 2:
		{
			return LangBigNumberMultiplication(A, B)
		}
	default:
		{
			panic(errors.New("unsupported algorithm"))
		}
	}
}

func BigNumberPower(Base, Exp int) string {
	base := strconv.Itoa(Base)
	res := "1"
	for Exp > 0 {
		if Exp&1 != 0 {
			res = BigNumberMultiply(res, base)
		}
		base = BigNumberMultiply(base, base)
		Exp >>= 1
	}
	return res
}

func LangBigNumberMultiplication(A, B string) string {
	X, _ := new(big.Int).SetString(A, 10)
	Y, _ := new(big.Int).SetString(B, 10)
	return X.Mul(X, Y).String()
}

func NaiveBigNumberMultiplication(A, B string) string {
	return BigNumberMultiplication(A, B, naiveMul)
}

func KaratsubaBigNumberMultiplication(A, B string) string {
	return BigNumberMultiplication(A, B, karatsubaMul)
}

// FFTBigNumberMultiplication
// bug still occur
func FFTBigNumberMultiplication(A, B string) string {
	return BigNumberMultiplication(A, B, fft.PolyMul)
}

// BigNumberMultiplication
// A,B is string by Decimal Integer order format
// the format also Polynomial Exponent DESC order
func BigNumberMultiplication(A, B string, algorithm func([]int, []int) []int) string {
	resSign := ""
	aSign, bSign := GetNumberStringSign(A), GetNumberStringSign(B)
	if aSign != bSign {
		resSign = "-"
	}
	A, B = truncateNumberStringSign(A), truncateNumberStringSign(B)
	if !numbersValidate(A, B) {
		panic(InvalidNumberError)
	}
	// A,B reverse
	A = lang.StringReverse(A)
	B = lang.StringReverse(B)
	aLen, bLen := len(A), len(B)
	destSize := lang.CeilBin(aLen + bLen)
	// A,B to int array
	AIntArray := String2IntArray(A)
	BIntArray := String2IntArray(B)
	// A,B int array trailing zero padding
	AIntArray = lang.Int1DArrayZeroPadding(AIntArray, destSize-aLen)
	BIntArray = lang.Int1DArrayZeroPadding(BIntArray, destSize-bLen)
	// A,B int array calculate result int array
	resIntArray := algorithm(AIntArray, BIntArray)
	// result int array bit carry
	resIntArray = bitCarry(resIntArray)
	// result int array trailing zero truncate
	resIntArray = lang.Int1DArrayTruncateTrailingZero(resIntArray)
	// result int array to string
	resString := IntArray2String(resIntArray)
	// result string reverse
	resString = lang.StringReverse(resString)
	if len(resString) == 0 {
		return "0"
	}
	return resSign + resString
}

// number A, B receive digit reversed
func naiveMul(A, B []int) []int {
	aLen, bLen := len(A), len(B)
	res := make([]int, aLen+bLen)

	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			res[i+j] += A[i] * B[j]
		}
	}

	return res
}

// number A, B receive digit reversed
func karatsubaMul(A, B []int) []int {
	return karatsubaRec(A, B, len(A), len(A)>>1)
}

// Karatsuba Algorithm:
// expand the original number: 0D123456789,0D987654321 -> 2ᵏ,2ᵏ bit
// number digit reverse
// zero padding for bits expanding.
// destLen := max(len(A), len(B))
// destLen -> CeilBin(destLen) with padding zero
// start algorithm:
// for example: 0D1234 * 0D5678
// A = 0D1234
// B = 0D5678
// a = 12, b = 34
// c = 56, d = 78
//(1). a*c = 672
//(2). b*d = 2652
//(3). (a+b)*(c+d) = 46*134 = 6164
//(4). (3)-(1)-(2) = 2840
//(5). path the number:
//		6720000
//+		0002652
//+		0284000
//=     7006652

// in a recursive level, each 4 mul is decreased to 3 mul
// in theorem, O(N¹`⁶)
func karatsubaRec(A, B []int, size, sz int) []int {
	if size <= naiveLen {
		return naiveMul(A, B)
	}
	res := make([]int, size<<1)

	Ar, Al := A[:sz], A[sz:]
	Br, Bl := B[:sz], B[sz:]
	P1 := karatsubaRec(Al, Bl, sz, sz>>1)
	P2 := karatsubaRec(Ar, Br, sz, sz>>1)

	Alr, Blr := make([]int, sz), make([]int, sz)
	for i := 0; i < sz; i++ {
		Alr[i] = Al[i] + Ar[i]
		Blr[i] = Bl[i] + Br[i]
	}
	P3 := karatsubaRec(Alr, Blr, sz, sz>>1)
	for i := 0; i < size; i++ {
		P3[i] -= P1[i] + P2[i]
	}
	for i := 0; i < size; i++ {
		res[i] = P2[i]
	}
	for i := size; i < size<<1; i++ {
		res[i] = P1[i-size]
	}
	for i := sz; i < size+sz; i++ {
		res[i] += P3[i-sz]
	}
	return res
}

func bitCarry(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		A[i+1] += A[i] / digitTen
		A[i] %= digitTen
	}
	return A
}

func digitsValidate(A string) bool {
	for _, aByte := range A {
		if aByte < '0' || aByte > '9' {
			return false
		}
	}
	return true
}

func numbersValidate(A ...string) bool {
	if aLen := len(A); aLen > 0 {
		for idx := 0; idx < aLen; idx++ {
			if !digitsValidate(A[idx]) {
				return false
			}
		}
	}
	return true
}

func String2IntArray(str string) []int {
	strLen := len(str)
	res := make([]int, strLen)
	for i := 0; i < strLen; i++ {
		res[i] = int(str[i] - '0')
	}
	return res
}

func IntArray2String(A []int) string {
	res := ""
	for _, aInt := range A {
		res += string(rune(aInt + '0'))
	}
	return res
}

func GetNumberStringSign(A string) string {
	if len(A) > 0 && A[0] == '-' {
		return "-"
	}
	return ""
}

func truncateNumberStringSign(A string) string {
	if len(A) > 0 && (A[0] == '+' || A[0] == '-') {
		return A[1:]
	}
	return A
}

func generateNumberSign() string {
	if rand.Intn(2) == 0 {
		return "-"
	} else {
		return "+"
	}
}

func generateRandomDigits(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = lang.Digits[rand.Intn(digitsLen)]
	}
	return string(result)
}

func GenerateNumberString(size int) string {
	return generateNumberSign() + generateRandomDigits(size)
}

func GetTimeOfBigNumberMultiplyByBit(bit int, runNaive ...bool) (int, time.Duration, time.Duration, time.Duration, time.Duration) {
	size := lang.GetRandomIntValue(1 << bit)
	if size == 0 {
		return size, 0, 0, 0, 0
	}
	var (
		time0               time.Time
		du1, du2, du3, du4  time.Duration
		_, res2, res3, res4 string
	)
	//time0 := time.Now()
	A, B := GenerateNumberString(size), GenerateNumberString(size)
	//fmt.Printf("generate two %d length number string time: %v\n", size, time.Since(time0))
	if len(runNaive) > 0 && runNaive[0] {
		time0 = time.Now()
		_ = NaiveBigNumberMultiplication(A, B)
		du1 = time.Since(time0)
	}
	//fmt.Printf("naive length %d multiply %d number string time: %v\n", size, size, time.Since(time1))
	time0 = time.Now()
	res2 = KaratsubaBigNumberMultiplication(A, B)
	du2 = time.Since(time0)
	//fmt.Printf("karatsuba length %d multiply %d number string time: %v\n", size, size, time.Since(time2))
	time0 = time.Now()
	res3 = FFTBigNumberMultiplication(A, B)
	du3 = time.Since(time0)
	//fmt.Printf("fft %d length multiply %d length number string time: %v\n", size, size, time.Since(time3))
	time0 = time.Now()
	res4 = LangBigNumberMultiplication(A, B)
	du4 = time.Since(time0)
	if res4 != res2 || res4 != res3 {
		panic(errors.New("algorithm is something wrong"))
	}
	return size, du1, du2, du3, du4
}

func RunBigNumberMultiply(maxBit, eachBitLoop int) {
	time0 := time.Now()
	sizeArray := make([]int, 0)
	//naiveArray := make([]int64, 0)
	karatsubaArray := make([]int64, 0)
	fftArray := make([]int64, 0)
	langArray := make([]int64, 0)
	for bit := 0; bit < maxBit; bit++ {
		for lp := 0; lp < eachBitLoop; lp++ {
			sizeN, _, karatsubaT, fftT, langT := GetTimeOfBigNumberMultiplyByBit(bit)
			sizeArray = append(sizeArray, sizeN)
			//naiveArray = append(naiveArray, lang.GetMS(naiveT))
			karatsubaArray = append(karatsubaArray, lang.GetUS(karatsubaT))
			fftArray = append(fftArray, lang.GetUS(fftT))
			langArray = append(langArray, lang.GetUS(langT))
		}
	}
	fmt.Println("size = ")
	lang.DisplayInt1DArrayInPythonFormat(sizeArray)
	//fmt.Println("naive = ")
	//lang.DisplayInt641DArrayInPythonFormat(naiveArray)
	fmt.Println("karatsuba = ")
	lang.DisplayInt641DArrayInPythonFormat(karatsubaArray)
	fmt.Println("fft = ")
	lang.DisplayInt641DArrayInPythonFormat(fftArray)
	fmt.Println("lang = ")
	lang.DisplayInt641DArrayInPythonFormat(langArray)
	fmt.Println("total time = ", time.Since(time0))
}
