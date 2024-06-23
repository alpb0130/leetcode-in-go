package zigzag_matrix

import "fmt"

/*
	[
		[1 2 3 x,
		[4 5 6 x],
		[7 8 9 x]
	]
	1 2 4 7 5 3 6 8 9

     / / /
      / / /
       / / /
*/

func ZigZagMatrix(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	i, j := 0, 0
	directions := [][]int{{-1, 1}, {1, -1}}
	level := 0
	for {
		fmt.Println(matrix[i][j])
		d := directions[level%2]
		nextI, nextJ := i+d[0], j+d[1]
		if isValid(m, n, nextI, nextJ) {
			i, j = nextI, nextJ
		} else {
			if level%2 == 0 {
				nextI = i
				nextJ = j + 1
				if nextJ >= n {
					nextI = i + 1
					nextJ = j
				}
			} else {
				nextI = i + 1
				nextJ = j
				if nextI >= m {
					nextI = i
					nextJ = j + 1
				}
			}
			i = nextI
			j = nextJ
			level++
		}
		if !isValid(m, n, i, j) {
			break
		}
	}
}

func isValid(m, n, i, j int) bool {
	return i >= 0 && i < m && j >= 0 && j < n
}
