package dp

import "math"

// The problems is similar to House Robber I but required that all houses are arranged in a circle.
// We can solve this problem by dividing it into two parts:
// 1. House 0 to house n-2
// 2. House 1 to house n-1
// For each parts, use dp solution in House Robber I and return the max value from those two part.
// Be very careful on the edge case: input slice is empty or length is 1
// Time: O(n)
// Space: O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return maxInt(robHelper(nums[:len(nums)-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	prev1, prev2 := 0, 0
	for _, num := range nums {
		prev2, prev1 = prev1, maxInt(num+prev2, prev1)
	}
	return prev1
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
