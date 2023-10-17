package la

import (
	"fmt"
	"math"
	"testing"
)

var (
	polyNode0 = ConstructPolyNode(1, 0)
	polyNode1 = ConstructPolyNode(2, 1)
	polyNode4 = ConstructPolyNode(5, 4)
	polyNode9 = ConstructPolyNode(10, 9)

	polyNode20  = ConstructPolyNode(1, 0)
	polyNode22  = ConstructPolyNode(3, 2)
	polyNode26  = ConstructPolyNode(7, 6)
	polyNode215 = ConstructPolyNode(16, 15)

	// (5x-89)(4x-33) = 20x^2-521x+2937 = 0, x1 = 17.8, x2 = 8.25
	polyNode30 = ConstructPolyNode(2937, 0)
	polyNode31 = ConstructPolyNode(-521, 1)
	polyNode32 = ConstructPolyNode(20, 2)

	// (3x-7)(5x-8)(9x-11)
	polyNode40 = ConstructPolyNode(-616, 0)
	polyNode41 = ConstructPolyNode(1153, 1)
	polyNode42 = ConstructPolyNode(-696, 2)
	polyNode43 = ConstructPolyNode(135, 3)

	// (x-2)^3
	polyNode50 = ConstructPolyNode(-8, 0)
	polyNode51 = ConstructPolyNode(12, 1)
	polyNode52 = ConstructPolyNode(-6, 2)
	polyNode53 = ConstructPolyNode(1, 3)

	// (x-1)(x-2)^2
	polyNode60 = ConstructPolyNode(-4, 0)
	polyNode61 = ConstructPolyNode(8, 1)
	polyNode62 = ConstructPolyNode(-5, 2)
	polyNode63 = ConstructPolyNode(1, 3)
	polyNode69 = ConstructPolyNode(3, 9)

	// (x-3)(x-(2+i))(x-(2-i))
	polyNode70 = ConstructPolyNode(-15, 0)
	polyNode71 = ConstructPolyNode(17, 1)
	polyNode72 = ConstructPolyNode(-7, 2)
	polyNode73 = ConstructPolyNode(1, 3)

	poly1 = &Poly{}
	poly2 = &Poly{}
)

func TestPolyNode_Display(t *testing.T) {
	polyNode0.swap(polyNode1)
	// 2x
	polyNode0.Display(true, 'x')
	// 1
	polyNode1.Display(true, 'x')
}

func TestPoly_Display(t *testing.T) {
	poly1.setNodeByOrder(polyNode0, polyNode1, polyNode4, polyNode9)
	poly2.setNodeByOrder(polyNode20, polyNode22, polyNode26, polyNode215)

	//1+2x+5x^4+10x^9
	poly1.Display(true, 0)
	//1+3x^2+7x^6+16x^15
	poly2.Display(true, 0)

	poly1.swap(poly2)

	//1+3x^2+7x^6+16x^15
	poly1.Display(true, 0)
	//1+2x+5x^4+10x^9
	poly2.Display(true, 0)
}

func TestPoly_makeCopy(t *testing.T) {
	poly1.setNodeByOrder(polyNode0, polyNode1, polyNode4, polyNode9)
	poly2.setNodeByOrder(polyNode20, polyNode22, polyNode26, polyNode215)
	poly1Copy := poly1.makeCopy()
	poly1Copy.Display(true, 0)
	// swap success is to show different address
	fmt.Println(poly1)
	fmt.Println(poly1Copy)
}

func TestPoly_Add_Integral_Derivative(t *testing.T) {
	poly1.setNodeByOrder(polyNode0, polyNode1, polyNode4, polyNode9)
	poly2.setNodeByOrder(polyNode20, polyNode22, polyNode26, polyNode215)
	// (2,0), (2,1), (3,2), (5,4), (7,6), (10,9), (16,15)
	//2+2x+3x^2+5x^4+7x^6+10x^9+16x^15
	poly1.Add(poly2).Display(true, 0)
	//2x+1x^2+1x^3+1x^5+1x^7+1x^10+1x^16
	poly1.Integral().Display(true, 0)
	//2+2x+3x^2+5x^4+7x^6+10x^9+16x^15
	poly1.Derivative().Display(true, 0)
	//10x^4+10x^5+15x^6+25x^8+35x^10+50x^13+80x^19
	poly1.MulNode(polyNode4).Display(true, 0)
}

