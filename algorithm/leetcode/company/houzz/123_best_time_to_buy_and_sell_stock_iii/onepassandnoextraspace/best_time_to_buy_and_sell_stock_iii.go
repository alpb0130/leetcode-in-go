package onepassandnoextraspace

import "math"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/39611/Is-it-Best-Solution-with-O(n)-O(1)
// Basically we just keep four variables to keep track of current min first stock cost, max first
// stock profit, second stock cost (relative to first stock profit), max second stock profit
// (relative to second stock profit).
// The final result will be second stock profit
// Time: O(n)
// Space: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	cost1, profit1, cost2, profit2 := math.MaxInt32, 0, math.MaxInt32, 0
	for _, price := range prices {
		cost1 = minInt(cost1, price)
		profit1 = maxInt(profit1, price-cost1)
		cost2 = minInt(cost2, price-profit1)
		profit2 = maxInt(profit2, price-cost2)
	}
	return profit2
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
