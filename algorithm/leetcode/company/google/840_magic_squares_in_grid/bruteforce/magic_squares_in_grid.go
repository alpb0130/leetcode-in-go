package bruteforce

// Traverse every 3*3 block and check whether every block is magic square.
// For every block, we need to check:
// 1. Only contains number 1-9 without repeat
// 2. Row sum, column sum and diagonal sum is 15.
func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}
	res := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 5 {
				if check(grid, i, j) {
					res++
				}
			}
		}
	}
	return res
}

func check(grid [][]int, a, b int) bool {
	rowSum := 0
	columnSum := [3]int{}
	diag1, diag2 := 0, 0
	numMap := map[int]int{}
	for i := a - 1; i <= a+1; i++ {
		for j := b - 1; j <= b+1; j++ {
			if grid[i][j] < 1 || grid[i][j] > 9 || numMap[grid[i][j]] > 0 {
				return false
			}
			numMap[grid[i][j]]++
			rowSum += grid[i][j]
			columnSum[j-(b-1)] += grid[i][j]
			if i-(a-1) == j-(b-1) {
				diag1 += grid[i][j]
			}
			if i-(a-1) == b+1-j {
				diag2 += grid[i][j]
			}
		}
		if rowSum != 15 {
			return false
		}
		rowSum = 0
	}
	return diag1 == 15 && diag2 == 15 && columnSum[0] == 15 &&
		columnSum[1] == 15 && columnSum[2] == 15
}
