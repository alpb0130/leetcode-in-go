package onepass

// In one time, we rotate 4 numbers.
// Time: O(n^2)
// Space: O(1)
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i], matrix[i][j] =
				matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i]
		}
	}
}
