package spiral_matrix_inside_out

import "fmt"

func SpiralOrder(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	i, j := (n-1)/2, (n-1)/2
	step := 1
	firstDir := [][]int{{0, 1}, {0, -1}}
	secondDir := [][]int{{1, 0}, {-1, 0}}
	fmt.Println(matrix[i][j])
	for {
		for k := 0; k < step; k++ {
			i, j = i+firstDir[(step-1)%2][0], j+firstDir[(step-1)%2][1]
			if !isValid(n, i, j) {
				return
			}
			fmt.Println(matrix[i][j])
		}
		for k := 0; k < step; k++ {
			i, j = i+secondDir[(step-1)%2][0], j+secondDir[(step-1)%2][1]
			if !isValid(n, i, j) {
				return
			}
			fmt.Println(matrix[i][j])
		}
		step++
	}
}

func isValid(n, i, j int) bool {
	return i >= 0 && i < n && j >= 0 && j < n
}
