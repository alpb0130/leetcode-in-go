package binarysearch

// The key point to solve this problem is to find some descending subarray and ascending subarray.
// We can use binary search to do that. If the mid point is smaller than its next, we know the
// current number is on the ascending array and the peak must be on the right. If the mid point is
// larger than its next, we know the current number is on the descending array and the peak must
// be on the left side (including itself).
// Time: O(logn)
// Space: O(n)
func findPeakElement(nums []int) int {
	i, j := 0, len(nums)-1
	for i < j {
		mid := (i + j) / 2
		if nums[mid] < nums[mid+1] {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return (i + j) / 2
}
