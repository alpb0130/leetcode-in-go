package squaretime

import "math"

// Fix the center bar and we use left and right to find the width (all bar within [left, right])
// should be equal to or higher than the current bar
// Time: O(n^2)
// Space: O(1)
func largestRectangleArea(heights []int) int {
	res := 0
	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for left-1 >= 0 && heights[left-1] >= heights[i] {
			left--
		}
		for right+1 < len(heights) && heights[right+1] >= heights[i] {
			right++
		}
		res = maxInt(res, heights[i]*(right-left+1))
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
