package lang

import (
	"fmt"
	"strconv"
)

const (
	complex128RealDefaultPrecision int = 2
	complex128ImagDefaultPrecision int = 2
)

type Complex128Slice []complex128

func (p Complex128Slice) Len() int { return len(p) }
func (p Complex128Slice) Less(i, j int) bool {
	return real(p[i]) < real(p[j]) || (real(p[i]) == real(p[j]) && imag(p[i]) < imag(p[j]))
}
func (p Complex128Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func ConstructComplex128Slice(slice []complex128) Complex128Slice {
	return slice
}

func DisplayOneComplex128(a complex128, precision ...int) {
	realPrec, imagPrec := complex128RealDefaultPrecision, complex128ImagDefaultPrecision
	if precLen := len(precision); precLen > 0 {
		if precLen >= 1 {
			realPrec = precision[0]
			imagPrec = precision[0]
		}
		if precLen >= 2 {
			imagPrec = precision[1]
		}
	}
	realDisplayed := false
	if realA := real(a); !EqualFloat64Zero(realA) {
		fmt.Printf("%."+strconv.Itoa(realPrec)+"v", realA)
		realDisplayed = true
	}
	if imagA := imag(a); !EqualFloat64Zero(imagA) {
		if realDisplayed && imagA > 0 {
			fmt.Printf("+")
		}
		fmt.Printf("%."+strconv.Itoa(imagPrec)+"vi", imagA)
	} else {
		if !realDisplayed {
			fmt.Printf("0")
		}
	}
}
func DisplayComplex128(realPrecision, imagPrecision int, a ...complex128) {
	if aLen := len(a); aLen > 0 {
		for idx := 0; idx < aLen; idx++ {
			DisplayOneComplex128(a[idx], realPrecision, imagPrecision)
		}
	}
}

func DisplayComplex1281DArray(realPrecision, imagPrecision int, a []complex128) {
	for idx := 0; idx < len(a); idx++ {
		DisplayOneComplex128(a[idx], realPrecision, imagPrecision)
		fmt.Printf(", ")
	}
	fmt.Println()
}

func IsOneComplex128PureReal(a complex128) bool {
	return EqualFloat64Zero(imag(a))
}

func IsOneComplex128PureImag(a complex128) bool {
	return EqualFloat64Zero(real(a))
}

func IsComplex128PureReal(a ...complex128) bool {
	if aLen := len(a); aLen > 0 {
		for aIdx := 0; aIdx < aLen; aIdx++ {
			if !IsOneComplex128PureReal(a[aIdx]) {
				return false
			}
		}
	}
	return true
}

func IsComplex128PureImag(a ...complex128) bool {
	if aLen := len(a); aLen > 0 {
		for aIdx := 0; aIdx < aLen; aIdx++ {
			if !IsOneComplex128PureImag(a[aIdx]) {
				return false
			}
		}
	}
	return true
}

func DivideComplex128ArrayEvenOddPart(complex128Array []complex128) ([]complex128, []complex128) {
	if complex128Array == nil ||
		len(complex128Array) == 0 {
		return nil, nil
	}
	size := len(complex128Array)
	evenPart, oddPart := make([]complex128, 0), make([]complex128, 0)
	for idx := 0; idx < size; idx++ {
		if Odd(idx) {
			oddPart = append(oddPart, complex128Array[idx])
		} else {
			evenPart = append(evenPart, complex128Array[idx])
		}
	}
	return evenPart, oddPart
}

// Complex1281DArrayHadamard
// change self
func Complex1281DArrayHadamard(p1, p2 []complex128) []complex128 {
	p1Size, p2Size := len(p1), len(p2)
	if p1Size != p2Size {
		return nil
	}
	for idx := 0; idx < p1Size; idx++ {
		p1[idx] *= p2[idx]
	}
	return p1
}

func Complex128ArrayToIntArray(p []complex128) []int {
	pInt := make([]int, len(p))
	for idx := 0; idx < len(p); idx++ {
		pInt[idx] = int(real(p[idx]))
	}
	return pInt
}

func IntArrayToComplex128Array(p []int) []complex128 {
	pC := make([]complex128, len(p))
	for idx := 0; idx < len(p); idx++ {
		pC[idx] = complex(float64(p[idx]), 0.0)
	}
	return pC
}

func RealArrayToComplex128Array(p []float64) []complex128 {
	pC := make([]complex128, len(p))
	for idx := 0; idx < len(p); idx++ {
		pC[idx] = complex(p[idx], 0.0)
	}
	return pC
}
