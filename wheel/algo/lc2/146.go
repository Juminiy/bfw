package lc2

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
	}
)

func Constructor(capacity int) LRUCache {
	c := LRUCache{}
	c.loc = make(map[int]*lruNode)
	return c
}

func (c *LRUCache) Get(key int) int {
	if keyLoc := c.get(key); keyLoc != nil {
		if keyLoc != c.head && keyLoc != c.tail {
			keyLoc.next = c.head
			c.head = keyLoc
		}
		return keyLoc.v
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if keyLoc := c.loc[key]; keyLoc != nil {

	}
}

func (c *LRUCache) get(key int) *lruNode {
	if keyLoc := c.loc[key]; keyLoc != nil {
		prev, next := keyLoc.prev, keyLoc.next
		if prev != nil {
			prev.next = next
		}
		if next != nil {
			next.prev = prev
		}
		keyLoc.prev = nil
		keyLoc.next = nil
		return keyLoc
	}
	return nil
}

func (c *LRUCache) put(key, value int) {

}
