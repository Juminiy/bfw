package fft

import (
	"bfw/wheel/lang"
	"math"
)

// fft stand for fast fourier transform

func PolyMul(p1, p2 []int) []int {
	return polyMul(p1, p2)
}

// the first case is:
//  1. high precision multiply
//     (1). convert to polynomial multiply
//     (2). for each bit[i], bit[i] %= 10, bit[i-1] += bit[i]/10
//  2. polynomial multiply
//     (1). coefficient polynomial stands
//     (2). trailing zero padding, 2^k - len(p1)+len(p2)
//
// when to call polyMul, all condition must fit
// 1. int array is digit reversed
// 2. int array is trialing zero padding
// 3. len(p1) == len(p2)
func polyMul(p1, p2 []int) []int {
	destSize := len(p1)
	// p1,p2 dft
	p1DFT := polyDFT(lang.IntArrayToComplex128Array(p1))
	p2DFT := polyDFT(lang.IntArrayToComplex128Array(p2))
	// p1,p2 hadamard resDFT
	resDFT := lang.Complex1281DArrayHadamard(p1DFT, p2DFT)
	// resDFT idft
	resDFT = polyIDFT(resDFT)
	// to res int array
	resIntArray := lang.Complex128ArrayToIntArray(resDFT)
	// res int array div destSize
	resIntArray = lang.Int1DArrayShiftBit(resIntArray, lang.CeilBinCnt(destSize))
	return resIntArray
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
	pi2Dn := math.Pi / float64(size>>1)
	inverseSign := 1.0
	if inverse {
		inverseSign = -1.0
	}
	for j := 0; j < (size >> 1); j++ {
		omegaPowerJ := complex(math.Cos(pi2Dn*float64(j)), inverseSign*math.Sin(pi2Dn*float64(j)))
		polyRes[j] = evenRes[j] + omegaPowerJ*oddRes[j]
		polyRes[j+(size>>1)] = evenRes[j] - omegaPowerJ*oddRes[j]
	}
	return polyRes
}

func FTComplex128Array(p *[]complex128, inverse bool) {
	*p = polyFT(*p, inverse)
}
