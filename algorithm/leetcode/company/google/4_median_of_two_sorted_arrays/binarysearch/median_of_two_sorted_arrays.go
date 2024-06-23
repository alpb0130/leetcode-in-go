package binarysearch

import "math"

// Use binary search. Detailed explanation is here:
// https://leetcode.com/problems/median-of-two-sorted-arrays/solution/
// We pick two cut points (i, j) in both array. Then each array is splited into
// two part: [0, i], [i+1, m] or [0, j], [j+1, n]. It's possible i is -1 which
// means the whole nums1 is belong the the large part.
// Time: O(log(min(m, n)))
// Space: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	}
	// We always iterate over short array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	left, right := -1, m-1
	for left <= right {
		i := (left + right) / 2
		j := (m+n)/2 - i - 2
		if j+1 < n && i >= 0 && nums1[i] > nums2[j+1] {
			right = i - 1
		} else if i+1 < m && j >= 0 && nums2[j] > nums1[i+1] {
			left = i + 1
		} else {
			if (m+n)%2 == 0 {
				// If there are two medians, we  know one of them master from the small part and
				// the other is from the large part.
				var lmax, rmin int
				// Get larger one from two smaller part. It's possible one of the idx is not valid.
				if i >= 0 && j >= 0 {
					lmax = maxInt(nums1[i], nums2[j])
				} else if i >= 0 {
					lmax = nums1[i]
				} else {
					lmax = nums2[j]
				}
				// Get smaller one from two larger part. It's possible one of the idx is not valid.
				if i+1 < m && j+1 < n {
					rmin = minInt(nums1[i+1], nums2[j+1])
				} else if i+1 < m {
					rmin = nums1[i+1]
				} else {
					rmin = nums2[j+1]
				}
				return (float64(lmax) + float64(rmin)) / 2
			} else {
				// If there is only on median, we know it's must be in the second part
				// by set the m and n use real number
				if i+1 < m && j+1 < n {
					return float64(minInt(nums1[i+1], nums2[j+1]))
				} else if i+1 < m {
					return float64(nums1[i+1])
				} else {
					return float64(nums2[j+1])
				}
			}
		}
	}
	return 0.0
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
