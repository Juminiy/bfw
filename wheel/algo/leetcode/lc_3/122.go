package lc_3

// 7,1,5,3,6,4
// 7,1,
func maxProfitQ2(prices []int) int {
	curMin, totProf, curProf := prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-curMin > curProf {
			totProf += prices[i] - curMin - curProf
			curProf = prices[i] - curMin
		} else {
			curProf = 0
			curMin = prices[i]
		}
	}
	return totProf
}
