package bruteforce

import "math"

// Process children one by one, for ascending order, we just keep increase the # of candies we want
// to give to this child. Once we found a drop, we only give 1 candy (greedy). But we also want to
// check whether the previous (larger) children also only has equal number of candies. If yes, we
// need to increase it by 1 until we see a drop when we visit back.
// Time: O(n^2)
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
			for j := i - 1; j >= 0; j-- {
				if ratings[j] > ratings[j+1] {
					candies[j] = maxInt(candies[j], candies[j+1]+1)
				} else {
					break
				}
			}
		}
	}
	res := 0
	for _, c := range candies {
		res += c
	}
	return res
}
func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
