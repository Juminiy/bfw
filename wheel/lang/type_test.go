package lang

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestConvertInterfaceArrayOrSliceToString(t *testing.T) {
	type a struct {
		A string
		B int
		C complex128
	}
	type aa struct {
		//aElem    a
		aElemPtr *a
	}
	var aElem a
	var aAElem aa

	fmt.Println(aElem, aAElem)

	var inta int
	pinta := unsafe.Pointer(&inta)
	fmt.Println(pinta)
	var intarra []int = make([]int, 10)
	pintarra := unsafe.Pointer(&intarra)
	fmt.Println(pintarra, &intarra[0], &intarra[1])
	pintarra = unsafe.Add(pintarra, 8)
	fmt.Println(pintarra)
}
