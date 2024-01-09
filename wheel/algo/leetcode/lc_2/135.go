package lc_2

// 10 9 8 7 6 5 6 7 8 9 10 9 8 7 6 5 4 3 2 1

// 10 9 8 3 4 2 1 0 5 6 3 1
// 4  3 2 1 4 3 2 1 2 3 2 1

// 1 2 3 4 5 2 1

// 1 2 5 4 3 2 1

// 5 4 3 2 1 2 3

// 5 4 1 2 3 4 5

// 恶心 我已经不想再考虑数据特征了
func candy(ratings []int) int {
	rl := len(ratings)
	tot := 0
	nextSeq := func(starti int) (si, ei int, desc bool) {
		si, ei = starti, starti
		if ei < rl && ei+1 < rl {
			desc = ratings[ei] >= ratings[ei+1]
		}
		for ei < rl && ei+1 < rl && ratings[ei] >= ratings[ei+1] == desc {
			ei++
		}
		return
	}
	si, ei, _ := nextSeq(0)
	for ei < rl {
		if ei == rl-1 {
			break
		}
		nsi, nei, _ := nextSeq(ei)
		if nei == rl-1 {
			break
		}
		llen, rlen := ei-si, nei-nsi
		if llen >= rlen {

		}
		si, ei = nsi, nei
		//fmt.Println(ei-si+1, ndesc)
	}
	return tot
}
