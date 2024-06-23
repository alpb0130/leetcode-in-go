package dp

import "math"

// Time: O(n)
// Space: O(n)
func rob(nums []int) int {
	l := len(nums)
	res := make([]int, l+2)
	for i := 0; i < l; i++ {
		res[i+2] = maxInt(res[i]+nums[i], res[i+1])
	}
	return res[l+1]
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
