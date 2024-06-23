package dp

import "math"

// DP. For problem coinChange(k), it can be solved by
// min(coinChange(k-c1), coinChange(k-c2), ..., coinChange(k-cn))
// Time: O(amount*len(coins))
// Space: O(amount)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = minInt(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
