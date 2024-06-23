package greedy

// We can buy and sell as many times as possible. Just do greedy method every time we see
// a position profit, we do the transaction.
// Time: O(n)
// Space: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}
