package dp

import "math"

// We could check number 1 to x and use is ugly number to get nth ugly number but it's very
// slow. What we do is use 2, 3, 5 to generate ugly numbers. Start from 1, we can go through
// three list which are existing ugly number * 2, 3 and 5 respectively. What we should take is
// different multiplier has different current ugly number. Every time a ugly number is processed,
// we should move forward the ptr for the multiplier one step ahead. Why's dp? We stop existing
// ugly number in a array and don't need to recompute them again. Just maintain a pointer for each
// multiplier group to point to current processing ugly number.
// Reference: https://www.geeksforgeeks.org/ugly-numbers/
// Time: O(n) - to generate kth ugly number, we need constant time
// Space: O(n) - to store existing ugly number
func nthUglyNumber(n int) int {
	nums := []int{1}
	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n; i++ {
		min := minOfThree(nums[i2]*2, nums[i3]*3, nums[i5]*5)
		if min == nums[i2]*2 {
			i2++
		}
		if min == nums[i3]*3 {
			i3++
		}
		if min == nums[i5]*5 {
			i5++
		}
		nums = append(nums, min)
	}
	return nums[n-1]
}

func minOfThree(i, j, k int) int {
	return int(math.Min(math.Min(float64(i), float64(j)), float64(k)))
}
