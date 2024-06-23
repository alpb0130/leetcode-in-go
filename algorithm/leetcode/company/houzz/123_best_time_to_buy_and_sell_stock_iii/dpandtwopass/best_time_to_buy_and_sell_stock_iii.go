package dpandtwopass

import "math"

// First pass to record the max profile cutting at each position (left to right)
// Second pass is to get the max profile cutting at each poistion but from right to left
// For second pass, we don't need extra space anymore because we can just keep updating the
// final results.
// Time: O(n)
// Space: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices))
	min := prices[0]
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		maxProfit = maxInt(maxProfit, prices[i]-min)
		dp[i] = maxProfit
		min = minInt(min, prices[i])
	}

	res := 0
	max := prices[len(prices)-1]
	for i := len(prices) - 1; i >= 0; i-- {
		res = maxInt(res, dp[i]+(max-prices[i]))
		max = maxInt(max, prices[i])
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
