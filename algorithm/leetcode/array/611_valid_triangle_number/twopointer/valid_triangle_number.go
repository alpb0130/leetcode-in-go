package twopointer

import "sort"


// Time: O(n)
// Space: O(logn) - sorting need O(logn) space
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := len(nums)-1; i > 1; i-- {
		j := 0
		k := i-1
		for j < k {
			if nums[j] + nums[k] > nums[i] {
				res += k-j
				k--
			} else {
				j++
			}
		}
	}
	return res
}
