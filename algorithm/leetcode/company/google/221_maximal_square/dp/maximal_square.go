package dp

import "math"

// dp[i][j] is the side length of max square locating at (i, j).
// By checking dp[i-1][j], dp[i][j-1] and dp[i-1][j-1]. We can get the
// side length of max square locating at (i, j).
// Time: O(m*n)
// Space: O(m*n) - can be improved to O(n)
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = minInt(dp[i][j+1], minInt(dp[i][j], dp[i+1][j])) + 1
				res = maxInt(res, dp[i+1][j+1])
			}
		}
	}
	return res * res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
