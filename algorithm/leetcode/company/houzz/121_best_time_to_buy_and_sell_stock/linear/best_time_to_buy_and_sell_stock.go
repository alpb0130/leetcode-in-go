package linear

import "math"

// We are only allowed to bug and sell stock once. We need to keep updating the min stock price
// if we find a smaller one. And we also need to update the max profit when we see a larger
// profit if we decide to sell stock at this position
// Time: O(n)
// Space: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	min := prices[0]
	for _, price := range prices {
		res = maxInt(res, price-min)
		min = minInt(min, price)
	}
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
