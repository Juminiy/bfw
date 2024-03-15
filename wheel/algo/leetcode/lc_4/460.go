package lc_4

// 2 1 2 1 2 3 4 , 3
// LRU: 4->3->2
// LRU是将一个当前访问的元素（无论是否存在）迅速放到队头
// 容量满时，移除队尾元素
// LFU: 2(3),1(2),
// LFU是将一个新元素迅速放到队尾,其他元素按照访问次数排好顺序，如果访问次数均等,最久未被访问应在后
// 容量满时，移除队尾元素

type (
	lfuNode struct {
		k, v, c int
	}
	LFUCache struct {
		m map[int]*lfuNode
		o int
	}
)

func Constructor(capacity int) LFUCache {
	return LFUCache{}
}

func (l *LFUCache) Get(key int) int {
	return -1
}

func (l *LFUCache) Put(key int, value int) {

}
