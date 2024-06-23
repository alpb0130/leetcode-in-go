package dp

import "math"

// DP. Use three different slice to cache the min cost for each color of a specific house.
// Time: O(n)
// Space: O(n)
func minCost(costs [][]int) int {
	dp := make([][]int, len(costs)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 3)
	}
	for i, cost := range costs {
		dp[i+1][0] = cost[0] + minInt(dp[i][1], dp[i][2])
		dp[i+1][1] = cost[1] + minInt(dp[i][0], dp[i][2])
		dp[i+1][2] = cost[2] + minInt(dp[i][0], dp[i][1])
	}
	res := minInt(dp[len(costs)][0], minInt(dp[len(costs)][1], dp[len(costs)][2]))
	return res
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
