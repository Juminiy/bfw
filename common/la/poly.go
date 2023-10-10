package la

import (
	lang2 "bfw/common/lang"
	"errors"
	"fmt"
)

const (
	polyNoSize        int    = 0
	polyInvalidMaxExp int    = -1
	polyInvalidAes    rune   = ' '
	polyDefaultAes    rune   = 'x'
	undefinedString   string = ""
)

var (
	polyIndexOutOfBound = errors.New("poly index is out of bound")
	polyInvalidError    = errors.New("poly is invalid")
	NullPoly            = &Poly{}
)

// Poly
// aes: algebraic expression symbol
// maxExp: max exponent in poly
// size(hidden): maxExp + 1
// assume the poly nodes is in ordered
// assume the poly nodes are in their own position, for example, nodes[idx] = (coe, idx)
// assume the poly idx == exp
// for simplify, we must set the poly nodes[idx].exp = idx
type Poly struct {
	nodes  []*PolyNode
	maxExp int
	aes    rune
}

func (poly *Poly) validate() bool {
	polySize := poly.maxExp + 1
	if polySize == polyNoSize ||
		poly.nodes == nil {
		return false
	}
	return true
}

func (poly *Poly) validateOneIndex(index int) bool {
	if index < 0 ||
		index > poly.maxExp {
		return false
	}
	return true
}

func (poly *Poly) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if !poly.validateOneIndex(index[indexIdx]) {
				return false
			}
		}
	}
	return true
}

// validateOneNode
// none sense
func (poly *Poly) validateOneNode(index int) bool {
	return poly.validateIndex(index) &&
		poly.nodes[index] != nil
}

// validateNode
// none sense
func (poly *Poly) validateNode(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if !poly.validateOneNode(index[indexIdx]) {
				return false
			}
		}
	}
	return true
}

// assign
// assigns the poly nodes for the maxExp+1 length nil value
func (poly *Poly) assign(maxExp int) {
	poly.setMaxExp(maxExp)
	polySize := poly.getSize()
	poly.setValues(make([]*PolyNode, polySize), maxExp, polyDefaultAes)
}

func (poly *Poly) isNull() bool {
	return !poly.validate()
}

func (poly *Poly) null() *Poly {
	return &Poly{}
}

func (poly *Poly) setNull() {
	poly.setValues(nil, polyInvalidMaxExp, polyInvalidAes)
}

func (poly *Poly) setZero() {
	poly.setValues(make([]*PolyNode, polyNoSize), polyInvalidMaxExp, polyDefaultAes)
}

func (poly *Poly) makeCopy() *Poly {
	polyCopy := &Poly{}
	polySize := poly.maxExp + 1
	polyCopy.setValues(make([]*PolyNode, polySize), poly.maxExp, poly.aes)
	for exp := 0; exp < polySize; exp++ {
		polyCopy.setElem(exp, poly.getElem(exp).makeCopy())
	}
	return polyCopy
}

func (poly *Poly) setSelf(p *Poly) {
	if !p.validate() {
		panic(polyInvalidError)
	}
	poly.setValues(p.nodes, p.maxExp, p.aes)
}

func (poly *Poly) setNodes(nodes []*PolyNode) {
	poly.nodes = nodes
}

func (poly *Poly) setMaxExp(maxExp int) {
	poly.maxExp = maxExp
}

func (poly *Poly) setAES(aes rune) {
	poly.aes = aes
}

func (poly *Poly) getSize() int {
	return poly.maxExp + 1
}

func (poly *Poly) setValues(nodes []*PolyNode, maxExp int, aes rune) {
	poly.setNodes(nodes)
	poly.setMaxExp(maxExp)
	poly.setAES(aes)
}

func (poly *Poly) setElem(index int, node *PolyNode) {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	poly.nodes[index] = node
}

func (poly *Poly) setElemValues(index int, node *PolyNode) {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if poly.nodes[index] == nil {
		poly.nodes[index] = ConstructValidPolyNode(index)
	}
	poly.nodes[index].setValues(node.coe, node.exp)
}

