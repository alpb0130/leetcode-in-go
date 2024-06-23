package dp

import "math"

// DP solution
// Time: O(m*n)
// Space: O(n)
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	left, right, heights := make([]int, len(matrix[0])), make([]int, len(matrix[0])), make([]int, len(matrix[0]))
	for i := range right {
		left[i] = -1
		right[i] = len(matrix[0])
	}
	for i := 0; i < len(matrix); i++ {
		curLeft, curRight := -1, len(matrix[0])
		// update heights
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		// update left
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				left[j] = maxInt(left[j], curLeft)
			} else {
				left[j] = -1
				curLeft = j
			}
		}
		// update right
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = minInt(right[j], curRight)
			} else {
				right[j] = len(matrix[0])
				curRight = j
			}
		}
		for j := 0; j < len(matrix[0]); j++ {
			res = maxInt(res, heights[j]*(right[j]-left[j]-1))
		}
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
