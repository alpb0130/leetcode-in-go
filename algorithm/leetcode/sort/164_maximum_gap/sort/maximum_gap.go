package sort

import "sort"

// Sort and get the max gap in sorted form array between successive elements
// Reference: https://leetcode.com/problems/maximum-gap/solution/
// Time: O(nlogn)
// Space: O(1)
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > max {
			max = diff
		}
	}
	return max
}
