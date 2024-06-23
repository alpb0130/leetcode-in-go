package dfsandmemo

import "math"

// DFS and memorization. For every position, we can use dfs to find the longest increasing path of it's
// neighbor position. But obviously, once we find the results for one position, it can be reused again
// when we process other positions.
// Time: O(m*n)
// Space: O(m*n)
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// dfs
			res = maxInt(res, dfs(matrix, m, n, i, j, dir, dp))
		}
	}
	return res
}

func dfs(matrix [][]int, m, n, i, j int, dir, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	dp[i][j] = 1
	max := 0
	for _, d := range dir {
		ii, jj := i+d[0], j+d[1]
		if isValid(m, n, ii, jj) && matrix[ii][jj] > matrix[i][j] {
			max = maxInt(max, dfs(matrix, m, n, ii, jj, dir, dp))
		}
	}
	dp[i][j] += max
	return dp[i][j]
}

func isValid(m, n, i, j int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
