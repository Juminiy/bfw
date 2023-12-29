package lang

import (
	"fmt"
)

type vType struct {
	a int
	b string
}

func NewFeaturesClear() {
	var int1dArray = []int{1, 2, 3, 4, 5}
	fmt.Println("arr = ", int1dArray, "; len(arr) = ", len(int1dArray), "; cap(arr) = ", cap(int1dArray))
	fmt.Println("max(1, 2, 3) = ", max(1, 2, 3), "\nmin(1, 2, 3) = ", min(1, 2, 3))
	clear(int1dArray)
	fmt.Println("after clear")
	fmt.Println("arr = ", int1dArray, "; len(arr) = ", len(int1dArray), "; cap(arr) = ", cap(int1dArray))

	var intStringMap = map[int]string{1: "v1", 2: "v2", 3: "v3", 4: "v4", 5: "v5"}
	fmt.Println("map = ", intStringMap, "; len(map) = ", len(intStringMap))
	clear(intStringMap)
	fmt.Println("map = ", intStringMap, "; len(map) = ", len(intStringMap))
}

func PointerNew() {
	vTypePtr := new(vType)
	vTypePtr.a = 1
	vTypePtr.b = "a"
	fmt.Println(vTypePtr)
}

type CommonInterface interface {
	func() int
	func(int, int) bool
	func(int, int)
}

type GenericStruct[T CommonInterface] struct {
}
