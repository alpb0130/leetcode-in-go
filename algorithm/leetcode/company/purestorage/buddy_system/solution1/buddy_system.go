package solution1

// Time: O(n) - number of nodes
func ClearBit(matrix [][]int, offset, length int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	cur := len(matrix) - 1
	left := offset
	right := offset + length - 1
	for cur >= 0 {
		for i := left; i <= right; i++ {
			matrix[cur][i] = 0
		}
		cur--
		left /= 2
		right /= 2
	}
}

func SetBit(matrix [][]int, offset, length int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	cur := len(matrix) - 1
	left := offset
	right := offset + length - 1
	for cur >= 0 && left <= right {
		for i := left; i <= right; i++ {
			matrix[cur][i] = 1
		}
		leftBit := 1
		if left%2 == 1 {
			leftBit = matrix[cur][left-1]
		}
		rightBit := 1
		if right%2 == 0 && right+1 < len(matrix[cur]) {
			rightBit = matrix[cur][right+1]
		}
		left /= 2
		if leftBit == 0 {
			left++
		}
		right /= 2
		if rightBit == 0 {
			right--
		}
		cur--
	}
}
