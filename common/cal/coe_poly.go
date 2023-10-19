package cal

import (
	"bfw/common/la"
	"math"
)

// CoePoly
// if exp coe is 0, set 0
type CoePoly struct {
	coe []float64
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

func (cp *CoePoly) TrailingZeroPaddingPower2() *CoePoly {
	zeroSize := math.Log(float64(cp.size()))
	cp.coe = append(cp.coe, make([]float64, int(zeroSize))...)
	return cp
}

// PointValuePoly
// a*X=b
// a = b*X^{-1}
// X[1] = [x1, x2, x3, ..., xn] must distinct
type PointValuePoly struct {
	x *la.Vandermonde
	b *la.Vector
}

func (pvp *PointValuePoly) CoePoly() *CoePoly {
	return ConstructCoePoly(pvp.b.MulMatrix(pvp.x.Matrix().GetInverse()).GetSlice())
}
