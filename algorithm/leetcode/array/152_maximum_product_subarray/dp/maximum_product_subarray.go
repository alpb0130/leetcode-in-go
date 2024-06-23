package dp

import "math"

// Dynamic programming. For max/min problem, consider dynamic programming.
// For the problem, we consider every sub-problem where the array is end up with k.
// If we want to know F(k), we need to know the max and min from F(k-1) where F(k-1)
// means the results that must contain num at index k-1. Why do we need the max and min?
// Because the number could be negative, and if it's negative, the max of this step could be
// num*min(from last step).
// Time: O(n)
// Space: O(1)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	min := nums[0]
	res := max
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}
		max = maxInt(nums[i], max*nums[i])
		min = minInt(nums[i], min*nums[i])
		res = maxInt(max, res)
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
