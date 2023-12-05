package la

import (
	"bfw/wheel/lang"
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"
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
	//m1 := &Matrix{rowSize: 2, columnSize: 3, slice: [][]float64{{1, 2, 3}, {4, 5, 6}}}
	m2 := &Matrix{rowSize: 2, columnSize: 3, slice: [][]float64{{1, 2, 3}, {4, 5, 6}}}
	m2.transpose().sub(m2.transpose()).Display().transpose()
	//m1.MulLambda(2).Display().mul(m2).Display()
	v1 := &Vector{size: 5, slice: []float64{1, 2, 3, 4, 5}, shape: true}
	v1.Display()
}

func TestMatrix_Convergence(t *testing.T) {
	m := &Matrix{rowSize: 3, columnSize: 3, slice: [][]float64{{0, 0.5, 0.5}, {1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}, {0, 1, 0}}}
	// from one state start, finally convergence to a static state
	//hamburger := &Vector{size: 3, shape: false, slice: []float64{1, 0, 0}}
	//pizza := &Vector{size: 3, shape: false, slice: []float64{0, 1, 0}}
	//hotdog := &Vector{size: 3, shape: false, slice: []float64{0, 0, 1}}
	//converM, converMT := m.Convergence()
	//fmt.Println("Matrix Self convergence time:", converMT)
	//converM.Display()
	//converHam, converHamT := hamburger.Convergence(m)
	//fmt.Println("Vector convergence time:", converHamT)
	//converHam.Display()
	//converPiz, converPizT := pizza.Convergence(m)
	//fmt.Println("Vector convergence time:", converPizT)
	//converPiz.Display()
	//converHot, converHotT := hotdog.Convergence(m)
	//fmt.Println("Vector convergence time:", converHotT)
	//converHot.Display()
	//fmt.Println()

	probabilityPi := &Vector{size: 3, shape: false, slice: []float64{0.5, 0, 0.5}}
	probabilityPi, piT := probabilityPi.Convergence(m)
	fmt.Println("pi convergence time:", piT)
	probabilityPi.Display()
	// now calculate the eigenvalue and eigenvectors
	//m.EigenValues().Display()
}

