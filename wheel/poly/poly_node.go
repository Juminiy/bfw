package poly

import (
	"bfw/wheel/char"
	"bfw/wheel/lang"
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
)

type Node struct {
	coe float64
	exp int
}

func ConstructPolyNode(coe float64, exp int) *Node {
	return constructPolyNode(coe, exp)
}

func ConstructNullPolyNode() *Node {
	return constructPolyNode(polyNodeInvalidCoefficient, polyNodeInvalidExponent, true)
}

func ConstructValidPolyNode(exp int) *Node {
	return constructPolyNode(polyNodeCoefficientZero, exp)
}

func constructPolyNode(coe float64, exp int, null ...bool) *Node {
	if len(null) != 0 && null[0] {
		return &Node{}
	}
	pn := &Node{}
	return pn.Construct(coe, exp)
}

func MakeNullPolyNode() *Node {
	return new(Node)
}

func (pn *Node) Construct(coe float64, exp int) *Node {
	pn.setValues(coe, exp)
	return pn
}

func (pn *Node) validate() bool {
	return pn.validateCoefficient() &&
		pn.validateExponent()
}

func (pn *Node) validateCoefficient() bool {
	return !lang.EqualFloat64ByAccuracy(polyNodeInvalidCoefficient, pn.coe)
}

func (pn *Node) validateExponent() bool {
	return pn.exp != polyNodeInvalidExponent
}

func (pn *Node) makeCopy() *Node {
	pnt := &Node{}
	pnt.setValues(pn.coe, pn.exp)
	return pnt
}

func (pn *Node) swap(pnt *Node) {
	pnTemp := &Node{}
	pnTemp.setSelf(pn)
	pn.setSelf(pnt)
	pnt.setSelf(pnTemp)
}

func (pn *Node) setSelf(pnt *Node) {
	pn.setValues(pnt.coe, pnt.exp)
}

func (pn *Node) setValues(coe float64, exp int) {
	pn.setCoe(coe)
	pn.setExp(exp)
}

func (pn *Node) setCoe(coe float64) {
	pn.coe = coe
}

func (pn *Node) setExp(exp int) {
	pn.exp = exp
}

func (pn *Node) null() *Node {
	return &Node{}
}

func (pn *Node) setNull() {
	pn.setValues(polyNodeInvalidCoefficient, polyNodeInvalidExponent)
}

func (pn *Node) setZero() {
	pn.setCoeZero()
}

func (pn *Node) setCoeZero() {
	pn.coe = polyNodeCoefficientZero
}

func (pn *Node) setOpt(opt rune, pnt *Node) {
	pn.one2OneOpt(opt, pnt)
}

func (pn *Node) isZero() bool {
	return pn.coe == polyNodeCoefficientZero
}

func (pn *Node) isConstant() bool {
	return pn.exp == polyNodeExponentZero
}
func (pn *Node) one2OneOptNotPanic(opt rune, pnt *Node) bool {
	if pnt == nil {
		return false
	}
	switch opt {
	case '+':
		{
			if pn.exp != pnt.exp {
				return false
			}
			pn.coe += pnt.coe
		}
	case '-':
		{
			if pn.exp != pnt.exp {
				return false
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
				return false
			}
			pn.coe /= pnt.coe
			pn.exp -= pnt.exp
		}
	default:
		{

		}
	}
	return true
}

