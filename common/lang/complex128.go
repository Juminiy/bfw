package lang

import (
	"fmt"
	"strconv"
)

const (
	complex128RealDefaultPrecision int = 2
	complex128ImagDefaultPrecision int = 2
)

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
	if realA := real(a); !EqualFloat64ByAccuracy(realA, 0.0) {
		fmt.Printf("%."+strconv.Itoa(realPrec)+"v", realA)
		realDisplayed = true
	}
	if imagA := imag(a); !EqualFloat64ByAccuracy(imagA, 0.0) {
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