func TestMatrix_Convergence2(t *testing.T) {
	m := ConstructMatrix([][]float64{{0.5, 0.2, 0.3}, {0.6, 0.2, 0.2}, {0.1, 0.8, 0.1}})
	m, timeT := m.Convergence()
	fmt.Println(timeT)
	m.Display()
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
	m.mul(m.GetInverse()).Display()
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
	m.mul(m.GetInverse()).Identity().Display()
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
	//m.GetPlus(m).Display()
	//m.GetMinus(m).Display()
	//m.GetTimes(m).Display()
	//m.GetRDivide(m).Display()
	//m.GetPower(m).Display()
	//m.MTimes(m).Display()
	//m.GetMRDivide(m).Approx().Display()
	//m.GetMPower(10).Display()
	//m.GetTranspose().Display()
	//m.GetInverse().Display()
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
	MatrixTestEigenValuesByRoundN(10)
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
	matrix1.GetFlip().Display()
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
	//fmt.Printf("%.5f\n", matrix1.Convolution(matrix1))
	//matrix1.Display()
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

func TestMatrix_PhalanxTranspose(t *testing.T) {
	matrix := GenMatrix(3, 3, "int", 100)
	matrix.Display().phalanxTranspose().Display()
}

func TestBlockMatrix_Transpose(t *testing.T) {
	genBTSize, genBSize := 3, 3
	bm := GenBlockMatrix(genBTSize, genBTSize, "i", genBSize, genBSize, 100)
	bm.Display()
	fmt.Println("----------------------------------------------------------------")
	bm.transpose().Display()
}

func TestBlockMatrix_Mul(t *testing.T) {
	genBTSize, genBSize := 4, 4
	time0 := time.Now()
	bm1 := GenBlockMatrix(genBTSize, genBTSize, "f", genBSize, genBSize, 100)
	bm2 := GenBlockMatrix(genBTSize, genBTSize, "f", genBSize, genBSize, 100)
	fmt.Printf("2 * %d*%d Matrix Generate time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time0))

	time1 := time.Now()
	bm1.Mul(bm2)
	fmt.Printf("%d*%d Matrix Multiply After Speed & DivBlock time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time1))

	time2 := time.Now()
	bm1.Matrix().MTimes(bm2.Matrix())
	fmt.Printf("%d*%d Matrix Multiply After speed time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time2))

	time2dot5 := time.Now()
	bm1.Matrix().mul(bm2.Matrix())
	fmt.Printf("%d*%d Matrix Multiply None speed time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time2dot5))

	time3 := time.Now()
	bm1.Matrix().Equal(bm2.Matrix())
	fmt.Printf("%d*%d Matrix Traverse time: %v\n", genBSize*genBTSize, genBSize*genBTSize, time.Since(time3))
	// validate answer same and correct
	//fmt.Println(blockMulRes.Matrix().Equal(simpleMulRes))
}

func TestMatrix_Intercept(t *testing.T) {
	matrix := GenMatrix(4, 4, "i", 100)
	//matrix.zeroPadding(5, 5).Display().getBlock(2, 2, 4, 4).Display()
	matrix.Display().zeroPadding(5, 5).Display().zeroTruncate(3, 3).Display()
}

func TestMatrix_MTimes(t *testing.T) {
	destSize, dataType, dataRange := 32, "i", []float64{1e2, 1e5}
	matrix1 := GenMatrix(destSize, destSize, dataType, dataRange...)
	matrix2 := GenMatrix(destSize, destSize, dataType, dataRange...)
	correctRes1 := matrix1.mulV1(matrix2)
	correctRes2 := matrix1.mulV2(matrix2)
	correctRes2Dot5 := matrix1.mulV2Dot5(matrix2)
	lookRes2Dot6 := matrix1.mulV2Dot6(matrix2)
	res3 := matrix1.mulV3(matrix2)
	//res4 := matrix1.mulV4(matrix2)
	//res5 := matrix1.mulV5(matrix2)
	fmt.Println("simple kij mul(cache speed):", correctRes1.Equal(correctRes2))
	fmt.Println("simple ij mul(A mul B.transpose, vector dot):", correctRes1.Equal(correctRes2Dot5))
	fmt.Println("parallel mul(A mul B.transpose, vector dot):", correctRes1.Equal(lookRes2Dot6))
	fmt.Println("recursive mul O(n^3):", correctRes1.Equal(res3))
	//fmt.Println("non recursive mul O(n^3):", correctRes1.Equal(res4))
	//fmt.Println("best mul(Strassen algorithm, none recursive, multiple thread, cache speed):", correctRes1.Equal(res5))
}

func TestMatrix_All(t *testing.T) {
	mArr := GenMatrixArray(4, 2, 2, "i", 100)
	mArr[0].Display()
	mArr[1].Display()
	mArr[2].Display()
	mArr[3].Display()
	phalanxBlock22Patch(mArr[0], mArr[1], mArr[2], mArr[3]).Display()
}

func TestMatrix_Mul2(t *testing.T) {
	size, dType, dRange := 64, "f", 1e+5
	ma := GenMatrix(size, size, dType, dRange)
	mb := GenMatrix(size, size, dType, dRange)

	time1 := time.Now()
	ma.mulV1(mb)
	fmt.Printf("%d*%d Matrix Multiply Simple: %v\n", size, size, time.Since(time1))

	time2 := time.Now()
	ma.mulV2(mb)
	fmt.Printf("%d*%d Matrix Multiply (by change loop order): %v\n", size, size, time.Since(time2))

	time2Dot5 := time.Now()
	ma.mulV2Dot5(mb)
	fmt.Printf("%d*%d Matrix Multiply (by transpose vector dot): %v\n", size, size, time.Since(time2Dot5))

	time2Dot6 := time.Now()
	ma.mulV2Dot6(mb)
	fmt.Printf("%d*%d Matrix Multiply (by transpose vector dot, parallel by go routine): %v\n", size, size, time.Since(time2Dot6))

	time3 := time.Now()
	ma.mulV3(mb)
	fmt.Printf("%d*%d Matrix Multiply After DivBlock Recursive O(n^3): %v\n", size, size, time.Since(time3))

	//time4 := time.Now()
	//ma.mulV4(mb)
	//fmt.Printf("%d*%d Matrix Multiply After DivBlock, None Recursive, O(n^3): %v\n", size, size, time.Since(time4))

	//time5 := time.Now()
	//ma.mulV5(mb)
	//fmt.Printf("%d*%d Matrix Multiply After DivBlock, None Recursive, multiple threads,Straseen Algorithm-O(n^2.807): %v\n", size, size, time.Since(time5))
}

func TestGenMatrix(t *testing.T) {
	m1 := &Matrix{}
	m2 := new(Matrix)
	fmt.Println(&m1, m1)
	fmt.Println(&m2, m2)

	m3 := &Matrix{}
	m3.setValues(make([][]float64, 20), 20, 20)
	m4 := new(Matrix)
	m4.setValues(make([][]float64, 20), 20, 20)
	fmt.Println(&m3, unsafe.Pointer(&m3.slice), &m3.rowSize, &m3.columnSize)
	fmt.Println(&m4, unsafe.Pointer(&m4.slice), &m4.rowSize, &m4.columnSize)
}

func TestGenerateTwoBlockFitFloat64Phalanx(t *testing.T) {
	n := lang.GetRandomIntValue(10)
	p := rand.Float64()
	q := 1 - p
	for i := 0; i < (n<<1 + 1); i++ {
		for j := 0; j < (n<<1 + 1); j++ {
			if j == i+1 {
				fmt.Printf("%f ", p)
			} else if j == i-1 {
				fmt.Printf("%f ", q)
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Println()
	}
}

func TestConstructMatrix(t *testing.T) {
	m := ConstructMatrix([][]float64{{0.5, 0.25, 0.125, 0.125}, {0, 0, 1, 0}, {0, 0, 0, 1}, {1, 0, 0, 0}})
	fmt.Println("P1")
	m.MPower(1).Display()
	fmt.Println("P2")
	m.MPower(2).Display()
	fmt.Println("P3")
	m.MPower(3).Display()
	fmt.Println("P4")
	m.MPower(4).Display()
	v := ConstructVector([]float64{0.25, 0.25, 0.25, 0.25})
	v, vT := v.Convergence(m)
	fmt.Println(vT)
	v.Display()
}

func TestAddRealMatrix(t *testing.T) {
	m := ConstructMatrix([][]float64{
		{0, 0.5, 0, 0.5},
		{0.5, 0, 0.5, 0},
		{0, 0.5, 0, 0.5},
		{0.5, 0, 0.5, 0}})
	m.MPower(2).Display()
	//v := ConstructVector([]float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}, true)
	//v, _ = v.Convergence(m)
	//v.Display()
	//m, _ = m.Convergence()
	//v.MulMatrix(m).Display()
}
