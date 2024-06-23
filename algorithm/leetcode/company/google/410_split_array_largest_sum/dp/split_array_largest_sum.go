package dp

import "math"

// DP. The brute force method is enum all possible split and find the answer which is
// Comb(n-1, m-1). But it has many duplicate process. Like for [subarray, [a1], [a2, a3]]
// and [subarray, [a1, a2], [a3]], we actually deal with subarray twice for the case where
// we slip array into m-2 groups. So we can use DP to record that when we first process that.
// Init the dp array so that all value is max int except we need to make dp[0][0] to 0
// so that we can start dp process.
// Time: O(n^2*m)
// Space: O(n*m)
func splitArray(nums []int, m int) int {
	// P(n, m) = min(sum(i)), max(P(i, m-1))
	n := len(nums)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = minInt(dp[i][j], maxInt(dp[k][j-1], prefixSum[i]-prefixSum[k]))
			}
		}
	}
	return dp[n][m]
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
