package lc_2

import (
	"errors"
)

var (
	chError = errors.New("char error")
)

type (
	trieNode struct {
		chi [26]*trieNode
		cnt int
		ter bool
	}
	Trie struct {
		root *trieNode
	}
)

func ConstructorTrie() Trie {
	return Trie{root: &trieNode{}}
}

func (t *Trie) Insert(word string) {
	search(t.root, word, 1)
}

func (t *Trie) Search(word string) bool {
	return search(t.root, word, 0).ter
}

func (t *Trie) StartsWith(prefix string) bool {
	return search(t.root, prefix, 0).cnt > 0
}

func search(root *trieNode, prefix string, cnt int) *trieNode {
	if root == nil || len(prefix) == 0 {
		return nil
	}
	var (
		ch  = prefix[0]
		pl  = len(prefix)
		ter = pl == 1 && cnt > 0
		cur = root.chi[getChIndex(ch)]
	)
	if cur == nil {
		cur = &trieNode{cnt: cnt, ter: ter}
	} else {
		cur.cnt += cnt
		if !cur.ter {
			cur.ter = ter
		}
	}
	root.chi[getChIndex(ch)] = cur
	if pl == 1 && cnt == 0 {
		return cur
	}
	return search(cur, prefix[1:], cnt)
}

func getChIndex(ch byte) int {
	if ch < 'a' || ch > 'z' {
		panic(chError)
	}
	return int(ch - 'a')
}
