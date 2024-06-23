package dfs

// Time: exponential
// Improve by dp
func NumOfPath(m, n int) int {
	return dfs(m, n, m-1, n-1)
}

func dfs(m, n, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if i == 0 && j == 0 {
		return 1
	}
	return dfs(m, n, i-1, j) + dfs(m, n, i, j-1)
}
