package lc_1

import "bfw/wheel/adt"

type LRUCache struct {
	lru adt.LRU[int, int]
}

func Constructor(capacity int) LRUCache {
	lru := adt.MakeLRU[int, int](capacity, -1)
	return LRUCache{lru: *lru}
}

func (c *LRUCache) Get(key int) int {
	return c.lru.Get(key)
}

func (c *LRUCache) Put(key int, value int) {
	c.lru.Put(key, value)
}
