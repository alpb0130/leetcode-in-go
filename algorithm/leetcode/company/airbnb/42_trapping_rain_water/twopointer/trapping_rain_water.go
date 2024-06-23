package twopointer

import "math"

// Two pointer. Initialize the two pointer and point to the start and end respectively.
// Use a global variable "min" to record current minimum boundary bar. It should be an
// incremental variable because no inner smaller boundary can replace the outer larger
// boundary. Move the pointer with smaller height inside for one step. If the height of
// new bar is larger than current min, update it else it must be smaller than the min
// height and I believe we can trap water in this case.
// Time: O(n)
// Space: O(1)
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	start, end := 0, len(height)-1
	min := minInt(height[start], height[end])
	res := 0
	for start < end {
		idx := 0
		if height[start] <= height[end] {
			start++
			idx = start
		} else {
			end--
			idx = end
		}
		if height[idx] <= min {
			res += min - height[idx]
		} else {
			min = minInt(height[start], height[end])
		}
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
