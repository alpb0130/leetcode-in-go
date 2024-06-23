package twopointer

import "math"

// Use two pointer. The brute-force method is for each left bar, we iterate over each right bar
// and update the max container size.
// By using two pointer, we start from two ends and compute the water we can hold at this point.
// Then we move the pointer with lower height one step inside because it's impossible for the low
// bar to have a better results.
// Reference: https://leetcode.com/problems/container-with-most-water/solution/
// One more: https://leetcode.com/problems/container-with-most-water/discuss/6099/Yet-another-way-to-see-what-happens-in-the-O(n)-algorithm
// Time: O(n)
// Space: O(1)
func maxArea(height []int) int {
	res := 0
	start, end := 0, len(height)-1
	for start < end {
		min := minInt(height[start], height[end])
		if min*(end-start) > res {
			res = min * (end - start)
		}
		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
