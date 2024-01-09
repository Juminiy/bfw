package lc_3

import "math/bits"

func subsets(nums []int) [][]int {
	cnt := len(nums)
	res := make([][]int, 1<<cnt)
	for i := 0; i < 1<<cnt; i++ {
		tmp := make([]int, 0)
		for j := 0; j < cnt; j++ {
			if i&(1<<j) > 0 {
				tmp = append(tmp, nums[j])
			}
		}
		res[i] = tmp
	}
	return res
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < 1<<n; i++ {
		if bits.OnesCount(uint(i)) == k {
			tmp := make([]int, 0)
			for j := 0; j < n; j++ {
				if i&(1<<j) > 0 {
					tmp = append(tmp, j+1)
				}
			}
			res = append(res, tmp)
		}
	}
	return res
}