func TestPoly_Sub(t *testing.T) {
	poly1.setNodeByOrder(polyNode50, polyNode51, polyNode52, polyNode53)
	//-8+12x-6x^2+1x^3
	poly1.Display(true, 0)

	poly2.setNodeByOrder(polyNode60, polyNode61, polyNode62, polyNode69)
	//-4+8x-5x^2+3x^9
	poly2.Display(true, 0)

	//-4+4x-1x^2+1x^3-3x^9
	poly1.Sub(poly2).Display(true, 0)
}

func TestPoly_Mul(t *testing.T) {
	poly1.setNodeByOrder(polyNode0, polyNode1, polyNode4, polyNode9)
	poly2.setNodeByOrder(polyNode20, polyNode22, polyNode26, polyNode215)
	//1+2x+5x^4+10x^9
	//poly1.MulNode(polyNode20).Display(true, 0)
	//3x^2+6x^3+15x^6+30x^11
	//poly1.MulNode(polyNode22).Display(true, 0)
	//7x^6+14x^7+35x^10+70x^15
	//poly1.MulNode(polyNode26).Display(true, 0)
	//16x^15+32x^16+80x^19+160x^24
	//poly1.MulNode(polyNode215).Display(true, 0)

	//simpleMul
	//1+2x+3x^2+6x^3+5x^4+22x^6+14x^7+10x^9+35x^10+30x^11+86x^15+32x^16+80x^19+160x^24
	//simpleMulV2
	//1+2x+3x^2+6x^3+5x^4+22x^6+14x^7+10x^9+35x^10+30x^11+86x^15+32x^16+80x^19+160x^24
	poly1.Mul(poly2).Display(true, 0)
}

func Test_Math(t *testing.T) {
	// 1
	fmt.Println(math.Pow(27.0, 1/3))
	// 2.9999999999999996
	fmt.Println(math.Pow(27.0, 1.0/3))
	// 3
	fmt.Println(math.Sqrt(9.0))
}

func TestPoly_Solve(t *testing.T) {
	//poly1.setNull()
	//poly1.setNodeByOrder(polyNode30, polyNode31, polyNode32)
	////[17.8, 8.25]
	//poly1.Solve().Display(5, 5)

	poly1.setNull()
	poly1.setNodeByOrder(polyNode40, polyNode41, polyNode42, polyNode43)
	//[2.33333, 1.6, 1.22222]
	poly1.Solve().Display(6, 6)
	fmt.Println()

	poly1.setNull()
	poly1.setNodeByOrder(polyNode50, polyNode51, polyNode52, polyNode53)
	//[2, 2, 2]
	poly1.Solve().Display(5, 5)
	fmt.Println()

	poly1.setNull()
	poly1.setNodeByOrder(polyNode60, polyNode61, polyNode62, polyNode63)
	//[1, 2, 2]
	poly1.Solve().Display(5, 5)
	fmt.Println()

	poly1.setNull()
	poly1.setNodeByOrder(polyNode70, polyNode71, polyNode72, polyNode73)
	//[3, 2+i, 2-i]
	poly1.Solve().Display(5, 5)
	fmt.Println()

}

func TestPoly_Factoring(t *testing.T) {
	poly1.setNull()
	poly1.setNodeByOrder(polyNode60, polyNode61, polyNode62, polyNode63)
	//[1, 2, 2]
	poly1.Factoring().Display(0)
	fmt.Println()

	//poly1.setNull()
	//poly1.setNodeByOrder(polyNode70, polyNode71, polyNode72, polyNode73)
	////[3, 2+i, 2-i]
	//poly1.Factoring().Display(1)
	//fmt.Println()
}

func TestPoly_DisplayV2(t *testing.T) {
	poly1.Display(false, 1)
}
