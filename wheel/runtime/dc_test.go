package runtime

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestDeepCopy(t *testing.T) {
	a := []int{1, 2, 3}
	b := runIntDeepCopy(a)
	fmt.Println("a:", unsafe.Pointer(&a), "b", unsafe.Pointer(&b))
	fmt.Println("a:", a, "b:", b)
}

func TestDeepCopy2(t *testing.T) {
	a := []s1{{a: 1, b: 2, c: "c", s2: s2{f: 4, m: s3{g: 9, k: "ss"}}, s2p: &s2{f: 5}}}
	b := runDeepCopy(a)
	fmt.Println("a:", unsafe.Pointer(&a), "b:", unsafe.Pointer(&b))
	fmt.Println("a:", a, "b:", b)
}
