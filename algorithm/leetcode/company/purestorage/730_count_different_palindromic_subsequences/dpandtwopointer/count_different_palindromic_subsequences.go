package dpandtwopointer

// https://leetcode.com/problems/count-different-palindromic-subsequences/discuss/109507/Java-96ms-DP-Solution-with-Detailed-Explanation
// Time: O(n^3)
// Space: O(n^2)
func countPalindromicSubsequences(S string) int {
	if len(S) == 0 {
		return 0
	}
	dp := make([][]int, len(S))
	for i := range dp {
		dp[i] = make([]int, len(S))
		dp[i][i] = 1
	}
	for i := 0; i < len(S); i++ {
		for j := i - 1; j >= 0; j-- {
			if S[i] != S[j] {
				dp[j][i] = dp[j][i-1] + dp[j+1][i] - dp[j+1][i-1]
			} else {
				left, right := j+1, i-1
				for left <= right && S[left] != S[j] {
					left++
				}
				for left <= right && S[right] != S[i] {
					right--
				}
				if left > right {
					dp[j][i] = 2*dp[j+1][i-1] + 2
				}
				if left == right {
					dp[j][i] = 2*dp[j+1][i-1] + 1
				}
				if left < right {
					dp[j][i] = 2*dp[j+1][i-1] - dp[left+1][right-1]
				}
			}
			if dp[j][i] < 0 {
				dp[j][i] += 1000000007
			}
			dp[j][i] %= 1000000007
		}
	}
	return dp[0][len(S)-1]
}
