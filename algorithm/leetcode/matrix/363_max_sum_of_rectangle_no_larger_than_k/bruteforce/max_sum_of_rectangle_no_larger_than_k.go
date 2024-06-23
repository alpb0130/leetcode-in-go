package bruteforce

import "math"

// Brute force + some dynamic programming (store the prefix sum)
// Time: O((m*n)^2)
// Space: O(m*n)
func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	max := math.MinInt32
	prefixSums := make([][]int, m)
	for i := range prefixSums {
		prefixSums[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// compute prefix sum
			prefixSums[i][j] += matrix[i][j]
			if i > 0 {
				prefixSums[i][j] += prefixSums[i-1][j]
			}
			if j > 0 {
				prefixSums[i][j] += prefixSums[i][j-1]
			}
			if i > 0 && j > 0 {
				prefixSums[i][j] -= prefixSums[i-1][j-1]
			}
			for h := 0; h <= i; h++ {
				for l := 0; l <= j; l++ {
					curSum := prefixSums[i][j]
					if h > 0 {
						curSum -= prefixSums[h-1][j]
					}
					if l > 0 {
						curSum -= prefixSums[i][l-1]
					}
					if h > 0 && l > 0 {
						curSum += prefixSums[h-1][l-1]
					}
					if curSum > max && curSum <= k {
						max = curSum
					}
				}
			}
		}
	}
	return max
}
