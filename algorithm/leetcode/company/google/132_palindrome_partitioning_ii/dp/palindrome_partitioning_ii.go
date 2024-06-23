package dp

import "math"

// Time: O(n^2)
// Space: O(n^2)
func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	minCut := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		minCut[i+1] = math.MaxInt32
		for j := i; j >= 0; j-- {
			if s[i] == s[j] {
				dp[j][i] = (j+1 > i-1) || dp[j+1][i-1]
			}
			if dp[j][i] {
				minCut[i+1] = minInt(minCut[i+1], minCut[j]+1)
			}
		}
	}

	return minCut[len(s)] - 1
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
