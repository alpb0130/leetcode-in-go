package layer

// The whole top line. The remain right line. Check 1 row or 1 column case.
// The remaining bottom line. The left line.
// Time: O(m*n)
// Space: O(1)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	left, right, up, bottom := 0, n-1, 0, m-1
	res := make([]int, m*n)
	index := 0
	for left <= right && up <= bottom {
		for i := left; i <= right; i++ {
			res[index] = matrix[up][i]
			index++

		}
		for i := up + 1; i <= bottom; i++ {
			res[index] = matrix[i][right]
			index++

		}
		if bottom == up || left == right {
			break
		}
		for i := right - 1; i >= left; i-- {
			res[index] = matrix[bottom][i]
			index++

		}
		for i := bottom - 1; i > up; i-- {
			res[index] = matrix[i][left]
			index++

		}
		left++
		right--
		up++
		bottom--
	}
	return res
}
