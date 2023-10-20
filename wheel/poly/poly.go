package poly

import (
	"bfw/wheel/lang"
	"errors"
	"fmt"
)

const (
	polyNoSize           int    = 0
	polyInvalidMaxExp    int    = -1
	polyInvalidAes       rune   = ' '
	polyDefaultAes       rune   = 'x'
	polyLambdaAes        rune   = 'λ'
	polyDefaultAesString string = "x"
	polyLambdaAesString  string = "λ"
	undefinedString      string = ""
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
	nodes  []*Node
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

func ValidatePoly(poly ...*Poly) bool {
	for idx := 0; idx < len(poly); idx++ {
		if poly[idx] == nil ||
			!poly[idx].validate() {
			return false
		}
	}
	return true
}

func ChainedAdd(poly ...*Poly) *Poly {
	polyRes := ConstructPolyNode(0.0, 0).Poly(polyDefaultAes)
	if polyLen := len(poly); polyLen > 0 {
		for polyIdx := 0; polyIdx < polyLen; polyIdx++ {
			if poly[polyIdx] != nil {
				polyRes.Add(poly[polyIdx])
			}
		}
	}
	return polyRes
}

func ChainedSub(destPoly *Poly, poly ...*Poly) *Poly {
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

func ChainedMul(poly ...*Poly) *Poly {
	if polyLen := len(poly); polyLen > 0 {
		polyRes := ConstructPolyNode(1.0, 0).Poly(polyDefaultAes)
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

func (p *Poly) Construct(real2DArray [][]float64, aes ...rune) *Poly {
	if real2DArray == nil ||
		len(real2DArray) == polyNoSize {
		return p
	}
	var (
		real2DArraySize int     = len(real2DArray)
		maxExp          int     = 0
		nodes           []*Node = make([]*Node, 0)
	)
	for idx := 0; idx < real2DArraySize; idx++ {
		node := real2DArray[idx]
		if len(node) >= 2 {
			maxExp = lang.MaxInt(maxExp, int(node[1]))
			nodes = append(nodes, ConstructPolyNode(node[0], int(node[1])))
		}
	}
	p.assign(maxExp)
	p.setNodeByOrder(nodes...)
	return p
}

func (p *Poly) setNullNodeByIndexRange(startIndex, endIndex int) {
	polyNode := &Node{}
	for idx := startIndex; idx <= endIndex; idx++ {
		p.setElem(idx, polyNode.null())
	}
}

func (p *Poly) setNodeByOrder(node ...*Node) {
	if nodeLen := len(node); nodeLen > 0 {
		maxExp := node[nodeLen-1].exp
		p.setValues(make([]*Node, maxExp+1), maxExp, polyDefaultAes)
		for idx := 0; idx < nodeLen; idx++ {
			nodeNode := node[idx]
			p.setElem(nodeNode.exp, nodeNode)
		}
	}
}

// Poly validate should be redefined
func (p *Poly) validate() bool {
	if p.getSize() == polyNoSize ||
		p.nodes == nil {
		return false
	}
	return true
}

func (p *Poly) validateSlice() bool {
	for exp := p.maxExp; exp >= polyNodeExponentZero; exp-- {
		if !p.getElem(exp).validate() {
			return false
		}
	}
	return true
}

func (p *Poly) validateOneIndex(index int) bool {
	if index < 0 ||
		index > p.maxExp {
		return false
	}
	return true
}

func (p *Poly) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if !p.validateOneIndex(index[indexIdx]) {
				return false
			}
		}
	}
	return true
}

// validateOneNode
// none sense
func (p *Poly) validateOneNode(index int) bool {
	return p.validateIndex(index) &&
		p.nodes[index] != nil
}

// validateNode
// none sense
func (p *Poly) validateNode(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		for indexIdx := 0; indexIdx < indexLen; indexIdx++ {
			if !p.validateOneNode(index[indexIdx]) {
				return false
			}
		}
	}
	return true
}

// Equal
// ignore the aes
func (p *Poly) Equal(pt *Poly) bool {
	if !p.sameShape(pt) {
		return false
	}
	for idx := 0; idx < p.getSize(); idx++ {
		if !p.getElem(idx).Equal(pt.getElem(idx)) {
			return false
		}
	}
	return true
}

func (p *Poly) sameShape(pt *Poly) bool {
	if !ValidatePoly(p, pt) {
		panic(polyInvalidError)
	}
	return p.getSize() == pt.getSize()
}

func (p *Poly) compareTo(pt *Poly) bool {
	if !ValidatePoly(p, pt) {
		panic(polyInvalidError)
	}
	return p.getSize() > pt.getSize()
}

