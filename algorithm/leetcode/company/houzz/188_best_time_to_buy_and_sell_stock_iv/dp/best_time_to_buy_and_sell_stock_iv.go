package dp

import "math"

// Explanation: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/39608/A-clean-DP-solution-which-generalizes-to-k-transactions
//              https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/135704/Detail-explanation-of-DP-solution
// Reference: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/discuss/54113/A-Concise-DP-Solution-in-Java
// Time: O(kn)
// Space: O(kn)
func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) <= 1 {
		return 0
	}
	if k >= len(prices)/2 {
		res := 0
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				res += prices[i] - prices[i-1]
			}
		}
		return res
	}
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(prices))
	}
	for i := 1; i < k+1; i++ {
		tmpProfit := dp[i-1][0] - prices[0]
		for j := 1; j < len(prices); j++ {
			dp[i][j] = maxInt(dp[i][j-1], prices[j]+tmpProfit)
			tmpProfit = maxInt(tmpProfit, dp[i-1][j]-prices[j])
		}
	}
	return dp[k][len(prices)-1]
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
