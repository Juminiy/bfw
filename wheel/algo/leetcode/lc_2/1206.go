package lc_2

import (
	"bfw/wheel/lang"
	"math/rand"
)

var (
	dummyHead = &skipNode{key: -1}
	dummyTail = &skipNode{key: 0xffffffff}
)

type (
	skipNode struct {
		key, cnt int
		next     []*skipNode
	}
	Skiplist struct {
		head, tail *skipNode
		height     int
		size       int
	}
)

// range from [0,20000]
func ConstructorSkiplist() Skiplist {
	dummyHead.next = append(dummyHead.next, dummyTail)
	return Skiplist{head: dummyHead, tail: dummyTail}
}

func (l *Skiplist) Add(num int) {
	prev := l.walk(num)
	if prev.next[0].key == num {
		prev.next[0].cnt++
	} else {

	}
}

func (l *Skiplist) Search(target int) bool {
	prev := l.walk(target)
	return prev.next[0] != nil &&
		prev.next[0].key == target
}

func (l *Skiplist) Erase(num int) bool {
	return false
}

func (l *Skiplist) walk(num int) *skipNode {
	walkNode := l.head
	curLevel := l.height
	for walkNode != nil && curLevel >= 0 {
		walkNext := walkNode.next[curLevel]
		for walkNext != nil && walkNext.key < num {
			walkNode = walkNext
			walkNext = walkNext.next[curLevel]
		}
		curLevel--
	}
	return walkNode
}

func getProb(size int) []bool {
	total := lang.CeilBinCnt(size)
	prob := make([]bool, total)
	for i := 0; i < total; i++ {
		prob[i] = rand.Intn(1<<i) == 1
	}
	return prob
}
