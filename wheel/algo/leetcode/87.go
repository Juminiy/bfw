package leetcode

func isScramble(s1 string, s2 string) bool {
	s1L, s2L := len(s1), len(s2)
	if s1L != s2L {
		return false
	}
	n := s1L
	dp := make([][][]bool, s1L)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
			dp[i][j][1] = s1[i] == s2[j]
		}
	}
	for k := 2; k <= n; k++ {
		for i := 0; i <= n-k; i++ {
			for j := 0; j <= n-k; j++ {
				for w := 1; w < k; w++ {
					if (dp[i][j][w] && dp[i+w][j+w][k-w]) ||
						(dp[i][j+k-w][w] && dp[i+w][j][k-w]) {
						dp[i][j][k] = true
						break
					}
				}
			}
		}
	}
	return dp[0][0][n]
}
