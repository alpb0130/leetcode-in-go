package transposeandreverserow

// Two pass. Transpose and then reverse row.
// Time: O(n^2)
// Space: O(1)
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
