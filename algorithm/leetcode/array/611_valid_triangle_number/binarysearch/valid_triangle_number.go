package binarysearch

import "sort"

// Fix the first number and second number and use binary search to find the third number.
// The problem here is n1 and n2 are not in the same direction as n3. This make it impossible
// to use two pointers.
// Time: O(n*n*logn)
// Space: O(logn) - sorting need O(logn) space
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			k := binarySearchHelper(nums[i]+nums[j], j+1, len(nums), nums)
			res += k - j
		}
	}
	return res
}

func binarySearchHelper(sum int, start, end int, nums []int) int {
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] < sum && (mid >= len(nums)-1 || nums[mid+1] >= sum) {
			return mid
		} else if nums[mid] < sum {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start - 1
}
