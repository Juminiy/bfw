package la

import (
	"bfw/common/lang"
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	polyNodeInvalidCoefficient float64 = 0.0
	polyNodeInvalidExponent    int     = -1
	polyNodeCoefficientZero    float64 = 0.0
	polyNodeCoefficientOne     float64 = 1.0
	polyNodeExponentZero       int     = 0
	polyNodeExponentOne        int     = 1
)

var (
	polyNodeCanNotOptError = errors.New("two poly nodes cannot be operated")
	polyNodeInvalidError   = errors.New("poly nodes is invalid")
	NullPolyNode           = &PolyNode{}
)

type PolyNode struct {
	coe float64
	exp int
}

func ConstructPolyNode(coe float64, exp int) *PolyNode {
	return constructPolyNode(coe, exp)
}

func ConstructNullPolyNode() *PolyNode {
	return constructPolyNode(polyNodeInvalidCoefficient, polyNodeInvalidExponent, true)
}

func ConstructValidPolyNode(exp int) *PolyNode {
	return constructPolyNode(polyNodeCoefficientZero, exp)
}

func constructPolyNode(coe float64, exp int, null ...bool) *PolyNode {
	if len(null) != 0 && null[0] {
		return &PolyNode{}
	}
	pn := &PolyNode{}
	return pn.Construct(coe, exp)
}

func (pn *PolyNode) Construct(coe float64, exp int) *PolyNode {
	pn.setValues(coe, exp)
	return pn
}

func (pn *PolyNode) validate() bool {
	return pn.validateCoefficient() &&
		pn.validateExponent()
}

func (pn *PolyNode) validateCoefficient() bool {
	return !lang.EqualFloat64ByAccuracy(polyNodeInvalidCoefficient, pn.coe)
}

func (pn *PolyNode) validateExponent() bool {
	return pn.exp != polyNodeInvalidExponent
}

func (pn *PolyNode) makeCopy() *PolyNode {
	pnt := &PolyNode{}
	pnt.setValues(pn.coe, pn.exp)
	return pnt
}

func (pn *PolyNode) swap(pnt *PolyNode) {
	pnTemp := &PolyNode{}
	pnTemp.setSelf(pn)
	pn.setSelf(pnt)
	pnt.setSelf(pnTemp)
}

func (pn *PolyNode) setSelf(pnt *PolyNode) {
	pn.setValues(pnt.coe, pnt.exp)
}

func (pn *PolyNode) setValues(coe float64, exp int) {
	pn.setCoe(coe)
	pn.setExp(exp)
}

func (pn *PolyNode) setCoe(coe float64) {
	pn.coe = coe
}

func (pn *PolyNode) setExp(exp int) {
	pn.exp = exp
}

func (pn *PolyNode) null() *PolyNode {
	return &PolyNode{}
}

func (pn *PolyNode) setNull() {
	pn.setValues(polyNodeInvalidCoefficient, polyNodeInvalidExponent)
}

func (pn *PolyNode) setZero() {
	pn.setCoeZero()
}

func (pn *PolyNode) setCoeZero() {
	pn.coe = polyNodeCoefficientZero
}

func (pn *PolyNode) setOpt(opt rune, pnt *PolyNode) {
	pn.one2OneOpt(opt, pnt)
}

func (pn *PolyNode) isZero() bool {
	return pn.coe == polyNodeCoefficientZero
}

func (pn *PolyNode) isConstant() bool {
	return pn.exp == polyNodeExponentZero
}

func (pn *PolyNode) one2OneOpt(opt rune, pnt *PolyNode) *PolyNode {
	if pnt == nil {
		panic(polyNodeInvalidError)
	}
	switch opt {
	case '+':
		{
			if pn.exp != pnt.exp {
				panic(polyNodeCanNotOptError)
			}
			pn.coe += pnt.coe
		}
	case '-':
		{
			if pn.exp != pnt.exp {
				panic(polyNodeCanNotOptError)
			}
			pn.coe -= pnt.coe
		}
	case '*':
		{
			pn.coe *= pnt.coe
			pn.exp += pnt.exp
		}
	case '/':
		{
			if pn.exp < pnt.exp {
				panic(polyNodeCanNotOptError)
			}
			pn.coe /= pnt.coe
			pn.exp -= pnt.exp
		}
	default:
		{

		}
	}
	return pn
}

