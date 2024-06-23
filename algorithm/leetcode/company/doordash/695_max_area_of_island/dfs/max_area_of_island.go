package dfs

// DFS
// Time: O(m*n)
// Space: O(m*n) - can remove the isVisited matrix by reset visited island to 0
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	isVisited := make([][]bool, m)
	for i := range isVisited {
		isVisited[i] = make([]bool, n)
	}
	res := 0
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !isVisited[i][j] && grid[i][j] == 1 {
				area := dfs(grid, i, j, isVisited, dir)
				if area > res {
					res = area
				}
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int, isVisited [][]bool, dir [][]int) int {
	isVisited[i][j] = true
	res := 1
	for _, d := range dir {
		ii := i + d[0]
		jj := j + d[1]
		if ii >= 0 && ii < len(grid) && jj >= 0 && jj < len(grid[0]) && !isVisited[ii][jj] && grid[ii][jj] == 1 {
			res += dfs(grid, ii, jj, isVisited, dir)
		}
	}
	return res
}
