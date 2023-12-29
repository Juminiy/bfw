package adt

import (
	"fmt"
)

const (
	defaultPairComma = ','
)

type Pair[K, V any] struct {
	key K
	val V
}

func MakePair[K, V any](k K, v V) *Pair[K, V] {
	return &Pair[K, V]{k, v}
}

func (p *Pair[K, V]) Key() K {
	return p.key
}

func (p *Pair[K, V]) Val() V {
	return p.val
}

func (p *Pair[K, V]) Get() (K, V) {
	return p.key, p.val
}

func (p *Pair[K, V]) SetKey(k K) {
	p.key = k
}

func (p *Pair[K, V]) SetVal(v V) {
	p.val = v
}

func (p *Pair[K, V]) Set(k K, v V) {
	p.SetKey(k)
	p.SetVal(v)
}

type IntPair Pair[int, int]

func MakeIntPair(a, b int) *IntPair {
	return &IntPair{a, b}
}

func (ip *IntPair) Self() *IntPair {
	return ip
}

func (ip *IntPair) Assign(ipt *IntPair) {
	ip.SetKV(ipt.GetKV())
}

func (ip *IntPair) GetKV() (int, int) {
	return ip.key, ip.val
}

func (ip *IntPair) GetKey() int {
	return ip.key
}

func (ip *IntPair) GetVal() int {
	return ip.val
}

func (ip *IntPair) SetKV(key, val int) *IntPair {
	return ip.SetKey(key).SetVal(val)
}

func (ip *IntPair) SetKey(key int) *IntPair {
	ip.key = key
	return ip
}

func (ip *IntPair) SetVal(val int) *IntPair {
	ip.val = val
	return ip
}

func (ip *IntPair) SetKVSwap() *IntPair {
	ip.key, ip.val = ip.val, ip.key
	return ip
}

func (ip *IntPair) Display(comma ...rune) *IntPair {
	destComma := defaultPairComma
	if len(comma) > 0 {
		destComma = comma[0]
	}
	fmt.Printf("(%v%c%v)", ip.key, destComma, ip.val)
	return ip
}

type IntPairSlice []*IntPair

func MakeIntPairSlice() IntPairSlice {
	return IntPairSlice{}
}

func (ips IntPairSlice) Less(i, j int) bool {
	return (ips[i].key == ips[j].key && ips[i].val < ips[j].val) ||
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
