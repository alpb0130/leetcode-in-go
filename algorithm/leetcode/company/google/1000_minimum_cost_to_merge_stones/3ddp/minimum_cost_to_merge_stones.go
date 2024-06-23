package _ddp

import "math"

// The brute force solution is try to fix the start position and search the who space to
// get all possible steps.
// For a problem F(i, j, k), it means the cost to split stones from i to j into k piles.
// This means the final results should be F(0, len(stones)-1, 1). In order to get result
// for F(i, j, k), what we need to do is try to consider merge [i, k] in one pile and
// the result of stones [k+1, j] in k-1 piles. But one thing to improve is only if len([i, k])
// is K, we can split merge them into 1 pile when we try to merge the first, part, we can
// move K-1 steps for every move.
// Another thing worth to be careful is every time when we finish processing all k different
// cases for a sub array. We should try to update F(i, j, 1) because if we can reach to
// F(i, j, K), this means the K piles could be merged into 1. BTW, we also need to add the total
// sum for this subarray.
// Reference: https://www.youtube.com/watch?v=FabkoUzs64o
// Time: O(n^3) - n is # of stones
// Space: O(n^3)
func mergeStones(stones []int, K int) int {
	if len(stones) == 1 {
		return 0
	}
	if len(stones) < K || (len(stones)-1)%(K-1) != 0 {
		return -1
	}
	dp := make([][][]int, len(stones))
	for i := range dp {
		dp[i] = make([][]int, len(stones))
		for j := range dp[i] {
			dp[i][j] = make([]int, K+1)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MaxInt32
			}
		}
		dp[i][i][1] = 0
	}
	prefixSum := make([]int, len(stones)+1)
	for i := 0; i < len(stones); i++ {
		prefixSum[i+1] = prefixSum[i] + stones[i]
	}
	for i := 0; i < len(stones); i++ {
		for j := i; j >= 0; j-- {
			for k := 2; k <= K; k++ {
				for l := j + 1; l <= i; l += K - 1 {
					dp[j][i][k] = minInt(dp[j][i][k], dp[j][l-1][1]+dp[l][i][k-1])
				}
			}
			if dp[j][i][K] != math.MaxInt32 {
				dp[j][i][1] = dp[j][i][K] + prefixSum[i+1] - prefixSum[j]
			}
		}
	}
	return dp[0][len(stones)-1][1]
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
