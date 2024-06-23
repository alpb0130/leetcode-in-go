package topright

// From top right to search the target.
// If equal, return true
// If matrix[i][j] < target, we don't need to check the current row and we do i++
// If matrix[i][j] > target, we don't need to check the current column and we do j++
// Time: O(m+n)
// Space: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
