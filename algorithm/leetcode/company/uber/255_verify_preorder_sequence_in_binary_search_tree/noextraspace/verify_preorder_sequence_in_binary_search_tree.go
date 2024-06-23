package noextraspace

import "math"

// Same as stack method but use original array as a stack
func verifyPreorder(preorder []int) bool {
	if len(preorder) == 0 {
		return true
	}
	min := math.MinInt32
	ptr := -1
	for i := 0; i < len(preorder); i++ {
		for ptr >= 0 && preorder[ptr] < preorder[i] {
			min = preorder[ptr]
			ptr--
		}
		if preorder[i] <= min {
			return false
		}
		ptr++
		preorder[ptr] = preorder[i]
	}
	return true
}
