package lc_3

// 7,1,5,3,6,4
// 7,1,
func maxProfitQ2(prices []int) int {
	pl := len(prices)
	dp := make([]int, pl)
	curMin := 0
	for i := range prices {
		dp[i] = max(dp[i], dp[curMin]+prices[i]-prices[curMin])
		if prices[i] < prices[curMin] {
			curMin = i
		}
	}
	return dp[pl-1]
}
