package twopointersfromend

// We can start from the end of both array because the high end of num1
// is not used and we don't need extra space.
// Time: O(n)
// Space: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr1 := m - 1
	ptr2 := n - 1
	ptr := m + n - 1
	for ptr1 >= 0 && ptr2 >= 0 {
		if nums1[ptr1] > nums2[ptr2] {
			nums1[ptr] = nums1[ptr1]
			ptr1--
		} else {
			nums1[ptr] = nums2[ptr2]
			ptr2--
		}
		ptr--
	}
	for ptr2 >= 0 {
		nums1[ptr] = nums2[ptr2]
		ptr--
		ptr2--
	}
}
