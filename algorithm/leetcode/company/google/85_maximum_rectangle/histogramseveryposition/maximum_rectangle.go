package histogramseveryposition

import "math"

// The kinda of brute force solution is precompute the prefix sum and then we iterate through
// matrix. For every point, we go over to left and up to find the number where the range sum is
// equal to the range size and then we update the results.
// A better way to do that is consider very point as horizontal histogram and at every point,
// we search upward to update the max rectangles
// Reference: https://leetcode.com/problems/maximal-rectangle/solution/
// Time: O(m^2*n)
// Space: O(m*n)
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	max := 0
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				if j-1 >= 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
			// process histgram
			width := dp[i][j]
			for k := i; k >= 0; k-- {
				width = minInt(width, dp[k][j])
				max = maxInt(max, (i-k+1)*width)
			}
		}
	}
	return max
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
