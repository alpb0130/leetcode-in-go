package dp

import "math"

// Optimization problem and use DP.
// Time: O(n)
// Space: O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	res := max
	for i := 1; i < len(nums); i++ {
		max = maxInt(nums[i], nums[i]+max)
		res = maxInt(max, res)
	}
	return res
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
