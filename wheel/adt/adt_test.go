package adt

import (
	"bfw/wheel/lang"
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
	fmt.Println("--------------------")
	list.ReverseTraverse(fmt.Println)
	fmt.Println("--------------------")
	list.PopFront()
	list.ForwardTraverse(fmt.Println)
	fmt.Println("--------------------")
	list.ReverseTraverse(fmt.Println)
	fmt.Println("--------------------")
	list.PopBack()
	list.ForwardTraverse(fmt.Println)
	fmt.Println("--------------------")
	list.ReverseTraverse(fmt.Println)
	fmt.Println("--------------------")
	//list.Assign(nil)
}

func TestGenericList_Assign(t *testing.T) {
	l1 := &GenericList[int]{}
	l1.ChainedPushBack(1).ChainedPushBack(2).ChainedPushBack(3)
	l2 := GenericList[int]{}
	l2.ChainedPushBack(4).ChainedPushBack(5).ChainedPushBack(6)
	l1.ForwardTraverse(fmt.Println)
}

func TestMakeBitMap(t *testing.T) {
	fmt.Println(bitValMax)
	fmt.Println(bitValMin)
	fmt.Println(int64(bitValSignMax))
	fmt.Println(bitValSignMin)
	bm := MakeBitMap(false, false)
	bm.clear()
}

func TestGenericList_At(t *testing.T) {
	fmt.Println(449>>6, 449&0b111111)
	fmt.Println(7<<6 + 1)
	val := 0xfeff
	//0xfeff
	//0xff8f
	val &= 0xffff - 1<<7
	fmt.Printf("%x, %x\n", 0xffff-1<<7, val)
	// 262144
	fmt.Println("2^18 = ", 1<<18)
	//65,536
	fmt.Println(1 << 16)
	fmt.Printf("%d\n", int64(1e7))
	//16,777,216
	fmt.Println(1 << 24)
	//4,294,967,296
	fmt.Println(1 << 32)
	//281,474,976,710,656
	fmt.Println(1 << 48)
	//18,446,744,073,709,551,615
	fmt.Println(bitValMax) // 1<<64-1
}

func TestBitMap_Array(t *testing.T) {
	bitMap := MakeBitMap(true, true)
	arr := lang.GenerateInt641DArray(1<<5, 1<<7)
	bitMap.batchInsert(arr...)
	fmt.Println(arr)
	fmt.Println(bitMap.Array())
	bitMap.Insert(8).Insert(8)
	fmt.Println(bitMap.Count(8))
}

func TestItemLess(t *testing.T) {
	//i1 := Int(5)
	//i2 := Int(8)
	fmt.Println(ItemLess(GetInf(-1), GetInf(+1)))
}

// Deprecated: Use TestMakeHeap instead.
func TestMakeIntHeap(t *testing.T) {
	//h := MakeIntHeap(true)
	//for i := 10; i >= 0; i-- {
	//	h.Push(i)
	//	//h.Print()
	//}
	//for i := 0; i < 10; i++ {
	//	e := h.Pop(1)[0]
	//	fmt.Printf("%d ", e)
	//	//h.Print()
	//}
}

func TestMakeHeap(t *testing.T) {
	h := MakeHeap[int](func(a int, b int) bool {
		return a < b
	})
	for i := 10; i >= 0; i-- {
		h.Push(i)
		//h.Print()
	}
	for i := 0; i < 10; i++ {
		e := h.Pop(1)[0]
		fmt.Printf("%d ", e)
		//h.Print()
	}
}

func TestMakeLRU(t *testing.T) {
	c := MakeLRU[int, int](2, -1)

	//case 1
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))
	c.Put(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))

	//case 2
	//c := MakeLRU[int, int](1, -1)
	//c.Put(2, 1)
	//fmt.Println(c.Get(2))
	//c.Put(3, 2)
	//fmt.Println(c.Get(2))
	//fmt.Println(c.Get(3))
}
