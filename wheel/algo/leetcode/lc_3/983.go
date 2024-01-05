package lc_3

const (
	dpInf = 0xffffffff
)

func mincostTickets(days []int, costs []int) int {
	next := getNext(days, []int{1, 7, 30})
	dp := make([]int, len(days))
	for i, _ := range dp {
		dp[i] = dpInf
	}
	//return mctV2(costs, dp, next, 0)
	//pre := getPre(days, []int{1, 7, 30})
	return mctV3(costs, dp, next)
}

func mctV3(costs, dp []int, next [][]int) int {
	ei, cur := len(dp)-1, 0
	for j := ei; j >= 0; j-- {
		for i, cost := range costs {
			if next[j][i] == -1 {
				cur = 0
			} else {
				cur = dp[next[j][i]]
			}
			dp[j] = min(dp[j], cur+cost)
		}
	}
	return dp[0]
}

func mctV2(costs, dp []int, next [][]int, cur int) int {
	if cur == -1 {
		return 0
	}
	if dp[cur] != dpInf {
		return dp[cur]
	}
	dp[cur] = dpInf
	for i, cost := range costs {
		dp[cur] = min(dp[cur], mctV2(costs, dp, next, next[cur][i])+cost)
	}
	return dp[cur]
}

// 1. 写出暴力的dfs算法
// days = [1,4,6,7,8,20], costs = [2,7,15]
func mctV1(days []int, c1, c2, c3 int, res int) int {
	dL := len(days)
	if dL == 0 {
		return res
	}
	r1 := mctV1(days[1:], c1, c2, c3, res+c1)
	end7, end30 := 0, 0
	for ; end7 < dL && days[end7]-days[0] < 7; end7++ {
	}
	r2 := mctV1(days[end7:], c1, c2, c3, res+c2)

	for ; end30 < dL && days[end30]-days[0] < 30; end30++ {
	}
	r3 := mctV1(days[end30:], c1, c2, c3, res+c3)
	return min(r1, r2, r3)
}

// days = [1,4,6,7,8,20], costs = [2,7,15]
// index= [0,1,2,3,4,5]
// pre =  [[0 0 0] [0 0 0] [1 0 0] [2 0 0] [3 0 0] [4 4 0]]
func getPre(days, dayDuration []int) [][]int {
	dL, ddL := len(days), len(dayDuration)
	pre := make([][]int, dL)
	for i := dL - 1; i >= 0; i-- {
		pre[i] = make([]int, ddL)
		for k := 0; k < ddL; k++ {
			for j := i - 1; j >= 0; j-- {
				if days[i]-days[j] >= dayDuration[k] {
					pre[i][k] = j
					break
				}
			}
		}
	}
	return pre
}

// days = [1,4,6,7,8,20], costs = [2,7,15]
// index= [0,1,2,3,4,5]
// next = [[1 4 0] [2 5 0] [3 5 0] [4 5 0] [5 5 0] [0 0 0]]
func getNext(days, dayDuration []int) [][]int {
	dL, ddL := len(days), len(dayDuration)
	next := make([][]int, dL)
	for i := 0; i < dL; i++ {
		next[i] = make([]int, ddL)
		for k := 0; k < ddL; k++ {
			for j := i + 1; j < dL; j++ {
				if days[j]-days[i] >= dayDuration[k] {
					next[i][k] = j
					break
				}
			}
			if next[i][k] == 0 {
				next[i][k] = -1
			}
		}
	}
	return next
}
