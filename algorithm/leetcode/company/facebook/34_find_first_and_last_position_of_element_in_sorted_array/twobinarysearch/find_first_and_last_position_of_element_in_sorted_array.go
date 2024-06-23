package twobinarysearch

// Apply two binary search
// Time: O(logn)
// Space: O(1)
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	l := len(nums)
	if len(nums) == 0 || nums[0] > target || nums[l-1] < target {
		return res
	}
	left, right := 0, l-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target && (mid == 0 || nums[mid-1] != target) {
			res[0] = mid
			break
		} else if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if res[0] == -1 {
		return res
	}
	left, right = 0, l-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target && (mid == l-1 || nums[mid+1] != target) {
			res[1] = mid
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
