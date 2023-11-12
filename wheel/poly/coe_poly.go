package poly

import (
	"bfw/wheel/char"
	"bfw/wheel/fft"
	"bfw/wheel/lang"
	"fmt"
	"math"
	"strconv"
)

// CoePoly
// direct operation in order to decrease const complexity
// if exp coe is 0, set 0
type CoePoly struct {
	coe []float64
	val []complex128
}

func ConstructCoePoly(coe []float64) *CoePoly {
	cp := &CoePoly{}
	return cp.Construct(coe)
}

func (cp *CoePoly) Construct(coe []float64) *CoePoly {
	cp.setValues(coe)
	return cp
}

func (cp *CoePoly) setValues(coe []float64) {
	cp.setCoe(coe)
}

func (cp *CoePoly) setCoe(coe []float64) {
	cp.coe = coe
}

func (cp *CoePoly) size() int {
	return len(cp.coe)
}

func (cp *CoePoly) GetCoe() []float64 {
	return cp.coe
}

func (cp *CoePoly) trailingZeroPaddingPower2() *CoePoly {
	zeroSize := math.Log(float64(cp.size()))
	cp.coe = append(cp.coe, make([]float64, int(zeroSize))...)
	return cp
}

func (cp *CoePoly) trailingZeroPadding(destSize int) *CoePoly {
	cp.coe = append(cp.coe, make([]float64, destSize-len(cp.coe))...)
	return cp
}

// Mul
// compare to 0D array is reverse
func (cp *CoePoly) Mul(cpt *CoePoly) *CoePoly {
	return cp.fftMul(cpt)
}

func (cp *CoePoly) fftMul(cpt *CoePoly) *CoePoly {
	destSize := lang.CeilBin(cp.size() + cpt.size())
	cp.trailingZeroPadding(destSize).setVal().dft()
	cpt.trailingZeroPadding(destSize).setVal().dft()
	cp.hadamardProduct(cpt).idft().setElemAndDiv(float64(destSize))
	cp.clearVal()
	cpt.clearVal()
	return cp
}

func (cp *CoePoly) dft() *CoePoly {
	cp.ft(false)
	return cp
}

func (cp *CoePoly) idft() *CoePoly {
	cp.ft(true)
	return cp
}

func (cp *CoePoly) ft(inverse bool) {
	if cp.val == nil {
		return
	}
	fft.FTComplex128Array(&cp.val, inverse)
}

func (cp *CoePoly) setVal() *CoePoly {
	cp.val = lang.RealArrayToComplex128Array(cp.coe)
	return cp
}

func (cp *CoePoly) clearVal() *CoePoly {
	cp.val = nil
	return cp
}

func (cp *CoePoly) hadamardProduct(cpt *CoePoly) *CoePoly {
	if cp.val == nil ||
		cpt.val == nil ||
		len(cp.val) != len(cpt.val) {
		return cp
	}
	for idx := 0; idx < len(cp.val); idx++ {
		cp.val[idx] *= cpt.val[idx]
	}
	return cp
}

func (cp *CoePoly) setElemAndDiv(lambda float64) *CoePoly {
	if cp.coe == nil ||
		cp.val == nil ||
		len(cp.coe) != len(cp.val) {
		return cp
	}
	for idx := 0; idx < len(cp.val); idx++ {
		cp.coe[idx] = real(cp.val[idx]) / lambda
		if lang.EqualFloat64Zero(cp.coe[idx]) {
			cp.coe[idx] = 0
		}
	}
	return cp
}

func (cp *CoePoly) setElemCon() *CoePoly {
	if cp.coe == nil {
		return cp
	}
	for idx := 0; idx < len(cp.coe)-1; idx++ {
		cp.coe[idx+1] += cp.coe[idx] / 10
		cp.coe[idx] = lang.Float64Mod(cp.coe[idx], 10)
	}
	return cp
}

func (cp *CoePoly) trailingZeroTruncate() *CoePoly {
	noneZeroIdx := len(cp.coe) - 1
	for ; noneZeroIdx >= 0 && lang.EqualFloat64Zero(cp.coe[noneZeroIdx]); noneZeroIdx-- {
	}
	cp.coe = cp.coe[:noneZeroIdx+1]
	return cp
}

func (cp *CoePoly) getElemString(index int, hasDisplayed bool, precision ...int) string {
	var (
		signString = undefinedString
		numString  = undefinedString
		expString  = undefinedString
	)
	if !lang.EqualFloat64Zero(cp.coe[index]) {
		if cp.coe[index] > 0 && hasDisplayed {
			signString = "+"
		}
		numString = lang.Float64ToString(cp.coe[index], precision...)
		if index > 1 {
			expString = char.GetExponent(strconv.Itoa(index))
		}
		return signString + numString + polyDefaultAesString + expString
	} else {
		return undefinedString
	}
}

func (cp *CoePoly) getValue10() string {
	cp.trailingZeroTruncate().setElemCon()
	value10 := undefinedString
	for idx := len(cp.coe) - 1; idx >= 0; idx-- {
		value10 += strconv.Itoa(int(cp.coe[idx]))
	}
	return value10
}

func (cp *CoePoly) Value(xVal float64) string {
	if lang.EqualFloat64Zero(xVal - 10) {
		return cp.getValue10()
	}
	return undefinedString
}

func (cp *CoePoly) Display(precision ...int) *CoePoly {
	if cp.coe == nil ||
		cp.size() == 0 {
		return cp
	}
	var (
		currentElemString = undefinedString
		hasDisplayed      = false
	)
	for idx := 0; idx < cp.size(); idx++ {
		currentElemString = cp.getElemString(idx, hasDisplayed, precision...)
		if !hasDisplayed &&
			currentElemString != undefinedString {
			hasDisplayed = true
		}
		fmt.Printf(currentElemString)
	}
	return cp
}

// PointValuePoly
// a*X=b
// a = b*X^{-1}
// X[1] = [x1, x2, x3, ..., xn] must distinct
//type PointValuePoly struct {
//	x *la.Vandermonde
//	b *la.Vector
//}
//
//func (pvp *PointValuePoly) CoePoly() *CoePoly {
//	return ConstructCoePoly(pvp.b.MulMatrix(pvp.x.Matrix().GetInverse()).GetSlice())
//}
