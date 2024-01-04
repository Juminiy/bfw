package beyond

// ST
// slice: st table
//
//	(
//	slice[index][step]
//	query
//	range: '[' index, index+(1<<step) ')'
//
// )
// fn: trinocular operator function (c = a opt b)
type ST struct {
	slice [][]int
	log2  []int
	fn    func(int, int) int
}

func MakeST(slice []int, fn func(int, int) int) *ST {
	st := &ST{}
	st.fn = fn
	st.make(slice)
	return st
}

func (st *ST) clear() {
	st.slice = nil
	st.log2 = nil
}

func (st *ST) make(slice []int) {
	sLen := len(slice)
	st.clear()
	st.slice = make([][]int, sLen)
	st.log2 = GenIntLog2(sLen)
	for step := 0; step <= st.log2[sLen]+1; step++ {
		for index := 0; index+(1<<step) <= sLen; index++ {
			if len(st.slice[index]) == 0 {
				st.slice[index] = make([]int, st.log2[sLen]+1)
			}
			if step == 0 {
				st.slice[index][step] = slice[index]
			} else {
				st.slice[index][step] = st.fn(st.slice[index][step-1], st.slice[index+1<<(step-1)][step-1])
			}
		}
	}
}

// Query
// range: '[' l, r ']'
func (st *ST) Query(l, r int) int {
	return st.query(l, r)
}

// query
// range: [l,r]
// step = k
// range: [l, l+(1<<step)] + [r-(1<<step), r]
func (st *ST) query(l, r int) int {
	step := st.log2[r-l+1]
	return st.fn(st.slice[l][step],
		st.slice[r-(1<<step)+1][step])
}

// GenIntLog2
// log2(1) = 0
// log2(2) = 1
// log2(3) = 1
// log2(4) = 2
func GenIntLog2(n int) []int {
	iLog2 := make([]int, n+1)
	iLog2[0], iLog2[1] = 0, 0
	bit, num, maxNum := 0, 0, n
	for ; (1<<bit) <= maxNum &&
		num <= maxNum; num++ {
		if (1 << (bit + 1)) == num {
			bit++
		}
		iLog2[num] = bit
		//fmt.Println("log2("+strconv.Itoa(num)+") = ", strconv.Itoa(bit))
	}
	return iLog2
}
