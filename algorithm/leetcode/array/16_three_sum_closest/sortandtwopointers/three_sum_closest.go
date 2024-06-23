package sortandtwopointers

import (
	"math"
	"sort"
)

// Almost the same as three sum but a little bit easier because three is not edge case.
// Time: O(n^2)
// Space: O(logn) - sorting need O(logn) space
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			diff := int(math.Abs(float64(target - sum)))
			resDiff := int(math.Abs(float64(target - res)))
			if diff < resDiff {
				res = sum
			}
			if target == sum {
				return sum
			} else if target > sum {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
