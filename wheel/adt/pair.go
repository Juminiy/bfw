package adt

import "fmt"

const (
	defaultPairComma = ','
)

type Pair[K, V any] struct {
	key   K
	value V
}

type IntPair Pair[int, int]

func MakeIntPair(a, b int) *IntPair {
	return &IntPair{a, b}
}

func (ip *IntPair) Display(comma ...rune) *IntPair {
	destComma := defaultPairComma
	if len(comma) > 0 {
		destComma = comma[0]
	}
	fmt.Printf("(%v%c%v)", ip.key, destComma, ip.value)
	return ip
}

type IntPairSlice []*IntPair

func MakeIntPairSlice() IntPairSlice {
	return IntPairSlice{}
}

func (ips IntPairSlice) Less(i, j int) bool {
	return (ips[i].key == ips[j].key && ips[i].value < ips[j].value) ||
		ips[i].key < ips[j].key
}

func (ips IntPairSlice) Swap(i, j int) {
	ips[i], ips[j] = ips[j], ips[i]
}

func (ips IntPairSlice) Len() int {
	return ips.size()
}

func (ips IntPairSlice) self() IntPairSlice {
	return ips
}

func (ips IntPairSlice) size() int {
	return len(ips)
}

func (ips IntPairSlice) Display() IntPairSlice {
	fmt.Printf("[")
	if ips.size() > 0 {
		ips[0].Display()
		for idx := 1; idx < ips.size(); idx++ {
			fmt.Printf(", ")
			ips[idx].Display()
		}
	} else {
		fmt.Printf("null")
	}
	fmt.Println("]")
	return ips
}
