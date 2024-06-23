package binarysearch2d

// Binary search row and then column
// Time: O(logm+logn)
// Space: O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	up, down := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	// set to -1
	row := -1
	for up <= down {
		mid := (up + down) / 2
		// should have =
		if matrix[mid][0] <= target {
			row = mid
			up = mid + 1
		} else {
			down = mid - 1
		}
	}
	if row == -1 {
		return false
	}
	for left <= right {
		mid := (left + right) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
