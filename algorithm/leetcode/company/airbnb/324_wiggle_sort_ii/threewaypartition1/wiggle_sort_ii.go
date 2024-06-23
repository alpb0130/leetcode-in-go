package threewaypartition1

import "math/rand"

// 1. Linear time to find the median (not exactly median but some middle number)
// 2. Use three-way partition to put the elements in correct position. We try to
// put the median number in the "middle", which is the left of even part and the
// right of odd part. Large number in the "right" which is the left of odd part.
// Small numbers in the "left" which is the right of even part.
// Reference: https://leetcode.com/problems/wiggle-sort-ii/discuss/77684/Summary-of-the-various-solutions-to-Wiggle-Sort-for-your-reference
// Time: O(n)
// Space: O(1)
func wiggleSort(nums []int) {
	// find the median in linear time. Not exactly median but the (len(nums)-1)/2 th largest number
	n := len(nums)
	mid := findMedian(nums, len(nums)/2)
	i, low, high := 0, 0, len(nums)-1
	for i <= high {
		if nums[indexMapping(i, n)] > mid {
			nums[indexMapping(i, n)], nums[indexMapping(low, n)] =
				nums[indexMapping(low, n)], nums[indexMapping(i, n)]
			i++
			low++
		} else if nums[indexMapping(i, n)] < mid {
			nums[indexMapping(i, n)], nums[indexMapping(high, n)] =
				nums[indexMapping(high, n)], nums[indexMapping(i, n)]
			high--
		} else {
			i++
		}
	}

}

func indexMapping(i, n int) int {
	return (2*i + 1) % (n | 1)
}

func findMedian(nums []int, k int) int {
	low, high := 0, len(nums)-1
	// shuffle
	for i := 0; i < len(nums)-1; i++ {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
	// quick selection to find median
	for {
		// if cannot find the median
		if low > high {
			break
		}
		l := partition(nums, low, high)
		if l > k {
			high = l
		} else if l < k {
			low = l + 1
		} else {
			break
		}
	}
	return nums[k]
}

func partition(nums []int, low, high int) int {
	i := low
	j, k := low+1, high
	for {
		for j <= high && nums[j] <= nums[i] {
			j++
		}
		for k >= low+1 && nums[k] >= nums[i] {
			k--
		}
		if j > k {
			nums[i], nums[k] = nums[k], nums[i]
			break
		}
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}

	return k
}
