package lc_3

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for j := 1; j <= n; j++ {
		dp[j] = 10001
		sj := int(math.Sqrt(float64(j)))
		for i := 1; i <= sj; i++ {
			dp[j] = min(dp[j], 1+dp[j-i*i])
		}
	}
	return dp[n]
}

func dfsNS(dp []int, n int) int {
	if dp[n] != 10001 {
		return dp[n]
	}
	for i := 1; i < n; i++ {
		dp[n] = min(dp[n], dfsNS(dp, i)+dfsNS(dp, n-i))
	}
	return dp[n]
}
