package binarysearchwithoutconsolidation

// Binary search but more cases. For every loop (left <= right)
// 1. if nums[mid] = target, return mid
// 2. if nums[mid] < target
//    2.1 if mid is in the first part, then target will always be on the right side and we do left = mid+1
//    2.2 if mid is in the second part
//        2.2.1 if target is in the first part, then we do right = mid-1
//        2.2.2 if target is in the second part, then we do left = mid+1
// 3. if nums[mid] > target
//    3.1 if mid is in the first part
//        3.1.1 if target is in the first part, then we do right = mid-1
//        3.1.2 if target is in the second part, then we do left = mid+1
//    3.2 if mid is in the second part, then target value will always be on the left side and we do right = mid-1
// Time: O(logn)
// Space: O(n)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			if nums[mid] >= nums[left] {
				left = mid + 1
			} else {
				if target >= nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		} else {
			if nums[mid] >= nums[left] {
				if target >= nums[left] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
