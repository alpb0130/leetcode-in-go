package dp

func NumOfPath(m, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	dp := make([][]int, m)
	for idx := range dp {
		dp[idx] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
