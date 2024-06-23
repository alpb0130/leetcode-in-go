package memorize

// Go through the matrix. The brute force method is for each position, we get the enemy number
// in O(n+m) time. But we notice this can be improved because if we calculate the row for
// grid[i][j], we don't need to do that again for grid[i][j+1] because the row count is the same.
// Every time we hit the first element in a row or the previous place is
// 'W', we recalculate the row enemy. Every time we hit the first element in a column or the
// previous place is 'W', we calculate the column enemy. We only need one variable for row because
// go through the matrix rowwise. We need an array for column value. Actually every value will be
// visit at most twice.
// Time: O(m*n)
// Space: O(n)
func maxKilledEnemies(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	row := 0
	columns := make([]int, n)
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'W' {
				continue
			}
			if j == 0 || grid[i][j-1] == 'W' {
				// calculate row enemy
				row = rowEnemy(grid, i, j)
			}
			if i == 0 || grid[i-1][j] == 'W' {
				// calculate column enemy
				columns[j] = colEnemy(grid, i, j)
			}
			if grid[i][j] == '0' {
				if row+columns[j] > res {
					res = row + columns[j]
				}
			}
		}
	}
	return res
}

func rowEnemy(grid [][]byte, i, j int) int {
	res := 0
	for j < len(grid[i]) && grid[i][j] != 'W' {
		if grid[i][j] == 'E' {
			res++
		}
		j++
	}
	return res
}

func colEnemy(grid [][]byte, i, j int) int {
	res := 0
	for i < len(grid) && grid[i][j] != 'W' {
		if grid[i][j] == 'E' {
			res++
		}
		i++
	}
	return res
}
