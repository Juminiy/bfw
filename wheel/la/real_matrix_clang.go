package la

/*
#cgo LDFLAGS: -L ../cgo/lib -lmatrix -lcommon
#include <stdio.h>
#include <stdlib.h>
#include "../cgo/include/matrix.h"
#include "../cgo/include/common.h"
*/
import "C"

func GenRealMatrix(rowSize, columnSize C.int, rangeStart, rangeEnd C.double) *C.real_matrix {
	return C.create_rand_matrix(rowSize, columnSize, rangeStart, rangeEnd)
}

func AddRealMatrix(A *C.real_matrix, B *C.real_matrix) *C.real_matrix {
	return C.add(A, B)
}

func MulRealMatrix(A *C.real_matrix, B *C.real_matrix) *C.real_matrix {
	return C.mulV2(A, B)
}

func DisplayRealMatrix(A *C.real_matrix) {
	C.print_real_matrix(A)
}

func TestGenMulDisplayRealMatrix() {
	var size C.int = C.int(2)
	var rangeStart, rangeEnd C.double = C.double(1e2), C.double(1e5)
	m1 := GenRealMatrix(size, size, rangeStart, rangeEnd)
	m2 := GenRealMatrix(size, size, rangeStart, rangeEnd)
	m3 := MulRealMatrix(m1, m2)
	DisplayRealMatrix(m3)
}
