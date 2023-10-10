package la

import (
	"bfw/common/lang"
	"fmt"
	"testing"
)

func TestMatrix_Add(t *testing.T) {
	m1 := &Matrix{rowSize: 2, lineSize: 3, slice: [][]float64{{1, 2, 3}, {4, 5, 6}}}
	m2 := &Matrix{rowSize: 2, lineSize: 3, slice: [][]float64{{1, 2, 3}, {4, 5, 6}}}
	m2.transpose().Sub(m2.transpose()).Display().transpose()
	m1.MulLambda(2).Display().Mul(m2).Display()
	v1 := &Vector{size: 5, slice: []float64{1, 2, 3, 4, 5}, shape: true}
	v1.Display()
}

func TestMatrix_Mul(t *testing.T) {
	m := &Matrix{rowSize: 3, lineSize: 3, slice: [][]float64{{0.2, 0.6, 0.2}, {0.3, 0, 0.7}, {0.5, 0, 0.5}}}
	hamburger := &Vector{size: 3, shape: false, slice: []float64{1, 0, 0}}
	pizza := &Vector{size: 3, shape: false, slice: []float64{0, 1, 0}}
	hotdog := &Vector{size: 3, shape: false, slice: []float64{0, 0, 1}}
	converM, converMT := m.Convergence()
	fmt.Println("Matrix Self convergence time:", converMT)
	converM.Display()
	converHam, converHamT := hamburger.Convergence(m)
	fmt.Println("Vector convergence time:", converHamT)
	converHam.Display()
	converPiz, converPizT := pizza.Convergence(m)
	fmt.Println("Vector convergence time:", converPizT)
	converPiz.Display()
	converHot, converHotT := hotdog.Convergence(m)
	fmt.Println("Vector convergence time:", converHotT)
	converHot.Display()
}

func TestMatrix_Adjoin(t *testing.T) {
	m := &Matrix{}
	m.assign(5, 5)
	m.setSlice([][]float64{
		{1, 8, 3, 5, 5},
		{6, 7, 8, 4, 10},
		{11, 12, 9, 14, 15},
		{16, 23, 19, 19, 20},
		{21, 43, 21, 1, 25}})
	m.Mul(m.GetInverse()).Identity().Display()
}

func TestGetCombinationSliceMap(t *testing.T) {
	fmt.Println(len(lang.GetCombinationSliceMap(10, 5)))
}

func TestMatrix_GetInverse(t *testing.T) {
	m := &Matrix{}
	m.assign(3, 3)
	m.setSlice([][]float64{
		{1, 2, 3},
		{3, 1, 2},
		{2, 9, 4},
	})
	m.Mul(m.GetInverse()).Identity().Display()
}

func TestMatrix_ReShape(t *testing.T) {
	m := &Matrix{}
	m.assign(3, 4)
	m.setSlice([][]float64{
		{1, 2, 3, 4},
		{3, 1, 2, 9},
		{2, 9, 4, 13},
	})
	m.ReShape(4, 3).Display()
}

func TestMatrix_Matlab(t *testing.T) {
	m := &Matrix{}
	m.assign(3, 3)
	m.setSlice([][]float64{
		{2, 3, 2},
		{1, 5, 1},
		{7, 2, 2},
	})
	m.GetPlus(m).Display()
	m.GetMinus(m).Display()
	m.GetTimes(m).Display()
	m.GetRDivide(m).Display()
	m.GetPower(m).Display()
	m.GetMTimes(m).Display()
	m.GetMRDivide(m).Approx().Display()
	m.GetMPower(10).Display()
	m.GetTranspose().Display()
	m.GetInverse().Display()
}

// [ 1  4 10 20 35 44 46 40 25]
func TestGenericMatrix_ALL(t *testing.T) {
	m := &Matrix{}
	m.setValues([][]float64{{0, 1, 0}, {0, 0, 1}, {-6, -11, -6}}, 3)
	p := &Matrix{}
	p.setValues([][]float64{{1, 1, 1}, {-1, -2, -3}, {1, 4, 9}}, 3)
	m.Similar(p).Display()
	//m.Display()
	//p.Display()
}

// {{-0.57735027, -0.21821789,  0.10482848},
// { 0.57735027,  0.43643578, -0.31448545},
// {-0.57735027, -0.87287156,  0.94345635}}
func TestMatrix_Approx(t *testing.T) {
	matrix := ConstructMatrix([][]float64{{0, 1, 0}, {0, 0, 1}, {-6, -11, -6}})
	eigenO := ConstructMatrix([][]float64{{1, 1, 1}, {-1, -2, -3}, {1, 4, 9}})
	//eigenSU := eigen.Schmidt().Unit()
	//eigenSame := ConstructMatrix([][]float64{{-0.57735027, -0.21821789, 0.10482848},
	//	{0.57735027, 0.43643578, -0.31448545},
	//	{-0.57735027, -0.87287156, 0.94345635}})
	//eigenSameSU := eigenSame.DisplayDelimiter().Schmidt().Unit().Display().DisplayDelimiter()
	//fmt.Println(eigenSU.Equal(eigenSameSU))
	//eigenSU.GetInverse().Display()
	//eigenVG := eigen.VectorGroup(true)
	//eigenVG.get(0).GetUnit().Display()
	//eigenVGOrthBasis := eigenVG.GetSchmidt().GetUnit().Display()
	//fmt.Println(eigenVGOrthBasis.ValidateOrthonormalBasis())
	matrix.Similar(eigenO).Display()
}

func TestVector_Add(t *testing.T) {
	vector := ConstructVector([]float64{1, 2, 3})
	vector.PNorm(1)
}

func TestMatrix_Total(t *testing.T) {
	mTest1 := ConstructMatrix([][]float64{{1, -2, 2}, {-2, -2, 4}, {2, 4, -2}})
	mTest2 := ConstructMatrix(nil)
	mTest1.Display()
	mTest2.Display()
	mTest3 := ConstructMatrix([][]float64{})
	mTest3.Display()
	fmt.Println(mTest1.validate(), mTest2.validate(), mTest3.validate())
	fmt.Println(mTest1.validateOneIndex(1), mTest1.validateOneIndex(3))
	fmt.Println(mTest1.validateIndex(1, 3), mTest1.validateIndex(4, 3))
	mTest1.Unit().Schmidt().Display()
	fmt.Println(mTest1.getLine(2))
	mTest4 := ConstructMatrix([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(mTest4.getRow(1), mTest4.getLine(2))
}
