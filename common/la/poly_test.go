package la

import "testing"

func TestPoly_Add(t *testing.T) {
	polyNode0 := ConstructPolyNode(1, 0)
	polyNode1 := ConstructPolyNode(2, 1)
	polyNode4 := ConstructPolyNode(5, 4)
	polyNode9 := ConstructPolyNode(10, 9)
	poly1 := &Poly{}
	poly1.setNodeByOrder(polyNode0, polyNode1, polyNode4, polyNode9)

	polyNode20 := ConstructPolyNode(1, 0)
	polyNode21 := ConstructPolyNode(3, 2)
	polyNode24 := ConstructPolyNode(7, 6)
	polyNode29 := ConstructPolyNode(16, 15)
	poly2 := &Poly{}
	poly2.setNodeByOrder(polyNode20, polyNode21, polyNode24, polyNode29)
	//polyNode0.swap(polyNode1)
	//polyNode0.Display()
	//polyNode1.Display()
	//polyNode3 := ConstructPolyNode(2, 0)
	//polyNode4 := ConstructPolyNode(4, 1)
	//poly2 := &Poly{}
	//poly2.setValues([]*PolyNode{polyNode3, polyNode4}, 1, 'r')
	//poly1.Display()
	//poly2.Display()
	//poly1.swap(poly2)
	//poly1.Display()
	//poly2.Display()

	//// (2,0), (2,1), (3,2), (5,4), (7,6), (10,9), (16,15)
	////2+2x+3x^2+5x^4+7x^6+10x^9+16x^15
	//poly1.Add(poly2).Display(0)
	////2x+1x^2+1x^3+1x^5+1x^7+1x^10+1x^16
	//poly1.Integral().Display(0)
	////2+2x+3x^2+5x^4+7x^6+10x^9+16x^15
	//poly1.Derivative().Display(0)
	////10x^4+10x^5+15x^6+25x^8+35x^10+50x^13+80x^19
	//poly1.MulNode(polyNode4).Display(0)
	//poly1.setNull()
	//poly1.setNodeByOrder(polyNode0, polyNode1, polyNode4, polyNode9)
	////5x^4+10x^5+25x^8+50x^13
	//poly1.MulNode(polyNode4).Display(0)

	poly1.Mul(poly2).Display()
}
