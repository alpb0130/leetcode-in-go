package constantspace

import "math"

// Time: O(n)
// Space: O(1) - reuse nums to record whether a number is in it.
func firstMissingPositive(nums []int) int {
	// all negative number and 0 to 1
	// because we use 1, we need to check whether 1 is in the original list
	// if not, 1 is the ans. Otherwise we are good to use 1
	hasOne := false
	for i := range nums {
		if nums[i] == 1 {
			hasOne = true
		}
		if nums[i] <= 0 {
			nums[i] = 1
		}
	}
	if !hasOne {
		return 1
	}
	// if len(nums) = k, then the result must be in [1, k+1]
	// use position info as map to record whether a number is in the list or not
	// we have convert all 0 and - to 1 so what we can do is if the number is in
	// the map, set it to be negative.
	for i := range nums {
		idx := int(math.Abs(float64(nums[i])))
		if idx <= len(nums) && nums[idx-1] > 0 {
			nums[idx-1] = -nums[idx-1]
		}
	}
	// Iterate over the list, once we see a positive number, the postion+1 must be the answer.
	res := 0
	for res < len(nums) && nums[res] < 0 {
		res++
	}
	return res + 1
}