func (pn *Node) one2OneOpt(opt rune, pnt *Node) *Node {
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

func (pn *Node) Equal(pnt *Node) bool {
	if pnt == nil {
		return false
	}
	if pn.exp == pnt.exp &&
		lang.EqualFloat64ByAccuracy(pn.coe, pnt.coe) {
		return true
	}
	return false
}

func (pn *Node) Add(pnt *Node) *Node {
	return pn.one2OneOpt('+', pnt)
}

func (pn *Node) GetPlus(pnt *Node) *Node {
	pnCopy := pn.makeCopy()
	return pnCopy.Add(pnt)
}

func (pn *Node) Sub(pnt *Node) *Node {
	return pn.one2OneOpt('-', pnt)
}

func (pn *Node) GetMinus(pnt *Node) *Node {
	pnCopy := pn.makeCopy()
	return pnCopy.Sub(pnt)
}

func (pn *Node) Mul(pnt *Node) *Node {
	return pn.one2OneOpt('*', pnt)
}

func (pn *Node) GetTimes(pnt *Node) *Node {
	pnCopy := pn.makeCopy()
	return pnCopy.Mul(pnt)
}

func (pn *Node) Div(pnt *Node) *Node {
	return pn.one2OneOpt('/', pnt)
}

func (pn *Node) GetQuotient(pnt *Node) *Node {
	pnCopy := pn.makeCopy()
	return pnCopy.Div(pnt)
}

func (pn *Node) Exp(exp int) *Node {
	pn.coe = math.Pow(pn.coe, float64(exp))
	pn.exp *= exp
	return pn
}

func (pn *Node) Sqrt(exp int) *Node {
	return pn.Exp(1.0 / exp)
}

func (pn *Node) Value(value float64) float64 {
	return pn.coe * math.Pow(value, float64(pn.exp))
}

func (pn *Node) Integral() *Node {
	if pn.validate() {
		if !pn.isZero() {
			pn.exp++
			pn.coe /= float64(pn.exp)
		}
	}
	return pn
}

func (pn *Node) Derivative() *Node {
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

func (pn *Node) convertToPoly(aes ...rune) *Poly {
	poly := &Poly{}
	AES := polyDefaultAes
	if len(aes) > 0 {
		AES = aes[0]
	}
	poly.setValues(make([]*Node, pn.exp+1), pn.exp, AES)
	poly.setElem(pn.exp, pn.makeCopy())
	return poly
}

func (pn *Node) Poly(aes ...rune) *Poly {
	return pn.convertToPoly(aes...)
}

// Display
// print the poly node
func (pn *Node) Display(isPrintln bool, aes rune) *Node {
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

func (pn *Node) CanDisplay(aes rune) bool {
	return pn.ToString(aes) != undefinedString
}

// ToString
func (pn *Node) ToString(aes rune, precision ...int) string {
	return pn.getCoefficientStr(precision...) + pn.getAESStr(aes) + pn.getExponentStr()
}

func (pn *Node) displayValidate() bool {
	return pn.validateCoefficient() &&
		pn.validateExponent()
}

func (pn *Node) displayCoefficient() *Node {
	if pn.displayValidate() {
		fmt.Printf("%.5v", pn.coe)
	}
	return pn
}

func (pn *Node) getCoefficientStr(precision ...int) string {
	coeStr := undefinedString
	if pn.displayValidate() {
		return lang.Float64ToString(pn.coe, precision...)
	}
	return coeStr
}

func (pn *Node) displayAES(aes rune) *Node {
	if pn.displayValidate() {
		if pn.exp != polyNodeExponentZero {
			fmt.Printf("%c", aes)
		}
	}
	return pn
}

func (pn *Node) getAESStr(aes rune) string {
	aesStr := undefinedString
	if pn.displayValidate() {
		if pn.exp != polyNodeExponentZero {
			aesStr = string(aes)
		}
	}
	return aesStr
}

// displayExponent
// should refer to numpy print method
func (pn *Node) displayExponent() *Node {
	if pn.displayValidate() {
		if pn.exp >= 2 {
			//fmt.Printf("^%v", pn.exp)
			fmt.Printf("%s", char.GetExponent(strconv.Itoa(pn.exp)))
		}
	}
	return pn
}

// getExponentStr
// should refer to numpy print method
func (pn *Node) getExponentStr() string {
	expStr := undefinedString
	if pn.displayValidate() {
		if pn.exp >= 2 {
			expStr = "^" + strconv.Itoa(pn.exp)
		}
	}
	return expStr
}
