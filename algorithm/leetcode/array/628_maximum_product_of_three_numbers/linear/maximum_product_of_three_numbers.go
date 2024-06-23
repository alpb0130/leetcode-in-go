package linear

import "math"

// The problem require us to get the max product of three numbers in the array. A straight
// forward idea is to get the three largest number and the product would be the maximum. But
// This is only true if at most 1 number is negative. If there are more than 1 numbers are
// negative like [-5, -10, 1, 2, 3], the max prod would be 150. So what we should is consider
// this edge case. Get the top-3 max (max1 <= max2 <= max3) number and top-2 min number (min1 <= min2)
// And the result would be max(max1*max2*max3, min1*min2*max3).
// Time: O(n)
// Space: O(1)
func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	var max1, max2, max3 *int
	var min1, min2 *int
	for i := 0; i < len(nums); i++ {
		if max1 == nil || nums[i] > *max1 {
			max1 = &nums[i]
		}
		if max2 == nil || nums[i] > *max2 {
			max1 = max2
			max2 = &nums[i]
		}
		if max3 == nil || nums[i] > *max3 {
			max2 = max3
			max3 = &nums[i]
		}
		if min1 == nil || nums[i] < *min1 {
			min1 = &nums[i]
		}
		if min2 == nil || nums[i] < *min2 {
			min1 = min2
			min2 = &nums[i]
		}
	}
	return maxInt((*max1)*(*max2)*(*max3), (*min1)*(*min2)*(*max3))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
