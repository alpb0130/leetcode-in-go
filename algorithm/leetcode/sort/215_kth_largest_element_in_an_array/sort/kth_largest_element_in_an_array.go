package sort

import "sort"

// Sort
// Time: O(nlogn)
// Space: O(1)
func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	sort.Ints(nums)
	return nums[len(nums)-k]
}
