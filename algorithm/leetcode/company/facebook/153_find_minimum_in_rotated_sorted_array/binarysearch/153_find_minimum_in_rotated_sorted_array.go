package binarysearch

// Binary search. Use nums[right] to determine which part the min value should be and don't use
// nums[left]!!! Because it's not guaranteed that the array is rotated.
// Our idea is still to decide which part we should go to find the min value. If mid is in the
// large part, we should go to right. If mid is in the small part, we should go to left. But the
// tricky thing is the array is not guaranteed to be rotated so we should use the right value
// to determine which part we should go. If mid==0 or nums[mid-1]>nums[mid],we know this is the
// min value.
// Time: O(logn)
// Space: O(1)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= nums[right] {
			if mid == 0 || nums[mid-1] > nums[mid] {
				break
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return nums[(left+right)/2]
}
