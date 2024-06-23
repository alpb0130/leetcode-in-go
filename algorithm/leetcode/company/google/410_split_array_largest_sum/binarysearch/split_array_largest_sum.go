package binarysearch

import "math"

// The final results must be in the range of [min(nums), sum(nums)]. We can use binary search
// to find a final answer. We know that for a target sum that is larger than res, we can always
// split the array into less or equal than m groups which the max sum is less then it. So we can
// use binary search. To find the min.
// Time: O(n*long(sum(nums))) - for each binary process, we need to process the whole array
//       to valid whether it can be splitted or not
// Space: O(1)
func splitArray(nums []int, m int) int {
	// assume len(nums) >= m
	if len(nums) == 0 {
		return 0
	}
	left, right := math.MinInt32, 0
	for _, num := range nums {
		left = minInt(left, num)
		right += num
	}
	res := 0
	for left <= right {
		mid := (left + right) / 2
		if validate(nums, mid, m) {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func validate(nums []int, max int, m int) bool {
	sum := 0
	numGroup := 0
	for _, num := range nums {
		if num > max {
			return false
		}
		if num+sum > max {
			numGroup++
			sum = num
		} else {
			sum += num
		}
	}
	numGroup++
	return m >= numGroup
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
