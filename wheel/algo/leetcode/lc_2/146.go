package lc_2

import (
	"errors"
)

var (
	lruNodeError = errors.New("lru node error")
)

// 4
// nil<->1<->nil
// nil<->2<->1<->nil
type (
	lruNode struct {
		k, v       int
		prev, next *lruNode
	}
	LRUCache struct {
		head, tail *lruNode
		loc        map[int]*lruNode
		cap        int
	}
)

func Constructor(capacity int) LRUCache {
	c := LRUCache{}
	c.loc = make(map[int]*lruNode)
	c.cap = capacity
	return c
}

func (c *LRUCache) Get(key int) int {
	if keyLoc := c.get(key); keyLoc != nil {
		return keyLoc.v
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if keyLoc := c.get(key); keyLoc != nil {
		keyLoc.v = value
	} else {
		if !c.canHold() {
			c.evict()
		}
		c.append(&lruNode{k: key, v: value})
	}
}

// 1, 1, #
// 2<->1, 2, #
// 3<->2<->1, 3, #

// 3<->2<->1, 2, 2<->3<->1

// 2<->1, 1
// 3<->2<->1, 1

func (c *LRUCache) get(key int) *lruNode {
	if keyLoc := c.loc[key]; keyLoc != nil {
		if keyLoc != c.head {
			if keyLoc == c.tail {
				c.tail = c.tail.prev
			}
			keyLoc.setSolitary()
			c.append(keyLoc)
		}
		return keyLoc
	}
	return nil
}

func (c *LRUCache) append(n ...*lruNode) {
	for _, nN := range n {
		c.appendOne(nN)
	}
}

// 2<->1
func (c *LRUCache) appendOne(n *lruNode) {
	if n != nil {
		if c.nValidate() {
			c.head = n
			c.tail = n
			c.loc[n.k] = n
		} else if c.validate() {
			n.next = c.head
			c.head.prev = n
			if len(c.loc) == 1 {
				c.tail = c.head
			}
			c.head = n
			c.loc[n.k] = n
		} else {
			panic(lruNodeError)
		}
	}
}

func (c *LRUCache) evict(n ...int) []*lruNode {
	var evicted []*lruNode
	nN := 1
	if len(n) > 0 && n[0] >= 0 {
		nN = n[0]
	}
	for c.head != nil && nN > 0 {
		nN--
		evicted = append(evicted, c.evictOne())
	}

	return evicted
}

// 3<->2<->1
// 3<->2
// 3
// nil
func (c *LRUCache) evictOne() *lruNode {
	var evicted *lruNode
	if c.tail != nil {
		evicted = c.tail
		c.tail = c.tail.prev
		evicted.setSolitary()
		delete(c.loc, evicted.k)
		if len(c.loc) <= 1 {
			c.head = c.tail
		}
	}
	return evicted
}

func (c *LRUCache) validate() bool {
	return c.head != nil && c.tail != nil
}

func (c *LRUCache) nValidate() bool {
	return c.head == nil && c.tail == nil
}

func (c *LRUCache) canHold() bool {
	return len(c.loc) < c.cap
}

func (c *lruNode) setSolitary() {
	if c.prev != nil {
		c.prev.next = c.next
	}
	if c.next != nil {
		c.next.prev = c.prev
	}
	c.prev = nil
	c.next = nil
}
