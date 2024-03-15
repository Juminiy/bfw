package lc_4

func totalNQueens(n int) int {
	tot := 0
	checkNQ := func(s []int, x, y int) bool {
		for _, si := range s {
			if si == y {
				return false
			}
		}
		m, k := x-y, x+y
		for mx := x - 1; mx >= 0; mx-- {
			if s[mx] == mx-m {
				return false
			}
		}
		for kx := x - 1; kx >= 0; kx-- {
			if s[kx] == k-kx {
				return false
			}
		}
		return true
	}
	dfsNQ(n, 0, &tot, make([]int, 0), checkNQ)
	return tot
}

/*
O X O O
O O 0 X
X O O O
O O X O
*/
func dfsNQ(n, c int, t *int, s []int, fn func([]int, int, int) bool) {
	if c == n {
		*t += 1
		return
	}
	for i := 0; i < n; i++ {
		if fn(s, c, i) {
			dfsNQ(n, c+1, t, append(s, i), fn)
		}
	}
}

// s = [1, 3, 0, 2]
// s'= [2, 4, 1, 3]
