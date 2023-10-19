package la

import (
	"bfw/common/lang"
	"fmt"
	"testing"
)

func MatrixTestEigenValuesByCallFunctionChain(real2DArray [][]float64) {
	matrix := ConstructMatrix(real2DArray)
	fmt.Println("matrix to poly matrix: ")
	polyMatrix := matrix.PolyMatrix()
	polyMatrix.Display(0)
	fmt.Println("matrix to eigen matrix: ")
	eigenMatrix := matrix.EigenMatrix().Display(0)
	fmt.Printf("poly matrix det poly: ")
	eigenMatrixDet := eigenMatrix.Det().Display(true, 1)
	fmt.Printf("equation solution: ")
	eigenMatrixDet.Solve().Display(3, 3)
	fmt.Printf("\neigen values: ")
	matrix.EigenValues().Display()
}

func MatrixTestEigenValuesByRoundN(n int) {
	if n > 1000 {
		fmt.Printf("Test %.5v case is enough\n", n)
	}
	var (
		idx    int = 1
		matrix *Matrix
	)
	for idx <= n {
		matrix = GenMatrix(3, 3, "float64", 100)

		eigenValues := matrix.EigenValues()
		if eigenValues.ValidateAllRealRoots() {
			fmt.Printf("Test [%.5v] ", idx)
			eigenValues.Display()
			fmt.Println()
		}
		idx++
	}
	fmt.Println("Test success, if any errors, program will panic")
}

func MatrixTestEigenValuesByRoundNToWriteFile(n int, fileName string) {
	if n > 500 {
		fmt.Printf("Test %.5v enough\n", n)
	}

	var (
		idx    int = 0
		matrix *Matrix
	)
	for idx < n {
		matrix = GenMatrix(3, 3, "int64", 10)
		matrix.EigenValues().Display()
		idx++
		fmt.Println()
	}
	fmt.Println("Test success, if any errors, program will panic")
}

func TestMatrix_Calculate1(t *testing.T) {
	m1 := &Matrix{rowSize: 2, columnSize: 3, slice: [][]float64{{1, 2, 3}, {4, 5, 6}}}
	m2 := &Matrix{rowSize: 2, columnSize: 3, slice: [][]float64{{1, 2, 3}, {4, 5, 6}}}
	m2.transpose().Sub(m2.transpose()).Display().transpose()
	m1.MulLambda(2).Display().Mul(m2).Display()
	v1 := &Vector{size: 5, slice: []float64{1, 2, 3, 4, 5}, shape: true}
	v1.Display()
}

