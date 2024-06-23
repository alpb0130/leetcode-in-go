package twopass

import "math"

// Two pass.
// First pass: from left to right, see increase, candy++; else candy=1
// Second pass: from right to left, see increase, candy = max(candy, prev_candy). We also
// calculate the final results.
// Ex:  1 6 10 8 7 3 2 9 11
// 1st: 1 2 3  1 1 1 1 2 3
// 2nd: 1 2 5  4 3 2 1 2 3
// Time: O(n)
// Space: O(n)
func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	res := candies[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = maxInt(candies[i], candies[i+1]+1)
		}
		res += candies[i]
	}
	return res
}
func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
