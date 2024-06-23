package infvalue

import "math"

// Binary search but we modify rotated part to be either Inf or -Inf so that we can do normal
// binary search.
// If target is in the second part, then all number in the first part should be -Inf.
// If target is in the first part, then all number in the second part should be Inf.
// Instead modify the before we do binary search (would be O(n)). We can do the change during
// binary search.
// Time: O(logn)
// Space: O(n)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		num := float64(nums[mid])
		if nums[mid] >= nums[left] && target < nums[left] {
			num = math.Inf(-1)
		} else if nums[mid] < nums[left] && target >= nums[left] {
			num = math.Inf(1)
		}
		if num == float64(target) {
			return mid
		} else if num > float64(target) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
