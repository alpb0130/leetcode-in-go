package bucketize

import "math"

// Bucketize the list. The smallest max gap for a array with n elements (min and max) are
// (max-min)/n. We can assign nums to a consecutive buckets with range less than (max-min)/n.
// In this case, we only need to compare the min and max across successive buckets and compute
// the diff. No need to compute the diff between elements within a bucket because we know the
// diff is less than (max-min)/n.
// Reference: https://leetcode.com/problems/maximum-gap/discuss/50643/bucket-sort-JAVA-solution-with-explanation-O(N)-time-and-space
// Time: O(n+b) - n is len(nums). b is # of bucket, which most of time equal to len(nums) but
// not always true if max-min > len(nums)
// Space: O(b) - two maps
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// find min and max
	// decide range
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	width := int(math.Max(float64(1), float64((max-min)/(len(nums)-1))))
	maxMap := map[int]int{}
	minMap := map[int]int{}
	for _, num := range nums {
		i := (num - min) / width
		curMax, ok := maxMap[i]
		if !ok || (ok && num > curMax) {
			maxMap[i] = num
		}
		curMin, ok := minMap[i]
		if !ok || (ok && num < curMin) {
			minMap[i] = num
		}
	}
	result := 0
	prevMax := min
	for i := 0; i < (max-min)/width+1; i++ {
		curMin, ok := minMap[i]
		if ok && curMin-prevMax > result {
			result = curMin - prevMax
		}
		curMax, ok := maxMap[i]
		if ok {
			prevMax = curMax
		}
	}
	return result
}
