package adt

import (
	"cmp"
	"errors"
)

var (
	lruError = errors.New("lru error")
)

type (
	lruNode[K cmp.Ordered, V any] struct {
		GenericListNode[*Pair[K, V]]
	}

	LRU[K cmp.Ordered, V any] struct {
		loc map[K]*Pair[K, V]
		lis *GenericList[*Pair[K, V]]
		cap int
		siz int
		zv0 V
	}
)

func MakeLRU[K cmp.Ordered, V any](cap int, zv0 V) *LRU[K, V] {
	return &LRU[K, V]{
		loc: make(map[K]*Pair[K, V]),
		lis: MakeGenericList[*Pair[K, V]](),
		cap: cap,
		siz: 0,
		zv0: zv0,
	}
}

func (lru *LRU[K, V]) Get(key K) V {
	if keyLoc := lru.get(key); keyLoc != nil {
		return keyLoc.Val()
	}
	return lru.zv0
}

func (lru *LRU[K, V]) Put(key K, value V) {
	if keyLoc := lru.get(key); keyLoc != nil {
		keyLoc.SetVal(value)
	} else {
		if !lru.hold() {
			lru.evict()
		}
		lru.append(MakePair[K, V](key, value))
	}
}

func (lru *LRU[K, V]) get(key K) *Pair[K, V] {
	if keyLoc := lru.loc[key]; keyLoc != nil {
		if keyLoc != lru.lis.Front() {

			lru.append(keyLoc)
		}
		return keyLoc
	}
	return nil
}

func (lru *LRU[K, V]) append(n ...*Pair[K, V]) {
	for _, nN := range n {
		lru.appendOne(nN)
	}
}

func (lru *LRU[K, V]) appendOne(n *Pair[K, V]) {
	if n != nil {
		lru.siz++
		lru.lis.PushFront(n)
		lru.loc[n.key] = n
	}
}

func (lru *LRU[K, V]) evict(n ...int) []*Pair[K, V] {
	var evicted []*Pair[K, V]
	nN := 1
	if len(n) > 0 && n[0] >= 0 {
		nN = n[0]
	}
	for !lru.lis.Empty() && nN > 0 {
		nN--
		evicted = append(evicted, lru.evictOne())
	}

	return evicted
}

func (lru *LRU[K, V]) evictOne() *Pair[K, V] {
	var evicted *Pair[K, V]
	if !lru.lis.Empty() {
		evicted = lru.lis.Back()
		delete(lru.loc, evicted.Key())
	}
	return evicted
}

func (lru *LRU[K, V]) hold() bool {
	return lru.siz < lru.cap
}