func (poly *Poly) setElemExp(index, exp int) {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	poly.nodes[index].exp = exp
}

func (poly *Poly) setElemCoe(index int, coe float64) {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	poly.nodes[index].coe = coe
}

func (poly *Poly) getElem(index int) *PolyNode {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if polyIndexNode := poly.nodes[index]; polyIndexNode != nil {
		poly.setElemExp(index, index)
		return polyIndexNode
	} else {
		return ConstructValidPolyNode(index)
	}
}

func (poly *Poly) getSetNullElem(index int) *PolyNode {
	return poly.getSetElem(index, index, polyNodeInvalidCoefficient)
}

func (poly *Poly) getSetZeroElem(index int) *PolyNode {
	return poly.getSetElem(index, index, polyNodeCoefficientZero)
}

func (poly *Poly) getSetInvalidElem(index, exp int) *PolyNode {
	return poly.getSetElem(index, exp, polyNodeInvalidCoefficient)
}

func (poly *Poly) getSetValidElem(index, exp int) *PolyNode {
	return poly.getSetElem(index, exp, polyNodeCoefficientOne)
}

func (poly *Poly) getSetElem(index, exp int, coe float64) *PolyNode {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if polyIndexNode := poly.nodes[index]; polyIndexNode != nil {
		return polyIndexNode
	} else {
		polyNode := constructPolyNode(coe, exp)
		poly.setElem(index, polyNode)
		return polyNode
	}
}

func (poly *Poly) setElemOne2OneOpt(index int, opt rune, p *Poly) {
	if !p.validate() {
		panic(polyInvalidError)
	}
	poly.getSetInvalidElem(index, index).one2OneOpt(opt, p.getSetInvalidElem(index, index))
}

func (poly *Poly) setGetElemOpt(index, exp int, opt rune, node *PolyNode) *PolyNode {
	return poly.getSetInvalidElem(index, exp).one2OneOpt(opt, node)
}

func (poly *Poly) setElemNull(index int) {
	poly.getElem(index).setNull()
}

func (poly *Poly) setElemZero(index int) {
	poly.getElem(index).setZero()
}

func (poly *Poly) resetMapExp(maxExp int) {
	if maxExp > poly.maxExp {
		appendNodes := make([]*PolyNode, maxExp-poly.maxExp)
		poly.nodes = append(poly.nodes, appendNodes...)
	} else if maxExp < poly.maxExp {
		for idx := maxExp + 1; idx <= poly.maxExp; idx++ {
			poly.setElem(idx, nil)
		}
		poly.nodes = poly.nodes[:maxExp+1]
	}
	poly.maxExp = maxExp
}

func (poly *Poly) Solve(value float64) float64 {
	return 0.0
}

// Equal
// ignore the aes
func (poly *Poly) Equal(p *Poly) bool {
	if !poly.validate() ||
		!p.validate() {
		panic(polyInvalidError)
	}
	polySize, pSize := poly.getSize(), p.getSize()
	if polySize == pSize {
		for idx := 0; idx < polySize; idx++ {
			if !poly.getElem(idx).Equal(p.getElem(idx)) {
				return false
			}
		}
		return true
	}
	return false
}

func (poly *Poly) Construct(real2DArray [][]float64, aes ...rune) *Poly {
	if real2DArray == nil ||
		len(real2DArray) == polyNoSize {
		return poly
	}
	var (
		real2DArraySize int         = len(real2DArray)
		maxExp          int         = 0
		nodes           []*PolyNode = make([]*PolyNode, 0)
	)
	for idx := 0; idx < real2DArraySize; idx++ {
		node := real2DArray[idx]
		if len(node) > 2 {
			maxExp = lang2.MaxInt(maxExp, int(node[1]))
			nodes = append(nodes, ConstructPolyNode(node[0], int(node[1])))
		}
	}
	poly.assign(maxExp)
	poly.setNodeByOrder(nodes...)
	return poly
}