func (p *Poly) swap(pt *Poly) {
	polyTemp := &Poly{}
	polyTemp.setSelf(p)
	p.setSelf(pt)
	pt.setSelf(polyTemp)
}

// assign
// assigns the poly nodes for the maxExp+1 length nil value
func (p *Poly) assign(maxExp int) {
	p.setMaxExp(maxExp)
	polySize := p.getSize()
	p.setValues(make([]*Node, polySize), maxExp, polyDefaultAes)
	for idx := 0; idx < polySize; idx++ {
		p.getSetInvalidElem(idx, idx)
	}
}

func (p *Poly) getSize() int {
	return p.maxExp + 1
}

func (p *Poly) isNull() bool {
	return !p.validate()
}

func (p *Poly) null() *Poly {
	return &Poly{}
}

func (p *Poly) setNull() {
	p.setValues(nil, polyInvalidMaxExp, polyInvalidAes)
}

func (p *Poly) setZero() {
	p.setValues(make([]*Node, polyNoSize), polyInvalidMaxExp, polyDefaultAes)
}

func (p *Poly) makeCopy() *Poly {
	polyCopy := &Poly{}
	polySize := p.getSize()
	polyCopy.setValues(make([]*Node, polySize), p.maxExp, p.aes)
	for exp := 0; exp < polySize; exp++ {
		polyCopy.setElem(exp, p.getSetElem(exp, exp).makeCopy())
	}
	return polyCopy
}

func (p *Poly) setSelf(pt *Poly) {
	if !pt.validate() {
		panic(polyInvalidError)
	}
	p.setValues(pt.nodes, pt.maxExp, pt.aes)
}

func (p *Poly) setValues(nodes []*Node, maxExp int, aes rune) {
	p.setNodes(nodes)
	p.setMaxExp(maxExp)
	p.setAES(aes)
}

func (p *Poly) setNodes(nodes []*Node) {
	p.nodes = nodes
}

func (p *Poly) setMaxExp(maxExp int) {
	p.maxExp = maxExp
}

func (p *Poly) resetMapExp(maxExp int) {
	if maxExp <= polyInvalidMaxExp {
		p.setNull()
	}
	if maxExp > p.maxExp {
		appendNodes := make([]*Node, maxExp-p.maxExp)
		p.nodes = append(p.nodes, appendNodes...)
	} else if maxExp < p.maxExp {
		for idx := maxExp + 1; idx <= p.maxExp; idx++ {
			p.setElem(idx, nil)
		}
		p.nodes = p.nodes[:maxExp+1]
	}
	p.setMaxExp(maxExp)
}

func (p *Poly) validateMaxExp() {
	validMaxExp := p.maxExp
	for exp := p.maxExp; exp >= polyNodeExponentZero; exp-- {
		elem := p.getSetInvalidElem(exp, exp)
		if elem.validateCoefficient() {
			validMaxExp = exp
			break
		}
	}
	p.resetMapExp(validMaxExp)
}

func (p *Poly) setAES(aes rune) {
	p.aes = aes
}

func (p *Poly) setElem(index int, node *Node) {
	if !p.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	p.nodes[index] = node
}

func (p *Poly) setElemValues(index int, node *Node) {
	if node != nil {
		p.setElemCoe(index, node.coe)
		p.setElemExp(index, node.exp)
	}
}

func (p *Poly) setElemCoe(index int, coe float64) {
	if !p.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if p.nodes[index] == nil {
		p.nodes[index] = ConstructValidPolyNode(index)
	}
	p.nodes[index].setValues(coe, index)
}

func (p *Poly) setElemExp(index, exp int) {
	if !p.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if p.nodes[index] == nil {
		p.nodes[index] = ConstructValidPolyNode(exp)
	}
	p.nodes[index].setExp(exp)
}

func (p *Poly) getElem(index int) *Node {
	if !p.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if polyIndexNode := p.nodes[index]; polyIndexNode != nil {
		p.setElemExp(index, index)
		return polyIndexNode
	} else {
		return ConstructValidPolyNode(index)
	}
}

func (p *Poly) getSetNullElem(index int) *Node {
	return p.getSetElem(index, index, polyNodeInvalidCoefficient)
}

func (p *Poly) getSetZeroElem(index int) *Node {
	return p.getSetElem(index, index, polyNodeCoefficientZero)
}

func (p *Poly) getSetInvalidElem(index, exp int) *Node {
	return p.getSetElem(index, exp, polyNodeInvalidCoefficient)
}

