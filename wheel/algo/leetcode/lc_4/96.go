package lc_4

// 0, 1
// 1, f[0]
// 2, f[0]*f[1]+f[1]*f[0]
// 3, f[0]*f[2]+f[1]*f[1]+f[2]*f[0]
func numTrees(n int) int {
	f := make([]int, n+1)
	f[0], f[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			f[i] += f[j] * f[i-j-1]
		}
	}
	return f[n]
}
