package dp

import "math"

// The brute force method is for each bar, iterate the array and identify the bound for each bar
// but we don't need to do that for each bar because we can store the value somewhere.
// What we should do is: iterate over from left to right and identify the left bound. Then
// iterate over from right to left and identify the right bound. Then compute the results.
// This method need two pass and an array to store the left boundary for each bar.
// Time: O(n)
// Space: O(n)
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	barMap := make([]int, len(height))
	max := height[0]
	for i := 0; i < len(height); i++ {
		h := height[i]
		if h > max {
			max = h
		}
		barMap[i] = max
	}
	max = height[len(height)-1]
	for i := len(height) - 1; i >= 0; i-- {
		h := height[i]
		if h > max {
			max = h
		}
		barMap[i] = minInt(barMap[i], max)
		res += (barMap[i] - h)
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
