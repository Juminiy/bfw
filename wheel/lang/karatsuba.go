package lang

import (
	"bfw/wheel/fft"
	"errors"
	"math/rand"
)

const (
	naiveLen int    = 1 << 5
	digitTen int    = 10
	Digits   string = "0123456789"
)

var (
	InvalidNumberError = errors.New("invalid number found")
)

func NaiveBigNumberMultiplication(A, B string) string {
	return bigNumberMultiplication(A, B, naiveMul)
}

func KaratsubaBigNumberMultiplication(A, B string) string {
	return bigNumberMultiplication(A, B, karatsubaMul)
}

func FFTBigNumberMultiplication(A, B string) string {
	return bigNumberMultiplication(A, B, fft.IntPolyMultiplication)
}

func bigNumberMultiplication(A, B string, algorithm func([]int, []int) []int) string {
	resSign := ""
	aSign, bSign := GetNumberStringSign(A), GetNumberStringSign(B)
	if aSign != bSign {
		resSign = "-"
	}
	A, B = truncateNumberStringSign(A), truncateNumberStringSign(B)
	if !numbersValidate(A, B) {
		panic(InvalidNumberError)
	}
	A = StringReverse(A)
	B = StringReverse(B)
	aLen, bLen := len(A), len(B)
	destSize := CeilBin(aLen + bLen)
	AIntArray := string2IntArray(A)
	BIntArray := string2IntArray(B)
	AIntArray = Int1DArrayZeroPadding(AIntArray, destSize-aLen)
	BIntArray = Int1DArrayZeroPadding(BIntArray, destSize-bLen)
	resIntArray := algorithm(AIntArray, BIntArray)
	resIntArray = bitCarry(resIntArray)
	resString := IntArray2String(resIntArray)
	resString = StringReverse(resString)
	resString = TruncateStringPrefixZero(resString)
	return resSign + resString
}

// number A, B is digit reversed
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

func karatsubaMul(A, B []int) []int {
	return karatsubaRec(A, B, len(A), len(A)>>1)
}

// history:
// 1800BC math, algebra
// 1946AD ENIAC
// BigNum Multiplication simple: 6N² ~ O(N²)
// 1960AD, BigNum Multiplication karatsuba: 60N¹`⁶ ~ O(Nˡᵒᵍ²⁽³⁾)
// 1805AD -> 1965AD (fft theorem) -> 1994AD, BigNum Multiplication fft: O(N*log(N)*log(log(N))) ~ O(N*log(N))

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

func string2IntArray(str string) []int {
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
	return undefinedString
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
		result[i] = Digits[rand.Intn(digitTen)]
	}
	return string(result)
}

func GenerateNumberString(size int) string {
	return generateNumberSign() + generateRandomDigits(size)
}