func ConstructPoly(real2DArray [][]float64, aes ...rune) *Poly {
	poly := &Poly{}
	return poly.Construct(real2DArray, aes...)
}

func ConstructNullPoly() *Poly {
	poly := &Poly{}
	poly.setZero()
	return poly
}

func (poly *Poly) setNullNodeByIndexRange(startIndex, endIndex int) {
	polyNode := &PolyNode{}
	for idx := startIndex; idx <= endIndex; idx++ {
		poly.setElem(idx, polyNode.null())
	}
}

func (poly *Poly) setNodeByOrder(node ...*PolyNode) {
	if nodeLen := len(node); nodeLen > 0 {
		maxExp := node[nodeLen-1].exp
		poly.setValues(make([]*PolyNode, maxExp+1), maxExp, polyDefaultAes)
		for idx := 0; idx < nodeLen; idx++ {
			nodeNode := node[idx]
			poly.setElem(nodeNode.exp, nodeNode)
		}
	}
}

func (poly *Poly) sameShape(p *Poly) bool {
	if !poly.validate() ||
		p.validate() {
		panic(polyInvalidError)
	}
	if poly.getSize() != p.getSize() {
		return false
	}
	return true
}

func (poly *Poly) compareTo(p *Poly) bool {
	if !poly.validate() ||
		!p.validate() {
		panic(polyInvalidError)
	}
	return poly.getSize() > p.getSize()
}

func (poly *Poly) swap(p *Poly) {
	polyTemp := &Poly{}
	polyTemp.setSelf(poly)
	poly.setSelf(p)
	p.setSelf(polyTemp)
}

func (poly *Poly) one2OneOpt(opt rune, p *Poly) *Poly {
	if !poly.compareTo(p) {
		poly.swap(p)
	}
	polyPSize := p.getSize()
	for idx := 0; idx < polyPSize; idx++ {
		poly.setElemOne2OneOpt(idx, opt, p)
	}
	return poly
}

func (poly *Poly) AddNode(node *PolyNode) *Poly {
	if node != nil &&
		node.validate() {
		poly.setGetElemOpt(node.exp, node.exp, '+', node)
	}
	return poly
}

// Add Test Pass
func (poly *Poly) Add(p *Poly) *Poly {
	return poly.one2OneOpt('+', p)
}

func (poly *Poly) GetPlus(p *Poly) *Poly {
	polyCopy := poly.makeCopy()
	return polyCopy.Add(p)
}

func (poly *Poly) Sub(p *Poly) *Poly {
	return poly.one2OneOpt('-', p)
}

func (poly *Poly) GetMinus(p *Poly) *Poly {
	polyCopy := poly.makeCopy()
	return polyCopy.Sub(p)
}

func (poly *Poly) Mul(p *Poly) *Poly {
	return poly.simpleMul(p)
}

func (poly *Poly) GetTimes(p *Poly) *Poly {
	polyCopy := poly.makeCopy()
	return polyCopy.Mul(p)
}

func (poly *Poly) MulNode(node *PolyNode) *Poly {
	if node == nil ||
		!node.validate() {
		panic(polyNodeInvalidError)
	}
	var (
		polyMaxExp int       = poly.maxExp
		nodeCopy   *PolyNode = node.makeCopy()
	)
	poly.resetMapExp(polyMaxExp + nodeCopy.exp)
	for idx := polyMaxExp; idx >= polyNodeExponentZero; idx-- {
		if poly.getElem(idx).validateCoefficient() {
			setExp := idx + nodeCopy.exp
			poly.setElemValues(setExp, poly.setGetElemOpt(idx, setExp, '*', nodeCopy))
			poly.setElemZero(idx)
		}
	}
	return poly
}

// judgeSparse
// if sparse
func (poly *Poly) judgeSparse(funcPtr func(params ...interface{}) bool) bool {
	return funcPtr()
}

func (poly *Poly) judgeMulAlgo() bool {
	return false
}

