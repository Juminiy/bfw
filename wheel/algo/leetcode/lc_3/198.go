package lc_3

// 3 1 1 9
func robQ1(nums []int) int {
	nl := len(nums)
	dp := make([]int, nl)

	if nl >= 1 {
		dp[0] = nums[0]
	}
	if nl >= 2 {
		dp[1] = max(nums[0], nums[1])
	}
	if nl >= 3 {
		for i := 2; i < nl; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
	}
	return dp[nl-1]
}
