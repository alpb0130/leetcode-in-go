package binarysearch

// Similar to 33. Search in Rotated Sorted Array but allow duplicate value. Then the challenge is
// it's hard for use to decide whether the mid is in the first part or second part. In this case,
// what we can do is don't decide which side the target value is but only move left and right into
// middle by 1 step so that we can reduce our search range because we can pretty sure in this case
// the target value is not both end of the range.
// Time: O(logn). Worst case O(n)
// Space: O(n)
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[mid] < target {
			if nums[mid] >= nums[left] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] < nums[left] || target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
