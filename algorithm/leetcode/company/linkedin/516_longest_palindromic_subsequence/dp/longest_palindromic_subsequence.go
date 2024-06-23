package dp

import "math"

// DP. Start with brute force solution - we can solve this problem recursively by
// keep move left and right point inside which we start from 0 and len(s)-1.
// If s[left] == s[right], the result is res(left+1, right-1)+2 (call 1).
// Else the result is max(res(left, right-1), res(left+1, right)) (call 2 and 3).
// We can find that in the recursive solution, we compute many sub problem many times.
// Like some value may be computed in call 1, 2, 3. We can memorize that. Then the solution
// can be evolved into DP. We can do that recursively and iteratively. If do that iteratively,
// we need to for loop. One is loop over the left index and the inner is to loop over the right
// index. We need to iterate from end of the string to start because we are look back the results
// from inside substring.
// Reference: https://leetcode.com/problems/longest-palindromic-subsequence/discuss/99111/Evolve-from-brute-force-to-dp
// Time: O(n^2)
// Space: O(n^2)
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = maxInt(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][len(s)-1]
}
func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
