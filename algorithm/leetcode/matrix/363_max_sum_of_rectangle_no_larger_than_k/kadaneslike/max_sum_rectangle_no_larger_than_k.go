package kadaneslike

import "math"

// Extend the problem of finding max subarray (max value) to 2D. Nested loop to iterate the
// column index and for each outer loop, use an array to store the prefix sum and get the max.
// Time: O((m*n)^2) - can be improved to O(m*logm*n*n) if we can use something like a treeset
// where insert and search are both O(logm) time
// Space: O(m)
func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	max := math.MinInt32
	for i := 0; i < n; i++ {
		// store the sum for each row from column i to column j
		sums := make([]int, m)
		for j := i; j < n; j++ {
			// update sums
			for k := 0; k < m; k++ {
				sums[k] += matrix[k][j]
			}
			// find the max sum in quadratic time and update max
			curMax := maxNoLargerThanKInLinear(sums, k)
			if curMax <= k && curMax > max {
				max = curMax
			}
		}
	}
	return max
}

// Time: O(m^2) - we can improve the complexity to O(m*logm). Use some structure like treeset
// where insert and search are both logarithm. Then we insert every prefix into the treeset (logm)
// and search for the first prefix sum larger than curprefixsum - k (logm).
// No need to compute the prefix sum in the first place.
// Space: O(1)
func maxNoLargerThanKInLinear(array []int, k int) int {
	if len(array) == 0 {
		return 0
	}
	// prefix sum
	prefixSums := make([]int, len(array)+1)
	for i, a := range array {
		prefixSums[i+1] = prefixSums[i] + a
	}
	max := math.MinInt32
	for i := 1; i < len(prefixSums); i++ {
		for j := 0; j < i; j++ {
			sum := prefixSums[i] - prefixSums[j]
			if sum > max && sum <= k {
				max = sum
			}
		}
	}
	return max
}