func (pn *PolyNode) Equal(pnt *PolyNode) bool {
	if pnt == nil {
		return false
	}
	if pn.exp == pnt.exp &&
		lang.EqualFloat64ByAccuracy(pn.coe, pnt.coe) {
		return true
	}
	return false
}

func (pn *PolyNode) Add(pnt *PolyNode) *PolyNode {
	return pn.one2OneOpt('+', pnt)
}

func (pn *PolyNode) GetPlus(pnt *PolyNode) *PolyNode {
	pnCopy := pn.makeCopy()
	return pnCopy.Add(pnt)
}

func (pn *PolyNode) Sub(pnt *PolyNode) *PolyNode {
	return pn.one2OneOpt('-', pnt)
}

func (pn *PolyNode) GetMinus(pnt *PolyNode) *PolyNode {
	pnCopy := pn.makeCopy()
	return pnCopy.Sub(pnt)
}

func (pn *PolyNode) Mul(pnt *PolyNode) *PolyNode {
	return pn.one2OneOpt('*', pnt)
}

func (pn *PolyNode) GetTimes(pnt *PolyNode) *PolyNode {
	pnCopy := pn.makeCopy()
	return pnCopy.Mul(pnt)
}

func (pn *PolyNode) Div(pnt *PolyNode) *PolyNode {
	return pn.one2OneOpt('/', pnt)
}

func (pn *PolyNode) GetQuotient(pnt *PolyNode) *PolyNode {
	pnCopy := pn.makeCopy()
	return pnCopy.Div(pnt)
}

func (pn *PolyNode) Exp(exp int) *PolyNode {
	pn.coe = math.Pow(pn.coe, float64(exp))
	pn.exp *= exp
	return pn
}

func (pn *PolyNode) Sqrt(exp int) *PolyNode {
	return pn.Exp(1.0 / exp)
}

func (pn *PolyNode) Value(value float64) float64 {
	return pn.coe * math.Pow(value, float64(pn.exp))
}

func (pn *PolyNode) Integral() *PolyNode {
	if pn.validate() {
		if !pn.isZero() {
			pn.exp++
			pn.coe /= float64(pn.exp)
		}
	}
	return pn
}

func (pn *PolyNode) Derivative() *PolyNode {
	if pn.validate() {
		if pn.isConstant() {
			pn.coe = polyNodeCoefficientZero
		} else {
			pn.coe *= float64(pn.exp)
			pn.exp--
		}
	}
	return pn
}

// Display
// print the poly node
func (pn *PolyNode) Display(isPrintln bool, aes rune) *PolyNode {
	if pn.CanDisplay(aes) {
		pn.displayCoefficient().
			displayAES(aes).
			displayExponent()
	}
	if isPrintln {
		fmt.Println()
	}
	return pn
}

func (pn *PolyNode) CanDisplay(aes rune) bool {
	return pn.ToString(aes) != undefinedString
}

func (pn *PolyNode) ToString(aes rune, precision ...int) string {
	return pn.getCoefficientStr(precision...) + pn.getAESStr(aes) + pn.getExponentStr()
}

func (pn *PolyNode) displayValidate() bool {
	return pn.validateCoefficient() &&
		pn.validateExponent()
}

func (pn *PolyNode) displayCoefficient() *PolyNode {
	if pn.displayValidate() {
		fmt.Printf("%.5v", pn.coe)
	}
	return pn
}

func (pn *PolyNode) getCoefficientStr(precision ...int) string {
	coeStr := undefinedString
	if pn.displayValidate() {
		return lang.Float64ToString(pn.coe, precision...)
	}
	return coeStr
}

func (pn *PolyNode) displayAES(aes rune) *PolyNode {
	if pn.displayValidate() {
		if pn.exp != polyNodeExponentZero {
			fmt.Printf("%c", aes)
		}
	}
	return pn
}

func (pn *PolyNode) getAESStr(aes rune) string {
	aesStr := undefinedString
	if pn.displayValidate() {
		if pn.exp != polyNodeExponentZero {
			aesStr = string(aes)
		}
	}
	return aesStr
}

func (pn *PolyNode) displayExponent() *PolyNode {
	if pn.displayValidate() {
		if pn.exp >= 2 {
			fmt.Printf("^%d", pn.exp)
		}
	}
	return pn
}

func (pn *PolyNode) getExponentStr() string {
	expStr := undefinedString
	if pn.displayValidate() {
		if pn.exp >= 2 {
			expStr = "^" + strconv.Itoa(pn.exp)
		}
	}
	return expStr
}
