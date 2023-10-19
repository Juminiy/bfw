package fft

import (
	"bfw/wheel/lang"
	"math"
)

// stand for Fast Fourier Transform

// the first case is:
//  1. high precision multiply
//     (1). convert to polynomial multiply
//     (2). for each bit[i], bit[i] %= 10, bit[i-1] += bit[i]/10
//  2. polynomial multiply
//     (1). coefficient polynomial stands
//     (2). trailing zero padding, 2^k - len(p1)+len(p2)
func polyMul(p1, p2 []int) []int {
	p1Len, p2Len := len(p1), len(p2)
	bitCnt := int(math.Ceil(math.Log2(float64(p1Len + p2Len))))
	p1 = lang.Int1DArrayZeroPadding(p1, 1<<(bitCnt)-p1Len)
	p2 = lang.Int1DArrayZeroPadding(p2, 1<<(bitCnt)-p2Len)
	p1DFT := polyDFT(lang.IntArrayToComplex128Array(p1))
	p2DFT := polyDFT(lang.IntArrayToComplex128Array(p2))
	p1DFT = lang.Complex1281DArrayHadamard(p1DFT, p2DFT)
	destIntArray := lang.Complex128ArrayToIntArray(polyIDFT(p1DFT))
	destIntArray = lang.Int1DArrayDivLambda(destIntArray, 1<<bitCnt)
	return lang.Int1DArrayContribute(destIntArray, false)
}

func polyDFT(p []complex128) []complex128 {
	return polyFT(p, false)
}

func polyIDFT(p []complex128) []complex128 {
	return polyFT(p, true)
}

func polyFT(p []complex128, inverse bool) []complex128 {
	if len(p) == 1 {
		return p
	}
	size := len(p)
	evenPart, oddPart := lang.DivideComplex128ArrayEvenOddPart(p)
	evenRes := polyFT(evenPart, inverse)
	oddRes := polyFT(oddPart, inverse)
	polyRes := make([]complex128, size)
	pi2dn := math.Pi / float64(size>>1)
	inverseSign := 1.0
	if inverse {
		inverseSign = -1.0
	}
	for j := 0; j < (size >> 1); j++ {
		omegaPowerJ := complex(math.Cos(pi2dn*float64(j)), inverseSign*math.Sin(pi2dn*float64(j)))
		polyRes[j] = evenRes[j] + omegaPowerJ*oddRes[j]
		polyRes[j+(size>>1)] = evenRes[j] - omegaPowerJ*oddRes[j]
	}
	return polyRes
}
