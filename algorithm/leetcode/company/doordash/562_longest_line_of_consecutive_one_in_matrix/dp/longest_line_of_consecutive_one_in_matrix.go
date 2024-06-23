package dp

import "math"

// Don't forget the anti-diagonal
// Time: O(m*n)
// Space: O(m*n)
func longestLine(M [][]int) int {
	if len(M) == 0 || len(M[0]) == 0 {
		return 0
	}
	dp := make([][][]int, len(M)+2)
	for i := range dp {
		dp[i] = make([][]int, len(M[0])+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4)
		}
	}
	res := 0
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); j++ {
			if M[i][j] == 1 {
				dp[i+1][j+1][0] = dp[i+1][j][0] + 1
				res = maxInt(res, dp[i+1][j+1][0])
				dp[i+1][j+1][1] = dp[i][j+1][1] + 1
				res = maxInt(res, dp[i+1][j+1][1])
				dp[i+1][j+1][2] = dp[i][j][2] + 1
				res = maxInt(res, dp[i+1][j+1][2])
				dp[i+1][j+1][3] = dp[i][j+2][3] + 1
				res = maxInt(res, dp[i+1][j+1][3])
			}
		}
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
