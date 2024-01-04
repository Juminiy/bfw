package lc_2

import (
	"bfw/wheel/lang"
	"errors"
	"math/rand"
)

var (
	dummyHead             = &skipNode{key: -1}
	dummyTail             = &skipNode{key: 0xffffffff}
	skipListLevelIncError = errors.New("skip list level inc error")
	skipListLevelError    = errors.New("skip list level error")
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
	if next := prev.next[0]; next.key == num {
		// found, inc cnt directly
		next.cnt++
	} else {
		newSkipNode := &skipNode{key: num, cnt: 1}
		l.size++

		// random the height
		h := getHeight(l.size)

		if inc := h - l.height; inc >= 2 {
			panic(skipListLevelIncError)
		} else if inc == 1 {
			l.growLevel()
		} else {

		}

		newSkipNode.next = make([]*skipNode, h)

		// insert into level h,h-1,...,0
		walkNode := l.head
		for h >= 0 {
			walkNode = l.walkInLevel(walkNode, h, num)
			newSkipNode.next[h] = walkNode.next[h]
			walkNode.next[h] = newSkipNode
			h--
		}

	}
}

func (l *Skiplist) Search(target int) bool {
	prev := l.walk(target)
	_, ex := l.exists(prev, target)
	return ex
}

func (l *Skiplist) Erase(num int) bool {
	var cur *skipNode
	prev := l.walk(num)
	cur, ex := l.exists(prev, num)
	if ex {
		cur.cnt--
		if cur.cnt == 0 {

		}
	}
	return ex
}

func (l *Skiplist) walk(num int) *skipNode {
	walkNode := l.head
	curLevel := l.height
	for walkNode != nil && curLevel >= 0 {
		walkNode = l.walkInLevel(walkNode, curLevel, num)
		curLevel--
	}
	return walkNode
}

func (l *Skiplist) walkInLevel(walkNode *skipNode, level int, target int) *skipNode {
	if level > l.height {
		panic(skipListLevelError)
	}
	walkNext := walkNode.next[level]
	for walkNext != nil && walkNext.key < target {
		walkNode = walkNext
		walkNext = walkNext.next[level]
	}
	return walkNode
}

func (l *Skiplist) growLevel() {
	l.height++
	l.head.next = append(l.head.next, l.tail)
}

func (l *Skiplist) exists(prev *skipNode, target int) (*skipNode, bool) {
	var cur *skipNode
	if prev != nil && len(prev.next) > 0 {
		cur = prev.next[0]
	}
	return cur, prev != nil &&
		cur != nil &&
		cur.key == target &&
		cur.next[0].cnt > 0
}

func getProb(size int) []bool {
	total := lang.CeilBinCnt(size)
	prob := make([]bool, total)
	for i := 0; i < total; i++ {
		prob[i] = rand.Intn(1<<i) == 1
	}
	return prob
}

// [0,h]
func getHeight(size int) int {
	return rand.Intn(lang.CeilBinCnt(size))
}
