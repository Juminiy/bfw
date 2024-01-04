package lc_3

func mincostTickets(days []int, costs []int) int {
	//next := getNext(days, []int{1, 7, 30})
	//return mctV2(costs, next, 0, 0)
	pre := getPre(days, []int{1, 7, 30})
	return mctV3(costs, pre)
}

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

func mctV2(costs []int, next [][]int, cur, res int) int {
	if cur == -1 {
		return res
	}
	minRes := 0xffffffff
	for i, cost := range costs {
		minRes = min(minRes, mctV2(costs, next, next[cur][i], res+cost))
	}
	return minRes
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

	//fmt.Println(next)

	return next
}

func mctV3(costs []int, pre [][]int) int {
	//dp[i][j] =
	//	min
	//dp[i][j-1] + cost[0],
	//	dp[i][j-per[j]] + cost[1]
	//dp[i][j-pre2[j]] + cost[2]
	//dp := make([]int, 0)
	//fmt.Println(pre)
	//for i, _ := range {}
	return 0
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
