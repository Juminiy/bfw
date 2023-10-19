package cal

import (
	lang "bfw/common/lang"
	"errors"
	"fmt"
)

const (
	polyNoSize            int    = 0
	polyInvalidMaxExp     int    = -1
	polyInvalidAes        rune   = ' '
	polyDefaultAes        rune   = 'x'
	polyStringDefaultZero string = "0"
	undefinedString       string = ""
)

var (
	polyIndexOutOfBound            = errors.New("poly index is out of bound")
	polyInvalidError               = errors.New("poly is invalid")
	polyEquationNoSolution         = errors.New("poly equation has no solution")
	polyEquationSolutionNotDevelop = errors.New("poly equation solution has not developed")
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

func ConstructPoly(real2DArray [][]float64, aes ...rune) *Poly {
	poly := &Poly{}
	return poly.Construct(real2DArray, aes...)
}

func ConstructNullPoly() *Poly {
	poly := &Poly{}
	poly.setZero()
	return poly
}

func PolyChainedAdd(poly ...*Poly) *Poly {
	polyRes := ConstructPolyNode(0.0, 0).Poly(eigenPolyMatrixDefaultAES)
	if polyLen := len(poly); polyLen > 0 {
		for polyIdx := 0; polyIdx < polyLen; polyIdx++ {
			if poly[polyIdx] != nil {
				polyRes.Add(poly[polyIdx])
			}
		}
	}
	return polyRes
}

func PolyChainedSub(destPoly *Poly, poly ...*Poly) *Poly {
	polyRes := destPoly.makeCopy()
	if polyLen := len(poly); polyLen > 0 {
		for polyIdx := 0; polyIdx < polyLen; polyIdx++ {
			if poly[polyIdx] != nil {
				polyRes.Sub(poly[polyIdx])
			}
		}
	}
	return polyRes
}

func PolyChainedMulV2(poly ...*Poly) *Poly {
	if polyLen := len(poly); polyLen > 0 {
		polyRes := ConstructPolyNode(1.0, 0).Poly(eigenPolyMatrixDefaultAES)
		for polyIdx := 0; polyIdx < polyLen; polyIdx++ {
			if poly[polyIdx] == nil ||
				!poly[polyIdx].validate() ||
				!poly[polyIdx].validateSlice() {
				return nil
			}
			polyRes = polyRes.Mul(poly[polyIdx])
		}
		return polyRes
	}
	return nil
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
		if len(node) >= 2 {
			maxExp = lang.MaxInt(maxExp, int(node[1]))
			nodes = append(nodes, ConstructPolyNode(node[0], int(node[1])))
		}
	}
	poly.assign(maxExp)
	poly.setNodeByOrder(nodes...)
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

// Poly validate should be redefined
func (poly *Poly) validate() bool {
	if poly.getSize() == polyNoSize ||
		poly.nodes == nil {
		return false
	}
	return true
}

func (poly *Poly) validateSlice() bool {
	for exp := poly.maxExp; exp >= polyNodeExponentZero; exp-- {
		if !poly.getElem(exp).validate() {
			return false
		}
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

// Equal
// ignore the aes
func (poly *Poly) Equal(p *Poly) bool {
	if !poly.validate() ||
		!p.validate() {
		panic(polyInvalidError)
	}
	if !poly.sameShape(p) {
		return false
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

func (poly *Poly) sameShape(p *Poly) bool {
	if !poly.validate() ||
		!p.validate() {
		panic(polyInvalidError)
	}
	if poly.getSize() != p.getSize() {
		return false
	}
	return true
}

func (poly *Poly) compareTo(p *Poly) bool {
	if p == nil ||
		!poly.validate() ||
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

// assign
// assigns the poly nodes for the maxExp+1 length nil value
func (poly *Poly) assign(maxExp int) {
	poly.setMaxExp(maxExp)
	polySize := poly.getSize()
	poly.setValues(make([]*PolyNode, polySize), maxExp, polyDefaultAes)
	for idx := 0; idx < polySize; idx++ {
		poly.getSetInvalidElem(idx, idx)
	}
}

func (poly *Poly) getSize() int {
	return poly.maxExp + 1
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
	polySize := poly.getSize()
	polyCopy.setValues(make([]*PolyNode, polySize), poly.maxExp, poly.aes)
	for exp := 0; exp < polySize; exp++ {
		polyCopy.setElem(exp, poly.getSetElem(exp, exp).makeCopy())
	}
	return polyCopy
}

func (poly *Poly) setSelf(p *Poly) {
	if !p.validate() {
		panic(polyInvalidError)
	}
	poly.setValues(p.nodes, p.maxExp, p.aes)
}

func (poly *Poly) setValues(nodes []*PolyNode, maxExp int, aes rune) {
	poly.setNodes(nodes)
	poly.setMaxExp(maxExp)
	poly.setAES(aes)
}

func (poly *Poly) setNodes(nodes []*PolyNode) {
	poly.nodes = nodes
}

func (poly *Poly) setMaxExp(maxExp int) {
	poly.maxExp = maxExp
}

func (poly *Poly) resetMapExp(maxExp int) {
	if maxExp <= polyInvalidMaxExp {
		poly.setNull()
	}
	if maxExp > poly.maxExp {
		appendNodes := make([]*PolyNode, maxExp-poly.maxExp)
		poly.nodes = append(poly.nodes, appendNodes...)
	} else if maxExp < poly.maxExp {
		for idx := maxExp + 1; idx <= poly.maxExp; idx++ {
			poly.setElem(idx, nil)
		}
		poly.nodes = poly.nodes[:maxExp+1]
	}
	poly.setMaxExp(maxExp)
}

func (poly *Poly) validateMaxExp() {
	validMaxExp := poly.maxExp
	for exp := poly.maxExp; exp >= polyNodeExponentZero; exp-- {
		elem := poly.getSetInvalidElem(exp, exp)
		if elem.validateCoefficient() {
			validMaxExp = exp
			break
		}
	}
	poly.resetMapExp(validMaxExp)
}

func (poly *Poly) setAES(aes rune) {
	poly.aes = aes
}

func (poly *Poly) setElem(index int, node *PolyNode) {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	poly.nodes[index] = node
}

func (poly *Poly) setElemValues(index int, node *PolyNode) {
	if node != nil {
		poly.setElemCoe(index, node.coe)
		poly.setElemExp(index, node.exp)
	}
}

func (poly *Poly) setElemCoe(index int, coe float64) {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if poly.nodes[index] == nil {
		poly.nodes[index] = ConstructValidPolyNode(index)
	}
	poly.nodes[index].setValues(coe, index)
}

func (poly *Poly) setElemExp(index, exp int) {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if poly.nodes[index] == nil {
		poly.nodes[index] = ConstructValidPolyNode(exp)
	}
	poly.nodes[index].setExp(exp)
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

func (poly *Poly) getSetUnitElem(index, exp int) *PolyNode {
	return poly.getSetElem(index, exp, polyNodeCoefficientOne)
}

func (poly *Poly) getSetElem(index, exp int, coe ...float64) *PolyNode {
	if !poly.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if polyIndexNode := poly.nodes[index]; polyIndexNode != nil {
		poly.nodes[index].setExp(exp)
		return polyIndexNode
	} else {
		destCoe := polyNodeCoefficientZero
		if len(coe) > 0 {
			destCoe = coe[0]
		}
		newPolyNode := constructPolyNode(destCoe, exp)
		poly.setElem(index, newPolyNode)
		return newPolyNode
	}
}

func (poly *Poly) setElemByOne2OneOpt(index int, opt rune, p *Poly) {
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

// one2OneOpt
// '+' Test Pass
// '-' Test Error
func (poly *Poly) one2OneOpt(opt rune, p *Poly) *Poly {
	if p == nil ||
		!poly.validate() ||
		!p.validate() {
		return poly
	}
	polyDestMaxExp := lang.MaxInt(poly.maxExp, p.maxExp)
	poly.resetMapExp(polyDestMaxExp)
	p.resetMapExp(polyDestMaxExp)
	for idx := 0; idx <= polyDestMaxExp; idx++ {
		poly.setElemByOne2OneOpt(idx, opt, p)
	}
	return poly
}

func (poly *Poly) AddNode(node *PolyNode) *Poly {
	if node != nil {
		poly.setGetElemOpt(node.exp, node.exp, '+', node)
	}
	return poly
}

func (poly *Poly) SubNode(node *PolyNode) *Poly {
	if node != nil {
		poly.setGetElemOpt(node.exp, node.exp, '-', node)
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
	//return poly.simpleMul(p)
	return poly.simpleMulV2(p)
}

func (poly *Poly) GetTimes(p *Poly) *Poly {
	polyCopy := poly.makeCopy()
	return polyCopy.Mul(p)
}

// simpleMul
// must use heap to accelerate the calculation
func (poly *Poly) simpleMul(p *Poly) *Poly {
	if !poly.validate() ||
		!p.validate() {
		panic(polyInvalidError)
	}
	var (
		pSize      int   = p.getSize()
		polyResult *Poly = &Poly{}
		polyCopy   *Poly
		polyPNode  *PolyNode
	)
	polyResult.assign(poly.maxExp + p.maxExp)
	for pIdx := 0; pIdx < pSize; pIdx++ {
		polyCopy = poly.makeCopy()
		polyPNode = p.getSetInvalidElem(pIdx, pIdx)
		if polyPNode.validate() {

			//debugSimpleMul(polyPNode, polyCopy)

			polyCopy.MulNode(polyPNode)
			polyResult.Add(polyCopy)
		}
	}
	return polyResult
}

func debugSimpleMul(polyPNode *PolyNode, polyCopy *Poly) {
	polyCopy.Display(false, 0)
	polyCopy.MulNode(polyPNode)
	fmt.Printf(" * ")
	polyPNode.Display(false, polyDefaultAes)
	fmt.Printf(" = ")
	polyCopy.Display(false, 0)
	fmt.Println()
}

// use heap to accelerate the polynomial multiply
// has not changed poly itself
func (poly *Poly) simpleMulV2(p *Poly) *Poly {
	if p == nil ||
		!poly.validate() ||
		!p.validate() {
		return nil
	}
	var (
		pSize int = p.getSize()
		//polyResult        *Poly = &Poly{}
		polyCopy          *Poly
		polyPNode         *PolyNode
		polyMiddleResults []*Poly
	)
	//polyResult.assign(poly.maxExp + p.maxExp)
	polyMiddleResults = make([]*Poly, pSize)

	// get polyMiddleResults
	for pIdx := 0; pIdx < pSize; pIdx++ {
		polyCopy = poly.makeCopy()
		polyPNode = p.getSetInvalidElem(pIdx, pIdx)
		if polyPNode.validate() {
			polyCopy.MulNode(polyPNode)
			polyMiddleResults[pIdx] = polyCopy
		}
	}

	// merge polyMiddleResults
	return binaryMergePoly(polyMiddleResults)
}

func binaryMergePoly(polys []*Poly) *Poly {
	polysLen := len(polys)
	if polysLen == 0 {
		return nil
	}
	if polysLen == 1 {
		return polys[0]
	}
	part1MergedPolys := binaryMergePoly(polys[:(polysLen >> 1)])
	part2MergedPolys := binaryMergePoly(polys[(polysLen >> 1):])
	if part1MergedPolys == nil {
		return part2MergedPolys
	} else {
		return part1MergedPolys.Add(part2MergedPolys)
	}
}

func (poly *Poly) fftMul(p *Poly) *Poly {
	return poly
}

// judgeSparse
// if sparse
func (poly *Poly) judgeSparse(funcPtr func(...interface{}) bool) bool {
	return funcPtr()
}

func (poly *Poly) judgeMulAlgo() bool {
	return false
}

func (poly *Poly) MulNode(node *PolyNode) *Poly {
	if node == nil {
		panic(polyNodeInvalidError)
	}
	var (
		polyMaxExp int       = poly.maxExp
		nodeCopy   *PolyNode = node.makeCopy()
	)
	if node.isConstant() {
		for idx := 0; idx <= polyMaxExp; idx++ {
			poly.setGetElemOpt(idx, idx, '*', nodeCopy)
		}
	} else {
		poly.resetMapExp(polyMaxExp + nodeCopy.exp)
		for idx := polyMaxExp; idx >= polyNodeExponentZero; idx-- {
			if poly.getSetInvalidElem(idx, idx).validateCoefficient() {
				setExp := idx + nodeCopy.exp
				poly.setElemValues(setExp, poly.setGetElemOpt(idx, setExp, '*', nodeCopy))
				poly.setElemZero(idx)
			}
		}
	}

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
		poly.nodes = append(poly.nodes, poly.getSetElem(polyMaxExp, polyMaxExp).Integral().makeCopy())
		for idx := polyMaxExp; idx > polyNodeExponentZero; idx-- {
			poly.setElemValues(idx, poly.getSetElem(idx-1, idx-1).Integral())
		}
		poly.getElem(polyNodeExponentZero).setZero()
		poly.setMaxExp(poly.maxExp + 1)
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
			poly.setElemValues(idx, poly.getSetElem(idx+1, idx+1).Derivative())
		}
		poly.resetMapExp(polyMaxExp - 1)
	}
	return poly
}

func (poly *Poly) SolvePoly(p *Poly) *Solution {
	return poly.Sub(p).Solve()
}

func (poly *Poly) SolveValue(value float64) *Solution {
	return poly.SubNode(ConstructPolyNode(value, polyNodeExponentZero)).Solve()
}

func (poly *Poly) getElemQuotient(indexI, indexJ int) *PolyNode {
	if !poly.validateIndex(indexI, indexJ) {
		panic(polyIndexOutOfBound)
	}
	return poly.getElem(indexI).GetQuotient(poly.getElem(indexJ))
}

func (poly *Poly) getElemQuotientCoe(indexI, indexJ int) float64 {
	return poly.getElemQuotient(indexI, indexJ).coe
}

// Solve
// poly equation 0.0
func (poly *Poly) Solve() *Solution {
	s := &Solution{}
	poly.validateMaxExp()
	switch poly.maxExp {
	case 0:
		{
			panic(polyEquationNoSolution)
		}
	case 1:
		{
			s.setValues(make([]complex128, 1), 1)
			exp0Value := poly.getSetElem(0, 0).coe
			exp1Coe := poly.getSetElem(1, 1).coe
			if lang.EqualFloat64Zero(exp1Coe) {
				panic(polyEquationNoSolution)
			}
			s.setElem(0, complex(exp0Value*-1.0/exp1Coe, 0.0))
		}
	case 2:
		{
			s.setValues(make([]complex128, 2), 2)
			exp0Coe := poly.getSetElem(0, 0).coe
			exp1Coe := poly.getSetElem(1, 1).coe
			exp2Coe := poly.getSetElem(2, 2).coe
			solution1, solution2, validateSolve :=
				lang.SolveQuadraticEquationOfOneVariable(exp2Coe, exp1Coe, exp0Coe)
			if !validateSolve {
				panic(polyEquationNoSolution)
			}
			s.setElem(0, solution1)
			s.setElem(1, solution2)
		}
	case 3:
		{
			s.setValues(make([]complex128, 3), 3)
			exp0Coe := poly.getSetElem(0, 0).coe
			exp1Coe := poly.getSetElem(1, 1).coe
			exp2Coe := poly.getSetElem(2, 2).coe
			exp3Coe := poly.getSetElem(3, 3).coe
			solution1, solution2, solution3, validateSolve :=
				lang.SolveCubicEquationOfOneVariableBySJ(exp3Coe, exp2Coe, exp1Coe, exp0Coe)
			if !validateSolve {
				panic(polyEquationNoSolution)
			}
			s.setElem(0, solution1)
			s.setElem(1, solution2)
			s.setElem(2, solution3)
		}
		// TODO: more power equation should be developed
	default:
		{
			panic(polyEquationSolutionNotDevelop)
		}
	}
	return s
}

func (poly *Poly) Factoring() *PolyFactors {
	solution := poly.Solve()
	return solution.PolyFactors()
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

// removeZeroNullInvalidNode
// get rid of invalid nodes
func (poly *Poly) removeZeroNullInvalidNode() *Poly {
	return poly
}

func (poly *Poly) Sort(predict ...bool) *Poly {
	if len(predict) > 0 && predict[0] {
		return poly.sortNodeByExponent()
	}
	return poly
}

// sortNodeByExponent
// sort the nodes by exp asc
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

func (poly *Poly) getValidNode() []*PolyNode {
	validNode := make([]*PolyNode, 0)
	for exp := 0; exp <= poly.maxExp; exp++ {
		if poly.getElem(exp).validate() {
			validNode = append(validNode, poly.getElem(exp).makeCopy())
		}
	}
	return validNode
}

func (poly *Poly) PolyNode() *PolyNode {
	if validNode := poly.getValidNode(); len(validNode) == 1 {
		return validNode[0]
	}
	return &PolyNode{}
}

func (poly *Poly) ToString(precisionBit ...int) string {
	if !poly.validate() {
		return lang.Float64ToString(0, precisionBit...)
	}
	var (
		polyString   string    = undefinedString
		polySize     int       = poly.getSize()
		preDisplayed bool      = false
		node         *PolyNode = poly.getElem(0)
	)
	if nodeStr := node.ToString(poly.aes, precisionBit...); !lang.StringIsNull(nodeStr) {
		polyString += nodeStr
		preDisplayed = true
	}
	for idx := 1; idx < polySize; idx++ {
		node = poly.getElem(idx)
		if nodeStr := node.ToString(poly.aes, precisionBit...); !lang.StringIsNull(nodeStr) {
			if !lang.Float64StringIsMinus(nodeStr) &&
				preDisplayed {
				polyString += "+"
			}
			polyString += nodeStr
			preDisplayed = true
		}
	}
	return polyString
}

func (poly *Poly) Display(isPrintln bool, precisionBit ...int) *Poly {
	polyString := poly.ToString(precisionBit...)
	if polyString == undefinedString {
		polyString = lang.Float64ToString(0, precisionBit...)
	}
	fmt.Printf(polyString)
	if isPrintln {
		fmt.Println()
	}
	return poly
}

func (poly *Poly) ToStringV2(reverse bool, precisionBit ...int) string {
	if !reverse {
		return poly.ToString(precisionBit...)
	}
	if !poly.validate() {
		return lang.Float64ToString(0, precisionBit...)
	}
	var (
		polyString   string    = undefinedString
		preDisplayed bool      = false
		node         *PolyNode = poly.getElem(poly.maxExp)
	)

	if nodeStr := node.ToString(poly.aes, precisionBit...); !lang.StringIsNull(nodeStr) {
		polyString += nodeStr
		preDisplayed = true
	}
	for idx := poly.maxExp - 1; idx >= polyNodeExponentZero; idx-- {
		node = poly.getElem(idx)
		if nodeStr := node.ToString(poly.aes, precisionBit...); !lang.StringIsNull(nodeStr) {
			if !lang.Float64StringIsMinus(nodeStr) &&
				preDisplayed {
				polyString += "+"
			}
			polyString += nodeStr
			preDisplayed = true
		}
	}
	return polyString
}

func (poly *Poly) DisplayV2(isPrintln bool, reverse bool, precisionBit ...int) *Poly {
	polyString := poly.ToStringV2(reverse, precisionBit...)
	if polyString == undefinedString {
		polyString = lang.Float64ToString(0, precisionBit...)
	}
	fmt.Printf(polyString)
	if isPrintln {
		fmt.Println()
	}
	return poly
}
