package dp

import "math"

// DP. For two words s and t, in order to get the edit distance for s[:i] and t[:j] and
// if we already know the shorted edit distance of a = (s[:i-1], t[:j-1]), b = (s[:i], t[:j-1])
// and c = (s[:i-1], t[:j]), we can the result for s[:i] and t[:j] is:
// 1. If s[i] == t[j], min(a, b+1, c+1)
// 2. Else min(a, b, c)+1
// Time: O(m*n)
// Space: O(m*n)
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			left := dp[i][j-1] + 1
			up := dp[i-1][j] + 1
			leftUp := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftUp++
			}
			dp[i][j] = minInt(leftUp, minInt(left, up))
		}
	}
	return dp[m][n]
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
