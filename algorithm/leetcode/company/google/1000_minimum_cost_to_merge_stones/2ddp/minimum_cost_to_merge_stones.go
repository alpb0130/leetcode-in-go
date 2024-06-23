package _ddp

import "math"

// We don't need the third dimension K actually. Use can use some greedy thoughts when we processing
// the DP. For every subarray, we try to merge stones as much as possible.
// Reference: https://www.youtube.com/watch?v=FabkoUzs64o
// Reference: https://leetcode.com/problems/minimum-cost-to-merge-stones/discuss/247567/JavaC%2B%2BPython-DP
// Time: O(n^3/K)
// Space: O(n^2)
func mergeStones(stones []int, K int) int {
	if len(stones) == 1 {
		return 0
	}
	if len(stones) < K || (len(stones)-1)%(K-1) != 0 {
		return -1
	}
	prefixSum := make([]int, len(stones)+1)
	for i := 0; i < len(stones); i++ {
		prefixSum[i+1] = prefixSum[i] + stones[i]
	}
	dp := make([][]int, len(stones))
	for i := range dp {
		dp[i] = make([]int, len(stones))
	}
	for i := 0; i < len(stones); i++ {
		// start from i-1 to 0. We don't need to visit dp[i][i] because they are already been set to 0
		// which is correct
		for j := i - 1; j >= 0; j-- {
			dp[j][i] = math.MaxInt32
			// The step is K-1. We don't need to look at dp[j][j+1]+dp[j+2][i] if K=3 because
			// we know [j, j+1] cannot be merged into one pile and it will be take care of when we
			// process dp[j][j]+dp[j+1][i]
			for k := j; k < i; k += K - 1 {
				dp[j][i] = minInt(dp[j][i], dp[j][k]+dp[k+1][i])
			}
			// If can merge into one pile, we also need to add the total # of stones within
			// the subarray
			if (i-j)%(K-1) == 0 {
				dp[j][i] += prefixSum[i+1] - prefixSum[j]
			}
		}
	}

	return dp[0][len(stones)-1]
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
