package leetcode

import "bfw/wheel/adt"

func nthUglyNumber(n int) int {
	//h := adt.MakeIntHeap(true)
	h := adt.MakeHeap[int](func(a int, b int) bool {
		return a < b
	})
	m := make(map[int]bool)
	e := 1
	h.Push(e)
	for n > 0 {
		n--
		e = h.Pop(1)[0]
		if !m[e<<1] {
			h.Push(e << 1)
			m[e<<1] = true
		}
		if !m[e<<1+e] {
			h.Push(e<<1 + e)
			m[e<<1+e] = true
		}
		if !m[e<<2+e] {
			h.Push(e<<2 + e)
			m[e<<2+e] = true
		}

	}
	return e
}

func getUglyNumberSlice(n int) []int {
	uglySlice := make([]int, 0)
	for i := 1; i <= n; i++ {
		u := nthUglyNumber(i)
		uglySlice = append(uglySlice, u)
	}
	return uglySlice
}

func nthUglyNumberV2(n int) int {
	a, p2, p3, p5 := make([]int, n+1), 1, 1, 1
	a[1] = 1
	for i := 2; i <= n; i++ {
		a2, a3, a5 := a[p2]*2, a[p3]*3, a[p5]*5
		a[i] = min(a2, a3, a5)
		if a[i] == a2 {
			p2++
		}
		if a[i] == a3 {
			p3++
		}
		if a[i] == a5 {
			p5++
		}
	}
	return a[n]
}
