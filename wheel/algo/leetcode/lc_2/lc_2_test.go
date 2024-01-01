package lc_2

import (
	"fmt"
	"testing"
)

func TestLC1(t *testing.T) {
	n := BuildTreeNodeByBFS(
		"1",
		"2", "3",
		"4", "#", "#", "6",
		"7", "#", "#", "#", "#", "#", "#", "8")
	//n := BuildTreeNodeByBFS(
	//	"1", "2", "3", "#", "#", "4", "5")
	connectV3(n)
	n.PrintNext()
}

// cap = 3
// 3<->2<->1
// get(3)
// 3<->2<->1
// get(2)
// 2<->3<->1
// get(1)
// 1<->2<->3

// cap = 2
// 2<->1
// get(1)
// 1<->2
func TestLC2(t *testing.T) {
	//case 1
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))
	c.Put(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
	//1
	//-1
	//-1
	//3
	//4

	//case 2
	//c := Constructor(1)
	//c.Put(2, 1)
	//fmt.Println(c.Get(2))
	//c.Put(3, 2)
	//fmt.Println(c.Get(2))
	//fmt.Println(c.Get(3))
}

func TestLC3(t *testing.T) {
	trie := ConstructorTrie()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.Search("cal"))
}