func (poly *Poly) simpleMul(p *Poly) *Poly {
	if !poly.validate() ||
		!p.validate() {
		panic(polyInvalidError)
	}
	var (
		pSize      int   = p.getSize()
		polyResult *Poly = &Poly{}
		polyCopy   *Poly
	)
	polyResult.assign(poly.maxExp + p.maxExp)
	for pIdx := 0; pIdx < pSize; pIdx++ {
		polyCopy = poly.makeCopy()
		polyResult.Add(polyCopy.MulNode(p.getSetInvalidElem(pIdx, pIdx)))
	}
	return poly
}

func (poly *Poly) fftMul(p *Poly) *Poly {
	return poly
}

// Div by Factoring
func (poly *Poly) Div(p *Poly) *Poly {
	return poly
}

func (poly *Poly) Value(value float64) float64 {
	var (
		polyValue float64 = 0.0
	)
	for idx := 0; idx < poly.getSize(); idx++ {
		polyValue += poly.getElem(idx).Value(value)
	}
	return polyValue
}

// Integral Test Pass
func (poly *Poly) Integral() *Poly {
	var (
		polyMaxExp int
	)
	if poly.validate() {
		polyMaxExp = poly.maxExp
		poly.nodes = append(poly.nodes, poly.getElem(polyMaxExp).Integral().makeCopy())
		for idx := polyMaxExp; idx > polyNodeExponentZero; idx-- {
			poly.setElemValues(idx, poly.getElem(idx-1).Integral())
		}
		poly.getElem(polyNodeExponentZero).setZero()
		poly.maxExp++
	}
	return poly
}

// Derivative Test Pass
func (poly *Poly) Derivative() *Poly {
	var (
		polyMaxExp int
	)
	if poly.validate() {
		polyMaxExp = poly.maxExp
		for idx := polyNodeExponentZero; idx < polyMaxExp; idx++ {
			poly.setElemValues(idx, poly.getElem(idx+1).Derivative())
		}
		poly.resetMapExp(polyMaxExp - 1)
	}
	return poly
}

func (poly *Poly) Factoring() *PolyFactors {
	return nil
}

func (poly *Poly) Normal() *Poly {
	return poly.Discard(true).
		Merge(true).
		Sort(true)
}

func (poly *Poly) Regular() *Poly {
	return poly.Discard().Merge().Sort()
}

func (poly *Poly) Discard(predict ...bool) *Poly {
	if len(predict) > 0 && predict[0] {
		return poly.removeZeroNullInvalidNode()
	}
	return poly
}

func (poly *Poly) removeZeroNullInvalidNode() *Poly {
	return poly
}

func (poly *Poly) Sort(predict ...bool) *Poly {
	if len(predict) > 0 && predict[0] {
		return poly.sortNodeByExponent()
	}
	return poly
}

func (poly *Poly) sortNodeByExponent() *Poly {
	return poly
}

func (poly *Poly) Merge(predict ...bool) *Poly {
	if len(predict) > 0 && predict[0] {
		return poly.mergeNodeBySameExponent()
	}
	return poly
}

// mergeNodes
// merge the nodes by same exp
func (poly *Poly) mergeNodeBySameExponent() *Poly {
	return poly
}

func (poly *Poly) Display(precisionBit ...int) *Poly {
	var (
		polySize     int       = poly.getSize()
		preDisplayed bool      = false
		node         *PolyNode = poly.getElem(0)
	)
	if nodeStr := node.ToString(poly.aes, precisionBit...); !lang2.StringIsNull(nodeStr) {
		fmt.Print(nodeStr)
		preDisplayed = true
	}
	for idx := 1; idx < polySize; idx++ {
		node = poly.getElem(idx)
		if nodeStr := node.ToString(poly.aes, precisionBit...); !lang2.StringIsNull(nodeStr) {
			if !lang2.Float64StringIsMinus(nodeStr) &&
				preDisplayed {
				fmt.Print("+")
			}
			fmt.Print(nodeStr)
			preDisplayed = true
		}
	}
	fmt.Println()
	return poly
}