func (p *Poly) getSetUnitElem(index, exp int) *Node {
	return p.getSetElem(index, exp, polyNodeCoefficientOne)
}

func (p *Poly) getSetElem(index, exp int, coe ...float64) *Node {
	if !p.validateIndex(index) {
		panic(polyIndexOutOfBound)
	}
	if polyIndexNode := p.nodes[index]; polyIndexNode != nil {
		p.nodes[index].setExp(exp)
		return polyIndexNode
	} else {
		destCoe := polyNodeCoefficientZero
		if len(coe) > 0 {
			destCoe = coe[0]
		}
		newPolyNode := constructPolyNode(destCoe, exp)
		p.setElem(index, newPolyNode)
		return newPolyNode
	}
}

func (p *Poly) setElemByOne2OneOpt(index int, opt rune, pt *Poly) {
	if !ValidatePoly(p, pt) {
		panic(polyInvalidError)
	}
	p.getSetInvalidElem(index, index).one2OneOpt(opt, pt.getSetInvalidElem(index, index))
}

func (p *Poly) setGetElemOpt(index, exp int, opt rune, node *Node) *Node {
	return p.getSetInvalidElem(index, exp).one2OneOpt(opt, node)
}

func (p *Poly) setElemNull(index int) {
	p.getElem(index).setNull()
}

func (p *Poly) setElemZero(index int) {
	p.getElem(index).setZero()
}

func (p *Poly) One2OneOpt(opt rune, pt *Poly) *Poly {
	return p.one2OneOpt(opt, pt)
}

// one2OneOpt
// '+' Test Pass
// '-' Test Error
func (p *Poly) one2OneOpt(opt rune, pt *Poly) *Poly {
	if !ValidatePoly(p, pt) {
		return p
	}
	polyDestMaxExp := lang.MaxInt(p.maxExp, pt.maxExp)
	p.resetMapExp(polyDestMaxExp)
	pt.resetMapExp(polyDestMaxExp)
	for idx := 0; idx <= polyDestMaxExp; idx++ {
		p.setElemByOne2OneOpt(idx, opt, pt)
	}
	return p
}

func (p *Poly) AddNode(node *Node) *Poly {
	if node != nil {
		p.setGetElemOpt(node.exp, node.exp, '+', node)
	}
	return p
}

func (p *Poly) SubNode(node *Node) *Poly {
	if node != nil {
		p.setGetElemOpt(node.exp, node.exp, '-', node)
	}
	return p
}

// Add Test Pass
func (p *Poly) Add(pt *Poly) *Poly {
	return p.one2OneOpt('+', pt)
}

func (p *Poly) GetPlus(pt *Poly) *Poly {
	polyCopy := p.makeCopy()
	return polyCopy.Add(pt)
}

func (p *Poly) Sub(pt *Poly) *Poly {
	return p.one2OneOpt('-', pt)
}

func (p *Poly) GetMinus(pt *Poly) *Poly {
	polyCopy := p.makeCopy()
	return polyCopy.Sub(pt)
}

func (p *Poly) Mul(pt *Poly) *Poly {
	//return p.simpleMul(p)
	return p.simpleMulV2(pt)
}

func (p *Poly) GetTimes(pt *Poly) *Poly {
	polyCopy := p.makeCopy()
	return polyCopy.Mul(pt)
}

// simpleMul
// must use heap to accelerate the calculation
func (p *Poly) simpleMul(pt *Poly) *Poly {
	if !ValidatePoly(p, pt) {
		panic(polyInvalidError)
	}
	var (
		pSize      int   = p.getSize()
		polyResult *Poly = &Poly{}
		polyCopy   *Poly
		polyPNode  *Node
	)
	polyResult.assign(p.maxExp + pt.maxExp)
	for pIdx := 0; pIdx < pSize; pIdx++ {
		polyCopy = p.makeCopy()
		polyPNode = p.getSetInvalidElem(pIdx, pIdx)
		if polyPNode.validate() {

			//debugSimpleMul(polyPNode, polyCopy)

			polyCopy.MulNode(polyPNode)
			polyResult.Add(polyCopy)
		}
	}
	return polyResult
}

