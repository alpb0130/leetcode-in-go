package bellmenford

import "math"

// Time: O(k*|E|) = O(k*n^2)
// Space: O(k*n)
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if src == dst {
		return 0
	}
	inf := 1000000000
	dp := make([][]int, K+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][src] = 0
	for i := 1; i <= K+1; i++ {
		dp[i][src] = 0
		for _, flight := range flights {
			dp[i][flight[1]] = minInt(dp[i][flight[1]], dp[i-1][flight[0]]+flight[2])
		}
	}
	if dp[K+1][dst] == inf {
		return -1
	}
	return dp[K+1][dst]
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