func TestMatrix_Convergence(t *testing.T) {
	m := &Matrix{rowSize: 3, columnSize: 3, slice: [][]float64{{0.2, 0.6, 0.2}, {0.3, 0, 0.7}, {0.5, 0, 0.5}}}
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
	m.SimilarityTransformation(p).Display()
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
	matrix.SimilarityTransformation(eigenO).Display()
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
	//fmt.Println(mTest1.validateOneIndex(1), mTest1.validateOneIndex(3))
	fmt.Println(mTest1.validateIndex(1, 3), mTest1.validateIndex(4, 3))
	mTest1.Unit().Schmidt().Display()
	fmt.Println(mTest1.getColumn(2))
	mTest4 := ConstructMatrix([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(mTest4.getRow(1), mTest4.getColumn(2))
}

// 1 -2 2
// -2 -2 4
// 2 4 -2

// 1-λ -2 2
// -2 -2-λ 4
// 2 4 -2-λ
func TestMatrix_EigenValues(t *testing.T) {
	MatrixTestEigenValuesByCallFunctionChain([][]float64{{1, -2, 2}, {-2, -2, 4}, {2, 4, -2}})
}

// 1 1 0
// 0 1 1
// -1 0 1

// 1-1λ 1 0
// 0 1-1λ 1
// -1 0 1-1λ
func TestMatrix_EigenValues2(t *testing.T) {
	MatrixTestEigenValuesByCallFunctionChain([][]float64{{1, 1, 0}, {0, 1, 1}, {-1, 0, 1}})
}

func TestMatrix_IsSymmetric(t *testing.T) {

	// true
	matrix := ConstructMatrix([][]float64{
		{2, -2, -8},
		{-2, -2, 1},
		{-8, 1, 4}})
	fmt.Println(matrix.IsSymmetric())

	// false
	matrix2 := ConstructMatrix([][]float64{
		{2, -2, -8},
		{9, -2, 1},
		{-8, 1, 4}})
	fmt.Println(matrix2.IsSymmetric())

	//true
	matrix3 := ConstructMatrix([][]float64{
		{2, -2, 8},
		{2, -2, 0},
		{-8, 0, 4}})
	fmt.Println(matrix3.IsAntiSymmetric())

	//false
	matrix4 := ConstructMatrix([][]float64{
		{2, -2, -8},
		{-2, -2, 1},
		{-8, 1, 4}})
	fmt.Println(matrix4.IsAntiSymmetric())

}

func TestMatrix_ET(t *testing.T) {
	matrix := ConstructMatrix([][]float64{
		{2, -2, -8},
		{-2, -2, 1},
		{-8, 1, 4}})
	//matrix.rowExchangeET(1, 2).Display()
	//matrix.columnExchangeET(1, 2).Display()
	//matrix.rowMulLambdaET(2, 2).Display()
	//matrix.columnMulLambdaET(2, 4).Display()
	//matrix.rowIMulLambdaAddRowJET(1, 2, 3).Display()
	matrix.columnIMulLambdaAddColumnJET(1, 2, 3).Display()
}

func TestMatrix_Gen(t *testing.T) {
	MatrixTestEigenValuesByRoundN(100)
}

func TestMatrix_Accu(t *testing.T) {
	matrix := GenMatrix(3, 3, "i", 5, 15)
	matrix.Display()
	fmt.Println(matrix.All(func(f float64) bool {
		return f >= 5
	}))
}

func TestMatrix_Vandermonde(t *testing.T) {
	ConstructVandermonde([]float64{1, 2}).Matrix().Display().GetInverse().Display()
}

func TestMatrix_Convolution(t *testing.T) {
	matrix1 := ConstructMatrix([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	//matrix2 := ConstructMatrix([][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
	//matrix3 := ConstructMatrix([][]float64{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}})

	//fmt.Println(matrix1.Convolution(matrix2))
	//matrix1.GetFlip().Display()
	//matrix2.GetFlip().Display()

	//matrix1.flipByMiddleRow().Display()
	//matrix2.flipByMiddleRow().Display()
	//matrix3.flipByMiddleRow().Display()

	//matrix1.flipByMainDiagonal().Display()
	//matrix2.flipByMainDiagonal().Display()
	//matrix3.flipByMainDiagonal().Display()

	//matrix1.rotate90(true).Display()
	//matrix2.rotate90(true).Display()
	//matrix3.rotate90(true).Display()

	//matrix1.rotate90(false).Display()
	//matrix2.rotate90(false).Display()
	//matrix3.rotate90(false).Display()

	//matrix1.rotate(true, 3).Display()
	//matrix2.rotate(true, 3).Display()
	//matrix3.rotate(true, 3).Display()

	//matrix1.flip().Display()
	//matrix2.flip().Display()
	//matrix3.flip().Display()

	//m1Flip := matrix1.GetFlip()
	//println(matrix1.rotate90(true).rotate90(true).Equal(m1Flip))
	//m2Flip := matrix2.GetFlip()
	//println(matrix2.rotate90(true).rotate90(true).Equal(m2Flip))
	//m3Flip := matrix3.GetFlip()
	//println(matrix3.rotate90(true).rotate90(true).Equal(m3Flip))

	//matrix1.flipByMainDiagonal().Display()
	//matrix1.flipBySubDiagonal().Display()

	//9+16+21+24+25+24+21+16+9=165
	fmt.Printf("%.5f\n", matrix1.Convolution(matrix1))
	matrix1.Display()
}

// 3 6 9
// 2 5 8
// 1 4 7

func TestMatrix_IsX(t *testing.T) {
	matrix := ConstructMatrix([][]float64{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}})
	println(matrix.IsX())
}

func TestMatrix_Union(t *testing.T) {
	matrix := ConstructMatrix([][]float64{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}})
	//[1 2 3 4 5 10 15 20 25 24 23 22 21 16 11 6 7 8 9 14 19 18 17 12 13]
	fmt.Println(matrix.getSpiralOrder())
}

func TestMatrix_Adjoins(t *testing.T) {
	matrix := ConstructMatrix([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	matrix.adjoin().Display()
}

// to make 500 turns test for matrix eigenvalues
func TestMatrix_EigenValues3(t *testing.T) {

}
