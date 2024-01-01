package adt

import (
	"cmp"
	"errors"
)

var (
	lruListError = errors.New("lru list error")
)

type LRU[K cmp.Ordered, V any] struct {
	loc map[K]*GenericListNode[*Pair[K, V]]
	lis *GenericList[*Pair[K, V]]
	cap int
	zv0 V
}

func MakeLRU[K cmp.Ordered, V any](cap int, zv0 V) *LRU[K, V] {
	return &LRU[K, V]{
		loc: make(map[K]*GenericListNode[*Pair[K, V]]),
		lis: MakeGenericList[*Pair[K, V]](),
		cap: cap,
		zv0: zv0,
	}
}

func (lru *LRU[K, V]) Get(key K) V {
	if keyLoc := lru.get(key); keyLoc != nil {
		return keyLoc.Data.Val()
	}
	return lru.zv0
}

func (lru *LRU[K, V]) Put(key K, value V) {
	if keyLoc := lru.get(key); keyLoc != nil {
		keyLoc.Data.SetVal(value)
	} else {
		if !lru.hold() {
			lru.evict()
		}
		lru.append(MakeGenericListNode[*Pair[K, V]](MakePair[K, V](key, value)))
	}
}

func (lru *LRU[K, V]) get(key K) *GenericListNode[*Pair[K, V]] {
	if keyLoc := lru.loc[key]; keyLoc != nil {
		if keyLoc != lru.lis.Head() {
			if keyLoc == lru.lis.Tail() {
				lru.lis.DelTail()
			}
			keyLoc.SetSolitary()
			lru.append(keyLoc)
		}
		return keyLoc
	}
	return nil
}

func (lru *LRU[K, V]) append(n ...*GenericListNode[*Pair[K, V]]) {
	for _, nN := range n {
		lru.appendOne(nN)
	}
}

func (lru *LRU[K, V]) appendOne(n *GenericListNode[*Pair[K, V]]) {
	if n != nil {
		lru.lis.AddHead(n)
		lru.loc[n.Data.Key()] = n
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
		delTail := lru.lis.DelTail()
		delTail.SetSolitary()
		evicted = delTail.Data
		delete(lru.loc, evicted.Key())
	}
	return evicted
}

func (lru *LRU[K, V]) hold() bool {
	return len(lru.loc) < lru.cap
}
