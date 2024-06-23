package _passdp

import "math"

// Apply DP twice. One is from up-left to down-right. The other is from down-right to up-left.
// What if there is no 0 in the matrix??
// There is also a BFS solution which take O(m*n) because every number should only been pushed into
// queue once. Reference: https://leetcode.com/problems/01-matrix/discuss/101021/Java-Solution-BFS
// Time: O(m*n)
// Space: O(1)
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	res := make([][]int, len(matrix))
	for i := range res {
		res[i] = make([]int, len(matrix[0]))
		for j := range res[i] {
			res[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[0]); j++ {
			if matrix[i][j] == 0 {
				res[i][j] = 0
			} else {
				if i > 0 {
					res[i][j] = minInt(res[i][j], res[i-1][j]+1)
				}
				if j > 0 {
					res[i][j] = minInt(res[i][j], res[i][j-1]+1)
				}
			}
		}
	}
	for i := len(res) - 1; i >= 0; i-- {
		for j := len(res[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				res[i][j] = 0
			} else {
				if i < len(res)-1 {
					res[i][j] = minInt(res[i][j], res[i+1][j]+1)
				}
				if j < len(res[0])-1 {
					res[i][j] = minInt(res[i][j], res[i][j+1]+1)
				}
			}
		}
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