func debugSimpleMul(polyPNode *Node, polyCopy *Poly) {
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
func (p *Poly) simpleMulV2(pt *Poly) *Poly {
	if !ValidatePoly(p, pt) {
		return nil
	}
	var (
		ptSize            = pt.getSize()
		polyCopy          *Poly
		polyPNode         *Node
		polyMiddleResults []*Poly
	)
	polyMiddleResults = make([]*Poly, ptSize)

	// get polyMiddleResults
	for pIdx := 0; pIdx < ptSize; pIdx++ {
		polyCopy = p.makeCopy()
		polyPNode = pt.getSetInvalidElem(pIdx, pIdx)
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

func (p *Poly) FFTMul(pt *Poly) *CoePoly {
	return p.fftMul(pt)
}

func (p *Poly) fftMul(pt *Poly) *CoePoly {
	return p.CoePoly().Mul(pt.CoePoly())
}

// judgeSparse
// if sparse
func (p *Poly) judgeSparse(funcPtr func(...interface{}) bool) bool {
	return funcPtr()
}

func (p *Poly) judgeMulAlgo() bool {
	return false
}

func (p *Poly) MulNode(node *Node) *Poly {
	if node == nil {
		panic(polyNodeInvalidError)
	}
	var (
		polyMaxExp int   = p.maxExp
		nodeCopy   *Node = node.makeCopy()
	)
	if node.isConstant() {
		for idx := 0; idx <= polyMaxExp; idx++ {
			p.setGetElemOpt(idx, idx, '*', nodeCopy)
		}
	} else {
		p.resetMapExp(polyMaxExp + nodeCopy.exp)
		for idx := polyMaxExp; idx >= polyNodeExponentZero; idx-- {
			if p.getSetInvalidElem(idx, idx).validateCoefficient() {
				setExp := idx + nodeCopy.exp
				p.setElemValues(setExp, p.setGetElemOpt(idx, setExp, '*', nodeCopy))
				p.setElemZero(idx)
			}
		}
	}

	return p
}

// Div by Factoring
func (p *Poly) Div(pt *Poly) *Poly {
	return p
}

func (p *Poly) Value(value float64) float64 {
	var (
		polyValue float64 = 0.0
	)
	for idx := 0; idx < p.getSize(); idx++ {
		polyValue += p.getElem(idx).Value(value)
	}
	return polyValue
}

// Integral Test Pass
func (p *Poly) Integral() *Poly {
	var (
		polyMaxExp int
	)
	if p.validate() {
		polyMaxExp = p.maxExp
		p.nodes = append(p.nodes, p.getSetElem(polyMaxExp, polyMaxExp).Integral().makeCopy())
		for idx := polyMaxExp; idx > polyNodeExponentZero; idx-- {
			p.setElemValues(idx, p.getSetElem(idx-1, idx-1).Integral())
		}
		p.getElem(polyNodeExponentZero).setZero()
		p.setMaxExp(p.maxExp + 1)
	}
	return p
}

// Derivative Test Pass
func (p *Poly) Derivative() *Poly {
	var (
		polyMaxExp int
	)
	if p.validate() {
		polyMaxExp = p.maxExp
		for idx := polyNodeExponentZero; idx < polyMaxExp; idx++ {
			p.setElemValues(idx, p.getSetElem(idx+1, idx+1).Derivative())
		}
		p.resetMapExp(polyMaxExp - 1)
	}
	return p
}

func (p *Poly) SolvePoly(pt *Poly) *Solution {
	return p.Sub(pt).Solve()
}

func (p *Poly) SolveValue(value float64) *Solution {
	return p.SubNode(ConstructPolyNode(value, polyNodeExponentZero)).Solve()
}

func (p *Poly) getElemQuotient(indexI, indexJ int) *Node {
	if !p.validateIndex(indexI, indexJ) {
		panic(polyIndexOutOfBound)
	}
	return p.getElem(indexI).GetQuotient(p.getElem(indexJ))
}

func (p *Poly) getElemQuotientCoe(indexI, indexJ int) float64 {
	return p.getElemQuotient(indexI, indexJ).coe
}

// Solve
// poly equation 0.0
func (p *Poly) Solve() *Solution {
	s := &Solution{}
	p.validateMaxExp()
	switch p.maxExp {
	case 0:
		{
			panic(polyEquationNoSolution)
		}
	case 1:
		{
			s.setValues(make([]complex128, 1), 1)
			exp0Value := p.getSetElem(0, 0).coe
			exp1Coe := p.getSetElem(1, 1).coe
			if lang.EqualFloat64Zero(exp1Coe) {
				panic(polyEquationNoSolution)
			}
			s.setElem(0, complex(exp0Value*-1.0/exp1Coe, 0.0))
		}
	case 2:
		{
			s.setValues(make([]complex128, 2), 2)
			exp0Coe := p.getSetElem(0, 0).coe
			exp1Coe := p.getSetElem(1, 1).coe
			exp2Coe := p.getSetElem(2, 2).coe
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
			exp0Coe := p.getSetElem(0, 0).coe
			exp1Coe := p.getSetElem(1, 1).coe
			exp2Coe := p.getSetElem(2, 2).coe
			exp3Coe := p.getSetElem(3, 3).coe
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

func (p *Poly) Factoring() *Factors {
	solution := p.Solve()
	return solution.PolyFactors()
}

func (p *Poly) Normal() *Poly {
	return p.Discard(true).
		Merge(true).
		Sort(true)
}

func (p *Poly) Regular() *Poly {
	return p.Discard().Merge().Sort()
}

func (p *Poly) Discard(predict ...bool) *Poly {
	if len(predict) > 0 && predict[0] {
		return p.removeZeroNullInvalidNode()
	}
	return p
}

// removeZeroNullInvalidNode
// get rid of invalid nodes
func (p *Poly) removeZeroNullInvalidNode() *Poly {
	return p
}

func (p *Poly) Sort(predict ...bool) *Poly {
	if len(predict) > 0 && predict[0] {
		return p.sortNodeByExponent()
	}
	return p
}

// sortNodeByExponent
// sort the nodes by exp asc
func (p *Poly) sortNodeByExponent() *Poly {
	return p
}

func (p *Poly) Merge(predict ...bool) *Poly {
	if len(predict) > 0 && predict[0] {
		return p.mergeNodeBySameExponent()
	}
	return p
}

// mergeNodes
// merge the nodes by same exp
func (p *Poly) mergeNodeBySameExponent() *Poly {
	return p
}

func (p *Poly) getValidNode() []*Node {
	validNode := make([]*Node, 0)
	for exp := 0; exp <= p.maxExp; exp++ {
		if p.getElem(exp).validate() {
			validNode = append(validNode, p.getElem(exp).makeCopy())
		}
	}
	return validNode
}

func (p *Poly) PolyNode() *Node {
	if validNode := p.getValidNode(); len(validNode) == 1 {
		return validNode[0]
	}
	return &Node{}
}

func (p *Poly) getCoeSlice() []float64 {
	coeSlice := make([]float64, p.getSize())
	for idx := 0; idx < p.getSize(); idx++ {
		coeSlice[idx] = p.getElem(idx).coe
	}
	return coeSlice
}

func (p *Poly) CoePoly() *CoePoly {
	return ConstructCoePoly(p.getCoeSlice())
}

func (p *Poly) ToString(precisionBit ...int) string {
	if !p.validate() {
		return lang.Float64ToString(0, precisionBit...)
	}
	var (
		polyString   string = undefinedString
		polySize     int    = p.getSize()
		preDisplayed bool   = false
		node         *Node  = p.getElem(0)
	)
	if nodeStr := node.ToString(p.aes, precisionBit...); !lang.StringIsNull(nodeStr) {
		polyString += nodeStr
		preDisplayed = true
	}
	for idx := 1; idx < polySize; idx++ {
		node = p.getElem(idx)
		if nodeStr := node.ToString(p.aes, precisionBit...); !lang.StringIsNull(nodeStr) {
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

func (p *Poly) Display(isPrintln bool, precisionBit ...int) *Poly {
	polyString := p.ToString(precisionBit...)
	if polyString == undefinedString {
		polyString = lang.Float64ToString(0, precisionBit...)
	}
	fmt.Printf(polyString)
	if isPrintln {
		fmt.Println()
	}
	return p
}

func (p *Poly) ToStringV2(reverse bool, precisionBit ...int) string {
	if !reverse {
		return p.ToString(precisionBit...)
	}
	if !p.validate() {
		return lang.Float64ToString(0, precisionBit...)
	}
	var (
		polyString   string = undefinedString
		preDisplayed bool   = false
		node         *Node  = p.getElem(p.maxExp)
	)

	if nodeStr := node.ToString(p.aes, precisionBit...); !lang.StringIsNull(nodeStr) {
		polyString += nodeStr
		preDisplayed = true
	}
	for idx := p.maxExp - 1; idx >= polyNodeExponentZero; idx-- {
		node = p.getElem(idx)
		if nodeStr := node.ToString(p.aes, precisionBit...); !lang.StringIsNull(nodeStr) {
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

func (p *Poly) DisplayV2(isPrintln bool, reverse bool, precisionBit ...int) *Poly {
	polyString := p.ToStringV2(reverse, precisionBit...)
	if polyString == undefinedString {
		polyString = lang.Float64ToString(0, precisionBit...)
	}
	fmt.Printf(polyString)
	if isPrintln {
		fmt.Println()
	}
	return p
}

func (p *Poly) DisplayV3() *CoePoly {
	return p.CoePoly().Display()
}
