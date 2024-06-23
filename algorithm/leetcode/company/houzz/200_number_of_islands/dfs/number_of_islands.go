package dfs

// DFS. No extra visit matrix
// Time: O(m*n)
// Space: O(1)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	res := 0
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				//dfs
				dfs(grid, i, j, dir)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, i, j int, dir [][]int) {
	grid[i][j] = '2'
	for _, d := range dir {
		ii := i + d[0]
		jj := j + d[1]
		if isValid(len(grid), len(grid[0]), ii, jj) {
			if grid[ii][jj] == '1' {
				dfs(grid, ii, jj, dir)
			}
		}
	}
}

func isValid(m, n, i, j int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
