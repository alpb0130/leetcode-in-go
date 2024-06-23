package binarysearcheasytounderstand

// Binary search but more cases. For every loop (left <= right)
// 1. if nums[mid] = target, return mid
// Then we need to figure out which part the mid is in (first or second)
// 2. if nums[mid] >= nums[left] (in the first part)
//    2.1 if nums[mid] > target (two cases)
//        2.1.1 target >= nums[left] (target in the first), right = mid - 1
//        2.1.2 target < nums[left] (target in the second), left = mid + 1
//    2.2 if nums[mid] < target, target must be in the first part and we do left = mid + 1
// 3. if nums[mid] < nums[left] (in the second part)
//    3.1 nums[mid] > target, target must be in the second part and we do right = mid - 1
//    3.2 nums[mid] < target (two case)
//        3.2.1 target >= nums[left], then target is in the first part and we do right = mid - 1
//        3.2.2 target < nums[left], then target is in the second part and we do left = mid + 1
// Time: O(logn)
// Space: O(n)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[left] {
			if nums[mid] > target {
				if target >= nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] > target {
				right = mid - 1
			} else {
				if target >= nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	return -1
}
