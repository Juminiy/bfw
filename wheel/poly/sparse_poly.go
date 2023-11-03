package poly

import (
	"errors"
	"fmt"
)

var (
	SparsePolyIteratorError = errors.New("sparse poly iterator or sparse poly is null")
)

// SparsePoly
// the linked shape do not fit multiply
// plus(add) and minus(sub) speed fast
type SparsePoly struct {
	node *Node
	next *SparsePoly
}

// ConstructSparsePolyByOrder
// assume in order
func ConstructSparsePolyByOrder(real2DArray [][]float64) *SparsePoly {
	var head *SparsePoly
	var prev *SparsePoly

	for idx := 0; idx < len(real2DArray); idx++ {
		node := ConstructPolyNode(real2DArray[idx][0], int(real2DArray[idx][1]))
		curr := constructSparsePoly(node)
		if head == nil {
			head = curr
			prev = curr
		} else {
			prev.next = curr
			prev = curr
		}
	}

	return head
}

func constructSparsePoly(node *Node) *SparsePoly {
	return &SparsePoly{node: node, next: nil}
}

func (sp *SparsePoly) Iterator() *SparsePolyIterator {
	return constructSPIterator(sp)
}

// make a new copy with duplicate head node
// return the next
func (sp *SparsePoly) makeCopy() *SparsePoly {
	spCopy := &SparsePoly{}
	spi := sp.Iterator()
	spI := spCopy.Iterator()
	for spi.validate() {
		newNode := spi.current.node.makeCopy()
		spI.link(constructSparsePoly(newNode))
		spi.next()
	}
	return spCopy.next
}

// make a new copy
func (sp *SparsePoly) one2OneOpt(opt rune, spt *SparsePoly) *SparsePoly {
	if spt == nil {
		return sp
	}
	resSp := &SparsePoly{}
	resSpi := resSp.Iterator()
	spi := sp.Iterator()
	spti := spt.Iterator()
	for validateSPIterator(spi, spti) {
		if compareRes := spi.compare(spti); compareRes > 0 {
			resSpi.link(spti.dupNode())
			spti.next()
		} else if compareRes < 0 {
			resSpi.link(spi.dupNode())
			spi.next()
		} else {
			resSpi.link(spi.optRes(opt, spti))
			spi.next()
			spti.next()
		}
	}
	if validateSPIterator(spi) {
		resSpi.link(spi.current)
	}

	if validateSPIterator(spti) {
		resSpi.link(spti.current)
	}

	return resSp.next
}

func (sp *SparsePoly) add(spt *SparsePoly) *SparsePoly {
	return sp.one2OneOpt('+', spt)
}

func (sp *SparsePoly) Plus(spt *SparsePoly) *SparsePoly {
	return sp.add(spt)
}

func (sp *SparsePoly) sub(spt *SparsePoly) *SparsePoly {
	return sp.one2OneOpt('-', spt)
}

func (sp *SparsePoly) Minus(spt *SparsePoly) *SparsePoly {
	return sp.sub(spt)
}

func (sp *SparsePoly) Display(isPrintln ...bool) *SparsePoly {
	spi := sp.Iterator()
	for spi.validate() {
		spi.display()
		fmt.Printf("->")
		spi.next()
	}
	fmt.Printf("null")
	if lenPrintln := len(isPrintln); lenPrintln == 0 ||
		(lenPrintln > 0 && isPrintln[0]) {
		fmt.Println()
	}
	return sp
}

type SparsePolyIterator struct {
	current *SparsePoly
}

func validateSPIterator(spi ...*SparsePolyIterator) bool {
	if spiLen := len(spi); spiLen > 0 {
		for idx := 0; idx < spiLen; idx++ {
			if spi[idx] == nil ||
				!spi[idx].validate() {
				return false
			}
		}
	}
	return true
}

func constructSPIterator(sp *SparsePoly) *SparsePolyIterator {
	if sp == nil {
		return nil
	}
	return &SparsePolyIterator{current: sp}
}

func (spi *SparsePolyIterator) validate() bool {
	return spi.current != nil
}

func (spi *SparsePolyIterator) next() *SparsePolyIterator {
	if spi.current != nil {
		spi.current = spi.current.next
		return spi
	}
	return nil
}

func (spi *SparsePolyIterator) link(sp *SparsePoly) *SparsePolyIterator {
	spi.current.next = sp
	spi.current = sp
	return spi
}

func (spi *SparsePolyIterator) dupNode() *SparsePoly {
	return constructSparsePoly(spi.current.node.makeCopy())
}

func (spi *SparsePolyIterator) compare(spti *SparsePolyIterator) int {
	if !validateSPIterator(spi, spti) {
		panic(SparsePolyIteratorError)
	}
	return spi.current.node.exp - spti.current.node.exp
}

func (spi *SparsePolyIterator) optBool(opt rune, spti *SparsePolyIterator) bool {
	return spi.current.node.one2OneOptNotPanic(opt, spti.current.node)
}

func (spi *SparsePolyIterator) optRes(opt rune, spti *SparsePolyIterator) *SparsePoly {
	spiDup := spi.dupNode()
	spiDup.node.one2OneOptNotPanic(opt, spti.current.node)
	return spiDup
}

func (spi *SparsePolyIterator) display() {
	spi.current.node.Display(false, polyDefaultAes)
}
