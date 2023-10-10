package adt

import (
	"fmt"
	"testing"
)

// TestGenericList_Back
// 3 1 2 4
func TestGenericList_Back(t *testing.T) {
	list := &GenericList[int]{}
	list.PushFront(1)
	list.PushBack(2)
	list.ForwardTraverse(fmt.Println)
	//fmt.Println("--------------------")
	list.ReverseTraverse(fmt.Println)
	//fmt.Println("--------------------")
	list.PopFront()
	list.ForwardTraverse(fmt.Println)
	//fmt.Println("--------------------")
	list.ReverseTraverse(fmt.Println)
	//fmt.Println("--------------------")
	list.PopBack()
	list.ForwardTraverse(fmt.Println)
	//fmt.Println("--------------------")
	list.ReverseTraverse(fmt.Println)
	//fmt.Println("--------------------")
	//list.Assign(nil)
}

func TestGenericList_Assign(t *testing.T) {
	l1 := &GenericList[int]{}
	l1.ChainedPushBack(1).ChainedPushBack(2).ChainedPushBack(3)
	l2 := GenericList[int]{}
	l2.ChainedPushBack(4).ChainedPushBack(5).ChainedPushBack(6)
	l1.ForwardTraverse(fmt.Println)
}
