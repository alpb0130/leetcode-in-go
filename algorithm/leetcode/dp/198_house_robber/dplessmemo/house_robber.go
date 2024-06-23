package dplessmemo

import "math"

// Similar to DP solution but we don't need to store all computed results. We only need to store
// previous 2 results
// Time: O(n)
// Space: O(1)
func rob(nums []int) int {
	prev1 := 0
	prev2 := 0
	for i := 0; i < len(nums); i++ {
		prev1, prev2 = maxInt(nums[i]+prev2, prev1), prev1
	}
	return prev1
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
